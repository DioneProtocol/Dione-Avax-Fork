// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/dioneprotocol/dionego/snow/engine/snowman/block"
	"github.com/dioneprotocol/dionego/version"
	"github.com/dioneprotocol/dionego/vms/rpcchainvm/grpcutils"
	"github.com/dioneprotocol/dionego/vms/rpcchainvm/gruntime"
	"github.com/dioneprotocol/dionego/vms/rpcchainvm/runtime"

	vmpb "github.com/dioneprotocol/dionego/proto/pb/vm"
	runtimepb "github.com/dioneprotocol/dionego/proto/pb/vm/runtime"
)

const defaultRuntimeDialTimeout = 5 * time.Second

// The address of the Runtime server is expected to be passed via ENV `runtime.EngineAddressKey`.
// This address is used by the Runtime client to send Initialize RPC to server.
//
// Serve starts the RPC Chain VM server and performs a handshake with the VM runtime service.
func Serve(ctx context.Context, vm block.ChainVM, opts ...grpcutils.ServerOption) error {
	shutdownHandler := make(chan os.Signal, 2)
	signal.Notify(shutdownHandler, os.Interrupt, syscall.SIGTERM)

	server := newVMServer(vm, opts...)

	go func(ctx context.Context) {
		select {
		case <-shutdownHandler:
			fmt.Println("runtime engine: received shutdown signal")
		case <-ctx.Done():
			fmt.Println("runtime engine: context has been cancelled")
		}
		server.GracefulStop()
		fmt.Println("vm server: graceful termination success")
	}(ctx)

	// address of Runtime server from ENV
	runtimeAddr := os.Getenv(runtime.EngineAddressKey)
	if runtimeAddr == "" {
		return fmt.Errorf("required env var missing: %q", runtime.EngineAddressKey)
	}

	clientConn, err := grpcutils.Dial(runtimeAddr)
	if err != nil {
		return fmt.Errorf("failed to create client conn: %w", err)
	}

	client := gruntime.NewClient(runtimepb.NewRuntimeClient(clientConn))

	listener, err := grpcutils.NewListener()
	if err != nil {
		return fmt.Errorf("failed to create new listener: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, defaultRuntimeDialTimeout)
	defer cancel()
	err = client.Initialize(ctx, version.RPCChainVMProtocol, listener.Addr().String())
	if err != nil {
		_ = listener.Close()
		return fmt.Errorf("failed to initialize vm runtime: %w", err)
	}

	// start RPC Chain VM server
	grpcutils.Serve(listener, server)

	return nil
}

// Returns an RPC Chain VM server serving health and VM services.
func newVMServer(vm block.ChainVM, opts ...grpcutils.ServerOption) *grpc.Server {
	server := grpcutils.NewServer(opts...)
	vmpb.RegisterVMServer(server, NewServer(vm))

	health := health.NewServer()
	health.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, health)

	return server
}
