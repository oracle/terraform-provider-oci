#!/bin/bash

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

#
# Find package dirs excluding vendor
#
all_source_files="$(find . -type f -name '*.go' ! -path '*/vendor/*')"


echo "go version $(go version | awk '{print $3}')"


echo "GoImports"

# Perform import group checking and ordering
goimports -w ${all_source_files}

echo "Simplify"

# Perform simplification
gofmt -s -w ${all_source_files}

echo "Done"
