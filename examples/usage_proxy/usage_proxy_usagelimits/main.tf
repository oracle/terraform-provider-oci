// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "region" {
}

variable "fingerprint" {
}

variable "tenancy_ocid" {
}

variable "compartment_ocid" {
}

variable "subscription_id" {
}

variable "user_ocid" {
}

variable "private_key_path" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_usage_proxy_usagelimits" "test_usagelimits" {
  #Required
  compartment_id = var.compartment_ocid
  subscription_id = var.subscription_id
}
