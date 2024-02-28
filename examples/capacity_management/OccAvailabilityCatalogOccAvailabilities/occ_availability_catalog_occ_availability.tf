// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "occ_availability_catalog_occ_availability_date_expected_capacity_handover" {
  default = "dateExpectedCapacityHandover"
}

variable "occ_availability_catalog_occ_availability_resource_type" {
  default = "SERVER_HW"
}

variable "occ_availability_catalog_occ_availability_workload_type" {
  default = "GENERIC"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occ_availability_catalog_occ_availabilities" "test_occ_availability_catalog_occ_availabilities" {
  #Required
  occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id

  #Optional
  date_expected_capacity_handover = var.occ_availability_catalog_occ_availability_date_expected_capacity_handover
  resource_name                   = oci_usage_proxy_resource.test_resource.name
  resource_type                   = var.occ_availability_catalog_occ_availability_resource_type
  workload_type                   = var.occ_availability_catalog_occ_availability_workload_type
}