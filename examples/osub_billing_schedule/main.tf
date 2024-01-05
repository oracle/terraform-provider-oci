// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscription_id" {}

variable "billing_schedule_x_one_origin_region" {
  default = "xOneOriginRegion"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osub_billing_schedule_billing_schedules" "test_billing_schedules" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id

  #Optional
  x_one_origin_region = var.region
  #subscribed_service_id = var.subscribed_service_id
}
