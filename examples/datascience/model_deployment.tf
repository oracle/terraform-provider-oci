
// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

variable "model_display_name" {
  default = "terraform-testing-model"
}

variable "model_description" {
  default = "Model for terraform testing"
}

variable "artifact_content_length" {
}

variable "model_artifact" {
}

variable "content_disposition" {
}

variable "shape" {
}

variable "model_defined_tags" {
}

variable "model_freeform_tag" {
}

variable "model_state" {
}


# A model resource configurations for creating a new model
resource "oci_datascience_model" "tf_model" {
  # Required
  artifact_content_length = var.artifact_content_length
  model_artifact          = var.model_artifact
  compartment_id          = var.compartment_ocid
  project_id              = var.project_ocid
  # Optional
  artifact_content_disposition = var.content_disposition
  defined_tags  = var.model_defined_tags
  description   = var.model_description
  display_name  = var.model_display_name
  freeform_tags = var.model_freeform_tag
}

# A data resource for the list of models in a specified compartment
data "oci_datascience_models" "tf_models" {
  # Required
  compartment_id = var.compartment_ocid
  # Optional
  created_by   = var.user_ocid
  display_name = oci_datascience_model.tf_model.display_name
  id           = oci_datascience_model.tf_model.id
  project_id   = var.project_ocid
  state        = var.model_state
}

# The data source for a list of model deployment shapes
data "oci_datascience_model_deployment_shapes" "tf_model_deployment_shapes" {
  # Required
  compartment_id = var.compartment_ocid
}

variable "model_deployment_display_name" {
  default = "terraform-testing-model-deployment"
}

variable "model_deployment_description" {
  default = "Model Deployment for terraform testing"
}

variable "shape" {
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

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_instance_count" {
  default = 1
}

variable "model_deployment_model_deployment_configuration_details_model_configuration_details_scaling_policy_policy_type" {
  default = "FIXED_SIZE"
}

variable "model_deployment_state" {
}

# A model deployment resource configurations for creating a new model deployment
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
      }
      model_id = oci_datascience_model.tf_model.id

      # Optional
      bandwidth_mbps = var.model_deployment_model_deployment_configuration_details_model_configuration_details_bandwidth_mbps
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
  defined_tags  = var.model_deployment_defined_tags
  description   = var.model_deployment_description
  display_name  = var.model_deployment_display_name
  freeform_tags = var.model_deployment_freeform_tag
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

