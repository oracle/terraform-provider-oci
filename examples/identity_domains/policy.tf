// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "policy_policy_count" {
  default = 10
}

variable "policy_policy_filter" {
  default = ""
}

variable "policy_active" {
  default = false
}

variable "policy_authorization" {
  default = "authorization"
}

variable "policy_description" {
  default = "description"
}

variable "policy_name" {
  default = "name"
}

variable "policy_policy_groovy" {
  default = "policyGroovy"
}

variable "policy_rules_name" {
  default = "name"
}

variable "policy_rules_sequence" {
  default = 10
}

variable "policy_start_index" {
  default = 1
}

variable "policy_tags_key" {
  default = "key"
}

variable "policy_tags_value" {
  default = "value"
}

resource "oci_identity_domains_rule" "test_policy_rule" {
  #Required
  condition     = "condition"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = "name"
  policy_type {
    #Required
    value = "IdentityProvider"
  }
  return {
    #Required
    name  = "LocalIDPs"
    value = "[\"UserNamePassword\"]"
  }
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Rule"]
}

resource "oci_identity_domains_policy" "test_policy" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = var.policy_name
  policy_type {
    #Required
    value = "IdentityProvider"
  }
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Policy"]

  #Optional
  active                       = var.policy_active
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.policy_authorization
  description                  = var.policy_description
  external_id                  = "externalId"
  policy_groovy                = var.policy_policy_groovy
  #use the latest if not provided
  # resource_type_schema_version = var.policy_resource_type_schema_version
  rules {
    #Required
    sequence = var.policy_rules_sequence
    value    = oci_identity_domains_rule.test_policy_rule.id
  }
  tags {
    #Required
    key   = var.policy_tags_key
    value = var.policy_tags_value
  }
}

data "oci_identity_domains_policies" "test_policies" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  policy_count                 = var.policy_policy_count
  policy_filter                = var.policy_policy_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.policy_authorization
  #use the latest if not provided
  # resource_type_schema_version = var.policy_resource_type_schema_version
  start_index                  = var.policy_start_index
}

