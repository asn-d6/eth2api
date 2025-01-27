package beaconapi

import (
	"context"
	"github.com/protolambda/eth2api"
	"github.com/protolambda/zrnt/eth2/beacon/common"
)

// Retrieves block header for given block id.
func BlockHeader(ctx context.Context, cli eth2api.Client, blockId eth2api.BlockId, dest *eth2api.BeaconBlockHeaderAndInfo) (exists bool, err error) {
	return eth2api.SimpleRequest(ctx, cli, eth2api.FmtGET("/eth/v1/beacon/headers/%s", blockId.BlockId()), eth2api.Wrap(dest))
}

// Retrieves block headers matching given query. By default it will fetch current head slot blocks.
func BlockHeaders(ctx context.Context, cli eth2api.Client, slot *common.Slot, parentRoot *common.Root, dest *[]eth2api.BeaconBlockHeaderAndInfo) (exists bool, err error) {
	var q eth2api.Query
	if slot != nil {
		if parentRoot != nil {
			q = eth2api.Query{"slot": *slot, "parent_root": *parentRoot}
		} else {
			q = eth2api.Query{"slot": *slot}
		}
	} else if parentRoot != nil {
		q = eth2api.Query{"parent_root": *parentRoot}
	}
	return eth2api.SimpleRequest(ctx, cli, eth2api.QueryGET(q, "/eth/v1/beacon/headers"), eth2api.Wrap(dest))
}
