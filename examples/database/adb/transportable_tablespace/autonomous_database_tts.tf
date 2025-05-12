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
  compute_count            = "2.0"
  compute_model            = "ECPU"
  data_storage_size_in_tbs = "1"
  db_name                  = "Xsk5jnfdl12423"
  db_version               = "19c"
  db_workload              = "AJD"
  license_model            = "LICENSE_INCLUDED"
  transportable_tablespace {
    tts_bundle_url = "https://<new_object_storage_url>.com"
  }
}