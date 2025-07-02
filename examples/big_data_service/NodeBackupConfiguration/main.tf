resource "oci_bds_bds_instance_node_backup_configuration" "test_bds_instance_node_backup_configuration" {
  backup_type     = "FULL"
  bds_instance_id = "${oci_bds_bds_instance.test_bds_instance.id}"
  level_type_details {
    level_type = "NODE_TYPE_LEVEL"
    node_type  = "MASTER"
  }
  number_of_backups_to_retain = "1"
  schedule                    = "FREQ=WEEKLY;BYDAY=MO;BYHOUR=10"
}
