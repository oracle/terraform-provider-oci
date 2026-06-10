// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "delegated_resource_access_request_request_status" {
  default = "CREATED"
}

variable "delegated_resource_access_request_state" {
  default = "SUCCEEDED"
}

variable "delegated_resource_access_request_time_end" {
  default = "timeEnd"
}

variable "delegated_resource_access_request_time_start" {
  default = "timeStart"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region  
}

data "oci_delegate_access_control_delegated_resource_access_requests" "test_delegated_resource_access_requests" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  request_status        = var.delegated_resource_access_request_request_status
  state                 = var.delegated_resource_access_request_state
}

