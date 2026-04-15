
variable "bds_instance_id" {}
variable "bootstrap_script_url" {}
variable "cluster_admin_password" {}

provider "oci" {
}

resource "oci_bds_bds_instance_execute_bootstrap_script_action" "test_bds_instance_execute_bootstrap_script_action" {
bds_instance_id = var.bds_instance_id
bootstrap_script_url = var.bootstrap_script_url
cluster_admin_password = var.cluster_admin_password
}
