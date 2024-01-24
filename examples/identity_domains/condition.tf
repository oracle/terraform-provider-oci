// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "condition_condition_count" {
  default = 10
}

variable "condition_condition_filter" {
  default = ""
}

variable "condition_attribute_name" {
  default = "attributeName"
}

variable "condition_attribute_sets" {
  default = []
}

variable "condition_attribute_value" {
  default = "attributeValue"
}

variable "condition_authorization" {
  default = "authorization"
}

variable "condition_description" {
  default = "description"
}

variable "condition_evaluate_condition_if" {
  default = "evaluateConditionIf"
}

variable "condition_name" {
  default = "name"
}

variable "condition_operator" {
  default = "eq"
}

variable "condition_resource_type_schema_version" {
  default = ""
}

variable "condition_start_index" {
  default = 1
}

variable "condition_tags_key" {
  default = "key"
}

variable "condition_tags_value" {
  default = "value"
}

resource "oci_identity_domains_condition" "test_condition" {
  #Required
  attribute_name  = var.condition_attribute_name
  attribute_value = var.condition_attribute_value
  idcs_endpoint   = data.oci_identity_domain.test_domain.url
  name            = var.condition_name
  operator        = var.condition_operator
  schemas         = ["urn:ietf:params:scim:schemas:oracle:idcs:Condition"]

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.condition_authorization
  description                  = var.condition_description
  evaluate_condition_if        = var.condition_evaluate_condition_if
  external_id                  = "externalId"
  #use the latest if not provided
  #resource_type_schema_version = var.condition_resource_type_schema_version
  tags {
    #Required
    key   = var.condition_tags_key
    value = var.condition_tags_value
  }
}

data "oci_identity_domains_conditions" "test_conditions" {
  #Required
  idcs_endpoint   = data.oci_identity_domain.test_domain.url

  #Optional
  condition_count              = var.condition_condition_count
  condition_filter             = var.condition_condition_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.condition_authorization
  #use the latest if not provided
  #resource_type_schema_version = var.condition_resource_type_schema_version
  start_index                  = var.condition_start_index
}

