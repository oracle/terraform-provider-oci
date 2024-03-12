// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_customer_group_display_name" {
  default = "displayName"
}

variable "occ_customer_group_id" {
  default = "id"
}

variable "occ_customer_group_status" {
  default = "ENABLED"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occ_customer_groups" "test_occ_customer_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.occ_customer_group_display_name
  id           = var.occ_customer_group_id
  status       = var.occ_customer_group_status
}