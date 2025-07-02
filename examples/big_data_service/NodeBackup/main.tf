resource "oci_bds_bds_instance_node_backup" "test_bds_instance_node_backup" {
  backup_type     = "FULL"
  bds_instance_id = "${oci_bds_bds_instance.test_bds_instance.id}"
  level_type_details {
    level_type = "NODE_TYPE_LEVEL"
    node_type  = "MASTER"
  }
}
