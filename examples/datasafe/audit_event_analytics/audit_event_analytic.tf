// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "audit_event_analytic_access_level" {
  default = "RESTRICTED"
}

variable "audit_event_analytic_compartment_id_in_subtree" {
  default = false
}

variable "audit_event_analytic_group_by" {
  default = []
}

variable "audit_event_analytic_query_time_zone" {
  default = "queryTimeZone"
}

variable "audit_event_analytic_scim_query" {
  default = "scimQuery"
}

variable "audit_event_analytic_summary_field" {
  default = []
}

variable "audit_event_analytic_time_ended" {
  default = "timeEnded"
}

variable "audit_event_analytic_time_started" {
  default = "timeStarted"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_audit_event_analytic" "test_audit_event_analytic" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  summary_field             = var.audit_event_analytic_summary_field
}

