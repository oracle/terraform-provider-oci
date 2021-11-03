// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "region" {
}

variable "fingerprint" {
}

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "private_key_path" {
}

variable "subscription_id" {
  default = "7368239"
}

variable "email_id" {
  default = "test@gmail.com"
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_usage_proxy_subscription_redeemable_user" "test_subscription_redeemable_user" {
  #Required
  items {
    email_id = var.email_id
  }
  subscription_id = var.subscription_id
  tenancy_id = var.tenancy_ocid
}

