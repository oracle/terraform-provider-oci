// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "iot_domain_group_ocid" {}

variable "iot_domain_defined_tags_value" {
  default = "value"
}

variable "iot_domain_description" {
  default = "description"
}

variable "iot_domain_display_name" {
  default = "displayName"
}

variable "iot_domain_freeform_tags" {
  default = { "Protocol" = "MQTT" }
}

variable "iot_domain_id" {
  default = "id"
}

variable "iot_domain_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_domain" "test_iot_domain" {
  #Required
  compartment_id      = var.compartment_ocid
  iot_domain_group_id = var.iot_domain_group_ocid

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.iot_domain_defined_tags_value)
  description   = var.iot_domain_description
  display_name  = var.iot_domain_display_name
  freeform_tags = var.iot_domain_freeform_tags
}

data "oci_iot_iot_domains" "test_iot_domains" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name        = var.iot_domain_display_name
  id                  = var.iot_domain_id
  iot_domain_group_id = var.iot_domain_group_ocid
  state               = var.iot_domain_state
}

