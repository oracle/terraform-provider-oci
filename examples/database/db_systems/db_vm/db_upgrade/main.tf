# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - DB System Upgrade Resources
#
#    USAGE
#      Use the following path for Example Test & Backward Compatibility Test: database/db_systems/db_vm/db_upgrade
#
#    NOTES
#      Terraform Example: TestDatabaseDatabaseUpgradeResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/23/2024 - Created



resource "oci_database_db_system" "test_db_system_for_upgrade" {
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
  display_name = "tfDbSystemForUpgradeExample"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  fault_domains = ["FAULT-DOMAIN-1"]
  hostname = "tfOracleDb"
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
    db_version = "19.24.0.0"
    source = "DB_VERSION"
  }
}

resource "oci_database_database_upgrade" "test_database_upgrade" {
  action = "UPGRADE"
  database_id = data.oci_database_databases.test_db_system_for_upgrade.databases.0.id
  database_upgrade_source_details {
    db_version = "19.24.0.0"
    options = "-upgradeTimezone false -keepEvents"
    source = "DB_VERSION"
  }
  depends_on = [oci_database_database_upgrade.test_database_precheck]
}