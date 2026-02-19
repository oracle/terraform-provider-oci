variable "bds_instance_id" {}
variable "bds_instance_cluster_admin_password" {}
variable "secret_id" {}

resource "oci_bds_bds_instance_resource_principal_configuration" "test_bds_instance_resource_principal_configuration" {
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.bds_instance_cluster_admin_password // Comment this if secret_id usage is preferred.
  // secret_id              = var.secret_id
  display_name           = "displayName"
  force_refresh_resource_principal_trigger = "0"
  session_token_life_span_duration_in_hours = "1"
}

data "oci_bds_bds_instance_resource_principal_configuration" "test_bds_instance_resource_principal_configuration" {
  bds_instance_id = var.bds_instance_id
  resource_principal_configuration_id = "${oci_bds_bds_instance_resource_principal_configuration.test_bds_instance_resource_principal_configuration.id}"
}
