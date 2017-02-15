#!/bin/bash

# This is called from pre-push.bash to do some verification checks on
# the Go code. The script will exit non-zero if any of these tests fail.
# However if environment variable IGNORE_VET_WARNINGS is a non-zero
# length string, go vet warnings will not exit non-zero. Also, if
# IGNORE_TEST_ERRORS is non-empty then test failures will not exit
# non-zero either.

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
all_packages="$(find . -type f -name *.go ! -path '*/vendor/*' -exec dirname {} \; | sort -u)"

echo "go version $(go version | awk '{print $3}')"

echo "checking: go fmt ..."
BADFMT=$(find . -type f -name '*.go' ! -path '*/vendor/*' | xargs gofmt -l)
if [ -n "$BADFMT" ]; then
    BADFMT=`echo "$BADFMT" | sed "s/^/  /"`
    echo -e "go fmt is sad:\n\n$BADFMT"
    exit 1
fi

echo "checking: go vet ..."
for dir in ${all_packages}; do echo " ... ${dir}"
	go tool vet ${dir}/*.go || [ -n "$IGNORE_VET_WARNINGS" ]
done


echo "checking: go build ..."
# check this branch builds cleanly
go build ${all_packages}


echo "checking: go test ..."
go test ${all_packages} || [ -n "$IGNORE_TEST_ERRORS" ]

echo "pass"
