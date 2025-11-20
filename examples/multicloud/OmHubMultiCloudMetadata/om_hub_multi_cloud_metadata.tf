// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "root_compartment_id" {}
variable "subscription_id" {}


provider "oci" {
  auth                  = var.auth
  config_file_profile   = var.config_file_profile
  region                = var.region
}

data "oci_multicloud_om_hub_multi_cloud_metadata" "test_om_hub_multi_cloud_metadata" {
  #Required
  compartment_id    = var.root_compartment_id
  subscription_id   = var.subscription_id
}

