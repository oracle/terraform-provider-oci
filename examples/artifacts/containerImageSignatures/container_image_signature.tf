// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "container_image_signature_compartment_id_in_subtree" {
  default = false
}

variable "container_image_signature_defined_tags_value" {
  default = "value"
}

variable "container_image_signature_display_name" {
  default = "displayName"
}

variable "container_image_signature_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "container_image_signature_image_digest" {
  default = "imageDigest"
}

variable "container_image_signature_message" {
  default = "message"
}

variable "container_image_signature_signature" {
  default = "signature"
}

variable "container_image_signature_signing_algorithm" {
  default = "SHA_224_RSA_PKCS_PSS"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_artifacts_container_image_signature" "test_container_image_signature" {
  #Required
  compartment_id     = var.compartment_id
  image_id           = oci_core_image.test_image.id
  kms_key_id         = oci_kms_key.test_key.id
  kms_key_version_id = oci_kms_key_version.test_key_version.id
  message            = var.container_image_signature_message
  signature          = var.container_image_signature_signature
  signing_algorithm  = var.container_image_signature_signing_algorithm

  #Optional
  defined_tags  = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.container_image_signature_defined_tags_value })
  freeform_tags = var.container_image_signature_freeform_tags
}

data "oci_artifacts_container_image_signatures" "test_container_image_signatures" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.container_image_signature_compartment_id_in_subtree
  display_name              = var.container_image_signature_display_name
  image_digest              = var.container_image_signature_image_digest
  image_id                  = oci_core_image.test_image.id
  kms_key_id                = oci_kms_key.test_key.id
  kms_key_version_id        = oci_kms_key_version.test_key_version.id
  repository_id             = oci_artifacts_repository.test_repository.id
  repository_name           = oci_artifacts_repository.test_repository.name
  signing_algorithm         = var.container_image_signature_signing_algorithm
}

