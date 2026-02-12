// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "email_ip_pool_defined_tags_value" {
  default = "value"
}

variable "email_ip_pool_description" {
  default = "description"
}

variable "email_ip_pool_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "email_ip_pool_id" {
  default = "id"
}

variable "email_ip_pool_name" {
  default = "name"
}

variable "email_ip_pool_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_email_email_ip_pool" "test_email_ip_pool" {
  #Required
  compartment_id = var.compartment_id
  name           = var.email_ip_pool_name
  outbound_ips {
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.email_ip_pool_defined_tags_value)
  description   = var.email_ip_pool_description
  freeform_tags = var.email_ip_pool_freeform_tags
}

data "oci_email_email_ip_pools" "test_email_ip_pools" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id    = var.email_ip_pool_id
  name  = var.email_ip_pool_name
  state = var.email_ip_pool_state
}

