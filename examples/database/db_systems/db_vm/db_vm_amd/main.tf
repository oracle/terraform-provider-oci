# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_vm_amd
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemAmdVM
#
#    FILE(S)
#      database_db_system_resource_amd_vm_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/12/2024 - Created


resource "oci_database_db_system" "test_amd_db_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_ocid
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = var.admin_password
      character_set = "AL32UTF8"
      db_backup_config {
        auto_backup_enabled = "false"
      }
      db_name = "tfDb"
      db_workload = "OLTP"
      kms_key_id = var.kms_key_id
      ncharacter_set = "AL16UTF16"
      pdb_name = "tfPdb"
      vault_id = var.vault_id
    }
    db_version = "19.0.0.0"
    display_name = "tfDbHome"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemAmd"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "oracle-db"
  kms_key_id = var.kms_key_id
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.E4.Flex"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}