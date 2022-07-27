package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/zetacore/types"
)

func (k msgServer) NonceVoter(goCtx context.Context, msg *types.MsgNonceVoter) (*types.MsgNonceVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validators := k.StakingKeeper.GetAllValidators(ctx)
	if !IsBondedValidator(msg.Creator, validators) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, fmt.Sprintf("signer %s is not a bonded validator", msg.Creator))
	}

	chain := msg.Chain
	chainNonce, isFound := k.GetChainNonces(ctx, chain)
	//if isDuplicateSigner(msg.Creator, chainNonce.Signers) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, fmt.Sprintf("signer %s double signing!!", msg.Creator))
	//}
	if isFound {
		chainNonce.Signers = append(chainNonce.Signers, msg.Creator)
		chainNonce.Nonce = msg.Nonce
	} else if !isFound {
		chainNonce = types.ChainNonces{
			Creator: msg.Creator,
			Index:   msg.Chain,
			Chain:   msg.Chain,
			Nonce:   msg.Nonce,
			Signers: []string{msg.Creator},
		}
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("chainNonce vote msg does not match state: %v vs %v", msg, chainNonce))
	}

	if hasSuperMajorityValidators(len(chainNonce.Signers), validators) {
		chainNonce.FinalizedHeight = uint64(ctx.BlockHeader().Height)
	}

	k.SetChainNonces(ctx, chainNonce)
	return &types.MsgNonceVoterResponse{}, nil
}

func isDuplicateSigner(creator string, signers []string) bool {
	for _, v := range signers {
		if creator == v {
			return true
		}
	}
	return false
}

func hasSuperMajorityValidators(numSigners int, validators []stakingtypes.Validator) bool {
	numValidValidators := 0
	for _, v := range validators {
		if v.IsBonded() {
			numValidValidators += 1
		}
	}
	threshold := numValidValidators * 2 / 3
	if threshold < 2 {
		threshold = 2
	}
	if numValidValidators == 1 {
		threshold = 1
	}
	return numSigners == threshold
}
