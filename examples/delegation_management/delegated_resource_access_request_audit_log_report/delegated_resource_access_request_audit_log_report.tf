// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "test_access_req_id" {}

variable "delegated_resource_access_request_audit_log_report_is_process_tree_enabled" {
  default = false
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_delegation_management_delegated_resource_access_request_audit_log_report" "test_delegated_resource_access_request_audit_log_report" {
  #Required
  delegated_resource_access_request_id = var.test_access_req_id

  #Optional
  is_process_tree_enabled = var.delegated_resource_access_request_audit_log_report_is_process_tree_enabled
}