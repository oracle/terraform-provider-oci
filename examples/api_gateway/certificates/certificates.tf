// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "certificate_certificate" {
  default = "-----BEGIN CERTIFICATE-----\n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\n-----END CERTIFICATE-----"
}

variable "certificate_defined_tags_value" {
  default = "value"
}

variable "certificate_display_name" {
  default = "displayName"
}

variable "certificate_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "certificate_intermediate_certificates" {
  default = "-----BEGIN CERTIFICATE-----\n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\n-----END CERTIFICATE-----"
}

variable "certificate_private_key" {
  default = "-----BEGIN PRIVATE KEY-----\n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\nVH5xRlhkwPht7KCedeiQpehEOOnA2A==\n-----END PRIVATE KEY-----"
}

variable "certificate_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_apigateway_certificate" "test_certificate" {
  #Required
  certificate    = "${var.certificate_certificate}"
  compartment_id = "${var.compartment_id}"
  private_key    = "${var.certificate_private_key}"

  #Optional
  defined_tags              = { "example-tag-namespace-all.example-tag" = "value" }
  display_name              = "${var.certificate_display_name}"
  freeform_tags             = "${var.certificate_freeform_tags}"
  intermediate_certificates = "${var.certificate_intermediate_certificates}"
  lifecycle {
      ignore_changes = [defined_tags]
  }
}

data "oci_apigateway_certificates" "test_certificates" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  display_name = "${var.certificate_display_name}"
  state        = "${var.certificate_state}"
}
