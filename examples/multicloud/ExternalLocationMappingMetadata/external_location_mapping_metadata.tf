// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "compartment_id" {}
variable "subscription_service_name_list" {
    type = list(string)
}


provider "oci" {
  auth                  = var.auth
  config_file_profile   = var.config_file_profile
  region                = var.region
}

data "oci_multicloud_external_location_mapping_metadata" "test_external_location_mapping_metadata" {
  #Required
  compartment_id            = var.compartment_id
  subscription_service_name = var.subscription_service_name_list
}
