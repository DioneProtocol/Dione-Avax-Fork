#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides

DIONE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd ) # Directory above this script

# Where DioneGo binary goes
dionego_path="$DIONE_PATH/build/dionego"
plugin_dir=${PLUGIN_DIR:-$HOME/.dionego/plugins}
evm_path=${EVM_PATH:-$plugin_dir/evm}
coreth_version=${CORETH_VERSION:-'v0.11.7-rc.3'}

# Set the PATHS
GOPATH="$(go env GOPATH)"
coreth_path=${CORETH_PATH:-"$GOPATH/pkg/mod/github.com/dioneprotocol/coreth@$coreth_version"}

# Avalabs docker hub
# dioneprotocol/dionego - defaults to local as to avoid unintentional pushes
# You should probably set it - export DOCKER_REPO='dioneprotocol/dionego'
dionego_dockerhub_repo=${DOCKER_REPO:-"dionego"}

# Current branch
# TODO: fix "fatal: No names found, cannot describe anything" in github CI
current_branch=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match || true)

git_commit=${DIONEGO_COMMIT:-$( git rev-list -1 HEAD )}

# Static compilation
static_ld_flags=''
if [ "${STATIC_COMPILATION:-}" = 1 ]
then
    export CC=musl-gcc
    which $CC > /dev/null || ( echo $CC must be available for static compilation && exit 1 )
    static_ld_flags=' -extldflags "-static" -linkmode external '
fi

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
