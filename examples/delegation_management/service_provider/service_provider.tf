// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "service_provider_name" {
  default = "name"
}

variable "service_provider_service_provider_type" {
  default = "ORACLE_PROVIDED"
}

variable "service_provider_state" {
  default = "ACTIVE"
}

variable "service_provider_supported_resource_type" {
  default = "VMCLUSTER"
}

variable "root_compartment_id" {
  default = "ocid1.tenancy.region1..aaaaaaaagyw5okosjg54csr3u5qgaxvtjufz55537h44mjy2umiqur4z5w3a"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_delegation_management_service_providers" "test_service_providers" {
  #Required
  compartment_id = var.root_compartment_id

  #Optional
  name                    = var.service_provider_name
  service_provider_type   = var.service_provider_service_provider_type
  state                   = var.service_provider_state
  supported_resource_type = var.service_provider_supported_resource_type
}