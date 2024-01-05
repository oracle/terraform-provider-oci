// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "account_recovery_setting_attribute_sets" {
  default = ["all"]
}

variable "account_recovery_setting_attributes" {
  default = ""
}

variable "account_recovery_setting_authorization" {
  default = "authorization"
}

variable "account_recovery_setting_compartment_ocid" {
  default = "compartmentOcid"
}

variable "account_recovery_setting_delete_in_progress" {
  default = false
}

variable "account_recovery_setting_domain_ocid" {
  default = "domainOcid"
}

variable "account_recovery_setting_factors" {
  default = ["email"]
}

variable "account_recovery_setting_id" {
  default = "AccountRecoverySettings"
}

variable "account_recovery_setting_idcs_created_by_display" {
  default = "display"
}

variable "account_recovery_setting_idcs_created_by_ocid" {
  default = "ocid"
}

variable "account_recovery_setting_idcs_created_by_ref" {
  default = "ref"
}

variable "account_recovery_setting_idcs_created_by_type" {
  default = "User"
}

variable "account_recovery_setting_idcs_created_by_value" {
  default = "value"
}

variable "account_recovery_setting_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "account_recovery_setting_idcs_last_modified_by_display" {
  default = "display"
}

variable "account_recovery_setting_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "account_recovery_setting_idcs_last_modified_by_ref" {
  default = "ref"
}

variable "account_recovery_setting_idcs_last_modified_by_type" {
  default = "User"
}

variable "account_recovery_setting_idcs_last_modified_by_value" {
  default = "value"
}

variable "account_recovery_setting_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "account_recovery_setting_idcs_prevented_operations" {
  default = []
}

variable "account_recovery_setting_lockout_duration" {
  default = 10
}

variable "account_recovery_setting_max_incorrect_attempts" {
  default = 10
}

variable "account_recovery_setting_meta_created" {
  default = "created"
}

variable "account_recovery_setting_meta_last_modified" {
  default = "lastModified"
}

variable "account_recovery_setting_meta_location" {
  default = "location"
}

variable "account_recovery_setting_meta_resource_type" {
  default = "resourceType"
}

variable "account_recovery_setting_meta_version" {
  default = "version"
}

# use the latest if not provided
variable "account_recovery_setting_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "account_recovery_setting_schemas" {
  default = ["urn:ietf:params:scim:schemas:oracle:idcs:AccountRecoverySettings"]
}

variable "account_recovery_setting_tags_key" {
  default = "key"
}

variable "account_recovery_setting_tags_value" {
  default = "value"
}

variable "account_recovery_setting_tenancy_ocid" {
  default = "tenancyOcid"
}


resource "oci_identity_domains_account_recovery_setting" "test_account_recovery_setting" {
  #Required
  account_recovery_setting_id = var.account_recovery_setting_id
  factors                     = var.account_recovery_setting_factors
  idcs_endpoint               = data.oci_identity_domain.test_domain.url

  lockout_duration       = var.account_recovery_setting_lockout_duration
  max_incorrect_attempts = var.account_recovery_setting_max_incorrect_attempts
  schemas                = ["urn:ietf:params:scim:schemas:oracle:idcs:AccountRecoverySettings"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.account_recovery_setting_authorization
  external_id    = "externalId"
  # resource_type_schema_version = var.account_recovery_setting_resource_type_schema_version
  tags {
    #Required
    key   = var.account_recovery_setting_tags_key
    value = var.account_recovery_setting_tags_value
  }
}

data "oci_identity_domains_account_recovery_settings" "test_account_recovery_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.account_recovery_setting_authorization
  # resource_type_schema_version = var.account_recovery_setting_resource_type_schema_version
}
