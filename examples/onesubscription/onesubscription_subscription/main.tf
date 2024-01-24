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

data "oci_onesubscription_subscriptions" "test_subscriptions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #buyer_email                   = var.subscription_buyer_email
  #is_commit_info_required       = var.subscription_is_commit_info_required
  #plan_number                   = var.subscription_plan_number
  subscription_id                = var.subscription_id
}
