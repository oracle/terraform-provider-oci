// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "application_id" {}
variable "application_name" {}
variable "host_id" {}
variable "list_jre_usage_time_end" {}
variable "list_jre_usage_time_start" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_list_jre_usage" "test_list_jre_usage" {

  #Optional
  application_id   = var.application_id
  application_name = var.application_name
  compartment_id   = var.compartment_id
  host_id          = var.host_id
  time_end         = var.list_jre_usage_time_end
  time_start       = var.list_jre_usage_time_start
}
