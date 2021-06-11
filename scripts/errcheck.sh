#!/usr/bin/env bash
# Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

echo "==> Checking for unchecked errors..."

if ! which errcheck > /dev/null; then
    echo "==> ignore errcheck if not installed"
    #go get -u github.com/kisielk/errcheck
    exit 0
fi

err_files=$(errcheck -ignoretests \
                     -ignore 'github.com/hashicorp/terraform-plugin-sdk/helper/schema:Set' \
                     -ignore 'bytes:.*' \
                     -ignore 'io:Close|Write' \
                     $(go list ./...| grep -v /vendor/))

if [[ -n ${err_files} ]]; then
    echo 'Unchecked errors found in the following places:'
    echo "${err_files}"
    echo "Please handle returned errors. You can check directly with \`make errcheck\`"
    exit 1
fi

exit 0
