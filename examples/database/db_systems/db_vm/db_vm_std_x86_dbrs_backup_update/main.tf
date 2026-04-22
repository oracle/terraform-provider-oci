# $Header$
#
# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Resources file
#
#    USAGE
#      Example & Backward Compatibility Path: database/db_systems/db_vm/db_vm_std_x86_dbrs_backup_update
#    NOTES
#      Mirrors the target update shape used by TestResourceDatabaseDBSystemVMStdx86BackupConfigNoRecreateUpdate.
#      This is the desired end-state config for an in-place update scenario.
#      For DBRS destinations, recovery_window_in_days is intentionally omitted.

resource "oci_recovery_protection_policy" "test_protection_policy" {
  display_name                    = "tfRecoveryServiceSubnetProtectionPolicyX86UpdateExample"
  backup_retention_period_in_days = "14"
  compartment_id                  = var.compartment_id
}

resource "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet_registration" {
  display_name   = "tfRecoveryServiceSubnetRegistrationX86UpdateExample"
  compartment_id = var.compartment_id
  subnets        = [oci_core_subnet.test_private_subnet.id]
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_database_db_system" "test_db_system" {
  display_name            = "tfDbSystemVmStdx86DBRSBackupUpdateExample"
  availability_domain     = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id          = var.compartment_id
  compute_count           = "4"
  compute_model           = "ECPU"
  data_storage_size_in_gb = "256"
  database_edition        = "ENTERPRISE_EDITION"
  disk_redundancy         = "NORMAL"
  domain                  = oci_core_subnet.test_subnet.subnet_domain_name
  hostname                = "myoracledb"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  shape                   = "VM.Standard.x86"
  ssh_public_keys         = [var.ssh_public_key]
  subnet_id               = oci_core_subnet.test_subnet.id

  db_system_options {
    storage_management = "LVM"
  }

  freeform_tags = {
    "Department" = "Admin"
  }

  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      character_set  = "AL32UTF8"
      db_name        = "tfDb"
      db_workload    = "OLTP"
      ncharacter_set = "AL16UTF16"
      pdb_name       = "tfPdb"

      db_backup_config {
        auto_backup_enabled    = true
        auto_full_backup_day   = "SUNDAY"
        backup_deletion_policy = "DELETE_AFTER_RETENTION_PERIOD"

        backup_destination_details {
          dbrs_policy_id            = oci_recovery_protection_policy.test_protection_policy.id
          is_zero_data_loss_enabled = false
          type                      = "DBRS"
        }
      }
    }

    db_version   = "19.0.0.0"
    display_name = "tfDbHome"
  }

  depends_on = [oci_recovery_recovery_service_subnet.test_recovery_service_subnet_registration]
}
