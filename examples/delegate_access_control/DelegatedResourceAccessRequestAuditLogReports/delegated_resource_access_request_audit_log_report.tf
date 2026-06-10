// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

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

data "oci_delegate_access_control_delegated_resource_access_request_audit_log_reports" "test_delegated_resource_access_request_audit_log_reports" {
  #Required
  delegated_resource_access_request_id = oci_delegate_access_control_delegated_resource_access_request.test_delegated_resource_access_request.id

  #Optional
  is_process_tree_enabled = var.delegated_resource_access_request_audit_log_report_is_process_tree_enabled
}

