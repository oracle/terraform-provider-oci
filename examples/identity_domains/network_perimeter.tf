// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "network_perimeter_network_perimeter_count" {
  default = 10
}

variable "network_perimeter_network_perimeter_filter" {
  default = ""
}

variable "network_perimeter_authorization" {
  default = "authorization"
}

variable "network_perimeter_compartment_ocid" {
  default = "compartmentOcid"
}

variable "network_perimeter_delete_in_progress" {
  default = false
}

variable "network_perimeter_description" {
  default = "description"
}

variable "network_perimeter_domain_ocid" {
  default = "domainOcid"
}

variable "network_perimeter_id" {
  default = "id"
}

variable "network_perimeter_idcs_created_by_display" {
  default = "display"
}

variable "network_perimeter_idcs_created_by_ocid" {
  default = "ocid"
}

variable "network_perimeter_idcs_created_by_ref" {
  default = "ref"
}

variable "network_perimeter_idcs_created_by_type" {
  default = "User"
}

variable "network_perimeter_idcs_created_by_value" {
  default = "value"
}

variable "network_perimeter_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "network_perimeter_idcs_last_modified_by_display" {
  default = "display"
}

variable "network_perimeter_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "network_perimeter_idcs_last_modified_by_ref" {
  default = "ref"
}

variable "network_perimeter_idcs_last_modified_by_type" {
  default = "User"
}

variable "network_perimeter_idcs_last_modified_by_value" {
  default = "value"
}

variable "network_perimeter_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "network_perimeter_idcs_prevented_operations" {
  default = []
}

variable "network_perimeter_ip_addresses_type" {
  default = "RANGE"
}

// need to provide ip addresses value
variable "network_perimeter_ip_addresses_value" {
  default = "value"
}

variable "network_perimeter_ip_addresses_version" {
  default = "IPV4"
}

variable "network_perimeter_meta_created" {
  default = "created"
}

variable "network_perimeter_meta_last_modified" {
  default = "lastModified"
}

variable "network_perimeter_meta_location" {
  default = "location"
}

variable "network_perimeter_meta_resource_type" {
  default = "resourceType"
}

variable "network_perimeter_meta_version" {
  default = "version"
}

variable "network_perimeter_name" {
  default = "name"
}

variable "network_perimeter_ocid" {
  default = "ocid"
}

variable "network_perimeter_resource_type_schema_version" {
  default = "resourceTypeSchemaVersion"
}

variable "network_perimeter_start_index" {
  default = 1
}

variable "network_perimeter_tags_key" {
  default = "key"
}

variable "network_perimeter_tags_value" {
  default = "value"
}

variable "network_perimeter_tenancy_ocid" {
  default = "tenancyOcid"
}


resource "oci_identity_domains_network_perimeter" "test_network_perimeter" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  ip_addresses {
    #Required
    value = var.network_perimeter_ip_addresses_value

    #Optional
    type    = var.network_perimeter_ip_addresses_type
    version = var.network_perimeter_ip_addresses_version
  }
  name    = var.network_perimeter_name
  schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:NetworkPerimeter"]

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.network_perimeter_authorization
  description    = var.network_perimeter_description
  external_id    = "externalId"
  ocid           = var.network_perimeter_ocid
  #use the latest if not provided
  #resource_type_schema_version = var.condition_resource_type_schema_version
  tags {
    #Required
    key   = var.network_perimeter_tags_key
    value = var.network_perimeter_tags_value
  }
}

data "oci_identity_domains_network_perimeters" "test_network_perimeters" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  network_perimeter_count  = var.network_perimeter_network_perimeter_count
  network_perimeter_filter = var.network_perimeter_network_perimeter_filter
  attribute_sets           = ["all"]
  attributes               = ""
  authorization            = var.network_perimeter_authorization
  #use the latest if not provided
  #resource_type_schema_version = var.condition_resource_type_schema_version
  start_index = var.network_perimeter_start_index
}

