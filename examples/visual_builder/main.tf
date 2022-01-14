// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

variable "idcs_open_id" {
}

variable "vb_instance_consumption_model" {
  default = "UCM"
}

variable "custom_endpoint_certificate_secret_id" {
}

variable "custom_endpoint_host_name" {
  default = "hostname.com"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_visual_builder_vb_instance" "test_vb_instance" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "displayName"
  is_visual_builder_enabled = "true"
  idcs_open_id              = var.idcs_open_id
  node_count                = "1"

  #Optional
  consumption_model = var.vb_instance_consumption_model
  custom_endpoint {
    hostname = var.custom_endpoint_host_name
    certificate_secret_id = var.custom_endpoint_certificate_secret_id
  }
  freeform_tags = {
    "bar-key" = "value"
  }

  state                  = "ACTIVE"
}

data "oci_visual_builder_vb_instances" "test_vb_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = "displayName"
  state        = "Active"
  filter {
    name = "id"
    values = [oci_visual_builder_vb_instance.test_vb_instance.id]
  }
}

data "oci_visual_builder_vb_instance" "test_vb_instance" {
  #Required
  vb_instance_id = oci_visual_builder_vb_instance.test_vb_instance.id
}

data "oci_visual_builder_vb_instance_applications" "test_vb_instance_applications" {
  #Required
  vb_instance_id  = oci_visual_builder_vb_instance.test_vb_instance.id
  idcs_open_id    = var.idcs_open_id
}