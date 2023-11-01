// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "sql_firewall_violation_analytic_compartment_id_in_subtree" {
  default = false
}

variable "sql_firewall_violation_analytic_access_level" {
  default = "RESTRICTED"
}

variable "sql_firewall_violation_analytic_summary_field"{
  default = []
}

variable "sql_firewall_violation_analytic_group_by"{
  default = []
}

variable "sql_firewall_violation_analytic_scim_query" {
  default = "violationCause eq \"SQL violation\""
}

variable "sql_firewall_violation_analytic_time_ended" {
  default = "2038-01-01T00:00:00.000Z"
}

variable "sql_firewall_violation_analytic_time_started" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "sql_firewall_violation_analytic_query_time_zone" {
  default = "UTC"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_sql_firewall_violation_analytics" "test_sql_firewall_violation_analytics" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  compartment_id_in_subtree = var.sql_firewall_violation_analytic_compartment_id_in_subtree
  access_level = var.sql_firewall_violation_analytic_access_level
  scim_query = var.sql_firewall_violation_analytic_scim_query
  query_time_zone = var.sql_firewall_violation_analytic_query_time_zone
}