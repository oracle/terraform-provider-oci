data "oci_database_exadb_vm_cluster_updates" "test_exadb_vm_cluster_updates" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  #Optional
  update_type = "GI_PATCH"
}

data "oci_database_exadb_vm_cluster_update" "test_exadb_vm_cluster_update" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  update_id           = data.oci_database_exadb_vm_cluster_updates.test_exadb_vm_cluster_updates.exadb_vm_cluster_updates[0].id
}