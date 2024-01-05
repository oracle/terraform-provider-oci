// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "environment_defined_tags_value" {
  default = "value"
}

variable "environment_display_name" {
  default = "displayName"
}

variable "environment_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "environment_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_environment" "test_environment" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.environment_defined_tags_value)
  display_name  = var.environment_display_name
  freeform_tags = var.environment_freeform_tags
}

data "oci_cloud_bridge_environments" "test_environments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name   = var.environment_display_name
  environment_id = oci_cloud_bridge_environment.test_environment.id
  state          = var.environment_state
}

