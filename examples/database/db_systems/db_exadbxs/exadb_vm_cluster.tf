// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_exadb_vm_cluster" "test_exadb_vm_cluster" {

  #Required
  availability_domain          = local.ad
  compartment_id               = var.compartment_ocid
  display_name                 = "ExampleExadbVmCluster"
  exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
  grid_image_id                = data.oci_database_gi_version_minor_versions.test_gi_minor_versions.gi_minor_versions[0].grid_image_id
  hostname                     = "apollo"
  cluster_name                 = "apollo"
  shape                        = "EXADBXS"
  ssh_public_keys              = [var.ssh_public_key]
  subnet_id                    = oci_core_subnet.exadbxs_client_subnet.id
  backup_subnet_id             = oci_core_subnet.exadbxs_backup_subnet.id

  node_config {
    enabled_ecpu_count_per_node              = "8"
    total_ecpu_count_per_node                = "16"
    vm_file_system_storage_size_gbs_per_node = "293"
  }

  node_resource {
    node_name = "node1"
  }

  node_resource {
    node_name = "node2"
  }

  node_resource {
    node_name = "node3"
  }

}

data "oci_database_exadb_vm_clusters" "test_exadb_vm_clusters" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
}

data "oci_database_exadb_vm_cluster" "test_exadb_vm_cluster" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
}