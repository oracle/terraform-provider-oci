// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_group_my_group_count" {
  default = 10
}

variable "my_group_my_group_filter" {
  default = ""
}

variable "my_group_authorization" {
  default = "authorization"
}

variable "my_group_start_index" {
  default = 1
}


data "oci_identity_domains_my_groups" "test_my_groups" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_group_count  = var.my_group_my_group_count
  my_group_filter = var.my_group_my_group_filter
  attribute_sets  = []
  attributes      = ""
  authorization   = var.my_group_authorization
  start_index     = var.my_group_start_index
}
