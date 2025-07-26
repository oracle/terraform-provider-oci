// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_ocid" {}
variable "config_file_profile" {}
variable "auth" {
  default = "SecurityToken"
}

variable "ssh_public_key" {
  default = ""
}

variable "db_name" {
  default = "TFDB"
}

variable "test_db_password" {
  default = "BEstrO0ng_#11"
}

variable "cpg_id" {
  default = null
}

variable "subscription_id" {
  default = null
}

variable "autoscale_limit_in_gbs" {
  default = null
}

variable "is_autoscale_enabled" {
  default = null
}