// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "namespace_occ_overview_from" {
  default = "2023-08-05T17:17:14.816Z"
}

variable "namespace_occ_overview_namespace" {
  default = "COMPUTE"
}

variable "namespace_occ_overview_to" {
  default = "2024-08-05T17:17:14.816Z"
}

variable "namespace_occ_overview_workload_type" {
  default = "workloadType"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_namespace_occ_overviews" "test_namespace_occ_overviews" {
  #Required
  compartment_id = var.compartment_id
  namespace      = var.namespace_occ_overview_namespace

  #Optional
  from          = var.namespace_occ_overview_from
  to            = var.namespace_occ_overview_to
  workload_type = var.namespace_occ_overview_workload_type
}
