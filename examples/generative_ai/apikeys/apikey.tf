// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {
  default = "SecurityToken"
}

variable "config_file_profile" {}
variable "region" {}
variable "compartment_ocid" {}

variable "apikey_display_name" {}
variable "apikey_id" {}

provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

data "oci_generative_ai_apikeys" "test_apikeys" {
  compartment_id = var.compartment_ocid
  display_name   = var.apikey_display_name
  id             = var.apikey_id
  state          = "ACTIVE"
}
