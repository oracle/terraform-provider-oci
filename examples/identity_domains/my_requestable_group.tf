// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_requestable_group_my_requestable_group_count" {
  default = 10
}

variable "my_requestable_group_my_requestable_group_filter" {
  default = ""
}

variable "my_requestable_group_authorization" {
  default = "authorization"
}

variable "my_requestable_group_start_index" {
  default = 1
}

data "oci_identity_domains_my_requestable_groups" "test_my_requestable_groups" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_requestable_group_count   = var.my_requestable_group_my_requestable_group_count
  my_requestable_group_filter  = var.my_requestable_group_my_requestable_group_filter
  authorization                = var.my_requestable_group_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_requestable_group_resource_type_schema_version
  start_index                  = var.my_requestable_group_start_index
}

