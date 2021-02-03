// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "is_repository_created_on_first_push" {
  default = true
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_artifacts_container_configuration" "test_container_configuration" {
  #Required
  compartment_id                      = var.tenancy_ocid
  is_repository_created_on_first_push = var.is_repository_created_on_first_push
}

data "oci_artifacts_container_configuration" "test_container_configuration" {
  #Required
  compartment_id = var.tenancy_ocid
}
