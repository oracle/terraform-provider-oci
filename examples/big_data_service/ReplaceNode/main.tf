variable "node_backup_id" {
}

variable "node_host_name" {
}

resource "oci_bds_bds_instance_replace_node_action" "test_bds_instance_replace_node_action" {
  bds_instance_id        = "${oci_bds_bds_instance.test_bds_instance.id}"
  cluster_admin_password = "T3JhY2xlVGVhbVVTQSExMjM="
  node_backup_id         = var.node_backup_id
  node_host_name         = var.node_host_name
}
