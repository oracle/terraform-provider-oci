#!/usr/bin/env bash
# Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofmt -l)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

# goimports checking is time consuming, so provide an escape hatch
while getopts "s" opt; do
  case $opt in
    s)
      echo "Skipping goimports checking"
      exit 0
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 3
      ;;
  esac
done

# Check goimports
#echo "==> Checking that code complies with goimports requirements..."
#goimports_files=$(find . -name '*.go' | grep -v vendor | xargs goimports -l -local github.com/terraform-providers/terraform-provider-oci)
#if [[ -n ${goimports_files} ]]; then
#    echo 'goimports needs running on the following files:'
#    echo "${goimports_files}"
#    echo "You can use the command: \`make fmt\` to reformat code."
#    exit 2
#fi

exit 0
