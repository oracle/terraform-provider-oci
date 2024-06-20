// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "config_file_profile" {
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

provider "local" {
  version = ">=1.3.0" # Need this version of the local provider to support base64 encoded inputs
}

