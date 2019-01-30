# Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

#!/usr/bin/env bash

TARGET_HOOK="$(git rev-parse --show-toplevel)/.git/hooks/pre-push"

if [ ! -L "${TARGET_HOOK}" ]; then
  echo "${TARGET_HOOK}" missing
  pushd "$(git rev-parse --show-toplevel)"
  ln -s ../../scripts/pre-push.bash .git/hooks/pre-push
  popd
else
  echo "hook already installed"
fi
