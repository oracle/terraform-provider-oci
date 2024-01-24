
// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "bds_instance_id" {}

variable "auto_scaling_configuration_cluster_admin_password" {
  default = "V2VsY29tZTE="
}

variable "auto_scaling_configuration_display_name" {
  default = "displayName"
}

variable "auto_scaling_configuration_is_enabled" {
  default = true
}

variable "auto_scaling_configuration_node_type" {
  default = "WORKER"
}

variable "auto_scaling_configuration_policy_policy_type" {
  default = "METRIC_BASED_VERTICAL_SCALING_POLICY"
}

variable "auto_scaling_configuration_policy_scale_up_rules_action" {
  default = "CHANGE_SHAPE_SCALE_UP"
}

variable "auto_scaling_configuration_policy_scale_up_rules_metric_metric_type" {
  default = "CPU_UTILIZATION"
}

variable "auto_scaling_configuration_policy_scale_up_rules_metric_threshold_duration_in_minutes" {
  default = 10
}

variable "auto_scaling_configuration_policy_scale_up_rules_metric_threshold_operator" {
  default = "GT"
}

variable "auto_scaling_configuration_policy_scale_up_rules_metric_threshold_value" {
  default = 90
}

variable "auto_scaling_configuration_policy_scale_down_rules_action" {
  default = "CHANGE_SHAPE_SCALE_DOWN"
}

variable "auto_scaling_configuration_policy_scale_down_rules_metric_metric_type" {
  default = "CPU_UTILIZATION"
}

variable "auto_scaling_configuration_policy_scale_down_rules_metric_threshold_duration_in_minutes" {
  default = 10
}

variable "auto_scaling_configuration_policy_scale_down_rules_metric_threshold_operator" {
  default = "LT"
}

variable "auto_scaling_configuration_policy_scale_down_rules_metric_threshold_value" {
  default = 10
}

variable "auto_scaling_configuration_state" {
  default = "ACTIVE"
}

variable "auto_scaling_configuration_policy_memory_step_size" {
  default = 2
}

variable "auto_scaling_configuration_policy_ocpu_step_size" {
  default = 2
}

variable "auto_scaling_configuration_policy_scale_down_config_min_memory_per_node" {
  default = 16
}

variable "auto_scaling_configuration_policy_scale_down_config_min_ocpus_per_node" {
  default = 3
}

variable "auto_scaling_configuration_policy_scale_down_config_max_memory_per_node" {
  default = 32
}

variable "auto_scaling_configuration_policy_scale_down_config_max_ocpus_per_node" {
  default = 6
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_bds_auto_scaling_configuration" "test_auto_scaling_configuration" {
  #Required
  bds_instance_id        = var.bds_instance_id
  cluster_admin_password = var.auto_scaling_configuration_cluster_admin_password
  is_enabled             = var.auto_scaling_configuration_is_enabled
  node_type              = var.auto_scaling_configuration_node_type
  policy_details {
    policy_type = var.auto_scaling_configuration_policy_policy_type
    scale_down_config {
      memory_step_size = var.auto_scaling_configuration_policy_memory_step_size
      metric {
        metric_type = var.auto_scaling_configuration_policy_scale_down_rules_metric_metric_type
        threshold {
          duration_in_minutes = var.auto_scaling_configuration_policy_scale_down_rules_metric_threshold_duration_in_minutes
          operator            = var.auto_scaling_configuration_policy_scale_down_rules_metric_threshold_operator
          value               = var.auto_scaling_configuration_policy_scale_down_rules_metric_threshold_value
        }
      }
      min_memory_per_node = var.auto_scaling_configuration_policy_scale_down_config_min_memory_per_node
      min_ocpus_per_node  = var.auto_scaling_configuration_policy_scale_down_config_min_ocpus_per_node
      ocpu_step_size      = var.auto_scaling_configuration_policy_ocpu_step_size
    }
    scale_up_config {
      memory_step_size = var.auto_scaling_configuration_policy_memory_step_size
      metric {
        metric_type = var.auto_scaling_configuration_policy_scale_down_rules_metric_metric_type
        threshold {
          duration_in_minutes = var.auto_scaling_configuration_policy_scale_up_rules_metric_threshold_duration_in_minutes
          operator            = var.auto_scaling_configuration_policy_scale_up_rules_metric_threshold_operator
          value               = var.auto_scaling_configuration_policy_scale_up_rules_metric_threshold_value
        }
      }
      max_memory_per_node = var.auto_scaling_configuration_policy_scale_down_config_max_memory_per_node
      max_ocpus_per_node  = var.auto_scaling_configuration_policy_scale_down_config_max_ocpus_per_node
      ocpu_step_size      = var.auto_scaling_configuration_policy_ocpu_step_size
    }
  }

  #Optional
  display_name = var.auto_scaling_configuration_display_name
}

data "oci_bds_auto_scaling_configurations" "test_auto_scaling_configuration" {
  #Required
  bds_instance_id = var.bds_instance_id
  compartment_id  = var.compartment_id

  #Optional
  display_name = var.auto_scaling_configuration_display_name
  state        = var.auto_scaling_configuration_state
}

data "oci_bds_auto_scaling_configuration" "test_auto_scaling_configuration" {
  #Required
  auto_scaling_configuration_id  = oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id
  bds_instance_id                = var.bds_instance_id
}

