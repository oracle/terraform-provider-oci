// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "sender_tenancy_id" {
  default = ""
}

variable "tenancy_ocid" {
  default = ""
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "sender_invitation_id" {
  default = ""
}

variable "sender_invitation_compartment_id" {
  default = ""
}

variable "compartment_ocid" {
  default = ""
}

locals {
  tenancy_ocid   = var.sender_tenancy_id != "" ? var.sender_tenancy_id : var.tenancy_ocid
  compartment_id = var.sender_invitation_compartment_id != "" ? var.sender_invitation_compartment_id : var.compartment_ocid
}

provider "oci" {
  tenancy_ocid        = local.tenancy_ocid
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

data "oci_tenantmanagercontrolplane_sender_invitation" "sender_invitation" {
  count                = var.sender_invitation_id != "" ? 1 : 0
  sender_invitation_id = var.sender_invitation_id
}

data "oci_tenantmanagercontrolplane_sender_invitations" "sender_invitations" {
  compartment_id = local.compartment_id
}
