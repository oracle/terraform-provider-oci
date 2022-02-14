// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "apm_domain_id" {}
variable "dvp_stack_id" {}
variable "dvp_stream_id" {}
variable "dvp_version" {}
variable "dvp_region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "dedicated_vantage_point_display_name" {
  default = "displayName"
}

variable "dedicated_vantage_point_name" {
  default = "name"
}

variable "dvp_stack_type" {
  default = "ORACLE_RM_STACK"
}

resource "oci_apm_synthetics_dedicated_vantage_point" "test_dedicated_vantage_point" {
  #Required
  apm_domain_id    = var.apm_domain_id
  display_name     = var.dedicated_vantage_point_display_name
  dvp_stack_details {
    #Required
    dvp_stack_id     = var.dvp_stack_id
    dvp_stream_id    = var.dvp_stream_id
    dvp_version      = var.dvp_version
    dvp_stack_type   = var.dvp_stack_type
  }
  region       = var.dvp_region

}

data "oci_apm_synthetics_dedicated_vantage_points" "test_dedicated_vantage_points" {
  #Required
  apm_domain_id    = var.apm_domain_id

  #Optional
  display_name     = var.dedicated_vantage_point_display_name
  name         = var.dedicated_vantage_point_name
}
