// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "instance_idcs_url" {}

variable "instance_defined_tags_value" {
  default = "value"
}

variable "instance_display_name" {
  default = "displayName"
}

variable "instance_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "instance_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_ddfs_instance" "test_instance" {
  #Required
  compartment_id = var.compartment_id
  idcs_url       = var.instance_idcs_url

  #Optional
  #defined_tags  = tomap({ "Operations.CostCenter" = var.instance_defined_tags_value })
  display_name  = var.instance_display_name
  freeform_tags = var.instance_freeform_tags
}

data "oci_ddfs_instances" "test_instances" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.instance_display_name
  id             = oci_ddfs_instance.test_instance.id
  state          = var.instance_state
}
