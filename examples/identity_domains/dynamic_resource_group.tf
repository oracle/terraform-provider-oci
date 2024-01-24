// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "dynamic_resource_group_dynamic_resource_group_count" {
  default = 10
}

variable "dynamic_resource_group_dynamic_resource_group_filter" {
  default = ""
}

variable "dynamic_resource_group_authorization" {
  default = "authorization"
}

variable "dynamic_resource_group_description" {
  default = "description"
}

variable "dynamic_resource_group_display_name" {
  default = "displayName"
}

variable "dynamic_resource_group_matching_rule" {
  default = "Any {Any {instance.id = \"instance.id\", instance.compartment.id = \"instance.compartment.id\"}}"
}

variable "dynamic_resource_group_start_index" {
  default = 1
}

variable "dynamic_resource_group_tags_key" {
  default = "key"
}

variable "dynamic_resource_group_tags_value" {
  default = "value"
}

variable "dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key" {
  default = "key"
}

variable "dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace" {
  default = "namespace"
}

variable "dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value" {
  default = "value"
}

variable "dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key" {
  default = "freeformKey"
}

variable "dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value" {
  default = "freeformValue"
}


resource "oci_identity_domains_dynamic_resource_group" "test_dynamic_resource_group" {
  #Required
  display_name  = var.dynamic_resource_group_display_name
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  matching_rule = var.dynamic_resource_group_matching_rule
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:DynamicResourceGroup"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.dynamic_resource_group_authorization
  description    = var.dynamic_resource_group_description
  #use the latest if not provided
  # resource_type_schema_version = var.dynamic_resource_group_resource_type_schema_version
  tags {
    #Required
    key   = var.dynamic_resource_group_tags_key
    value = var.dynamic_resource_group_tags_value
  }
  urnietfparamsscimschemasoracleidcsextension_oci_tags {

    #Optional
    /* #create tagNamespace to use defined tags
    defined_tags {
      #Required
      key       = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
      namespace = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
      value     = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
    }
    */
    freeform_tags {
      #Required
      key   = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
      value = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
    }
  }
  lifecycle {
    ignore_changes = [schemas]
  }
}

data "oci_identity_domains_dynamic_resource_groups" "test_dynamic_resource_groups" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  dynamic_resource_group_count  = var.dynamic_resource_group_dynamic_resource_group_count
  dynamic_resource_group_filter = var.dynamic_resource_group_dynamic_resource_group_filter
  attribute_sets                = []
  attributes                    = ""
  authorization                 = var.dynamic_resource_group_authorization
  #use the latest if not provided
  # resource_type_schema_version  = var.dynamic_resource_group_resource_type_schema_version
  start_index = var.dynamic_resource_group_start_index
}

