variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}
variable "bds_instance_id" {}


variable "bds_instance_os_patch_action_cluster_admin_password" {
  default = "T3JhY2xlVGVhbVVTQSExMjM="
}

variable "bds_instance_os_patch_action_os_patch_version" {
  default = "ol7.9-x86_64-1.27.0.696-0.0"
}

variable "bds_instance_os_patch_action_patching_config_strategy" {
  default = "BATCHING_BASED"
}

variable "bds_instance_os_patch_action_batch_size" {
  default = "1"
}

variable "bds_instance_os_patch_action_wait_time_between_batch_in_seconds" {
  default = "120"
}

variable "bds_instance_os_patch_action_tolerance_threshold_per_batch" {
  default = "0"
}

variable "bds_instance_os_patch_action_wait_time_between_domain_in_seconds" {
  default = "120"
}

variable "bds_instance_os_patch_action_tolerance_threshold_per_domain" {
  default = "1"
}

variable "bds_instance_get_os_patch_os_patch_version" {
 default = "ol7.9-x86_64-1.27.0.696-0.0"
}

resource "oci_bds_bds_instance_os_patch_action" "test_bds_instance_os_patch_action" {
  #Required
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.bds_instance_os_patch_action_cluster_admin_password
  os_patch_version       = var.bds_instance_os_patch_action_os_patch_version
  patching_configs {
    patching_config_strategy              = var.bds_instance_os_patch_action_patching_config_strategy
    batch_size                            = var.bds_instance_os_patch_action_batch_size
    wait_time_between_batch_in_seconds    = var.bds_instance_os_patch_action_wait_time_between_batch_in_seconds
    tolerance_threshold_per_batch         = var.bds_instance_os_patch_action_tolerance_threshold_per_batch

    wait_time_between_domain_in_seconds = var.bds_instance_os_patch_action_wait_time_between_domain_in_seconds
    tolerance_threshold_per_domain      = var.bds_instance_os_patch_action_tolerance_threshold_per_domain
  }

  timeouts {
   create = "24h"
   update = "24h"
   delete = "24h"
   }

}

data "oci_bds_bds_instance_get_os_patch" "test_bds_instance_get_os_patch" {
  bds_instance_id   = var.bds_instance_id
  os_patch_version    = var.bds_instance_get_os_patch_os_patch_version
}

data "oci_bds_bds_instance_list_os_patches" "test_bds_instance_list_os_patches" {
  bds_instance_id   = var.bds_instance_id
}