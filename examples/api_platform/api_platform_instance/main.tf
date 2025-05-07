// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "api_platform_instance_description" {
  default = "description"
}

variable "api_platform_instance_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "api_platform_instance_name" {
  default = "name"
}

variable "api_platform_instance_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_api_platform_api_platform_instance" "test_api_platform_instance" {
  #Required
  compartment_id = var.compartment_id
  name = var.api_platform_instance_name

  #Optional
  defined_tags = { "example-tag-namespace-all.example-tag" = "value" }
  description   = var.api_platform_instance_description
  freeform_tags = var.api_platform_instance_freeform_tags
}

data "oci_api_platform_api_platform_instances" "test_api_platform_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id    = oci_api_platform_api_platform_instance.test_api_platform_instance.id
  name  = var.api_platform_instance_name
  state = var.api_platform_instance_state
}

data "oci_api_platform_api_platform_instance" "test_api_platform_instance_data" {
  #Required
  api_platform_instance_id = oci_api_platform_api_platform_instance.test_api_platform_instance.id
}