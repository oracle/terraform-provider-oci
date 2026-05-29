// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "recipient_tenancy_id" {
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

variable "recipient_invitation_id" {
  default = ""
}

variable "recipient_invitation_compartment_id" {
  default = ""
}

variable "compartment_ocid" {
  default = ""
}

locals {
  tenancy_ocid   = var.recipient_tenancy_id != "" ? var.recipient_tenancy_id : var.tenancy_ocid
  compartment_id = var.recipient_invitation_compartment_id != "" ? var.recipient_invitation_compartment_id : var.compartment_ocid
}

provider "oci" {
  tenancy_ocid        = local.tenancy_ocid
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

data "oci_tenantmanagercontrolplane_recipient_invitation" "recipient_invitation" {
  count                   = var.recipient_invitation_id != "" ? 1 : 0
  recipient_invitation_id = var.recipient_invitation_id
}

data "oci_tenantmanagercontrolplane_recipient_invitations" "recipient_invitations" {
  compartment_id = local.compartment_id
}
