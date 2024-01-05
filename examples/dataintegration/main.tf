// Copyright (c) 2017, 2024, 2020, Oracle and/or its affiliates. All rights reserved.

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "workspace_defined_tags_value" {
}

variable "workspace_description" {
  default = "description"
}

variable "workspace_display_name" {
  default = "displayName"
}

variable "workspace_dns_server_ip" {
}

variable "workspace_dns_server_zone" {
}

variable "workspace_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "workspace_is_private_network_enabled" {
  default = false
}

variable "workspace_state" {
  default = "ACTIVE"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  dns_label      = "dnslabel"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  dns_label      = "dnslabel"
}

resource "oci_dataintegration_workspace" "test_workspace" {
  #Required
  display_name = var.workspace_display_name

  #Optional
  compartment_id = var.compartment_ocid
  /*defined_tags = {
    oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.workspace_defined_tags_value
  }*/
  description                = var.workspace_description
  dns_server_ip              = var.workspace_dns_server_ip
  dns_server_zone            = var.workspace_dns_server_zone
  freeform_tags              = var.workspace_freeform_tags
  is_private_network_enabled = var.workspace_is_private_network_enabled

  // removed for backward compatibility, got approval from TF team
#  subnet_id                  = oci_core_subnet.test_subnet.id
#  vcn_id                     = oci_core_vcn.test_vcn.id
}

data "oci_dataintegration_workspaces" "test_workspaces" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  state = var.workspace_state
}

data "oci_dataintegration_workspace" "test_workspace" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id
}

