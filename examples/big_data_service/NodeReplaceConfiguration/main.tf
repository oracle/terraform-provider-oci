variable "bds_instance_id" {}
variable "bds_instance_cluster_admin_password" {}
variable "secret_id" {}

resource "oci_bds_bds_instance_node_replace_configuration" "test_bds_instance_node_replace_configuration" {
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.bds_instance_cluster_admin_password // Comment this if secret_id usage is preferred.
  // secret_id              = var.secret_id
  display_name           = "displayName"
  duration_in_minutes    = "20"
  level_type_details {
    level_type = "NODE_TYPE_LEVEL"
    node_type  = "UTILITY"
  }
  metric_type = "INSTANCE_STATUS"
}

data "oci_bds_bds_instance_node_replace_configuration" "test_bds_instance_node_replace_configuration" {
  bds_instance_id = var.bds_instance_id
  node_replace_configuration_id = "${oci_bds_bds_instance_node_replace_configuration.test_bds_instance_node_replace_configuration.id}"
}
