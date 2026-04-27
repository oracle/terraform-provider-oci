// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "database_tools_endpoint_service_id" {}
variable "subnet_id" {}

variable "database_tools_private_endpoint_defined_tags_value" {
  default = "value"
}

variable "database_tools_private_endpoint_description" {
  default = "Private Endpoint for mySubnet"
}

variable "database_tools_private_endpoint_display_name" {
  default = "MyPE"
}

variable "database_tools_private_endpoint_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_private_endpoint_locks_message" {
  default = "message"
}

variable "database_tools_private_endpoint_locks_time_created" {
  default = "timeCreated"
}

variable "database_tools_private_endpoint_locks_type" {
  default = "FULL"
}

variable "database_tools_private_endpoint_nsg_ids" {
  default = []
}

variable "database_tools_private_endpoint_private_endpoint_ip" {
  default = "10.0.0.4"
}

variable "database_tools_private_endpoint_security_attributes" {
  default = { "Oracle-ZPR" = "MaxEgressCount=42" }
}

variable "database_tools_private_endpoint_state" {
  default = "ACTIVE"
}

provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_private_endpoint" "test_database_tools_private_endpoint" {
  #Required
  compartment_id      = var.compartment_id
  display_name        = var.database_tools_private_endpoint_display_name
  endpoint_service_id = var.database_tools_endpoint_service_id
  subnet_id           = var.subnet_id

  #Optional
  defined_tags  = {}
  description   = var.database_tools_private_endpoint_description
  freeform_tags = var.database_tools_private_endpoint_freeform_tags

  nsg_ids             = var.database_tools_private_endpoint_nsg_ids
  private_endpoint_ip = var.database_tools_private_endpoint_private_endpoint_ip
  security_attributes = var.database_tools_private_endpoint_security_attributes
}

data "oci_database_tools_database_tools_private_endpoints" "test_database_tools_private_endpoints" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name        = var.database_tools_private_endpoint_display_name
  endpoint_service_id = var.database_tools_endpoint_service_id
  state               = var.database_tools_private_endpoint_state
  subnet_id           = var.subnet_id
}

