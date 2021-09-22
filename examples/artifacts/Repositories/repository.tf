// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

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
  is_immutable    = false
  repository_type = "GENERIC"
}

resource "oci_generic_artifacts_content_artifact_by_path" "test_artifact" {
  #Required
  artifact_path  = "artifact_path"
  repository_id    = oci_artifacts_repository.test_repository.id
  version = "1.0"
  content = "<a1>content</a1>"
}
resource "oci_artifacts_generic_artifact" "test_artifact" {
  artifact_id = oci_generic_artifacts_content_artifact_by_path.test_artifact.id
}


resource "oci_generic_artifacts_content_artifact_by_path" "test_artifact_by_source" {
  #Required
  artifact_path  = "artifact_path"
  repository_id    = oci_artifacts_repository.test_repository.id
  version = "2.0"
  source = "index.html"
}
resource "oci_artifacts_generic_artifact" "test_artifact_by_source" {
  artifact_id = oci_generic_artifacts_content_artifact_by_path.test_artifact_by_source.id
}

data "oci_artifacts_repositories" "test_repositories" {
  #Required
  compartment_id = var.compartment_id
  state = "AVAILABLE"
}

