// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "root_compartment_id" {}


provider "oci" {
  auth                  = var.auth
  config_file_profile   = var.config_file_profile
  region                = var.region
}

data "oci_multicloud_multicloudsubscriptions" "test_multicloudsubscriptions" {
  #Required
  compartment_id           = var.root_compartment_id
}
