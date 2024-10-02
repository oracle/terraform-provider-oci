// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "configuration_defined_tags" {
  default = {}
}

variable "configuration_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "configuration_zpr_status" {
  default = "ENABLED"
}

variable "zpr_policy_description" {
  default = "description"
}

variable "zpr_policy_defined_tags" {
  default = {}
}

variable "zpr_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "zpr_policy_name" {
  default = "name"
}

variable "zpr_policy_state" {
  default = "ACTIVE"
}

variable "zpr_policy_statements" {
  default = []
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_zpr_configuration" "test_configuration" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  defined_tags  = var.configuration_defined_tags
  freeform_tags = var.configuration_freeform_tags
}

data "oci_zpr_configuration" "test_configuration" {
  #Required
  compartment_id = var.tenancy_ocid
}

resource "oci_zpr_zpr_policy" "test_zpr_policy" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.zpr_policy_description
  name           = var.zpr_policy_name
  statements     = var.zpr_policy_statements

  #Optional
  defined_tags  = var.zpr_policy_defined_tags
  freeform_tags = var.zpr_policy_freeform_tags
}

data "oci_zpr_zpr_policies" "test_zpr_policies" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  name  = var.zpr_policy_name
  state = var.zpr_policy_state
}

data "oci_zpr_zpr_policy" "test_zpr_policy" {
	#Required
	zpr_policy_id = oci_zpr_zpr_policy.test_zpr_policy.id
}