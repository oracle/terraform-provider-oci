// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "environment_display_name" {
  default = "displayName"
}

variable "environment_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_environment" "test_environment" {
  compartment_id = var.compartment_id
  display_name  = var.environment_display_name
}

data "oci_cloud_bridge_environments" "test_environments" {
  compartment_id = var.compartment_id
  display_name   = var.environment_display_name
  environment_id = oci_cloud_bridge_environment.test_environment.id
  state          = var.environment_state
}
