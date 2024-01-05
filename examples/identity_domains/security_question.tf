// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "security_question_security_question_count" {
  default = 1
}

variable "security_question_security_question_filter" {
  default = ""
}

variable "security_question_active" {
  default = false
}

variable "security_question_attribute_sets" {
  default = ["all"]
}

variable "security_question_attributes" {
  default = "attributes"
}

variable "security_question_authorization" {
  default = "authorization"
}

variable "security_question_compartment_ocid" {
  default = "compartmentOcid"
}

variable "security_question_delete_in_progress" {
  default = false
}

variable "security_question_domain_ocid" {
  default = "domainOcid"
}

variable "security_question_id" {
  default = "SecurityQuestions"
}

variable "security_question_idcs_created_by_display" {
  default = "display"
}

variable "security_question_idcs_created_by_ocid" {
  default = "ocid"
}

variable "security_question_idcs_created_by_ref" {
  default = "ref"
}

variable "security_question_idcs_created_by_type" {
  default = "User"
}

variable "security_question_idcs_created_by_value" {
  default = "value"
}

variable "security_question_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "security_question_idcs_last_modified_by_display" {
  default = "display"
}

variable "security_question_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "security_question_idcs_last_modified_by_ref" {
  default = "ref"
}

variable "security_question_idcs_last_modified_by_type" {
  default = "User"
}

variable "security_question_idcs_last_modified_by_value" {
  default = "value"
}

variable "security_question_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "security_question_idcs_prevented_operations" {
  default = []
}

variable "security_question_meta_created" {
  default = "created"
}

variable "security_question_meta_last_modified" {
  default = "lastModified"
}

variable "security_question_meta_location" {
  default = "location"
}

variable "security_question_meta_resource_type" {
  default = "resourceType"
}

variable "security_question_meta_version" {
  default = "version"
}

variable "security_question_ocid" {
  default = "ocid"
}

variable "security_question_question_text_default" {
  default = true
}

variable "security_question_question_text_locale" {
  default = "en"
}

variable "security_question_question_text_value" {
  default = "value"
}

# use the latest if not provided
variable "security_question_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "security_question_schemas" {
  default = []
}

variable "security_question_start_index" {
  default = 1
}

variable "security_question_tags_key" {
  default = "key"
}

variable "security_question_tags_value" {
  default = "value"
}

variable "security_question_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "security_question_type" {
  default = "custom"
}

resource "oci_identity_domains_security_question" "test_security_question" {
  #Required
  active        = var.security_question_active
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  /* One and only one "question_text" needs to have "default" set to true */
  question_text {
    #Required
    locale = var.security_question_question_text_locale
    value  = var.security_question_question_text_value

    #Optional
    default = var.security_question_question_text_default
  }
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:SecurityQuestion"]
  type    = var.security_question_type

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.security_question_authorization
  external_id    = "externalId"
  #resource_type_schema_version = var.security_question_resource_type_schema_version
  tags {
    #Required
    key   = var.security_question_tags_key
    value = var.security_question_tags_value
  }
}

data "oci_identity_domains_security_questions" "test_security_questions" {

  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  security_question_count  = var.security_question_security_question_count
  security_question_filter = var.security_question_security_question_filter
  attribute_sets           = ["all"]
  attributes               = ""
  authorization            = var.security_question_authorization
  #resource_type_schema_version = var.security_question_resource_type_schema_version
  start_index = var.security_question_start_index
}
