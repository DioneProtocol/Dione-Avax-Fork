// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"github.com/dioneprotocol/dionego/snow"
	"github.com/dioneprotocol/dionego/vms"
	"github.com/dioneprotocol/dionego/vms/platformvm/config"
)

var _ vms.Factory = (*Factory)(nil)

// Factory can create new instances of the Platform Chain
type Factory struct {
	config.Config
}

// New returns a new instance of the Platform Chain
func (f *Factory) New(*snow.Context) (interface{}, error) {
	return &VM{Factory: *f}, nil
}