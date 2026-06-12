// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "config_file_profile" {
}

variable "region" {}
variable "compartment_id" {}

variable "hosted_application_storage_description" {
  default = "description"
}

variable "hosted_application_storage_display_name" {
  default = "displayName"
}

variable "hosted_application_storage_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "hosted_application_storage_hosted_application_storage_type" {
  default = "CACHE"
}

variable "hosted_application_storage_state" {
  default = "ACTIVE"
}

variable "hosted_application_storage_storage_type" {
  default = "CACHE"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_generative_ai_hosted_application_storage" "test_hosted_application_storage" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.hosted_application_storage_display_name
  storage_type   = var.hosted_application_storage_storage_type

  #Optional
  description   = var.hosted_application_storage_description
  freeform_tags = var.hosted_application_storage_freeform_tags
}

data "oci_generative_ai_hosted_application_storages" "test_hosted_application_storages" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                    = var.hosted_application_storage_display_name
  hosted_application_storage_type = var.hosted_application_storage_hosted_application_storage_type
  id                              = oci_generative_ai_hosted_application_storage.test_hosted_application_storage.id
  state                           = var.hosted_application_storage_state
}
