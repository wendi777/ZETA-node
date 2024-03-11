package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

func TestNewMsgMigrateTssFunds_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgMigrateTssFunds
		error bool
	}{
		{
			name: "invalid creator",
			msg: types.NewMsgMigrateTssFunds(
				"invalid address",
				common.DefaultChainsList()[0].ChainId,
				sdkmath.NewUintFromString("100000"),
			),
			error: true,
		},
		{
			name: "invalid chain id",
			msg: types.NewMsgMigrateTssFunds(
				"zeta15ruj2tc76pnj9xtw64utktee7cc7w6vzaes73z",
				999,
				sdkmath.NewUintFromString("100000"),
			),
			error: true,
		},
		{
			name: "invalid amount",
			msg: types.NewMsgMigrateTssFunds(
				"zeta15ruj2tc76pnj9xtw64utktee7cc7w6vzaes73z",
				common.DefaultChainsList()[0].ChainId,
				sdkmath.NewUintFromString("0"),
			),
			error: true,
		},
		{
			name: "valid msg",
			msg: types.NewMsgMigrateTssFunds(
				"zeta15ruj2tc76pnj9xtw64utktee7cc7w6vzaes73z",
				common.DefaultChainsList()[0].ChainId,
				sdkmath.NewUintFromString("100000"),
			),
			error: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keeper.SetConfig(false)
			err := tt.msg.ValidateBasic()
			if tt.error {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewMsgMigrateTssFunds_GetSigners(t *testing.T) {
	signer := sample.AccAddress()
	tests := []struct {
		name   string
		msg    types.MsgMigrateTssFunds
		panics bool
	}{
		{
			name: "valid signer",
			msg: types.MsgMigrateTssFunds{
				Creator: signer,
				ChainId: common.DefaultChainsList()[0].ChainId,
				Amount:  sdkmath.NewUintFromString("100000"),
			},
			panics: false,
		},
		{
			name: "invalid signer",
			msg: types.MsgMigrateTssFunds{
				Creator: "invalid_address",
				ChainId: common.DefaultChainsList()[0].ChainId,
				Amount:  sdkmath.NewUintFromString("100000"),
			},
			panics: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panics {
				signers := tt.msg.GetSigners()
				assert.Equal(t, []sdk.AccAddress{sdk.MustAccAddressFromBech32(signer)}, signers)
			} else {
				assert.Panics(t, func() {
					tt.msg.GetSigners()
				})
			}
		})
	}
}

func TestNewMsgMigrateTssFunds_Type(t *testing.T) {
	msg := types.MsgMigrateTssFunds{
		Creator: sample.AccAddress(),
		ChainId: common.DefaultChainsList()[0].ChainId,
		Amount:  sdkmath.NewUintFromString("100000"),
	}
	assert.Equal(t, types.TypeMsgMigrateTssFunds, msg.Type())
}

func TestNewMsgMigrateTssFunds_Route(t *testing.T) {
	msg := types.MsgMigrateTssFunds{
		Creator: sample.AccAddress(),
		ChainId: common.DefaultChainsList()[0].ChainId,
		Amount:  sdkmath.NewUintFromString("100000"),
	}
	assert.Equal(t, types.RouterKey, msg.Route())
}

func TestNewMsgMigrateTssFunds_GetSignBytes(t *testing.T) {
	msg := types.MsgMigrateTssFunds{
		Creator: sample.AccAddress(),
		ChainId: common.DefaultChainsList()[0].ChainId,
		Amount:  sdkmath.NewUintFromString("100000"),
	}
	assert.NotPanics(t, func() {
		msg.GetSignBytes()
	})
}
