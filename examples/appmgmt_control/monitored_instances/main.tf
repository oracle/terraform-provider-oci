// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "monitored_instance_id" {}

variable "monitored_instance_display_name" {
  default = "displayName"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_appmgmt_control_monitored_instances" "test_monitored_instances" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.monitored_instance_display_name
}

data "oci_appmgmt_control_monitored_instance" "test_monitored_instance" {
  #Required
  monitored_instance_id = var.monitored_instance_id
}
