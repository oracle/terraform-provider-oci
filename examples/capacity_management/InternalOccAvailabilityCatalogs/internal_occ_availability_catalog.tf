// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "internal_occ_availability_catalog_catalog_state" {
  default = "STAGED"
}

variable "internal_occ_availability_catalog_display_name" {
  default = "displayName"
}

variable "internal_occ_availability_catalog_id" {
  default = "id"
}

variable "internal_occ_availability_catalog_namespace" {
  default = "COMPUTE"
}

variable "occ_customer_group_id" {
  default = "customerGroupId"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_internal_occ_availability_catalogs" "test_internal_occ_availability_catalogs" {
  #Required
  compartment_id        = var.compartment_id
  occ_customer_group_id = var.occ_customer_group_id

  #Optional
  catalog_state = var.internal_occ_availability_catalog_catalog_state
  display_name  = var.internal_occ_availability_catalog_display_name
  id            = var.internal_occ_availability_catalog_id
  namespace     = var.internal_occ_availability_catalog_namespace
}
