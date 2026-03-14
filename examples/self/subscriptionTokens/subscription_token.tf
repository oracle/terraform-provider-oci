// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {
  default = "tenancy_ocid"
}
variable "user_ocid" {
  default = "user_id"
}
variable "fingerprint" {
  default = ""
}
variable "private_key_path" {
  default = "private_key"
}
variable "region" {
  default = "region"
}

variable "subscriptionId" {
  default = "subscription_id"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_self_subscription_token" "test_subscription_token" {
  #Required
  subscription_id = var.subscriptionId
}

