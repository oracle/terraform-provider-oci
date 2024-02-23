// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {  
  default = "<compartment.ocid>"
}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "executions_statuses_start_time" {
   default = "2024-01-06T11:34:00.000Z"
}

variable "executions_statuses_end_time" {
   default = "2024-01-06T12:34:00.000Z"
}

data "oci_database_management_job_executions_statuses" "test_job_executions_status" {
  #Required
  compartment_id = var.compartment_id
  #start_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timeadd(timestamp(), "-12h")) 
  #end_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timestamp())
  start_time = var.executions_statuses_start_time 
  end_time = var.executions_statuses_end_time

  #Optional
  managed_database_id = var.managed_database_id
}
