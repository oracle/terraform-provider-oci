// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "container_image_ocid" {}
variable "container_image_repository_ocid" {}

variable "container_image_compartment_id_in_subtree" {
  default = false
}

variable "container_image_is_versioned" {
  default = true
}

variable "container_image_state" {
  default = "AVAILABLE"
}

variable "container_image_repository_name" {
  default = "dont-delete-used-by-terraform-test/busybox"
}

variable "container_image_display_name" {
  default = "dont-delete-used-by-terraform-test/busybox:latest"
}

variable "image_version" {
  default = "latest"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_artifacts_container_image" "test_container_image" {
  #Required
  image_id = var.container_image_ocid
}

data "oci_artifacts_container_images" "test_container_images" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  compartment_id_in_subtree = var.container_image_compartment_id_in_subtree
  is_versioned = var.container_image_is_versioned
  state = var.container_image_state
  image_id = var.container_image_ocid
  display_name = var.container_image_display_name
  repository_id = var.container_image_repository_ocid
  repository_name = var.container_image_repository_name
  version = var.image_version
}