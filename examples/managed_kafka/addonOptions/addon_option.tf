// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}
variable "compartment_id" {}

variable "addon_option_name" {
  default = "name"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  region           = var.region
}

data "oci_managed_kafka_addon_options" "test_addon_options" {

  #Optional
  compartment_id = var.compartment_id
  name           = var.addon_option_name
}

