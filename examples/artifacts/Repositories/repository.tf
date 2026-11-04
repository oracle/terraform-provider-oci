// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "repository_defined_tags_value" {
  default = "value"
}

variable "repository_description" {
  default = "description"
}

variable "repository_display_name" {
  default = "displayName"
}

variable "repository_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "repository_id" {
  default = "id"
}

variable "repository_is_immutable" {
  default = false
}

variable "repository_repository_type" {
  default = "GENERIC"
}

variable "repository_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_artifacts_repository" "test_repository" {
  #Required
  compartment_id  = var.compartment_id
  is_immutable    = var.repository_is_immutable
  repository_type = var.repository_repository_type

  #Optional
  defined_tags  = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.repository_defined_tags_value })
  description   = var.repository_description
  display_name  = var.repository_display_name
  freeform_tags = var.repository_freeform_tags
}

data "oci_artifacts_repositories" "test_repositories" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.repository_display_name
  id           = var.repository_id
  is_immutable = var.repository_is_immutable
  state        = var.repository_state
}

