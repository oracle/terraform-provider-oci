// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "security_policy_ocid" {}
variable "data_safe_target_ocid" {}

variable "description" {
  default = "description"
}

variable "display_name" {
  default = "security_policy_deployment_updated"
}

variable "security_policy_deployment_access_level" {
  default = "ACCESSIBLE"
}

variable "security_policy_deployment_compartment_id_in_subtree" {
  default = false
}

variable "target_type" {
  default = "TARGET_DATABASE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_security_policy_deployment" "test_security_policy_deployment" {
  #Required
  compartment_id     = var.compartment_ocid
  security_policy_id = var.security_policy_ocid
  target_id          = var.data_safe_target_ocid
  target_type        = var.target_type

  #Optional
  description  = var.description
  display_name = var.display_name
}

data "oci_data_safe_security_policy_deployments" "test_security_policy_deployments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  security_policy_deployment_id = oci_data_safe_security_policy_deployment.test_security_policy_deployment.id
  access_level                  = var.security_policy_deployment_access_level
  compartment_id_in_subtree     = var.security_policy_deployment_compartment_id_in_subtree
}