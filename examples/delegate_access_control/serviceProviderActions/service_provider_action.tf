// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "service_provider_action_name" {
  default = "name"
}

variable "service_provider_action_resource_type" {
  default = "VMCLUSTER"
}

variable "service_provider_action_service_provider_service_type" {
  default = []
}

variable "service_provider_action_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_delegate_access_control_service_provider_actions" "test_service_provider_actions" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  name                          = var.service_provider_action_name
  resource_type                 = var.service_provider_action_resource_type
  service_provider_service_type = var.service_provider_action_service_provider_service_type
  state                         = var.service_provider_action_state
}

