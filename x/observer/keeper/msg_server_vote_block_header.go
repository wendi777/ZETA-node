package keeper

import (
	"context"

	cosmoserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

// VoteBlockHeader vote for a new block header to the storers
func (k msgServer) VoteBlockHeader(goCtx context.Context, msg *types.MsgVoteBlockHeader) (*types.MsgVoteBlockHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the chain is enabled
	chain := k.GetSupportedChainFromChainID(ctx, msg.ChainId)
	if chain == nil {
		return nil, cosmoserrors.Wrapf(types.ErrSupportedChains, "chain id: %d", msg.ChainId)
	}

	// check if observer
	if ok := k.IsNonTombstonedObserver(ctx, msg.Creator); !ok {
		return nil, types.ErrNotObserver
	}

	// check the new block header is valid
	parentHash, err := k.lightclientKeeper.CheckNewBlockHeader(ctx, msg.ChainId, msg.BlockHash, msg.Height, msg.Header)
	if err != nil {
		return nil, cosmoserrors.Wrap(types.ErrInvalidBlockHeader, err.Error())
	}

	// add vote to ballot
	ballot, _, err := k.FindBallot(ctx, msg.Digest(), chain, types.ObservationType_InBoundTx)
	if err != nil {
		return nil, cosmoserrors.Wrap(err, "failed to find ballot")
	}
	ballot, err = k.AddVoteToBallot(ctx, ballot, msg.Creator, types.VoteType_SuccessObservation)
	if err != nil {
		return nil, cosmoserrors.Wrap(err, "failed to add vote to ballot")
	}
	_, isFinalized := k.CheckIfFinalizingVote(ctx, ballot)
	if !isFinalized {
		return &types.MsgVoteBlockHeaderResponse{}, nil
	}

	// add the new block header to the store
	k.lightclientKeeper.AddBlockHeader(ctx, msg.ChainId, msg.Height, msg.BlockHash, msg.Header, parentHash)

	return &types.MsgVoteBlockHeaderResponse{}, nil
}
