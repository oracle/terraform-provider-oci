#!/usr/bin/env bash
# Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
.googlesource.com,TRUE,/,TRUE,2147483647,o,git-paul.hashicorp.com=1/z7s05EYPudQ9qoe6dMVfmAVwgZopEkZBb1a2mA5QtHE
__END__
