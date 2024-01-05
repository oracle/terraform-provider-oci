// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "alert_access_level" {
  default = "RESTRICTED"
}

variable "alert_comment" {
  default = "comment"
}

variable "alert_compartment_id_in_subtree" {
  default = true
}

variable "alert_defined_tags_value" {
  default = "value"
}

variable "alert_field" {
  default = []
}

variable "alert_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "data_safe_alert_ocid" {}

variable "alert_scim_query" {
  default = "severity eq 'HIGH'"
}

variable "alert_status" {
  default = "OPEN"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_alert" "test_alert" {
  #Required
  alert_id = var.data_safe_alert_ocid

  #Optional
  comment       = var.alert_comment
  freeform_tags = var.alert_freeform_tags
  status        = var.alert_status
}

data "oci_data_safe_alerts" "test_alerts" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level              = var.alert_access_level
  compartment_id_in_subtree = var.alert_compartment_id_in_subtree
  id                        = var.data_safe_alert_ocid
  scim_query                = var.alert_scim_query
}

