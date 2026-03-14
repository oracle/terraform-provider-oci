// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "partner_subscription_display_name" {
  default = "displayName"
}

variable "listing_id" {
  type = string
  default = "listing_id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_self_partner_subscriptions" "test_partner_subscriptions" {
  #Required
  listing_id = var.listing_id

  #Optional
  display_name = var.partner_subscription_display_name
}

