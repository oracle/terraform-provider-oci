// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "vbs_instance_defined_tags_value" {
  default = "value"
}

variable "vbs_instance_display_name" {
  default = "displayName"
}

variable "vbs_instance_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "vbs_instance_id" {
  default = "id"
}

variable "vbs_instance_idcs_access_token" {
  default = "idcsAccessToken"
}

variable "vbs_instance_is_resource_usage_agreement_granted" {
  default = true
}

variable "vbs_instance_name" {
  default = "name"
}

variable "vbs_instance_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_vbs_inst_vbs_instance" "test_vbs_instance" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.vbs_instance_display_name
  name           = var.vbs_instance_name

  #Optional
  #defined_tags                        = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.vbs_instance_defined_tags_value)
  freeform_tags                       = var.vbs_instance_freeform_tags
  #idcs_access_token                   = var.vbs_instance_idcs_access_token
  is_resource_usage_agreement_granted = var.vbs_instance_is_resource_usage_agreement_granted
  resource_compartment_id             = var.resource_compartment_id
}

data "oci_vbs_inst_vbs_instances" "test_vbs_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id    = var.vbs_instance_id
  name  = var.vbs_instance_name
  state = var.vbs_instance_state
}

