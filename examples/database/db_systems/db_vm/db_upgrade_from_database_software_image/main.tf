# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_upgrade_from_database_software_image
#    NOTES
#      Terraform Integration Test: TestDatabaseDatabaseUpgradeResource_DbSoftwareImage
#
#    FILE(S)
#      database_database_upgrade_resource_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   12/16/2024 - Created


resource "oci_database_db_system" "test_db_system_for_upgrade" {
  display_name = "tfDbSystemForUpgradeFromDatabaseSoftwareImageExample"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      character_set = "AL32UTF8"
      db_name = "tfDb"
      db_workload = "OLTP"
      ncharacter_set = "AL16UTF16"
      pdb_name = "tfPdb"
    }
    db_version = "12.2.0.1"
    display_name = "tfDbHome"
  }
  db_system_options {
    storage_management = "LVM"
  }
  disk_redundancy = "NORMAL"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  fault_domains = ["FAULT-DOMAIN-1"]
  hostname = "oracle-db"
  license_model = "LICENSE_INCLUDED"
  lifecycle {
    ignore_changes = [db_home[0].db_version, defined_tags]
  }
  node_count = "1"
  nsg_ids = [oci_core_network_security_group.test_nsg.id]
  shape = "VM.Standard2.2"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_database_upgrade" "test_database_precheck" {
  action = "PRECHECK"
  database_id = data.oci_database_databases.test_db_system_for_upgrade.databases.0.id
  database_upgrade_source_details {
    database_software_image_id = var.database_software_image_id
    options = "-upgradeTimezone false -keepEvents"
    source = "DB_SOFTWARE_IMAGE"
  }
  depends_on = [oci_database_db_system.test_db_system_for_upgrade]
}

resource "oci_database_database_upgrade" "test_database_upgrade" {
  action = "UPGRADE"
  database_id = data.oci_database_databases.test_db_system_for_upgrade.databases.0.id
  database_upgrade_source_details {
    database_software_image_id = var.database_software_image_id
    options = "-upgradeTimezone false -keepEvents"
    source = "DB_SOFTWARE_IMAGE"
  }
  depends_on = [oci_database_database_upgrade.test_database_precheck]
}