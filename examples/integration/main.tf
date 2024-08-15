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

variable "region" {
}

variable "compartment_id" {
}

variable "integration_instance_idcs_access_token" {
  default = ""
}

variable "integration_instance_consumption_model" {
  default = "UCM"
}

variable allow_listed_http_vcn {
  default = ""
}

variable certificate_secret_id {
  default = ""
}

variable domain_id {
  default = ""
}

variable subnet_id {
  default = ""
}

variable nsg_id {
  default = ""
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  compartment_id            = var.compartment_id
  integration_instance_type = "HEALTHCARE"
  shape                     = "DEVELOPMENT"
  display_name              = "DataRetention"
  is_byol                   = "false"
  message_packs             = "10"
  domain_id                 = var.domain_id
  # idcs_at                 = var.integration_instance_idcs_access_token

  #Optional
# For stand / enterprise type only
#  consumption_model = "${var.integration_instance_consumption_model}"
#  custom_endpoint {
#    certificate_secret_id = var.certificate_secret_id
#    hostname = "hostname.com"
#  }
#  freeform_tags = {
#    "bar-key" = "value"
#  }

#  is_file_server_enabled = true
#  is_visual_builder_enabled = true
#  state                  = "ACTIVE"
}

data "oci_integration_integration_instances" "test_integration_instances" {
  #Required
  compartment_id = var.compartment_id

  display_name = "displayName"
  state        = "Active"
}

data "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}

# resource "oci_integration_integration_instance" "test_integration_instance_idcs" {
#   #Required
#   compartment_id            = var.compartment_id
#   display_name              = "instance4643_idcs"
#   integration_instance_type = "STANDARDX"
#   shape                     = "DEVELOPMENT"
#   # shape                     = "PRODUCTION"
#   is_byol                   = "false"
#   message_packs             = "10"
#   idcs_at                   = var.integration_instance_idcs_access_token
# }

resource "oci_integration_private_endpoint_outbound_connection" "integration_private_endpoint" {
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
  nsg_ids = [var.nsg_id]
  subnet_id = var.subnet_id
}
