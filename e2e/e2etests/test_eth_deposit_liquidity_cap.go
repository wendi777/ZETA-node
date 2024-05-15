package e2etests

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/zeta-chain/zetacore/e2e/runner"
	"github.com/zeta-chain/zetacore/e2e/utils"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	fungibletypes "github.com/zeta-chain/zetacore/x/fungible/types"
)

// TestDepositEtherLiquidityCap tests depositing Ethers in a context where a liquidity cap is set
func TestDepositEtherLiquidityCap(r *runner.E2ERunner, args []string) {
	if len(args) != 1 {
		panic("TestDepositEtherLiquidityCap requires exactly one argument for the liquidity cap.")
	}

	liquidityCapArg := math.NewUintFromString(args[0])
	supply, err := r.ETHZRC20.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	liquidityCap := math.NewUintFromBigInt(supply).Add(liquidityCapArg)
	amountLessThanCap := liquidityCapArg.BigInt().Div(liquidityCapArg.BigInt(), big.NewInt(10)) // 1/10 of the cap
	amountMoreThanCap := liquidityCapArg.BigInt().Mul(liquidityCapArg.BigInt(), big.NewInt(10)) // 10 times the cap
	msg := fungibletypes.NewMsgUpdateZRC20LiquidityCap(
		r.ZetaTxServer.GetAccountAddress(0),
		r.ETHZRC20Addr.Hex(),
		liquidityCap,
	)
	res, err := r.ZetaTxServer.BroadcastTx(utils.FungibleAdminName, msg)
	if err != nil {
		panic(err)
	}
	r.Logger.Info("set liquidity cap tx hash: %s", res.TxHash)

	r.Logger.Info("Depositing more than liquidity cap should make cctx reverted")
	signedTx, err := r.SendEther(r.TSSAddress, amountMoreThanCap, nil)
	if err != nil {
		panic(err)
	}
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.EVMClient, signedTx, r.Logger, r.ReceiptTimeout)
	if receipt.Status == 0 {
		panic("deposit eth tx failed")
	}
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, signedTx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	if cctx.CctxStatus.Status != types.CctxStatus_Reverted {
		panic(fmt.Sprintf("expected cctx status to be Reverted; got %s", cctx.CctxStatus.Status))
	}
	r.Logger.Info("CCTX has been reverted")

	r.Logger.Info("Depositing less than liquidity cap should still succeed")
	initialBal, err := r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
	if err != nil {
		panic(err)
	}
	signedTx, err = r.SendEther(r.TSSAddress, amountLessThanCap, nil)
	if err != nil {
		panic(err)
	}
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.EVMClient, signedTx, r.Logger, r.ReceiptTimeout)
	if receipt.Status == 0 {
		panic("deposit eth tx failed")
	}
	cctx = utils.WaitCctxMinedByInboundHash(r.Ctx, signedTx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	if cctx.CctxStatus.Status != types.CctxStatus_OutboundMined {

		panic(fmt.Sprintf(
			"expected cctx status to be Success; got %s; message: %s; supply: %s; liquidity cap: %s, amountLessThanCap: %s, amountMoreThanCap: %s",
			cctx.CctxStatus.Status,
			cctx.CctxStatus.StatusMessage,
			supply.String(),
			liquidityCap.String(),
			amountLessThanCap.String(),
			amountMoreThanCap.String(),
		))
	}

	expectedBalance := big.NewInt(0).Add(initialBal, amountLessThanCap)

	bal, err := r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
	if err != nil {
		panic(err)
	}
	if bal.Cmp(expectedBalance) != 0 {

		panic(fmt.Sprintf("expected balance to be %s; got %s", expectedBalance.String(), bal.String()))
	}
	r.Logger.Info("Deposit succeeded")

	r.Logger.Info("Removing the liquidity cap")
	msg = fungibletypes.NewMsgUpdateZRC20LiquidityCap(
		r.ZetaTxServer.GetAccountAddress(0),
		r.ETHZRC20Addr.Hex(),
		math.ZeroUint(),
	)
	res, err = r.ZetaTxServer.BroadcastTx(utils.FungibleAdminName, msg)
	if err != nil {
		panic(err)
	}
	r.Logger.Info("remove liquidity cap tx hash: %s", res.TxHash)
	initialBal, err = r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
	if err != nil {
		panic(err)
	}
	signedTx, err = r.SendEther(r.TSSAddress, amountMoreThanCap, nil)
	if err != nil {
		panic(err)
	}
	receipt = utils.MustWaitForTxReceipt(r.Ctx, r.EVMClient, signedTx, r.Logger, r.ReceiptTimeout)
	if receipt.Status == 0 {
		panic("deposit eth tx failed")
	}
	utils.WaitCctxMinedByInboundHash(r.Ctx, signedTx.Hash().Hex(), r.CctxClient, r.Logger, r.CctxTimeout)
	expectedBalance = big.NewInt(0).Add(initialBal, amountMoreThanCap)

	bal, err = r.ETHZRC20.BalanceOf(&bind.CallOpts{}, r.DeployerAddress)
	if err != nil {
		panic(err)
	}
	if bal.Cmp(expectedBalance) != 0 {
		panic(fmt.Sprintf("expected balance to be %s; got %s", expectedBalance.String(), bal.String()))
	}
	r.Logger.Info("New deposit succeeded")
}
