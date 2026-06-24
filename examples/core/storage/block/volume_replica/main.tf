// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "auth" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "config_file_profile" {
}

provider "oci" {
  #  version          = "6.9.0"
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region

}
