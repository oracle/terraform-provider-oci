# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Resources file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_backup
#    NOTES
#      Terraform Integration Test: TestDatabaseBackupResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   11/1/2024 - Created


resource "oci_recovery_protection_policy" "test_protection_policy" {
  display_name = "tfRecoveryServiceSubnetProtectionPolicyExample"
  backup_retention_period_in_days = "14"
  compartment_id = var.compartment_id
}

resource "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet_registration" {
  display_name = "tfRecoveryServiceSubnetRegistrationExample"
  compartment_id = var.compartment_id
  subnets = [oci_core_subnet.test_private_subnet.id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_database_db_system" "test_db_system" {
  display_name = "tfDbSystemWithDatabaseBackupExample"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      db_backup_config {
        auto_backup_enabled = "true"
        auto_backup_window = "SLOT_TWO"
        backup_deletion_policy = "DELETE_IMMEDIATELY"
        backup_destination_details {
          dbrs_policy_id = oci_recovery_protection_policy.test_protection_policy.id
          type = "DBRS"
        }
        run_immediate_full_backup = "true"
      }
      db_name = "tfDb"
    }
    db_version = "19.0.0.0"
    display_name = "tfDbHome"
  }
  disk_redundancy = "NORMAL"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "tf-oracle-db"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard2.2"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
  depends_on = [oci_recovery_recovery_service_subnet.test_recovery_service_subnet_registration]
}