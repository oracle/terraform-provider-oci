// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "mapped_attribute_mapped_attribute_count" {
  default = 10
}

variable "mapped_attribute_mapped_attribute_filter" {
  default = ""
}

variable "mapped_attribute_attribute_mappings_applies_to_actions" {
  default = []
}

variable "mapped_attribute_attribute_mappings_idcs_attribute_name" {
  default = "$(user.userName)"
}

variable "mapped_attribute_attribute_mappings_managed_object_attribute_name" {
  default = "name"
}

variable "mapped_attribute_attribute_mappings_required" {
  default = false
}

variable "mapped_attribute_attribute_mappings_saml_format" {
  default = "samlFormat"
}

variable "mapped_attribute_attribute_sets" {
  default = []
}

variable "mapped_attribute_attributes" {
  default = "attributes"
}

variable "mapped_attribute_authorization" {
  default = "authorization"
}

variable "mapped_attribute_compartment_ocid" {
  default = "compartmentOcid"
}

variable "mapped_attribute_delete_in_progress" {
  default = false
}

variable "mapped_attribute_direction" {
  default = "direction"
}

variable "mapped_attribute_domain_ocid" {
  default = "domainOcid"
}

variable "mapped_attribute_id" {
  default = "id"
}

variable "mapped_attribute_idcs_created_by__ref" {
  default = "ref"
}

variable "mapped_attribute_idcs_created_by_display" {
  default = "display"
}

variable "mapped_attribute_idcs_created_by_ocid" {
  default = "ocid"
}

variable "mapped_attribute_idcs_created_by_type" {
  default = "User"
}

variable "mapped_attribute_idcs_created_by_value" {
  default = "value"
}

variable "mapped_attribute_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "mapped_attribute_idcs_last_modified_by__ref" {
  default = "ref"
}

variable "mapped_attribute_idcs_last_modified_by_display" {
  default = "display"
}

variable "mapped_attribute_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "mapped_attribute_idcs_last_modified_by_type" {
  default = "User"
}

variable "mapped_attribute_idcs_last_modified_by_value" {
  default = "value"
}

variable "mapped_attribute_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "mapped_attribute_idcs_prevented_operations" {
  default = []
}

variable "mapped_attribute_idcs_resource_type" {
  default = "User"
}

variable "mapped_attribute_meta_created" {
  default = "created"
}

variable "mapped_attribute_meta_last_modified" {
  default = "lastModified"
}

variable "mapped_attribute_meta_location" {
  default = "location"
}

variable "mapped_attribute_meta_resource_type" {
  default = "resourceType"
}

variable "mapped_attribute_meta_version" {
  default = "version"
}

variable "mapped_attribute_ocid" {
  default = "ocid"
}

variable "mapped_attribute_ref_resource_id" {
  default = "refResourceID"
}

variable "mapped_attribute_ref_resource_type" {
  default = "App"
}

variable "mapped_attribute_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "mapped_attribute_schemas" {
  default = []
}

variable "mapped_attribute_start_index" {
  default = 1
}

variable "mapped_attribute_tags_key" {
  default = "key"
}

variable "mapped_attribute_tags_value" {
  default = "value"
}

variable "mapped_attribute_tenancy_ocid" {
  default = "tenancyOcid"
}


resource "oci_identity_domains_mapped_attribute" "test_mapped_attribute" {
  #Required
  direction           = "outbound"
  idcs_endpoint       = data.oci_identity_domain.test_domain.url
  idcs_resource_type  = "User"
  mapped_attribute_id = data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.id
  ref_resource_id     = var.mapped_attribute_ref_resource_id
  ref_resource_type   = "App"
  schemas             = ["urn:ietf:params:scim:schemas:oracle:idcs:MappedAttribute"]

  #Optional
  attribute_mappings {
    #Required
    idcs_attribute_name           = "$(user.userName)"
    managed_object_attribute_name = "name"

    #Optional
    applies_to_actions = ["create"]
    #required           = var.mapped_attribute_attribute_mappings_required
    #saml_format        = var.mapped_attribute_attribute_mappings_saml_format
  }
  attribute_sets = ["all"]
  attributes     = ""
  #authorization  = var.mapped_attribute_authorization
  #id             = var.mapped_attribute_id
  #ocid           = var.mapped_attribute_ocid
  #resource_type_schema_version = var.mapped_attribute_resource_type_schema_version
  tags {
    #Required
    key   = var.mapped_attribute_tags_key
    value = var.mapped_attribute_tags_value
  }
}

data "oci_identity_domains_mapped_attributes" "test_mapped_attributes" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  mapped_attribute_count  = var.mapped_attribute_mapped_attribute_count
  mapped_attribute_filter = var.mapped_attribute_mapped_attribute_filter
  attribute_sets          = ["all"]
  attributes              = ""
  authorization           = var.mapped_attribute_authorization
  #resource_type_schema_version = var.mapped_attribute_resource_type_schema_version
  start_index = var.mapped_attribute_start_index
}

output "mapped_attribute_resource" {
  value = {
    mapped_attribute_id = data.oci_identity_domains_mapped_attributes.test_mapped_attributes.mapped_attributes.0.id
  }
}
