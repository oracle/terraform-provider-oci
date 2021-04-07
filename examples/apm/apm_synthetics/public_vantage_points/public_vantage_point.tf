// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}

variable "public_vantage_point_display_name" {
  default = "displayName"
}

variable "public_vantage_point_name" {
  default = "name"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_apm_synthetics_public_vantage_points" "test_public_vantage_points" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  display_name = var.public_vantage_point_display_name
  name         = var.public_vantage_point_name
}

