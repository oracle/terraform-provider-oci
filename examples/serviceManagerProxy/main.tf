// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "service_environment_display_name" {
  default = "displayName"
}

variable "service_environment_service_environment_type" {
  default = "serviceEnvironmentType"
}
variable "service_environment_id" {
  default = "123456"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_service_manager_proxy_service_environments" "test_service_environments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name             = var.service_environment_display_name
  service_environment_id   = var.service_environment_id
  service_environment_type = var.service_environment_service_environment_type
}


