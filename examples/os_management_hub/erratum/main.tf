// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


data "oci_os_management_hub_errata" "test_errata" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  name = ["ELBA-2024-12244"]
  name_contains = "ELBA-2024-12244"
}

data "oci_os_management_hub_errata" "test_errata_2" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  classification_type = ["BUGFIX"]
  name = ["ELBA-2024-12244"]
  name_contains = "ELBA-2024-12244"
  os_family = "ORACLE_LINUX_9"
  time_issue_date_end = "2024-04-01T00:00:00.000Z"
  time_issue_date_start = "2024-03-20T00:00:00.000Z"
}


data "oci_os_management_hub_erratum" "test_erratum" {
  compartment_id = "${var.compartment_id}"
  name = "ELBA-2024-12244"
}
