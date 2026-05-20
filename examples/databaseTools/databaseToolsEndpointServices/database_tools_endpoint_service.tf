// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "database_tools_endpoint_service_display_name" {
  default = "displayName"
}

variable "database_tools_endpoint_service_name" {
  default = "name"
}

variable "database_tools_endpoint_service_state" {
  default = "ACTIVE"
}

provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}


data "oci_database_tools_database_tools_endpoint_services" "test_database_tools_endpoint_services" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.database_tools_endpoint_service_display_name
  name         = var.database_tools_endpoint_service_name
  state        = var.database_tools_endpoint_service_state
}

