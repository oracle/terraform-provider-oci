// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "cost_anomaly_event_analytic_target_tenant_id" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_costad_cost_anomaly_event_analytics" "test_cost_anomaly_event_analytics" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  region           = [var.region]
  target_tenant_id = length(var.cost_anomaly_event_analytic_target_tenant_id) > 0 ? var.cost_anomaly_event_analytic_target_tenant_id : [var.compartment_id]
}
