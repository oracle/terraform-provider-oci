// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "rule_rule_count" {
  default = 10
}

variable "rule_rule_filter" {
  default = ""
}

variable "rule_active" {
  default = false
}

variable "rule_attribute_sets" {
  default = []
}

variable "rule_attributes" {
  default = "attributes"
}

variable "rule_authorization" {
  default = "authorization"
}

variable "rule_condition" {
  default = "condition"
}

variable "rule_condition_group_type" {
  default = "Condition"
}

variable "rule_description" {
  default = "description"
}

variable "rule_locked" {
  default = false
}

variable "rule_name" {
  default = "name"
}

variable "rule_policy_type_value" {
  default = "SignOn"
}

variable "rule_return_name" {
  default = "name"
}

variable "rule_return_return_groovy" {
  default = "returnGroovy"
}

variable "rule_return_value" {
  default = "value"
}

variable "rule_rule_groovy" {
  default = "ruleGroovy"
}

variable "rule_start_index" {
  default = 1
}

variable "rule_tags_key" {
  default = "key"
}

variable "rule_tags_value" {
  default = "value"
}

resource "oci_identity_domains_rule" "test_rule" {
  #Required
  condition     = var.rule_condition
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = var.rule_name
  policy_type {
    #Required
    value = var.rule_policy_type_value
  }
  return {
    #Required
    name  = var.rule_return_name
    value = var.rule_return_value

    #Optional
    return_groovy = var.rule_return_return_groovy
  }
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Rule"]

  #Optional
  active         = var.rule_active
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.rule_authorization
  condition_group {
    #Required
    type = var.rule_condition_group_type

    #Optional
    value = oci_identity_domains_condition.test_rule_condition.id
  }
  description                  = var.rule_description
  external_id                  = "externalId"
  locked                       = var.rule_locked
#  resource_type_schema_version = var.rule_resource_type_schema_version
  rule_groovy                  = var.rule_rule_groovy
  tags {
    #Required
    key   = var.rule_tags_key
    value = var.rule_tags_value
  }
}

data "oci_identity_domains_rules" "test_rules" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  rule_count                   = var.rule_rule_count
  rule_filter                  = var.rule_rule_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.rule_authorization
#  resource_type_schema_version = var.rule_resource_type_schema_version
  start_index                  = var.rule_start_index
}

resource "oci_identity_domains_condition" "test_rule_condition" {
  #Required
  idcs_endpoint   = data.oci_identity_domain.test_domain.url
  name            = "name"
  operator        = "eq"
  schemas         = ["urn:ietf:params:scim:schemas:oracle:idcs:Condition"]
  attribute_value = "attributeValue"
  attribute_name  = "attributeName"

  #Optional
  description                  = "description"
  evaluate_condition_if        = "evaluateConditionIf"
  external_id                  = "externalId"
}

