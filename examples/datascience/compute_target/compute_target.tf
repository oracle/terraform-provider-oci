// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "user_ocid" {
}

variable "region" {
}

variable config_file_profile {
}

provider "oci" {
  region              = var.region
  user_ocid           = var.user_ocid
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
}

variable "compartment_ocid" {
}

variable "compute_target_description" {
  default = "Compute Target for terraform testing"
}

variable "compute_target_display_name" {
  default = "terraform-testing-compute-target"
}

variable "compute_target_description_fixed_size" {
  default = "test"
}

variable "compute_target_display_name_fixed_size" {
  default = "ct-test"
}

variable "compute_target_display_name_autoscaling" {
  default = "terraform-testing-compute-target-autoscaling"
}

variable "compute_target_freeform_tags" {
  default = {
    Department = "Finance"
  }
}

variable "compute_target_freeform_tags_fixed_size" {
  default = {
    DEV = "true"
  }
}

variable "compute_target_metadata" {
  default = {
    skipImageVerification = "true"
  }
}

variable "compute_target_compute_configuration_details_compute_type" {
  default = "MANAGED_COMPUTE_CLUSTER"
}

variable "compute_target_compute_configuration_details_instance_configuration_instance_shape" {
  default = "VM.Standard.E4.Flex"
}

variable "compute_target_compute_configuration_details_instance_configuration_boot_volume_size_in_gbs" {
  default = 100
}

variable "compute_target_compute_configuration_details_instance_configuration_instance_shape_details_memory_in_gbs_fixed_size" {
  default = 20
}

variable "compute_target_compute_configuration_details_instance_configuration_instance_shape_details_ocpus_fixed_size" {
  default = 10
}

variable "compute_target_compute_configuration_details_instance_configuration_instance_shape_details_memory_in_gbs" {
  default = 10
}

variable "compute_target_compute_configuration_details_instance_configuration_instance_shape_details_ocpus" {
  default = 10
}

variable "compute_target_compute_configuration_details_scaling_policy_policy_type" {
  default = "FIXED_SIZE"
}

variable "compute_target_compute_configuration_details_scaling_policy_fixed_instance_count" {
  default = 1
}

variable "compute_target_compute_configuration_details_scaling_policy_policy_type_autoscaling" {
  default = "AUTOSCALING"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_auto_scaling_policy_type" {
  default = "THRESHOLD"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_initial_instance_count" {
  default = 1
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_maximum_instance_count" {
  default = 2
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_minimum_instance_count" {
  default = 1
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_expression_rule_type" {
  default = "CUSTOM_EXPRESSION"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_type" {
  default = "CPU_UTILIZATION"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_instance_count_adjustment" {
  default = 1
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_pending_duration" {
  default = "PT3M"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_scaling_configuration_type" {
  default = "QUERY"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_threshold" {
  default = 10
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_query" {
  default = "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_instance_count_adjustment" {
  default = 1
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_pending_duration" {
  default = "PT3M"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_scaling_configuration_type" {
  default = "QUERY"
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_threshold" {
  default = 60
}

variable "compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_query" {
  default = "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 99"
}

variable "compute_target_compute_configuration_details_scaling_policy_cool_down_in_seconds" {
  default = 300
}

variable "compute_target_compute_configuration_details_scaling_policy_is_enabled" {
  default = true
}

# A compute target resource configuration for creating a new compute target with scaling policy type = FIXED_SIZE
resource "oci_datascience_compute_target" "tf_compute_target" {
  compartment_id = var.compartment_ocid

  compute_configuration_details {
    compute_type = var.compute_target_compute_configuration_details_compute_type

    instance_configuration {
      instance_shape          = var.compute_target_compute_configuration_details_instance_configuration_instance_shape
      boot_volume_size_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_boot_volume_size_in_gbs

      instance_shape_details {
        memory_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_memory_in_gbs_fixed_size
        ocpus         = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_ocpus_fixed_size
      }
    }

    scaling_policy {
      policy_type    = var.compute_target_compute_configuration_details_scaling_policy_policy_type
      instance_count = var.compute_target_compute_configuration_details_scaling_policy_fixed_instance_count
    }
  }

  description   = var.compute_target_description_fixed_size
  display_name  = var.compute_target_display_name_fixed_size
  freeform_tags = var.compute_target_freeform_tags_fixed_size
  metadata      = var.compute_target_metadata
}

# A compute target resource configuration for creating a new compute target with scaling policy type = AUTOSCALING
resource "oci_datascience_compute_target" "tf_compute_target_autoscaling" {
  compartment_id = var.compartment_ocid

  compute_configuration_details {
    compute_type = var.compute_target_compute_configuration_details_compute_type

    instance_configuration {
      instance_shape          = var.compute_target_compute_configuration_details_instance_configuration_instance_shape
      boot_volume_size_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_boot_volume_size_in_gbs

      instance_shape_details {
        memory_in_gbs = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_memory_in_gbs
        ocpus         = var.compute_target_compute_configuration_details_instance_configuration_instance_shape_details_ocpus
      }
    }

    scaling_policy {
      policy_type          = var.compute_target_compute_configuration_details_scaling_policy_policy_type_autoscaling
      cool_down_in_seconds = var.compute_target_compute_configuration_details_scaling_policy_cool_down_in_seconds
      is_enabled           = var.compute_target_compute_configuration_details_scaling_policy_is_enabled

      auto_scaling_policies {
        auto_scaling_policy_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_auto_scaling_policy_type
        initial_instance_count   = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_initial_instance_count
        maximum_instance_count   = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_maximum_instance_count
        minimum_instance_count   = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_minimum_instance_count

        rules {
          metric_expression_rule_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_expression_rule_type

          scale_in_configuration {
            instance_count_adjustment  = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_instance_count_adjustment
            pending_duration           = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_pending_duration
            query                      = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_query
            scaling_configuration_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_scaling_configuration_type
          }

          scale_out_configuration {
            instance_count_adjustment  = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_instance_count_adjustment
            pending_duration           = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_pending_duration
            query                      = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_query
            scaling_configuration_type = var.compute_target_compute_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_scaling_configuration_type
          }
        }
      }
    }
  }

  description   = var.compute_target_description
  display_name  = var.compute_target_display_name_autoscaling
  freeform_tags = var.compute_target_freeform_tags
  metadata      = var.compute_target_metadata
}

data "oci_datascience_compute_targets" "tf_compute_targets" {
  compartment_id = var.compartment_ocid
}

data "oci_datascience_compute_target" "tf_compute_target" {
  compute_target_id = oci_datascience_compute_target.tf_compute_target_autoscaling.id
}
