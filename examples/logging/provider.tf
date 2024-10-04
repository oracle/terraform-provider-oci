// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file shows how to configure the oci provider to target a single region.
 */

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1..aaaaaaaa4s2hncj4oaulmf5tz4yfeska6fya4gkd5jsg3fmlgq7pprgr7wiq"
}

variable "user_ocid" {
  default = "ocid1.user.oc1..aaaaaaaark6yo7jgevogxohlgerphpr6lreunmmsovjdkhmujnuj2urix5aq"
}

variable "fingerprint" {
  default = "16:9a:cf:f4:78:3f:ba:fd:67:fc:74:30:72:e8:e7:11"
}

variable "private_key_path" {
  default = "/Users/zhenyao/.oci/oci_api_key.pem"
}

variable "compartment_ocid" {
  default = "ocid1.compartment.oc1..aaaaaaaajdgiuoxrwem3326sihqitq3rf62hg4bq255hzchqwszx2xz4zega"
}

variable "region" {
  default = "us-phoenix-1"
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  # https://registry.terraform.io/providers/oracle/oci/latest/docs
#  version = "5.32.0"
}
