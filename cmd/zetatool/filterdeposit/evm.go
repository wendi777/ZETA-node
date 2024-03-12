package filterdeposit

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/erc20custody.sol"
	"github.com/zeta-chain/protocol-contracts/pkg/contracts/evm/zetaconnector.non-eth.sol"
	"github.com/zeta-chain/zetacore/cmd/zetatool/config"
)

const (
	TopicsDeposited = 2
	TopicsZetaSent  = 3
	DonationMessage = "I am rich!"
)

var evmCmd = &cobra.Command{
	Use:   "evm",
	Short: "Filter inbound eth deposits",
	Run:   FilterEVMTransactions,
}

func init() {
	Cmd.AddCommand(evmCmd)
}

func FilterEVMTransactions(cmd *cobra.Command, _ []string) {
	configFile, err := cmd.Flags().GetString(config.Flag)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	list := GetEthHashList(cfg)
	CheckForCCTX(list, cfg)
}

func GetEthHashList(cfg *config.Config) []deposit {
	startBlock := cfg.EvmStartBlock
	client, err := ethclient.Dial(cfg.EthRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successful")

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	latestBlock := uint64(header.Number.Int64())
	fmt.Println("latest Block: ", latestBlock)

	endBlock := startBlock + cfg.EvmMaxRange
	deposits := make([]deposit, 0)
	segment := 0
	for startBlock < latestBlock {
		fmt.Printf("adding segment: %d, startblock: %d\n", segment, startBlock)
		deposits = append(deposits, GetHashListSegment(client, startBlock, endBlock, cfg)...)
		startBlock = endBlock
		endBlock = endBlock + cfg.EvmMaxRange
		if endBlock > latestBlock {
			endBlock = latestBlock
		}
		segment++
	}
	return deposits
}

func GetHashListSegment(client *ethclient.Client, startBlock uint64, endBlock uint64, cfg *config.Config) []deposit {
	deposits := make([]deposit, 0)

	connectorAddress := common.HexToAddress(cfg.ConnectorAddress)
	connectorContract, err := zetaconnector.NewZetaConnectorNonEth(connectorAddress, client)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	erc20CustodyAddress := common.HexToAddress(cfg.CustodyAddress)
	erc20CustodyContract, err := erc20custody.NewERC20Custody(erc20CustodyAddress, client)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	custodyIter, err := erc20CustodyContract.FilterDeposited(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: context.TODO(),
	}, []common.Address{})
	if err != nil {
		fmt.Println("error loading filter: ", err.Error())
		return deposits
	}

	connectorIter, err := connectorContract.FilterZetaSent(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: context.TODO(),
	}, []common.Address{}, []*big.Int{})
	if err != nil {
		fmt.Println("error loading filter: ", err.Error())
		return deposits
	}

	// ********************** Get ERC20 Custody deposit events
	for custodyIter.Next() {
		// sanity check tx event
		err := CheckEvmTxLog(&custodyIter.Event.Raw, erc20CustodyAddress, "", TopicsDeposited)
		if err == nil {
			//fmt.Println("adding deposits")
			deposits = append(deposits, deposit{
				txId:   custodyIter.Event.Raw.TxHash.Hex(),
				amount: float64(custodyIter.Event.Amount.Int64()),
			})
		}
	}

	// ********************** Get Connector ZetaSent events
	for connectorIter.Next() {
		// sanity check tx event
		err := CheckEvmTxLog(&connectorIter.Event.Raw, connectorAddress, "", TopicsZetaSent)
		if err == nil {
			//fmt.Println("adding deposits")
			deposits = append(deposits, deposit{
				txId:   connectorIter.Event.Raw.TxHash.Hex(),
				amount: float64(connectorIter.Event.ZetaValueAndGas.Int64()),
			})
		}
	}

	//********************** Get Transactions sent directly to TSS address
	tssDeposits, err := getTSSDeposits(cfg.TssAddressEVM, int(startBlock), int(endBlock))
	if err != nil {
		fmt.Printf("getTSSDeposits returned err: %s", err.Error())
	}
	deposits = append(deposits, tssDeposits...)

	return deposits
}

func getTSSDeposits(tssAddress string, startBlock int, endBlock int) ([]deposit, error) {
	client := etherscan.New(etherscan.Mainnet, "S3AVTNXDJQZQQUVXJM4XVIPBRYECGK88VX")
	deposits := make([]deposit, 0)

	txns, err := client.NormalTxByAddress(tssAddress, &startBlock, &endBlock, 0, 0, true)
	if err != nil {
		return deposits, err
	}

	fmt.Println("getTSSDeposits - Num of transactions: ", len(txns))

	for _, tx := range txns {
		if tx.To == tssAddress {
			if strings.Compare(tx.Input, DonationMessage) == 0 {
				continue // skip donation tx
			}
			if tx.TxReceiptStatus != "1" {
				continue
			}
			//fmt.Println("getTSSDeposits - adding deposit")
			deposits = append(deposits, deposit{
				txId:   tx.Hash,
				amount: float64(tx.Value.Int().Int64()),
			})
		}
	}

	return deposits, nil
}

func CheckEvmTxLog(vLog *ethtypes.Log, wantAddress common.Address, wantHash string, wantTopics int) error {
	if vLog.Removed {
		return fmt.Errorf("log is removed, chain reorg?")
	}
	if vLog.Address != wantAddress {
		return fmt.Errorf("log emitter address mismatch: want %s got %s", wantAddress.Hex(), vLog.Address.Hex())
	}
	if vLog.TxHash.Hex() == "" {
		return fmt.Errorf("log tx hash is empty: %d %s", vLog.BlockNumber, vLog.TxHash.Hex())
	}
	if wantHash != "" && vLog.TxHash.Hex() != wantHash {
		return fmt.Errorf("log tx hash mismatch: want %s got %s", wantHash, vLog.TxHash.Hex())
	}
	if len(vLog.Topics) != wantTopics {
		return fmt.Errorf("number of topics mismatch: want %d got %d", wantTopics, len(vLog.Topics))
	}
	return nil
}
