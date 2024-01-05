// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "private_endpoint_defined_tags_value" {
  default = "value"
}

variable "private_endpoint_description" {
  default = "Example Private Endpoint"
}

variable "private_endpoint_display_name" {
  default = "My Private Endpoint"
}

variable "private_endpoint_dns_zones" {
  default = []
}

variable "private_endpoint_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "private_endpoint_is_used_with_configuration_source_provider" {
  default = false
}

variable "private_endpoint_nsg_id_list" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_resourcemanager_private_endpoint" "test_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.private_endpoint_display_name
  subnet_id      = oci_core_subnet.test_subnet.id
  vcn_id         = oci_core_vcn.test_vcn.id

  #Optional
  defined_tags                               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.private_endpoint_defined_tags_value)
  description                                = var.private_endpoint_description
  dns_zones                                  = var.private_endpoint_dns_zones
  freeform_tags                              = var.private_endpoint_freeform_tags
  is_used_with_configuration_source_provider = var.private_endpoint_is_used_with_configuration_source_provider
  nsg_id_list                                = var.private_endpoint_nsg_id_list
}

data "oci_resourcemanager_private_endpoints" "test_private_endpoints" {

  #Optional
  compartment_id      = var.compartment_id
  display_name        = var.private_endpoint_display_name
  private_endpoint_id = oci_resourcemanager_private_endpoint.test_private_endpoint.id
  vcn_id              = oci_core_vcn.test_vcn.id
}

