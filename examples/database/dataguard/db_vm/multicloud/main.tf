# Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/dataguard/db_vm/multicloud
#    NOTES
#      Terraform Integration Test: TestDatabaseDataGuardAssociationResourceMultiCloud
#
#    FILE(S)
#      database_data_guard_association_multicloud_test.go
#
#    MODIFIED   MM/DD/YY
#    escabrer   08/28/2025 - Created


resource "oci_database_data_guard_association" "test_multicloud_dataguard_association" {
  availability_domain = data.oci_identity_availability_domain.test_multicloud_availability_domain.name
  cluster_placement_group_id = var.multicloud_cluster_placement_group_id
  compute_count = "4"
  compute_model = "ECPU"
  creation_type = "NewDbSystem"
  data_collection_options {
    is_diagnostics_events_enabled = "false"
    is_health_monitoring_enabled = "false"
    is_incident_logs_enabled = "false"
  }
  database_admin_password = var.admin_password
  database_id = data.oci_database_databases.test_multicloud_databases.databases.0.id
  delete_standby_db_home_on_delete = "true"
  display_name = "tfDataguardAssociationMultiCloud"
  domain = var.multicloud_domain
  hostname = "tfpeerdb311"
  is_active_data_guard_enabled = "false"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  nsg_ids = [var.multicloud_nsg_id]
  protection_mode = "MAXIMUM_PERFORMANCE"
  shape = "VM.Standard.x86"
  storage_volume_performance_mode = "HIGH_PERFORMANCE"
  subnet_id = var.multicloud_subnet_id
  subscription_id = var.multicloud_subscription_id
  transport_type = "ASYNC"
}
