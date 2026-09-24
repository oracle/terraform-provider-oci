// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "generic_artifact_artifact_path" {
  default = "artifactPath"
}

variable "generic_artifact_defined_tags_value" {
  default = "value"
}

variable "generic_artifact_display_name" {
  default = "displayName"
}

variable "generic_artifact_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "generic_artifact_id" {
  default = "id"
}

variable "generic_artifact_sha256" {
  default = "sha256"
}

variable "generic_artifact_state" {
  default = "AVAILABLE"
}

variable "generic_artifact_version" {
  default = "version"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_artifacts_generic_artifact" "test_generic_artifact" {
  #Required
  artifact_id = oci_artifacts_artifact.test_artifact.id

  #Optional
  defined_tags  = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.generic_artifact_defined_tags_value })
  freeform_tags = var.generic_artifact_freeform_tags
}

data "oci_artifacts_generic_artifacts" "test_generic_artifacts" {
  #Required
  compartment_id = var.compartment_id
  repository_id  = oci_artifacts_repository.test_repository.id

  #Optional
  artifact_path = var.generic_artifact_artifact_path
  display_name  = var.generic_artifact_display_name
  id            = var.generic_artifact_id
  sha256        = var.generic_artifact_sha256
  state         = var.generic_artifact_state
  version       = var.generic_artifact_version
}

