// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable config_file_profile {
}

variable "region" {
}

variable "compartment_ocid" {

}

provider "oci" {
  auth = "SecurityToken"
  config_file_profile = var.config_file_profile
  region = var.region
}


