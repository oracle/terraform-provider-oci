// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "account_mgmt_info_account_mgmt_info_count" {
  default = 10
}

variable "account_mgmt_info_account_mgmt_info_filter" {
  default = ""
}

variable "account_mgmt_info_attribute_sets" {
  default = ["all"]
}

variable "account_mgmt_info_attributes" {
  default = "attributes"
}

variable "account_mgmt_info_authorization" {
  default = "authorization"
}

variable "account_mgmt_info_idcs_endpoint" {
  default = "idcsEndpoint"
}

# use the latest if not provided
variable "account_mgmt_info_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "account_mgmt_info_start_index" {
  default = 10
}

data "oci_identity_domains_account_mgmt_infos" "test_account_mgmt_infos" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  account_mgmt_info_count      = var.account_mgmt_info_account_mgmt_info_count
  account_mgmt_info_filter     = var.account_mgmt_info_account_mgmt_info_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.account_mgmt_info_authorization
  # resource_type_schema_version = var.account_mgmt_info_resource_type_schema_version
  start_index                  = var.account_mgmt_info_start_index
}
