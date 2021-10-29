// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "operator_action_name" {
  default = "name"
}

variable "operator_action_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "operator_action_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_operator_access_control_operator_actions" "test_operator_actions" {
  #Required
  compartment_id = var.compartment_id
  state         = var.operator_action_state

  #Optional
  name          = var.operator_action_name
  resource_type = var.operator_action_resource_type
  
}

