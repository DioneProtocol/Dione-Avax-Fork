// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package verify

import (
	"context"
	"errors"
	"fmt"

	"github.com/dioneprotocol/dionego/ids"
	"github.com/dioneprotocol/dionego/snow"
)

var (
	errSameChainID         = errors.New("same chainID")
	errMismatchedSubnetIDs = errors.New("mismatched subnetIDs")
)

// SameSubnet verifies that the provided [ctx] was provided to a chain in the
// same subnet as [peerChainID], but not the same chain. If this verification
// fails, a non-nil error will be returned.
func SameSubnet(ctx context.Context, chainCtx *snow.Context, peerChainID ids.ID) error {
	if peerChainID == chainCtx.ChainID {
		return errSameChainID
	}

	subnetID, err := chainCtx.ValidatorState.GetSubnetID(ctx, peerChainID)
	if err != nil {
		return fmt.Errorf("failed to get subnet of %q: %w", peerChainID, err)
	}
	if chainCtx.SubnetID != subnetID {
		return fmt.Errorf("%w; expected %q got %q", errMismatchedSubnetIDs, chainCtx.SubnetID, subnetID)
	}
	return nil
}
