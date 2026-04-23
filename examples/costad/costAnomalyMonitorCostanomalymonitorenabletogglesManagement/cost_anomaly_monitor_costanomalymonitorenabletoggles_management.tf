// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management" "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management" {
  #Required
  cost_anomaly_monitor_id               = oci_costad_cost_anomaly_monitor.test_cost_anomaly_monitor.id
  enable_costanomalymonitorenabletoggle = var.enable_costanomalymonitorenabletoggle
}


