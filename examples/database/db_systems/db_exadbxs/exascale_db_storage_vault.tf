// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_exascale_db_storage_vault" "test_exascale_db_storage_vault" {
  #Required
  availability_domain = local.ad
  compartment_id      = var.compartment_ocid
  display_name        = "ExampleExascaleDbStorageVault"
  high_capacity_database_storage {
    total_size_in_gbs = 800
  }
  additional_flash_cache_in_percent = 20
}

data "oci_database_exascale_db_storage_vaults" "test_exascale_db_storage_vaults" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_database_exascale_db_storage_vault" "test_exascale_db_storage_vault" {
  #Required
  exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
}