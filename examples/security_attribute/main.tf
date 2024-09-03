// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "security_attribute_namespace_compartment_id_in_subtree" {
  default = false
}

variable "security_attribute_namespace_defined_tags_value" {
  default = "value"
}

variable "security_attribute_namespace_description" {
  default = "This is the Zero Trust Packet Routing security attribute namespace description sample."
}

variable "security_attribute_namespace_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "security_attribute_namespace_name" {
  default = "example-security-attribute-namespace"
}

variable "security_attribute_namespace_state" {
  default = "ACTIVE"
}

variable "security_attribute_description" {
  default = "This is a sample security attribute description."
}

variable "security_attribute_name" {
  default = "TFTestSecurityAttribute"
}

variable "security_attribute_state" {
  default = "ACTIVE"
}

variable "security_attribute_validator_validator_type" {
  default = "ENUM"
}

variable "security_attribute_validator_values" {
  default = []
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint       = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_security_attribute_security_attribute_namespace" "test_security_attribute_namespace" {
  #Required
  compartment_id = var.compartment_id
  description    = var.security_attribute_namespace_description
  name           = var.security_attribute_namespace_name

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.security_attribute_namespace_defined_tags_value)
  freeform_tags = var.security_attribute_namespace_freeform_tags
}

data "oci_security_attribute_security_attribute_namespaces" "test_security_attribute_namespaces" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.security_attribute_namespace_compartment_id_in_subtree
  name                      = var.security_attribute_namespace_name
  state                     = var.security_attribute_namespace_state
}

resource "oci_security_attribute_security_attribute" "test_security_attribute" {
  #Required
  description                     = var.security_attribute_description
  name                            = var.security_attribute_name
  security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id

  #Optional
  validator {
    #Required
    validator_type = var.security_attribute_validator_validator_type

    #Optional
    values = var.security_attribute_validator_values
  }
}

data "oci_security_attribute_security_attributes" "test_security_attributes" {
  #Required
  security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id

  #Optional
  state = var.security_attribute_state
}

data "oci_security_attribute_security_attribute" "test_security_attribute" {
	#Required
	security_attribute_name = oci_security_attribute_security_attribute.test_security_attribute.name
	security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id
}