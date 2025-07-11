# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/multicloud
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemMultiCloud
#
#    FILE(S)
#      database_db_system_resource_multicloud_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   08/28/2025 - Created


resource "oci_database_db_system" "test_multicloud_db_system" {
  availability_domain = data.oci_identity_availability_domain.test_multicloud_availability_domain.name
  cluster_placement_group_id = var.multicloud_cluster_placement_group_id
  compartment_id = var.multicloud_compartment_id
  compute_count = "8"
  compute_model = "ECPU"
  data_collection_options {
    is_diagnostics_events_enabled = "false"
    is_health_monitoring_enabled = "false"
    is_incident_logs_enabled = "false"
  }
  data_storage_percentage = "80"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = var.admin_password
      character_set = "AL32UTF8"
      db_backup_config {
        auto_backup_enabled = "false"
      }
      db_name = "tfdb312"
      db_workload = "OLTP"
      ncharacter_set = "AL16UTF16"
    }
    db_version = "19.26.0.0"
    display_name = "tfDbHomeMultiCloud"
  }
  db_system_options {
    storage_management = "LVM"
  }
  display_name = "tfDBSystemMultiCloud"
  domain = var.multicloud_domain
  hostname = "tfdbhost312"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  nsg_ids = [var.multicloud_nsg_id]
  shape = "VM.Standard.x86"
  source = "NONE"
  ssh_public_keys = [var.ssh_public_key]
  storage_volume_performance_mode = "HIGH_PERFORMANCE"
  subnet_id = var.multicloud_subnet_id
  subscription_id = var.multicloud_subscription_id
}
