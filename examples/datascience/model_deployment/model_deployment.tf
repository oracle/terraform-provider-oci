// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "compartment_ocid" {
}

variable "project_ocid" {
}

variable "shape" {
}

variable "model_id" {
}

variable "model_deployment_display_name" {
  default = "terraform-testing-model-deployment"
}

variable "model_deployment_description" {
  default = "Model Deployment for terraform testing"
}

variable "log_group_id" {
}

variable "model_deployment_defined_tags" {
}

variable "model_deployment_freeform_tag" {
}

variable "access_log_id" {
}

variable "predict_log_id" {
}

variable "model_deployment_model_deployment_configuration_details_deployment_type" {
  default = "SINGLE_MODEL"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps" {
  default = 10
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_maximum_bandwidth_mbps" {
  default = 10
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_instance_count" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_policy_type" {
  default = "FIXED_SIZE"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_cpu_baseline" {
  default = "BASELINE_1_8"
}

variable "model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_ocpus" {
  default = 1.0
}

variable "model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_memory_in_gbs" {
  default = 6.0
}

variable "model_deployment_state" {
  default = "ACTIVE"
}

# these variables are for scaling policy type = AUTOSCALING
variable "model_deployment_display_name_for_autoscaling_deployment" {
  default = "terraform-testing-autoscaling-model-deployment"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_type_autoscaling" {
  default = "AUTOSCALING"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_auto_scaling_policy_type" {
  default = "THRESHOLD"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_initial_instance_count" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_maximum_instance_count" {
  default = 2
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_minimum_instance_count" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_expression_rule_type" {
  default = "PREDEFINED_EXPRESSION"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_type" {
  default = "CPU_UTILIZATION"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_instance_count_adjustment" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_pending_duration" {
  default = "PT5M"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_scaling_configuration_type" {
  default = "THRESHOLD"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_threshold" {
  default = 10
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_instance_count_adjustment" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_pending_duration" {
  default = "PT3M"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_scaling_configuration_type" {
  default = "THRESHOLD"
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_threshold" {
  default = 70
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_cool_down_in_seconds" {
  default = 600
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_is_enabled" {
  default = true
}

# these variables for BYOC option
variable "model_byoc_id" {
}

variable "model_deployment_byoc_display_name" {
  default = "terraform-testing-model-deployment_byoc"
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_configuration_type" {
  default = "OCIR_CONTAINER"
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_cmd" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_entrypoint" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_variables" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_health_check_port" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_image" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_image_digest" {
}

variable "model_deployment_model_deployment_configuration_details_environment_configuration_details_server_port" {
}

# A model deployment resource configurations for creating a new model deployment with scaling policy type = FIXED SIZE
resource "oci_datascience_model_deployment" "tf_model_deployment" {
  # Required
  compartment_id = var.compartment_ocid
  model_deployment_configuration_details {
    # Required
    deployment_type = var.model_deployment_model_deployment_configuration_details_deployment_type
    model_configuration_details {
      # Required
      instance_configuration {
        # Required
        instance_shape_name = var.shape

        #Optional
        model_deployment_instance_shape_config_details {

          #Optional
          cpu_baseline  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_cpu_baseline
          memory_in_gbs = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_memory_in_gbs
          ocpus         = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_ocpus
        }
      }
      model_id = var.model_id

      # Optional
      bandwidth_mbps         = var.model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps
      maximum_bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_maximum_bandwidth_mbps

      scaling_policy {
        # Required
        instance_count = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_instance_count
        policy_type    = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_policy_type
      }
    }
  }
  project_id = var.project_ocid

  # Optional
  category_log_details {

    # Optional
    access {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.access_log_id
    }
    predict {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.predict_log_id
    }
  }
  # Optional
//  defined_tags  = var.model_deployment_defined_tags
  description   = var.model_deployment_description
  display_name  = var.model_deployment_display_name
//  freeform_tags = var.model_deployment_freeform_tag
}

# A model deployment resource configurations for creating a new model deployment with scaling policy type = AUTOSCALING
resource "oci_datascience_model_deployment" "tf_model_deployment_autoscaling" {
  # Required
  compartment_id = var.compartment_ocid
  model_deployment_configuration_details {
    # Required
    deployment_type = var.model_deployment_model_deployment_configuration_details_deployment_type
    model_configuration_details {
      # Required
      instance_configuration {
        # Required
        instance_shape_name = var.shape

        #Optional
        model_deployment_instance_shape_config_details {

          #Optional
          cpu_baseline  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_cpu_baseline
          memory_in_gbs = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_memory_in_gbs
          ocpus         = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_ocpus
        }
      }
      model_id = var.model_id

      # Optional
      bandwidth_mbps         = var.model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps
      maximum_bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_maximum_bandwidth_mbps

      scaling_policy {
        # Required
        policy_type    = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_type_autoscaling

        #Optional
        auto_scaling_policies {
          #Required
          auto_scaling_policy_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_auto_scaling_policy_type
          initial_instance_count   = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_initial_instance_count
          maximum_instance_count   = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_maximum_instance_count
          minimum_instance_count   = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_minimum_instance_count
          rules {
            #Required
            metric_expression_rule_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_expression_rule_type
            scale_in_configuration {

              #Optional
              instance_count_adjustment  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_instance_count_adjustment
              pending_duration           = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_pending_duration
              scaling_configuration_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_scaling_configuration_type
              threshold                  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_in_configuration_threshold
            }
            scale_out_configuration {

              #Optional
              instance_count_adjustment  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_instance_count_adjustment
              pending_duration           = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_pending_duration
              scaling_configuration_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_scaling_configuration_type
              threshold                  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_scale_out_configuration_threshold
            }

            #Optional
            metric_type = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_auto_scaling_policies_rules_metric_type
          }
        }
        cool_down_in_seconds = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_cool_down_in_seconds
        is_enabled           = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_is_enabled
      }
    }
  }
  project_id = var.project_ocid

  # Optional
  category_log_details {

    # Optional
    access {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.access_log_id
    }
    predict {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.predict_log_id
    }
  }
  # Optional
  description   = var.model_deployment_description
  display_name  = var.model_deployment_display_name_for_autoscaling_deployment
}

resource "oci_datascience_model_deployment" "tf_model_deployment_byoc" {
  # Required
  compartment_id = var.compartment_ocid
  model_deployment_configuration_details {
    # Required
    deployment_type = var.model_deployment_model_deployment_configuration_details_deployment_type
    model_configuration_details {
      # Required
      instance_configuration {
        # Required
        instance_shape_name = var.shape

        #Optional
        model_deployment_instance_shape_config_details {

          #Optional
          cpu_baseline  = var.model_deployment_model_deployment_configuration_details_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_cpu_baseline
          memory_in_gbs = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_memory_in_gbs
          ocpus         = var.model_deployment_model_configuration_details_instance_configuration_model_deployment_instance_shape_config_details_ocpus
        }
      }
      model_id = var.model_byoc_id

      # Optional
      bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps
      maximum_bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_maximum_bandwidth_mbps

      scaling_policy {
        # Required
        instance_count = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_instance_count
        policy_type    = var.model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_policy_type
      }
    }

    # Optional for BYOC
    environment_configuration_details {
      #Required
      environment_configuration_type = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_configuration_type

      #Optional
      cmd = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_cmd
      entrypoint = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_entrypoint
      environment_variables = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_environment_variables
      health_check_port = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_health_check_port
      image = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_image
      image_digest = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_image_digest
      server_port = var.model_deployment_model_deployment_configuration_details_environment_configuration_details_server_port
    }
  }
  project_id = var.project_ocid

  # Optional
  category_log_details {

    # Optional
    access {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.access_log_id
    }
    predict {
      # Required
      log_group_id = var.log_group_id
      log_id       = var.predict_log_id
    }
  }
  # Optional
  //  defined_tags  = var.model_deployment_defined_tags
  description   = var.model_deployment_description
  display_name  = var.model_deployment_byoc_display_name
  //  freeform_tags = var.model_deployment_freeform_tag
}

# The data resource for a list of model deployments in a specified compartment
data "oci_datascience_model_deployments" "tf_model_deployments" {
  # Required
  compartment_id = var.compartment_ocid

  # Optional
  created_by   = var.user_ocid
  display_name = oci_datascience_model_deployment.tf_model_deployment.display_name
  id           = oci_datascience_model_deployment.tf_model_deployment.id
  project_id   = var.project_ocid
  state        = var.model_deployment_state
}