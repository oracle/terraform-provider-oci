// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_pending_approval_my_pending_approval_count" {
  default = 10
}

variable "my_pending_approval_my_pending_approval_filter" {
  default = ""
}

variable "my_pending_approval_authorization" {
  default = "authorization"
}

variable "my_pending_approval_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "my_pending_approval_start_index" {
  default = 1
}

data "oci_identity_domains_my_pending_approvals" "test_my_pending_approvals" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_pending_approval_count    = var.my_pending_approval_my_pending_approval_count
  my_pending_approval_filter   = var.my_pending_approval_my_pending_approval_filter
  authorization                = var.my_pending_approval_authorization
#  resource_type_schema_version = var.my_pending_approval_resource_type_schema_version
  start_index                  = var.my_pending_approval_start_index
}

