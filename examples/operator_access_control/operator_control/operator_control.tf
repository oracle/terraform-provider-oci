// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

data "oci_identity_groups" "get_admin_approver_group" {
    #Required
    compartment_id = var.tenancy_ocid

    #Optional
    name = "Administrators"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_operator_access_control_operator_control" "test_operator_control" {
  #Required
  compartment_id        = var.compartment_ocid
  operator_control_name = "Tersi-Example"
  approver_groups_list = [data.oci_identity_groups.get_admin_approver_group.groups[0].id]
  is_fully_pre_approved = true
  resource_type = "EXADATAINFRASTRUCTURE"
}

data "oci_operator_access_control_operator_controls" "test_operator_controls" {
  compartment_id = var.compartment_ocid
}
