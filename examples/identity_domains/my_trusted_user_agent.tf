// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_trusted_user_agent_my_trusted_user_agent_count" {
  default = 10
}

variable "my_trusted_user_agent_my_trusted_user_agent_filter" {
  default = ""
}

variable "my_trusted_user_agent_authorization" {
  default = "authorization"
}

variable "my_trusted_user_agent_start_index" {
  default = 1
}


data "oci_identity_domains_my_trusted_user_agents" "test_my_trusted_user_agents" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_trusted_user_agent_count  = var.my_trusted_user_agent_my_trusted_user_agent_count
  my_trusted_user_agent_filter = var.my_trusted_user_agent_my_trusted_user_agent_filter
  attribute_sets               = []
  attributes                   = ""
  authorization                = var.my_trusted_user_agent_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.my_trusted_user_agent_resource_type_schema_version
  start_index = var.my_trusted_user_agent_start_index
}

