// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_core_vcn" "job" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "dsmljobs"
  dns_label      = "dsmljobs"
}

resource "oci_core_nat_gateway" "job" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.job.id
  display_name   = "job"
}

resource "oci_core_subnet" "regional_with_natgw" {
  cidr_block        = "10.0.1.0/24"
  display_name      = "regional_with_natgw"
  dns_label         = "regwithnatgw"
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.job.id
  security_list_ids = [oci_core_security_list.regional_with_natgw.id]
  route_table_id    = oci_core_route_table.regional_with_natgw.id
}

resource "oci_core_route_table" "regional_with_natgw" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.job.id
  display_name   = "regional_with_natgw"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_nat_gateway.job.id
  }
}

resource "oci_core_security_list" "regional_with_natgw" {
  compartment_id = var.compartment_ocid
  display_name   = "regional_with_natgw"
  vcn_id         = oci_core_vcn.job.id

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
    stateless = false
  }
}

resource "oci_logging_log_group" "job" {
  compartment_id = var.compartment_ocid
  display_name = "jobs"
}

resource "oci_datascience_project" "job" {
  compartment_id = var.compartment_ocid
}

resource "oci_datascience_job" "job" {
  compartment_id = var.compartment_ocid
  project_id = oci_datascience_project.job.id
  job_artifact = "${path.root}/job-artifact.py"
  artifact_content_length = 1380
  artifact_content_disposition = "attachment; filename=job_artifact.py"
  delete_related_job_runs = true

  job_configuration_details {
    job_type = "DEFAULT"
    maximum_runtime_in_minutes = 30
  }

  job_infrastructure_configuration_details {
    job_infrastructure_type = "STANDALONE"
    shape_name     = "VM.Standard2.2"
    subnet_id      = oci_core_subnet.regional_with_natgw.id
    block_storage_size_in_gbs = 66
  }

  job_log_configuration_details {
    enable_logging = true
    enable_auto_log_creation = true
    log_group_id = oci_logging_log_group.job.id
  }
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
  project_id = oci_datascience_project.job.id
  job_id = oci_datascience_project.job.id
  asynchronous = false
}

# terraform will NOT wait for completion before returning.
# it will move on once the JobRun has reached "In Progress" status.
resource "oci_datascience_job_run" "async" {
  compartment_id = var.compartment_ocid
  project_id = oci_datascience_project.job.id
  job_id = oci_datascience_project.job.id
  asynchronous = true
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
# }
