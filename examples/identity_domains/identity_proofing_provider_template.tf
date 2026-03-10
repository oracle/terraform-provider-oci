// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "identity_proofing_provider_template_identity_proofing_provider_template_count" {
  default = 10
}

variable "identity_proofing_provider_template_identity_proofing_provider_template_filter" {
  default = ""
}

variable "identity_proofing_provider_template_identity_proofing_provider_template_provider" {
  default = "identityProofingProviderTemplateProvider"
}

variable "identity_proofing_provider_template_authorization" {
  default = "authorization"
}

variable "identity_proofing_provider_template_compartment_ocid" {
  default = "compartmentOcid"
}

variable "identity_proofing_provider_template_configuration_name" {
  default = "name"
}

variable "identity_proofing_provider_template_configuration_sensitivity" {
  default = false
}

variable "identity_proofing_provider_template_configuration_type" {
  default = "type"
}

variable "identity_proofing_provider_template_delete_in_progress" {
  default = false
}

variable "identity_proofing_provider_template_domain_ocid" {
  default = "domainOcid"
}

variable "identity_proofing_provider_template_id" {
  default = "id"
}

variable "identity_proofing_provider_template_idcs_created_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_template_idcs_created_by_display" {
  default = "display"
}

variable "identity_proofing_provider_template_idcs_created_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_template_idcs_created_by_type" {
  default = "User"
}

variable "identity_proofing_provider_template_idcs_created_by_value" {
  default = "value"
}

variable "identity_proofing_provider_template_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "identity_proofing_provider_template_idcs_last_modified_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_template_idcs_last_modified_by_display" {
  default = "display"
}

variable "identity_proofing_provider_template_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_template_idcs_last_modified_by_type" {
  default = "User"
}

variable "identity_proofing_provider_template_idcs_last_modified_by_value" {
  default = "value"
}

variable "identity_proofing_provider_template_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "identity_proofing_provider_template_idcs_locked_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_template_idcs_locked_by_display" {
  default = "display"
}

variable "identity_proofing_provider_template_idcs_locked_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_template_idcs_locked_by_type" {
  default = "User"
}

variable "identity_proofing_provider_template_idcs_locked_by_value" {
  default = "value"
}

variable "identity_proofing_provider_template_idcs_locked_on" {
  default = "idcsLockedOn"
}

variable "identity_proofing_provider_template_idcs_locked_operations" {
  default = []
}

variable "identity_proofing_provider_template_idcs_prevented_operations" {
  default = []
}

variable "identity_proofing_provider_template_meta_created" {
  default = "created"
}

variable "identity_proofing_provider_template_meta_last_modified" {
  default = "lastModified"
}

variable "identity_proofing_provider_template_meta_location" {
  default = "location"
}

variable "identity_proofing_provider_template_meta_resource_type" {
  default = "resourceType"
}

variable "identity_proofing_provider_template_meta_version" {
  default = "version"
}

variable "identity_proofing_provider_template_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_template_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "identity_proofing_provider_template_schemas" {
  default = []
}

variable "identity_proofing_provider_template_service_type" {
  default = []
}

variable "identity_proofing_provider_template_start_index" {
  default = 1
}

variable "identity_proofing_provider_template_tags_key" {
  default = "key"
}

variable "identity_proofing_provider_template_tags_value" {
  default = "value"
}

variable "identity_proofing_provider_template_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "identity_proofing_provider_template_verification_url" {
  default = "verificationUrl"
}

resource "oci_identity_domains_identity_proofing_provider_template" "test_identity_proofing_provider_template" {
  #Required
  identity_proofing_provider_template_provider = var.identity_proofing_provider_template_identity_proofing_provider_template_provider
  idcs_endpoint                                = data.oci_identity_domain.test_domain.url
  schemas                                      = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityProofingProviderTemplate"]
  service_type                                 = ["serviceType"]
  verification_url                             = var.identity_proofing_provider_template_verification_url

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.identity_proofing_provider_template_authorization
  configuration {
    #Required
    name        = var.identity_proofing_provider_template_configuration_name
    sensitivity = var.identity_proofing_provider_template_configuration_sensitivity
    type        = var.identity_proofing_provider_template_configuration_type
  }
}

data "oci_identity_domains_identity_proofing_provider_templates" "test_identity_proofing_provider_templates" {
  #Required
  idcs_endpoint                                = data.oci_identity_domain.test_domain.url

  #Optional
  identity_proofing_provider_template_count  = var.identity_proofing_provider_template_identity_proofing_provider_template_count
  identity_proofing_provider_template_filter = var.identity_proofing_provider_template_identity_proofing_provider_template_filter
  attribute_sets = ["all"]
  attributes     = ""
  authorization                              = var.identity_proofing_provider_template_authorization
  # resource_type_schema_version               = var.identity_proofing_provider_template_resource_type_schema_version
  start_index                                = var.identity_proofing_provider_template_start_index
}

