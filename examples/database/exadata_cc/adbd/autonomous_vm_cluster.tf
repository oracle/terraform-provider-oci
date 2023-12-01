// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.


resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
  #Required
  compartment_id            = var.compartment_ocid
  db_servers                = [data.oci_database_db_servers.test_db_servers.db_servers.0.id, data.oci_database_db_servers.test_db_servers.db_servers.1.id]
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id
  cpu_core_count_per_node   = "20"
  autonomous_data_storage_size_in_tbs = "2.0"
  memory_per_oracle_compute_unit_in_gbs = "12"
  total_container_databases             = "2"
  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  is_mtls_enabled         = "true"
  scan_listener_port_tls  = "3600"
  scan_listener_port_non_tls = "1600"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_autonomous_vm_cluster" "autonomous_vm_cluster_2" {
  #Required
  compartment_id            = var.compartment_ocid
  display_name              = "peerAutonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network2.id
  cpu_core_count_per_node   = "20"
  autonomous_data_storage_size_in_tbs = "2.0"
  memory_per_oracle_compute_unit_in_gbs = "12"
  total_container_databases             = "2"
  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Accounts"
  }
}





