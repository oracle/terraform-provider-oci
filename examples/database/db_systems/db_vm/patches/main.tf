# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Main file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/db_systems/db_vm/patches
#    NOTES
#      Terraform Integration Test: TestDatabaseDbSystemPatchResource_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   06/05/2025 - Created


resource "oci_database_db_system" "test_db_system" {
  display_name            = "tfDbSystem"
  availability_domain     = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id          = var.compartment_id
  subnet_id               = oci_core_subnet.test_subnet.id
  database_edition        = "ENTERPRISE_EDITION"
  shape                   = "VM.Standard.E4.Flex"
  cpu_core_count          = "4"
  ssh_public_keys         = [var.ssh_public_key]
  domain                  = oci_core_subnet.test_subnet.subnet_domain_name
  hostname                = "tfHost"
  data_storage_size_in_gb = "256"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  db_system_options {
    storage_management = "LVM"
  }
  db_home {
    db_version   = "19.0.0.0"
    display_name = "tfDbHome"
    database {
      admin_password = var.admin_password
      db_name        = "tfDb"
    }
  }
  lifecycle {
    ignore_changes = [defined_tags["Oracle-Tags.CreatedBy"], defined_tags["Oracle-Tags.CreatedOn"]]
  }
}
