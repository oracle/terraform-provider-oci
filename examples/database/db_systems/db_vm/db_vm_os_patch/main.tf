# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.

resource "oci_database_db_system" "test_db_system_os_patch" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_id
  subnet_id           = oci_core_subnet.test_subnet.id

  database_edition = "ENTERPRISE_EDITION"
  disk_redundancy  = "NORMAL"
  shape            = "VM.Standard.x86"
  compute_model    = "ECPU"
  compute_count    = 4

  ssh_public_keys = [var.ssh_public_key]
  domain          = oci_core_subnet.test_subnet.subnet_domain_name
  hostname        = "tfDbVmOsPatch"
  display_name    = "tfDbSystemVmOsPatchExample"

  data_storage_size_in_gb = 256
  license_model           = "LICENSE_INCLUDED"
  node_count              = 1

  db_system_options {
    storage_management = "LVM"
  }

  db_home {
    display_name = "tfDbHome"
    db_version   = "19.0.0.0"
    database {
      admin_password = var.admin_password
      db_name        = "tfDb"
      pdb_name       = "tfPdb"
      db_workload    = "OLTP"
      character_set  = "AL32UTF8"
      ncharacter_set = "AL16UTF16"

      db_backup_config {
        auto_backup_enabled = false
      }
    }
  }

  nsg_ids = [oci_core_network_security_group.test_nsg.id]

  os_patch_action  = var.os_patch_action
  os_patch_trigger = var.os_patch_trigger
}
