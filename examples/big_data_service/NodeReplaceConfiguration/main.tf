resource "oci_bds_bds_instance_node_replace_configuration" "test_bds_instance_node_replace_configuration" {
  bds_instance_id        = "${oci_bds_bds_instance.test_bds_instance.id}"
  cluster_admin_password = "T3JhY2xlVGVhbVVTQSExMjM="
  display_name           = "displayName"
  duration_in_minutes    = "20"
  level_type_details {
    level_type = "NODE_TYPE_LEVEL"
    node_type  = "UTILITY"
  }
  metric_type = "INSTANCE_STATUS"
}
