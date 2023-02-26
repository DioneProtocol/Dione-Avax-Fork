// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/dioneprotocol/dionego/snow/validators"
	"github.com/dioneprotocol/dionego/utils/crypto/bls"
	"github.com/dioneprotocol/dionego/utils/set"
)

var (
	_ Signature = (*BitSetSignature)(nil)

	ErrInvalidBitSet      = errors.New("bitset is invalid")
	ErrInsufficientWeight = errors.New("signature weight is insufficient")
	ErrInvalidSignature   = errors.New("signature is invalid")
	ErrParseSignature     = errors.New("failed to parse signature")
)

type Signature interface {
	// Verify that this signature was signed by at least [quorumNum]/[quorumDen]
	// of the validators of [msg.SourceChainID] at [pChainHeight].
	//
	// Invariant: [msg] is correctly initialized.
	Verify(
		ctx context.Context,
		msg *UnsignedMessage,
		pChainState validators.State,
		pChainHeight uint64,
		quorumNum uint64,
		quorumDen uint64,
	) error
}

type BitSetSignature struct {
	// Signers is a big-endian byte slice encoding which validators signed this
	// message.
	Signers   []byte                 `serialize:"true"`
	Signature [bls.SignatureLen]byte `serialize:"true"`
}

func (s *BitSetSignature) Verify(
	ctx context.Context,
	msg *UnsignedMessage,
	pChainState validators.State,
	pChainHeight uint64,
	quorumNum uint64,
	quorumDen uint64,
) error {
	subnetID, err := pChainState.GetSubnetID(ctx, msg.SourceChainID)
	if err != nil {
		return err
	}

	vdrs, totalWeight, err := GetCanonicalValidatorSet(ctx, pChainState, pChainHeight, subnetID)
	if err != nil {
		return err
	}

	// Parse signer bit vector
	signerIndices := set.BitsFromBytes(s.Signers)
	if len(signerIndices.Bytes()) != len(s.Signers) {
		return ErrInvalidBitSet
	}

	// Get the validators that (allegedly) signed the message.
	signers, err := FilterValidators(signerIndices, vdrs)
	if err != nil {
		return err
	}

	// Because [signers] is a subset of [vdrs], this can never error.
	sigWeight, _ := SumWeight(signers)

	// Make sure the signature's weight is sufficient.
	err = VerifyWeight(
		sigWeight,
		totalWeight,
		quorumNum,
		quorumDen,
	)
	if err != nil {
		return err
	}

	// Parse the aggregate signature
	aggSig, err := bls.SignatureFromBytes(s.Signature[:])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrParseSignature, err)
	}

	// Create the aggregate public key
	aggPubKey, err := AggregatePublicKeys(signers)
	if err != nil {
		return err
	}

	// Verify the signature
	unsignedBytes := msg.Bytes()
	if !bls.Verify(aggPubKey, aggSig, unsignedBytes) {
		return ErrInvalidSignature
	}
	return nil
}

// VerifyWeight returns [nil] if [sigWeight] is at least [quorumNum]/[quorumDen]
// of [totalWeight].
// If [sigWeight >= totalWeight * quorumNum / quorumDen] then return [nil]
func VerifyWeight(
	sigWeight uint64,
	totalWeight uint64,
	quorumNum uint64,
	quorumDen uint64,
) error {
	// Verifies that quorumNum * totalWeight <= quorumDen * sigWeight
	scaledTotalWeight := new(big.Int).SetUint64(totalWeight)
	scaledTotalWeight.Mul(scaledTotalWeight, new(big.Int).SetUint64(quorumNum))
	scaledSigWeight := new(big.Int).SetUint64(sigWeight)
	scaledSigWeight.Mul(scaledSigWeight, new(big.Int).SetUint64(quorumDen))
	if scaledTotalWeight.Cmp(scaledSigWeight) == 1 {
		return fmt.Errorf(
			"%w: %d*%d > %d*%d",
			ErrInsufficientWeight,
			quorumNum,
			totalWeight,
			quorumDen,
			sigWeight,
		)
	}
	return nil
}
