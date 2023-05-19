resource "oci_database_autonomous_vm_cluster_ssl_certificate_management" "test_avm_db_mgmt_res"{
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
  certificate_generation_type = "SYSTEM"
}

