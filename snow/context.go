// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snow

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/dioneprotocol/dionego/api/keystore"
	"github.com/dioneprotocol/dionego/api/metrics"
	"github.com/dioneprotocol/dionego/chains/atomic"
	"github.com/dioneprotocol/dionego/ids"
	"github.com/dioneprotocol/dionego/snow/validators"
	"github.com/dioneprotocol/dionego/utils"
	"github.com/dioneprotocol/dionego/utils/logging"
	"github.com/dioneprotocol/dionego/vms/platformvm/warp"
)

// ContextInitializable represents an object that can be initialized
// given a *Context object
type ContextInitializable interface {
	// InitCtx initializes an object provided a *Context object
	InitCtx(ctx *Context)
}

// Context is information about the current execution.
// [NetworkID] is the ID of the network this context exists within.
// [ChainID] is the ID of the chain this context exists within.
// [NodeID] is the ID of this node
type Context struct {
	NetworkID uint32
	SubnetID  ids.ID
	ChainID   ids.ID
	NodeID    ids.NodeID

	XChainID    ids.ID
	CChainID    ids.ID
	DIONEAssetID ids.ID

	Log          logging.Logger
	Lock         sync.RWMutex
	Keystore     keystore.BlockchainKeystore
	SharedMemory atomic.SharedMemory
	BCLookup     ids.AliaserReader
	Metrics      metrics.OptionalGatherer

	WarpSigner warp.Signer

	// snowman++ attributes
	ValidatorState validators.State // interface for P-Chain validators
	// Chain-specific directory where arbitrary data can be written
	ChainDataDir string
}

// Expose gatherer interface for unit testing.
type Registerer interface {
	prometheus.Registerer
	prometheus.Gatherer
}

type ConsensusContext struct {
	*Context

	Registerer Registerer

	// DecisionAcceptor is the callback that will be fired whenever a VM is
	// notified that their object, either a block in snowman or a transaction
	// in dione, was accepted.
	DecisionAcceptor Acceptor

	// ConsensusAcceptor is the callback that will be fired whenever a
	// container, either a block in snowman or a vertex in dione, was
	// accepted.
	ConsensusAcceptor Acceptor

	// State indicates the current state of this consensus instance.
	State utils.Atomic[EngineState]

	// True iff this chain is executing transactions as part of bootstrapping.
	Executing utils.Atomic[bool]

	// True iff this chain is currently state-syncing
	StateSyncing utils.Atomic[bool]
}

func DefaultContextTest() *Context {
	return &Context{
		NetworkID:    0,
		SubnetID:     ids.Empty,
		ChainID:      ids.Empty,
		NodeID:       ids.EmptyNodeID,
		Log:          logging.NoLog{},
		BCLookup:     ids.NewAliaser(),
		Metrics:      metrics.NewOptionalGatherer(),
		ChainDataDir: "",
	}
}

func DefaultConsensusContextTest() *ConsensusContext {
	return &ConsensusContext{
		Context:           DefaultContextTest(),
		Registerer:        prometheus.NewRegistry(),
		DecisionAcceptor:  noOpAcceptor{},
		ConsensusAcceptor: noOpAcceptor{},
	}
}
