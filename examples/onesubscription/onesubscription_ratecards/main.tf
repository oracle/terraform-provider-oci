// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscription_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_onesubscription_ratecards" "test_ratecards" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id

  #Optional
  #part_number         = var.ratecard_part_number
  #time_from           = var.ratecard_time_from
  #time_to             = var.ratecard_time_to
}
