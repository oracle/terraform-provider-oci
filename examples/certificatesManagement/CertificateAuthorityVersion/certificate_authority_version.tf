// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "certificate_authority_version_certificate_authority_version_number" {
  default = 10
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_certificates_management_certificate_authority_version" "test_certificate_authority_version" {
  #Required
  certificate_authority_id             = oci_certificates_management_certificate_authority.test_certificate_authority.id
  certificate_authority_version_number = var.certificate_authority_version_certificate_authority_version_number
}

