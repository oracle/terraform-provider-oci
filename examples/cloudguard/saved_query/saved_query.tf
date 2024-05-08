// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "saved_query_access_level" {
  default = "ACCESSIBLE"
}

variable "saved_query_compartment_id_in_subtree" {
  default = true
}

variable "saved_query_defined_tags_value" {
  default = "value"
}

variable "saved_query_description" {
  default = "description"
}

variable "saved_query_display_name" {
  default = "displayName"
}

variable "saved_query_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "saved_query_query" {
  default = "select pid from processes"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  //version             = "5.39.0"
  /*tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region*/
}

resource "oci_cloud_guard_saved_query" "test_saved_query" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.saved_query_display_name
  query          = var.saved_query_query

  #Optional
  defined_tags  = { "example-tag-namespace-all.example-tag" = var.saved_query_defined_tags_value}
  description   = var.saved_query_description
  freeform_tags = var.saved_query_freeform_tags
}

data "oci_cloud_guard_saved_queries" "test_saved_queries" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  access_level              = var.saved_query_access_level
  compartment_id_in_subtree = var.saved_query_compartment_id_in_subtree
  display_name              = var.saved_query_display_name
}
