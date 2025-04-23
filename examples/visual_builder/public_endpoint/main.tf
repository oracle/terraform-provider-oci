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

variable "idcs_open_id" {
}

variable "vb_instance_consumption_model" {
  default = "UCM"
}

variable "custom_endpoint_certificate_secret_id" {
}

variable "vb_instance_network_endpoint_details_allowlisted_http_ips" {
  default = ["0.0.0.0/32"]
}

variable "vb_instance_network_endpoint_details_allowlisted_http_vcns_allowlisted_ip_cidrs" {
  default = []
}

variable "vb_instance_network_endpoint_details_allowlisted_http_vcns_id" {
  default = "id"
}

variable "vb_instance_network_endpoint_details_network_endpoint_type" {
  default = "PUBLIC"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_visual_builder_vb_instance" "test_vb_instance_acl" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "displayName"
  is_visual_builder_enabled = "true"
  idcs_open_id              = var.idcs_open_id
  node_count                = "1"

  #Optional
  consumption_model = var.vb_instance_consumption_model
  #Optional
  network_endpoint_details {
    #Required
    network_endpoint_type = var.vb_instance_network_endpoint_details_network_endpoint_type

    #Optional
    allowlisted_http_ips = var.vb_instance_network_endpoint_details_allowlisted_http_ips
    allowlisted_http_vcns {
      #Required
      id = var.vb_instance_network_endpoint_details_allowlisted_http_vcns_id

      #Optional
      allowlisted_ip_cidrs = var.vb_instance_network_endpoint_details_allowlisted_http_vcns_allowlisted_ip_cidrs
    }
  }
}

data "oci_visual_builder_vb_instances" "test_vb_instances_acl" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = "displayName"
  state        = "Active"
  filter {
    name = "id"
    values = [oci_visual_builder_vb_instance.test_vb_instance_acl.id]
  }
}

data "oci_visual_builder_vb_instance" "test_vb_instance_acl" {
  #Required
  vb_instance_id = oci_visual_builder_vb_instance.test_vb_instance_acl.id
}

data "oci_visual_builder_vb_instance_applications" "test_vb_instance_applications" {
  #Required
  vb_instance_id  = oci_visual_builder_vb_instance.test_vb_instance_acl.id
  idcs_open_id    = var.idcs_open_id
}
