// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "java_license_acceptance_record_id" {
  default = "id"
}

variable "java_license_acceptance_record_license_type" {
  default = "OTN"
}

variable "java_license_acceptance_record_status" {
  default = "ACCEPTED"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_java_downloads_java_license_acceptance_record" "test_java_license_acceptance_record" {
  #Required
  compartment_id            = var.tenancy_ocid
  license_acceptance_status = var.java_license_acceptance_record_status
  license_type              = var.java_license_acceptance_record_license_type
  lifecycle {
    ignore_changes = [defined_tags, system_tags]
  }
}

data "oci_jms_java_downloads_java_license_acceptance_records" "test_java_license_acceptance_records" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  id             = var.java_license_acceptance_record_id
  license_type   = var.java_license_acceptance_record_license_type
  search_by_user = var.user_ocid
  status         = var.java_license_acceptance_record_status
}

