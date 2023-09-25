package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/observer/keeper"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// MigrateStore migrates the x/observer module state from the consensus version 2 to 3
// This migration update the policy group
func MigrateStore(ctx sdk.Context, k keeper.Keeper) error {
	// Get first admin policy group
	p := k.GetParams(ctx)
	if len(p.AdminPolicy) == 0 || p.AdminPolicy[0] == nil {
		return nil
	}

	admin := p.AdminPolicy[0].Address
	p.AdminPolicy = []*types.Admin_Policy{
		{
			Address:    admin,
			PolicyType: types.Policy_Type_group1,
		},
		{
			Address:    admin,
			PolicyType: types.Policy_Type_group2,
		},
	}
	k.SetParams(ctx, p)

	return nil
}
