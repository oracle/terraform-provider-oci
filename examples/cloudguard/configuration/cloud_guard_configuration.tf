// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

//The reporting region needs to be a valid reporting region where the tenancy is subscribed to.
//In most cases the home-region of the tenancy is its reporting region.
//In a single region tenancy, the home region, reporting region and the monitoring region are all same.
variable "cloud_guard_configuration_reporting_region" {
  default = "us-phoenix-1"
}

//The acceptable values for status are `ENABLED` and `DISABLED`.
//DISABLING the tenancy is equivalent to off-boarding resulting in deletion of all the Control Plane entities, also disallowing most of the CloudGuard Operations.
//Once ENABLED, the reporting region can't be switched unless it is DISABLED and then ENABLED again for another region.
//However, The reporting region needs to be a valid reporting region where the tenancy is subscribed to.
variable "cloud_guard_configuration_status" {
  default = "ENABLED"
}

//Setting this variable to true lets the user seed the oracle managed entities with minimal changes to the original entities.
//False will delegate this responsibility to CloudGuard for seeding the oracle managed entities.
variable "cloud_guard_configuration_self_manage_resources" {
  default = false
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

//CloudGuard enabling and disabling is a tenant-level operation so the compartment-id needs to be a tenant-ocid.
resource "oci_cloud_guard_cloud_guard_configuration" "test_cloud_guard_configuration" {
  #Required
  compartment_id   = "${var.tenancy_ocid}"
  reporting_region = "${var.cloud_guard_configuration_reporting_region}"
  status           = "${var.cloud_guard_configuration_status}"

  #Optional
  self_manage_resources = "${var.cloud_guard_configuration_self_manage_resources}"
}

//You can inspect the details of a tenant (whether CloudGuard is enabled/disabled) through any of its child compartments.
data "oci_cloud_guard_cloud_guard_configuration" "test_cloud_guard_configuration" {
  #Required
  compartment_id = "${var.compartment_id}"
}