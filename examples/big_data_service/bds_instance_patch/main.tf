variable "bds_instance_patch_action_cluster_admin_password" {
  default = "V2VsY29tZTE="
}

variable "bds_instance_patch_action_version" {
  default = "ODH-1.1.0.379"
}

variable "bds_instance_patch_history_patch_version" {}

variable "bds_instance_patch_history_state" {}

variable "bds_instance_id" {}

resource "oci_bds_bds_instance_patch_action" "test_bds_instance_patch_action" {
  #Required
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.bds_instance_patch_action_cluster_admin_password
  version                = var.bds_instance_patch_action_version
}

data "oci_bds_bds_instance_patches" "test_bds_instance_patches" {
  #Required
  bds_instance_id = var.bds_instance_id
}

data "oci_bds_bds_instance_patch_histories" "test_bds_instance_patch_histories" {
  #Required
  bds_instance_id = var.bds_instance_id

  #Optional
  patch_version = var.bds_instance_patch_history_patch_version
  state         = var.bds_instance_patch_history_state
}