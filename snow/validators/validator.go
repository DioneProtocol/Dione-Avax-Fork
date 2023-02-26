// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"github.com/dioneprotocol/dionego/ids"
	"github.com/dioneprotocol/dionego/utils/crypto/bls"
)

// Validator is a struct that contains the base values representing a validator
// of the Dione Network.
type Validator struct {
	NodeID    ids.NodeID
	PublicKey *bls.PublicKey
	TxID      ids.ID
	Weight    uint64

	// index is used to efficiently remove validators from the validator set. It
	// represents the index of this validator in the vdrSlice and weights
	// arrays.
	index int
}

// GetValidatorOutput is a struct that contains the publicly relevant values of
// a validator of the Dione Network for the output of GetValidator.
type GetValidatorOutput struct {
	NodeID    ids.NodeID
	PublicKey *bls.PublicKey
	Weight    uint64
}
