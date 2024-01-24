// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "time_start" {}

variable "time_end" {}

variable "summary_contains" {
  default = "example summary value"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_announcements" "test_jms_announcements" {
	#Optional
	summary_contains = var.summary_contains
	time_start = var.time_start
	time_end = var.time_end
}
