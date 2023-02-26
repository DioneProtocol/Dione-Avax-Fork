# Dione gRPC

Now Serving: **Protocol Version 23**

Protobuf files are hosted at [https://buf.build/dioneprotocol/dionetps://buf.build/dioneprotocol/dione) and can be used as dependencies in other projects.

Protobuf linting and generation for this project is managed by [buf](https://github.com/bufbuild/buf).

Please find installation instructions on [https://docs.buf.build/installation/](https://docs.buf.build/installation/) or use `Dockerfile.buf` provided in the `proto/` directory of DioneGo.

Any changes made to proto definition can be updated by running `protobuf_codegen.sh` located in the `scripts/` directory of DioneGo.

Introduction to `buf` [https://docs.buf.build/tour/introduction](https://docs.buf.build/tour/introduction)

## Protocol Version Compatibility

The protobuf definitions and generated code are versioned based on the [RPCChainVMProtocol](../version/version.go#L13) defined for the RPCChainVM.
Many versions of an Dione client can use the same [RPCChainVMProtocol](../version/version.go#L13). But each Dione client and subnet vm must use the same protocol version to be compatible.