// Copyright (c) 2017, 2022, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "standby_tenancy_ocid" {
}

variable "user_ocid" {
}

variable "standby_user_ocid" {
}

variable "fingerprint" {
}

variable "standby_fingerprint" {
}

variable "private_key_path" {
}

variable "standby_private_key_path" {
}

variable "compartment_id" {
}

variable "standby_compartment_id" {
}

variable "standby_region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

provider "oci" {
  alias            = "standby_tenancy"
  region           = var.standby_region
  tenancy_ocid     = var.standby_tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.standby_fingerprint
  private_key_path = var.standby_private_key_path
}