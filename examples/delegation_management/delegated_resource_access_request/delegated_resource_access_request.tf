// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "test_delegation_control_id" {}

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

data "oci_delegation_management_delegated_resource_access_requests" "test_delegated_resource_access_requests" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  delegation_control_id = var.test_delegation_control_id
  request_status        = var.delegated_resource_access_request_request_status
  #resource_id           = oci_usage_proxy_resource.test_resource.id
  state                 = var.delegated_resource_access_request_state
  #time_end              = var.delegated_resource_access_request_time_end
  #time_start            = var.delegated_resource_access_request_time_start
}