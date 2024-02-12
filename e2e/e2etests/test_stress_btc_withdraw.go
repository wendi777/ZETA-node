package e2etests

import (
	"fmt"
	"math/big"
	"time"

	"github.com/btcsuite/btcutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
	crosschaintypes "github.com/zeta-chain/zetacore/x/crosschain/types"
	"golang.org/x/sync/errgroup"
)

// TestStressBTCWithdraw tests the stressing withdraw of btc
func TestStressBTCWithdraw(r *runner.E2ERunner) {
	// number of withdraws to perform
	numWithdraws := 100

	r.Logger.Print("starting stress test of %d withdraws", numWithdraws)

	// create a wait group to wait for all the withdraws to complete
	var eg errgroup.Group

	// send the withdraws
	for i := 0; i < numWithdraws; i++ {
		i := i
		tx, err := r.BTCZRC20.Withdraw(
			r.ZevmAuth,
			[]byte(r.BTCDeployerAddress.EncodeAddress()),
			big.NewInt(0.01*btcutil.SatoshiPerBitcoin),
		)
		if err != nil {
			panic(err)
		}
		receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZevmClient, tx, r.Logger, r.ReceiptTimeout)
		if receipt.Status == 0 {
			//r.Logger.Info("index %d: withdraw evm tx failed", index)
			panic(fmt.Sprintf("index %d: withdraw btc tx %s failed", i, tx.Hash().Hex()))
		}
		r.Logger.Print("index %d: starting withdraw, tx hash: %s", i, tx.Hash().Hex())

		eg.Go(func() error {
			return MonitorBTCWithdraw(r, tx, i, time.Now())
		})
	}

	// wait for all the withdraws to complete
	if err := eg.Wait(); err != nil {
		panic(err)
	}

	r.Logger.Print("all withdraws completed")
}

// MonitorBTCWithdraw monitors the withdraw of BTC, returns once the withdraw is complete
func MonitorBTCWithdraw(r *runner.E2ERunner, tx *ethtypes.Transaction, index int, startTime time.Time) error {
	cctx := utils.WaitCctxMinedByInTxHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
	if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
		return fmt.Errorf(
			"index %d: withdraw cctx failed with status %s, message %s, cctx index %s",
			index,
			cctx.CctxStatus.Status,
			cctx.CctxStatus.StatusMessage,
			cctx.Index,
		)
	}
	timeToComplete := time.Now().Sub(startTime)
	r.Logger.Print("index %d: withdraw cctx success in %s", index, timeToComplete.String())

	return nil
}
