// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "container_image_compartment_id_in_subtree" {
  default = false
}

variable "container_image_display_name" {
  default = "displayName"
}

variable "container_image_image_digest" {
  default = "imageDigest"
}

variable "container_image_is_versioned" {
  default = false
}

variable "container_image_state" {
  default = "AVAILABLE"
}

variable "container_image_version" {
  default = "version"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_artifacts_container_images" "test_container_images" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.container_image_compartment_id_in_subtree
  display_name              = var.container_image_display_name
  image_digest              = var.container_image_image_digest
  image_id                  = oci_core_image.test_image.id
  is_versioned              = var.container_image_is_versioned
  repository_id             = oci_artifacts_repository.test_repository.id
  repository_name           = oci_artifacts_repository.test_repository.name
  state                     = var.container_image_state
  version                   = var.container_image_version
}

