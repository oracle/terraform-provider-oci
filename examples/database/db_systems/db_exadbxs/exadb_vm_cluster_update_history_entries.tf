data "oci_database_exadb_vm_cluster_update_history_entries" "test_exadb_vm_cluster_update_history_entries" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  #Optional
  update_type = "OS_UPDATE"
}

data "oci_database_exadb_vm_cluster_update_history_entry" "test_exadb_vm_cluster_update_history_entry" {
  #Required
  exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  #Optional
  update_history_entry_id = data.oci_database_exadb_vm_cluster_update_history_entries.test_exadb_vm_cluster_update_history_entries.exadb_vm_cluster_update_history_entries[0].id
}