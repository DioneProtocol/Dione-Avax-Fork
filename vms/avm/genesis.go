// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/dioneprotocol/dionego/utils"
	"github.com/dioneprotocol/dionego/vms/avm/txs"
)

var _ utils.Sortable[*GenesisAsset] = (*GenesisAsset)(nil)

type Genesis struct {
	Txs []*GenesisAsset `serialize:"true"`
}

type GenesisAsset struct {
	Alias             string `serialize:"true"`
	txs.CreateAssetTx `serialize:"true"`
}

func (g *GenesisAsset) Less(other *GenesisAsset) bool {
	return g.Alias < other.Alias
}
