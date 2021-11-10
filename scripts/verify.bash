#!/usr/bin/env bash
# Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
local_go_version="$(go version | { read _ _ v _; echo ${v#go}; })"
provider_go_version=$(<.go-version)

if [[ "$local_go_version" == "$provider_go_version" ]]; then
 echo "Go version match!"
else
    echo "GO version mismatch! You are currently on GO version ${local_go_version} and the expected GO version is ${provider_go_version}"
    read -p "Do you wish to continue with the Push before synchronizing the GO version (y/n)?" input < /dev/tty

    if [[ "$input" == "y" ]]; then
        :
    else
        echo "Aborting Push"
        exit 1
    fi
fi

check_make_target 2 'vet'
check_make_target 3 'errcheck'
check_make_target 5 'test-compile' 'TEST=./internal/integrationtest'
check_make_target 6 'ocicheck'
#check_make_target 7 'website-test'
check_make_target 7 'test-docscheck'


echo "checking: make build ..."
# check this branch builds cleanly (which internally also takes care of the formatting check)
make build

echo "pass"
