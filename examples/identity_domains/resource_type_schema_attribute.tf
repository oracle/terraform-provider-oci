// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "resource_type_schema_attribute_resource_type_schema_attribute_count" {
  default = 10
}

variable "resource_type_schema_attribute_resource_type_schema_attribute_filter" {
  default = "resourcetype eq \"User\""
}

variable "resource_type_schema_attribute_authorization" {
  default = "authorization"
}

variable "resource_type_schema_attribute_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "resource_type_schema_attribute_start_index" {
  default = 1
}

data "oci_identity_domains_resource_type_schema_attributes" "test_resource_type_schema_attributes" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  resource_type_schema_attribute_count                  = var.resource_type_schema_attribute_resource_type_schema_attribute_count
  resource_type_schema_attribute_filter                 = var.resource_type_schema_attribute_resource_type_schema_attribute_filter
  attribute_sets                                        = ["all"]
  attributes                                            = ""
  authorization                                         = var.resource_type_schema_attribute_authorization
  start_index                                           = var.resource_type_schema_attribute_start_index
  #use the latest version if not provided
  #resource_type_schema_version = var.resource_type_schema_attribute_resource_type_schema_version
}

