package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

type ObserverKeeper interface {
	GetCrosschainFlags(ctx sdk.Context) (val types.CrosschainFlags, found bool)
	SetCrosschainFlags(ctx sdk.Context, crosschainFlags types.CrosschainFlags)
}

func MigrateStore(ctx sdk.Context, observerStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	newCrossChainFlags := types.DefaultCrosschainFlags()
	var val types.LegacyCrosschainFlags
	store := prefix.NewStore(ctx.KVStore(observerStoreKey), types.KeyPrefix(types.CrosschainFlagsKey))
	b := store.Get([]byte{0})
	if b != nil {
		cdc.MustUnmarshal(b, &val)
		if val.GasPriceIncreaseFlags != nil {
			newCrossChainFlags.GasPriceIncreaseFlags = val.GasPriceIncreaseFlags
		}
		newCrossChainFlags.IsOutboundEnabled = val.IsOutboundEnabled
		newCrossChainFlags.IsInboundEnabled = val.IsInboundEnabled
	}
	b = cdc.MustMarshal(newCrossChainFlags)
	store.Set([]byte{0}, b)
	return nil
}
