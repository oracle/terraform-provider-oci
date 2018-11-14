#!/usr/bin/env bash

# This is called from pre-push.bash to do some verification checks on
# the Go code. The script will exit non-zero if any of these tests fail.

set -e

#
# Put ourselves at the known project root location, and go back where
# we were on exit.
#
PROJECT_ROOT="$(git rev-parse --show-toplevel)"
function _trap_exit() {
  popd &> /dev/null
}
trap _trap_exit EXIT
pushd "${PROJECT_ROOT}" &> /dev/null

function check_make_target() {
    local return_code=${1}
    local target=${2}
    local target_args=${3}

    echo "checking: make $target ..."
    if ! make ${target} ${target_args}; then
        echo -e "$target is sad"
        exit ${return_code}
    fi
}

echo "go version $(go version | awk '{print $3}')"

check_make_target 2 'vet'
check_make_target 3 'errcheck'
check_make_target 4 'vendor-status'
check_make_target 5 'test-compile' 'TEST=./oci'
check_make_target 6 'website-test'
check_make_target 7 'ocicheck'

echo "checking: make build ..."
# check this branch builds cleanly (which internally also takes care of the formatting check)
make build

echo "pass"
