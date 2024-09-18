// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "occ_customer_group_occ_customer_description" {
  default = "Customer Tenancy"
}

variable "occ_customer_group_occ_customer_display_name" {
  default = "tenancy"
}

variable "occ_customer_group_occ_customer_status" {
  default = "ENABLED"
}

variable "occ_customer_group_id" {
  default = "customerGroupId"
}

variable "occ_customer_group-occ_customer_tenancy_id" {
  default = "customerId"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occ_customer_group_occ_customer" "test_occ_customer_group_occ_customer" {
  #Required
  display_name          = var.occ_customer_group_occ_customer_display_name
  occ_customer_group_id = var.occ_customer_group_id
  tenancy_id            = var.occ_customer_group-occ_customer_tenancy_id

  #Optional
  description = var.occ_customer_group_occ_customer_description
  status      = var.occ_customer_group_occ_customer_status
}

