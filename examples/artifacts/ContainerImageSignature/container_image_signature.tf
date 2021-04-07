// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "container_image_signature_compartment_id_in_subtree" {}

// specify the container image to upload to
variable "container_image_id" {}

// specify the management endpoint for the key
variable "crypto_endpoint" {}

// specify the kms key to sign the message
variable "kms_rsa_key_id" {}

// specify the kms key version to sign the message
variable "kms_rsa_key_version_id" {}

// the algorithm to sign with the key
variable "kms_signing_algorithm" {}

// user inputted description to include in the signature
variable "description" {}

// user defined a json string to include in the signature (eg. use escape character for the key/value string)
// ex. "{\\\"buildNumber\\\":\\\"123\\\"}"
variable "metadata" {}

data "oci_artifacts_container_image" "test_container_image" {
  image_id = var.container_image_id
}

output "oci_test_container_image" {
  value = data.oci_artifacts_container_image.test_container_image.repository_name
}

locals {
  message = templatefile("./artifacts_container_image_signature_message_json.tmpl", {
    "description" = var.description
    "digest" = data.oci_artifacts_container_image.test_container_image.digest
    "kms_key_id" = var.kms_rsa_key_id
    "kms_key_version_id" = var.kms_rsa_key_version_id
    "metadata" = var.metadata
    "region" = var.region
    "repository_name" = data.oci_artifacts_container_image.test_container_image.repository_name
    "signing_algorithm" = var.kms_signing_algorithm
  })
}

resource "oci_kms_sign" "test_sign" {
  crypto_endpoint      = var.crypto_endpoint
  key_id               = var.kms_rsa_key_id
  message              = base64encode(local.message)
  signing_algorithm    = var.kms_signing_algorithm

  key_version_id = var.kms_rsa_key_version_id
  message_type = "RAW"
}

resource "oci_artifacts_container_image_signature" "test_container_image_signature" {
  #Required
  compartment_id = var.compartment_ocid
  image_id = var.container_image_id
  kms_key_id = var.kms_rsa_key_id
  kms_key_version_id = var.kms_rsa_key_version_id
  message = base64encode(local.message)
  signature = oci_kms_sign.test_sign.signature
  signing_algorithm = var.kms_signing_algorithm
}

data "oci_artifacts_container_image_signatures" "test_container_image_signatures" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  compartment_id_in_subtree = var.container_image_signature_compartment_id_in_subtree
  image_id = var.container_image_id
  kms_key_id = var.kms_rsa_key_id
}
