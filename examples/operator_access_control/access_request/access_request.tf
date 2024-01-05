// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "access_request_resource_name" {
  default = "resourceName"
}

variable "access_request_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "access_request_state" {
  default = "APPROVED"
}

variable "access_request_time_end" {
  default = "2020-01-02T15:04:05Z"
}

variable "access_request_time_start" {
  default = "2006-01-02T15:04:05Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_operator_access_control_access_requests" "test_access_requests" {
  #Required
  compartment_id = var.compartment_id
  resource_type = var.access_request_resource_type


  #Optional
  #resource_name = var.access_request_resource_name
  #resource_type = var.access_request_resource_type
  #state         = var.access_request_state
  #time_end      = var.access_request_time_end
  #time_start    = var.access_request_time_start
}

