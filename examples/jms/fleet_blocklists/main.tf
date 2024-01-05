// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fleet_id" {
  default = "example-fleet-id"
}
variable "managed_instance_id" {}
variable "fleet_blocklist_operation" {
  default = "DELETE_JAVA_INSTALLATION"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_fleet_blocklists" "test_fleet_blocklists" {
  #Required
  fleet_id = var.fleet_id

  #Optional
  managed_instance_id = var.managed_instance_id
  operation           = var.fleet_blocklist_operation
}

