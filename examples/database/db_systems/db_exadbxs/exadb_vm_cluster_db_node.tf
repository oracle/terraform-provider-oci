// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Get db node list
data "oci_database_db_nodes" "test_exadb_vm_cluster_db_nodes" {
  compartment_id = var.compartment_ocid
  vm_cluster_id   = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
}

# Get db node details
data "oci_database_db_node" "test_exadb_vm_cluster_db_node" {
  db_node_id = data.oci_database_db_nodes.test_exadb_vm_cluster_db_nodes.db_nodes[0].id
}