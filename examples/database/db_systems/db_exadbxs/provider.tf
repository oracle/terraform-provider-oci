// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}
