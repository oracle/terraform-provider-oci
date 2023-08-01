resource "oci_database_autonomous_vm_cluster_ords_certificate_management" "test_avm_ords_mgmt_res"{
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
  certificate_generation_type = "BYOC"
  certificate_id = var.avm_certificate_id
}

