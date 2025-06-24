# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Main file
#
#    USAGE
#      Use the following path for Example and Backward-Compatibility Tests: database/db_systems/db_vm/db_management
#    NOTES
#      Associated Integration Test: TestDatabaseCloudDatabaseManagementResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/23/2025 - Created


resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "4"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = var.admin_password
      character_set = "AL32UTF8"
      db_name = "tfDb"
      db_workload = "OLTP"
      kms_key_id = var.kms_key_id
      kms_key_version_id = var.kms_key_version_id
      ncharacter_set = "AL16UTF16"
      pdb_name = "tfPdb"
      vault_id = var.vault_id
    }
    db_version = "19.0.0.0"
    display_name = "tfDbHome"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfDbSystem"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  fault_domains = ["FAULT-DOMAIN-1"]
  hostname = "tfDbHost"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.E4.Flex"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id

  lifecycle {
    ignore_changes = [defined_tags["Oracle-Tags.CreatedBy"], defined_tags["Oracle-Tags.CreatedOn"]]
  }
}

resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  compartment_id = var.compartment_id
  name = "tfPrivateEndpoint"
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_cloud_database_management" "test_database_cloud_database_management" {
  credentialdetails {
    password_secret_id = var.ssl_secret_id
    user_name = "sys"
  }
  database_id = data.oci_database_databases.test_databases.databases.0.id
  enable_management = "true"
  management_type = "BASIC"
  private_end_point_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
  service_name = "${data.oci_database_databases.test_databases.databases.0.db_unique_name}.${oci_core_subnet.test_subnet.subnet_domain_name}"
}
