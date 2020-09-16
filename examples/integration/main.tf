// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
  default = "idcsAt"
}

variable "integration_instance_consumption_model" {
  default = "UCM"
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
  display_name              = "displayName"
  integration_instance_type = "STANDARD"
  is_byol                   = "false"
  message_packs             = "10"

  #Optional
  consumption_model = "${var.integration_instance_consumption_model}"

  freeform_tags = {
    "bar-key" = "value"
  }

  idcs_at                = var.integration_instance_idcs_access_token
  is_file_server_enabled = true
  state                  = "ACTIVE"
}

data "oci_integration_integration_instances" "test_integration_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = "displayName"
  state        = "Active"
}

data "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}

