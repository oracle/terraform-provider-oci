// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_osp_gateway_subscriptions" "test_subscriptions" {
  compartment_id = var.tenancy_ocid
  osp_home_region = var.region
}

data "oci_osp_gateway_subscription" "test_subscription" {
  compartment_id = var.tenancy_ocid
  osp_home_region = var.region
  subscription_id = "${lookup(data.oci_osp_gateway_subscriptions.test_subscriptions.subscription_collection.0.items[0], "id")}"
}
