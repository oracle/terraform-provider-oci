// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "internal_namespace_occ_overview_from" {
  default = "2023-08-05T17:17:14.816Z"
}

variable "internal_namespace_occ_overview_namespace" {
  default = "COMPUTE"
}

variable "internal_namespace_occ_overview_to" {
  default = "2023-10-05T17:17:14.816Z"
}

variable "internal_namespace_occ_overview_workload_type" {
  default = "workloadType"
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

data "oci_capacity_management_internal_namespace_occ_overviews" "test_internal_namespace_occ_overviews" {
  #Required
  compartment_id        = var.compartment_id
  namespace             = var.internal_namespace_occ_overview_namespace
  occ_customer_group_id = var.occ_customer_group_id

  #Optional
  from          = var.internal_namespace_occ_overview_from
  to            = var.internal_namespace_occ_overview_to
  workload_type = var.internal_namespace_occ_overview_workload_type
}
