// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {

}
variable "user_ocid" {

}
variable "fingerprint" {

}
variable "private_key_path" {

}
variable "region" {
    default = "us-ashburn-1"
}
variable "compartment_ocid" {
    default = "ocid1.tenancy.oc1..aaaaaaaahzy3x4boh7ipxyft2rowu2xeglvanlfewudbnueugsieyuojkldq"
}

variable "pipeline_configuration_details_command_line_arguments" {
  default ="commandLineArguments"

}

variable "pipeline_configuration_details_environment_variables" {
  default = {
        "environmentVariables": "environmentVariables"
  }
}

variable "pipeline_configuration_details_maximum_runtime_in_minutes" {
  default = 20
}

variable "pipeline_configuration_details_type" {
  default = "DEFAULT"
}

variable "pipeline_defined_tags_value" {
  default = "value"
}

variable "pipeline_description" {
  default = "description"
}

variable "pipeline_display_name" {
  default = "displayName"
}

variable "pipeline_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "pipeline_id" {
  default = "id"
}

variable "pipeline_infrastructure_configuration_details_block_storage_size_in_gbs" {
  default = 50
}

variable "pipeline_infrastructure_configuration_details_shape_config_details_memory_in_gbs" {
  default = 1.0
}

variable "pipeline_infrastructure_configuration_details_shape_config_details_ocpus" {
  default = 1.0
}

variable "pipeline_log_configuration_details_enable_auto_log_creation" {
  default = true
}

variable "pipeline_log_configuration_details_enable_logging" {
  default = true
}

variable "pipeline_state" {
  default = "AVAILABLE"
}

variable "pipeline_step_details_depends_on" {
  default = []
}

variable "pipeline_step_details_description" {
  default = "description"
}

variable "pipeline_step_details_is_artifact_uploaded" {
  default = false
}

variable "pipeline_step_details_step_configuration_details_command_line_arguments" {
  default = "commandLineArguments"
}

variable "pipeline_step_details_step_configuration_details_environment_variables" {
  default ={
        "environmentVariables": "environmentVariables"
  }
}

variable "pipeline_step_details_step_configuration_details_maximum_runtime_in_minutes" {
  default = 10
}

variable "pipeline_step_details_step_infrastructure_configuration_details_block_storage_size_in_gbs" {
  default = 50
}

variable "pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_memory_in_gbs" {
  default = 1.0
}

variable "pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_ocpus" {
  default = 1.0
}

variable "pipeline_step_details_step_name" {
  default = "stepName"
}

variable "pipeline_step_details_step_type" {
  default = "CUSTOM_SCRIPT"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_datascience_project" "pipeline" {
  compartment_id = var.compartment_ocid
}

resource "oci_logging_log_group" "pipeline" {
  compartment_id = var.compartment_ocid
  display_name   = "pipelines"
}


resource "oci_datascience_pipeline" "test_pipeline" {
  #Required
  compartment_id = var.compartment_ocid
  project_id     = oci_datascience_project.pipeline.id
  delete_related_pipeline_runs = true
  step_details {
    #Required
    step_name = var.pipeline_step_details_step_name
    step_type = var.pipeline_step_details_step_type

    #Optional
    depends_on           = var.pipeline_step_details_depends_on
    description          = var.pipeline_step_details_description
    is_artifact_uploaded = var.pipeline_step_details_is_artifact_uploaded
    # No Job Id for CUSTOM_SCRIPT
    # job_id               = oci_datascience_job.test_job.id
    step_configuration_details {

      #Optional
      command_line_arguments     = var.pipeline_step_details_step_configuration_details_command_line_arguments
      environment_variables      = var.pipeline_step_details_step_configuration_details_environment_variables
      maximum_runtime_in_minutes = var.pipeline_step_details_step_configuration_details_maximum_runtime_in_minutes
    }
    step_infrastructure_configuration_details {

      #Optional
      block_storage_size_in_gbs = var.pipeline_step_details_step_infrastructure_configuration_details_block_storage_size_in_gbs
    #   shape_config_details {

    #     #Optional
    #     memory_in_gbs = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_memory_in_gbs
    #     ocpus         = var.pipeline_step_details_step_infrastructure_configuration_details_shape_config_details_ocpus
    #   }
      shape_name = "VM.Standard2.1"
    }
  }

  #Optional in ML_JOB step type, mandatory for CUSTOM_SCRIPT step type
  step_artifact {
      step_name                    = "stepName"
      pipeline_step_artifact       = "${path.root}/pipeline-artifact.py"
      artifact_content_length      = 1380
      artifact_content_disposition = "attachment; filename=pipeline_artifact.py"
	}

  #Optional
  configuration_details {
    #Required
    type = var.pipeline_configuration_details_type

    #Optional
    command_line_arguments     = var.pipeline_configuration_details_command_line_arguments
    environment_variables      = var.pipeline_configuration_details_environment_variables
    maximum_runtime_in_minutes = var.pipeline_configuration_details_maximum_runtime_in_minutes
  }
#   defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.pipeline_defined_tags_value)
  description   = var.pipeline_description
  display_name  = var.pipeline_display_name
#   freeform_tags = var.pipeline_freeform_tags
  infrastructure_configuration_details {
    #Required
    block_storage_size_in_gbs = var.pipeline_infrastructure_configuration_details_block_storage_size_in_gbs
    shape_name                = "VM.Standard2.1"
    #Optional ONLY required if the shape is a flex shape
    # shape_config_details {

    #   #Optional
    #   memory_in_gbs = var.pipeline_infrastructure_configuration_details_shape_config_details_memory_in_gbs
    #   ocpus         = var.pipeline_infrastructure_configuration_details_shape_config_details_ocpus
    # }
  }
  log_configuration_details {

    #Optional
    enable_auto_log_creation = var.pipeline_log_configuration_details_enable_auto_log_creation
    enable_logging           = var.pipeline_log_configuration_details_enable_logging
    log_group_id             = oci_logging_log_group.pipeline.id
    # log_id                   = oci_logging_log.test_log.id
  }
}

resource "oci_logging_log_group" "pipeline_run" {
  compartment_id = var.compartment_ocid
  display_name   = "pipeline_run"
}

resource "oci_datascience_pipeline_run" "test_pipeline_run" {
  #Required
  compartment_id = var.compartment_ocid
  pipeline_id    = oci_datascience_pipeline.test_pipeline.id
  delete_related_job_runs = true

  #Optional
  configuration_override_details {
    #Required
    type = "DEFAULT"

    #Optional
    command_line_arguments     = "CommandLineArgumentsOverride"
    environment_variables      = {"environmentVariablesOverride":"environmentVariablesOverride"}
    maximum_runtime_in_minutes = 30
  }
  display_name  = "DisplayName1"
  log_configuration_override_details {

    #Optional
    enable_auto_log_creation = true
    enable_logging           = true
    log_group_id             = oci_logging_log_group.pipeline_run.id
    # log_id                   = oci_logging_log.test_log.id
  }
  project_id = oci_datascience_project.pipeline.id
  step_override_details {
    #Required
    step_configuration_details {

      #Optional
     command_line_arguments     = "CommandLineArgumentsStepOverride"
     environment_variables      = {"environmentVariablesStepOverride":"environmentVariablesStepOverride"}
     maximum_runtime_in_minutes = 30
    }
    step_name = "stepName"
  }
}