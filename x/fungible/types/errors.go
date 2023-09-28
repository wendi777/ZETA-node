package types

// DONTCOVER

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/core/vm"
)

// x/fungible module sentinel errors
var (
	ErrSample                 = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrABIPack                = sdkerrors.Register(ModuleName, 1101, "abi pack error")
	ErrABIGet                 = sdkerrors.Register(ModuleName, 1102, "abi get error")
	ErrUnexpectedEvent        = sdkerrors.Register(ModuleName, 1103, "unexpected event")
	ErrABIUnpack              = sdkerrors.Register(ModuleName, 1104, "abi unpack error")
	ErrBlanceQuery            = sdkerrors.Register(ModuleName, 1105, "balance query error")
	ErrBalanceInvariance      = sdkerrors.Register(ModuleName, 1106, "balance invariance error")
	ErrContractNotFound       = sdkerrors.Register(ModuleName, 1107, "contract not found")
	ErrChainNotFound          = sdkerrors.Register(ModuleName, 1108, "chain not found")
	ErrContractCall           = sdkerrors.Register(ModuleName, 1109, "contract call error")
	ErrSystemContractNotFound = sdkerrors.Register(ModuleName, 1110, "system contract not found")
	ErrInvalidAddress         = sdkerrors.Register(ModuleName, 1111, "invalid address")
	ErrStateVariableNotFound  = sdkerrors.Register(ModuleName, 1112, "state variable not found")
	ErrDeployContract         = sdkerrors.Register(ModuleName, 1113, "deploy contract error")
	ErrEmitEvent              = sdkerrors.Register(ModuleName, 1114, "emit event error")
	ErrInvalidDecimals        = sdkerrors.Register(ModuleName, 1115, "invalid decimals")
	ErrGasPriceNotFound       = sdkerrors.Register(ModuleName, 1116, "gas price not found")
	ErrUpdateNonce            = sdkerrors.Register(ModuleName, 1117, "update nonce error")
	ErrInvalidGasLimit        = sdkerrors.Register(ModuleName, 1118, "invalid gas limit")
	ErrSetBytecode            = sdkerrors.Register(ModuleName, 1119, "set bytecode error")
	ErrInvalidContract        = sdkerrors.Register(ModuleName, 1120, "invalid contract")
	ErrPausedZRC20            = sdkerrors.Register(ModuleName, 1121, "ZRC20 is paused")
	ErrForeignCoinNotFound    = sdkerrors.Register(ModuleName, 1122, "foreign coin not found")
)

// IsRevertError checks if an error is a evm revert error
func IsRevertError(err error) bool {
	return strings.Contains(err.Error(), vm.ErrExecutionReverted.Error())
}
