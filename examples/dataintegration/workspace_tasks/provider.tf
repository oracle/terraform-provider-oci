// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
default = "us-ashburn-1"
}

variable "tenancy_ocid" {
default = ""
}

variable "user_ocid" {
default = ""
}

variable "fingerprint" {
default = ""
}

variable "private_key_path" {
default = ""
}



provider "oci" {
  region = var.region
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
}