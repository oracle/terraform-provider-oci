// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "identity_proofing_provider_identity_proofing_provider_count" {
  default = 10
}

variable "identity_proofing_provider_identity_proofing_provider_filter" {
  default = ""
}

variable "identity_proofing_provider_identity_proofing_provider_provider" {
  default = "identityProofingProviderProvider"
}

variable "identity_proofing_provider_authorization" {
  default = "authorization"
}

variable "identity_proofing_provider_claim_mapping_attr_match" {
  default = "attrMatch"
}

variable "identity_proofing_provider_claim_mapping_verifiable_claim" {
  default = "verifiableClaim"
}

variable "identity_proofing_provider_compartment_ocid" {
  default = "compartmentOcid"
}

variable "identity_proofing_provider_configuration_name" {
  default = "name"
}

variable "identity_proofing_provider_configuration_value" {
  default = "value"
}

variable "identity_proofing_provider_delete_in_progress" {
  default = false
}

variable "identity_proofing_provider_description" {
  default = "description"
}

variable "identity_proofing_provider_domain_ocid" {
  default = "domainOcid"
}

variable "identity_proofing_provider_id" {
  default = "id"
}

variable "identity_proofing_provider_idcs_created_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_idcs_created_by_display" {
  default = "display"
}

variable "identity_proofing_provider_idcs_created_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_idcs_created_by_type" {
  default = "User"
}

variable "identity_proofing_provider_idcs_created_by_value" {
  default = "value"
}

variable "identity_proofing_provider_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "identity_proofing_provider_idcs_last_modified_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_idcs_last_modified_by_display" {
  default = "display"
}

variable "identity_proofing_provider_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_idcs_last_modified_by_type" {
  default = "User"
}

variable "identity_proofing_provider_idcs_last_modified_by_value" {
  default = "value"
}

variable "identity_proofing_provider_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "identity_proofing_provider_idcs_locked_by__ref" {
  default = "ref"
}

variable "identity_proofing_provider_idcs_locked_by_display" {
  default = "display"
}

variable "identity_proofing_provider_idcs_locked_by_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_idcs_locked_by_type" {
  default = "User"
}

variable "identity_proofing_provider_idcs_locked_by_value" {
  default = "value"
}

variable "identity_proofing_provider_idcs_locked_on" {
  default = "idcsLockedOn"
}

variable "identity_proofing_provider_idcs_locked_operations" {
  default = []
}

variable "identity_proofing_provider_idcs_prevented_operations" {
  default = []
}

variable "identity_proofing_provider_meta_created" {
  default = "created"
}

variable "identity_proofing_provider_meta_last_modified" {
  default = "lastModified"
}

variable "identity_proofing_provider_meta_location" {
  default = "location"
}

variable "identity_proofing_provider_meta_resource_type" {
  default = "resourceType"
}

variable "identity_proofing_provider_meta_version" {
  default = "version"
}

variable "identity_proofing_provider_name" {
  default = "name"
}

variable "identity_proofing_provider_ocid" {
  default = "ocid"
}

variable "identity_proofing_provider_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "identity_proofing_provider_runtime_data_attr_name" {
  default = "attrName"
}

variable "identity_proofing_provider_runtime_data_attr_value" {
  default = "attrValue"
}

variable "identity_proofing_provider_schemas" {
  default = []
}

variable "identity_proofing_provider_start_index" {
  default = 1
}

variable "identity_proofing_provider_status" {
  default = "INACTIVE"
}

variable "identity_proofing_provider_tags_key" {
  default = "key"
}

variable "identity_proofing_provider_tags_value" {
  default = "value"
}

variable "identity_proofing_provider_tenancy_ocid" {
  default = "tenancyOcid"
}

# dependency identity_proofing_provider_template resource
resource "oci_identity_domains_identity_proofing_provider_template" "test_identity_proofing_provider_template_dependency" {
    #Required
    identity_proofing_provider_template_provider = var.identity_proofing_provider_identity_proofing_provider_provider
    idcs_endpoint                                = data.oci_identity_domain.test_domain.url
    schemas                                      = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityProofingProviderTemplate"]
    service_type                                 = ["serviceType"]
    verification_url                             = var.identity_proofing_provider_template_verification_url

}

resource "oci_identity_domains_identity_proofing_provider" "test_identity_proofing_provider" {
  #Required
  identity_proofing_provider_provider = oci_identity_domains_identity_proofing_provider_template.test_identity_proofing_provider_template_dependency.identity_proofing_provider_template_provider
  claim_mapping {
    #Required
    attr_match       = var.identity_proofing_provider_claim_mapping_attr_match
    verifiable_claim = var.identity_proofing_provider_claim_mapping_verifiable_claim
  }
  configuration {
    #Required
    name  = var.identity_proofing_provider_configuration_name
    value = var.identity_proofing_provider_configuration_value
  }
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  name          = var.identity_proofing_provider_name
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityProofingProvider"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.identity_proofing_provider_authorization
  description    = var.identity_proofing_provider_description
  runtime_data {
    #Required
    attr_name  = var.identity_proofing_provider_runtime_data_attr_name
    attr_value = var.identity_proofing_provider_runtime_data_attr_value
  }
  status = var.identity_proofing_provider_status
}

data "oci_identity_domains_identity_proofing_providers" "test_identity_proofing_providers" {
  #Required
  idcs_endpoint                                = data.oci_identity_domain.test_domain.url

  #Optional
  identity_proofing_provider_count  = var.identity_proofing_provider_identity_proofing_provider_count
  identity_proofing_provider_filter = var.identity_proofing_provider_identity_proofing_provider_filter
  attribute_sets = ["all"]
  attributes     = ""
  authorization                     = var.identity_proofing_provider_authorization
  # resource_type_schema_version      = var.identity_proofing_provider_resource_type_schema_version
  start_index                       = var.identity_proofing_provider_start_index
}

