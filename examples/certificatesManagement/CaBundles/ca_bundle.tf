// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ca_bundle_ca_bundle_pem" {
  default = "-----BEGIN PUBLIC KEY-----<var>&lt;keycontents&gt;</var>n-----END PUBLIC KEY-----n"
}

variable "ca_bundle_defined_tags_value" {
  default = "value"
}

variable "ca_bundle_description" {
  default = "description"
}

variable "ca_bundle_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "ca_bundle_name" {
  default = "test-ca-bundle"
}

variable "ca_bundle_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_certificates_management_ca_bundle" "test_ca_bundle" {
  #Required
  ca_bundle_pem  = var.ca_bundle_ca_bundle_pem
  compartment_id = var.compartment_id
  name           = var.ca_bundle_name

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.ca_bundle_defined_tags_value)
  description   = var.ca_bundle_description
  freeform_tags = var.ca_bundle_freeform_tags
}

data "oci_certificates_management_ca_bundles" "test_ca_bundles" {

  #Optional
  ca_bundle_id   = oci_certificates_management_ca_bundle.test_ca_bundle.id
  compartment_id = var.compartment_id
  name           = var.ca_bundle_name
  state          = var.ca_bundle_state
}

