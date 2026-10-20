// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_id" {}
variable "crypto_assessment_id" {}

variable "crypto_assessment_display_name" {
  default = "displayName"
}

variable "crypto_assessment_description" {
  default = "description"
}

variable "is_assessment_scheduled" {
  default = false
}

variable "crypto_assessment_schedule" {
  default = "v1; 00 30 15 * *"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  tenancy_ocid        = var.tenancy_ocid
  region              = var.region
}

resource "oci_data_safe_crypto_assessment_management" "test_crypto_assessment" {
  crypto_assessment_id    = var.crypto_assessment_id
  compartment_id          = var.compartment_id
  display_name            = var.crypto_assessment_display_name
  description             = var.crypto_assessment_description
  schedule                = var.crypto_assessment_schedule
  is_assessment_scheduled = var.is_assessment_scheduled

  lifecycle {
    ignore_changes = [defined_tags, freeform_tags, system_tags]
  }
}
