// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "job_environment_configuration_details_job_environment_type" {
  default = "OCIR_CONTAINER"
}

variable "job_environment_configuration_details_cmd" {
}

variable "job_environment_configuration_details_entrypoint" {
}

variable "job_environment_configuration_details_image" {
}

variable "job_environment_configuration_details_image_digest" {
}

variable "job_environment_configuration_details_image_signature_id" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_logging_log_group" "job" {
  compartment_id = var.compartment_ocid
  display_name   = "jobs"
}

resource "oci_datascience_project" "job" {
  compartment_id = var.compartment_ocid
}

#resource "oci_core_subnet" "tf_subnet" {
#  cidr_block     = "10.0.1.0/24"
#  compartment_id = var.compartment_ocid
#  vcn_id         = oci_core_vcn.tf_vcn.id
#}

#resource "oci_core_vcn" "tf_vcn" {
#  cidr_block     = "10.0.0.0/16"
#  compartment_id = var.compartment_ocid
#}

resource "oci_datascience_job" "job" {
  compartment_id               = var.compartment_ocid
  project_id                   = oci_datascience_project.job.id
  delete_related_job_runs      = true
  job_run_retention_period_in_days = 10
  # job_storage_mount_configuration_details_list {
  #   storage_type               = "OBJECT_STORAGE"
  #   namespace                  = "ociodscdev"
  #   bucket                     = "Jobs-Dev-tools-Script-Bkt"
  #   destination_directory_name = "oss1"
  #   destination_path           = "/"
  # }

  # job_storage_mount_configuration_details_list {
  #   storage_type               = "OBJECT_STORAGE"
  #   namespace                  = "ociodscdev"
  #   bucket                     = "Jobs-Dev-tools-Script-Bkt"
  #   destination_directory_name = "oss2"
  #   prefix                     = "dir"
  # }
  # optional parameter
  # opc_parent_rpt_url = ""

   job_configuration_details {
     job_type                   = "DEFAULT"
     maximum_runtime_in_minutes = 30
   }

  job_environment_configuration_details {
    job_environment_type = var.job_environment_configuration_details_job_environment_type
    image                = var.job_environment_configuration_details_image
  }

  job_infrastructure_configuration_details {
    job_infrastructure_type = "MANAGED_COMPUTE_CLUSTER"
    compute_target_id       = "compute_target_id"
    network_configuration {
      network_type = "CUSTOM_NETWORK"
      subnet_id    = "subnet_id"
    }
    resource_configuration {
      resource_request_configuration {
        ocpus         = 1
        memory_in_gbs = 2
      }
    }
  }

  # New Optional parameter for Multi Node
#  job_node_configuration_details {
#    job_node_type =                            "MULTI_NODE"
#    job_network_configuration  {
#      job_network_type = "CUSTOM_NETWORK"
#      subnet_id=        "<subnet_id>"
#    }
#    job_node_group_configuration_details_list {
#      name=                      "replica1"
#      job_configuration_details {
#        job_type=               "DEFAULT"
#        command_line_arguments=  "commandLineArguments"
#      }
#      job_infrastructure_configuration_details {
#        job_infrastructure_type   = "MULTI_NODE"
#        shape_name                = "VM.Standard3.Flex"
#        block_storage_size_in_gbs = 50
#        # subnet_id = "<subnet_id>"
#
#        # Optional
#        job_shape_config_details {
#          memory_in_gbs = 16
#          ocpus         = 3
#          cpu_baseline = "BASELINE_1_2"
#        }
#      }
#      minimum_success_replicas=1
#      replicas=1
#    }
#    startup_order = "IN_ORDER"
#  }

  #  job_environment_configuration_details {
  #		cmd                  = var.job_environment_configuration_details_cmd
  #		entrypoint           = var.job_environment_configuration_details_entrypoint
  #		image                = var.job_environment_configuration_details_image
  #		image_digest         = var.job_environment_configuration_details_image_digest
  #    image_signature_id   = var.job_environment_configuration_details_image_signature_id
  #		job_environment_type = var.job_environment_configuration_details_job_environment_type
  #  }

  #  job_log_configuration_details {
  #    enable_logging           = true
  #    enable_auto_log_creation = true
  #    log_group_id             = oci_logging_log_group.job.id
  #  }
}

data "oci_datascience_job" "read_job" {
  job_id = oci_datascience_job.job.id
}

output "job_run_retention_period_in_days" {
  value = data.oci_datascience_job.read_job.job_run_retention_period_in_days
}

data "oci_datascience_jobs" "by_compartment" {
  #Required
  compartment_id = var.compartment_ocid

  # #Optional
  # created_by   = var.job_created_by
  # display_name = var.job_display_name
  # id           = var.job_id
  # project_id   = oci_datascience_project.test_project.id
  # state        = var.job_state
}

# terraform will wait until the JobRun reaches completion.
resource "oci_datascience_job_run" "sync" {
  compartment_id = var.compartment_ocid
  project_id     = oci_datascience_project.job.id
  job_id         = oci_datascience_job.job.id
  asynchronous   = false
}

# terraform will NOT wait for completion before returning.
# it will move on once the JobRun has reached "In Progress" status.
resource "oci_datascience_job_run" "async" {
  compartment_id = var.compartment_ocid
  project_id     = oci_datascience_project.job.id
  job_id         = oci_datascience_job.job.id
  asynchronous   = true
}

data "oci_datascience_job_runs" "by_compartment" {
  #Required
  compartment_id = var.compartment_ocid

  #   #Optional
  #   created_by   = var.job_run_created_by
  #   display_name = var.job_run_display_name
  #   id           = var.job_run_id
  #   job_id       = oci_datascience_job.job.id
  #   state        = var.job_run_state
}
