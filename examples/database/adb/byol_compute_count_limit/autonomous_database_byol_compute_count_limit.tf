// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  admin_password           = "BEstrO0ng_#11"
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "Xsk523nccsdf23"
  db_version               = "19c"
  db_workload              = "OLTP"
  license_model            = "BRING_YOUR_OWN_LICENSE"
  database_edition         = "STANDARD_EDITION"
  byol_compute_count_limit  = "10"
}
