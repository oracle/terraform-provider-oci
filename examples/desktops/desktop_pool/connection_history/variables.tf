// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "desktop_pool_id" {
  description = "Desktop Pool OCID"
  #default = "<desktop_pool_ocid>"
}

variable "desktop_id" {
  description = "Desktop OCID"
  #default = "<desktop_ocid>"
}
