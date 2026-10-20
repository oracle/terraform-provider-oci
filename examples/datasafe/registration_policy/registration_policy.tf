# Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
# Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

# The database or CDB OCID for which Data Safe should create the policy.
variable "registration_policy_resource_id" {}

variable "registration_policy_features" {
  type    = list(string)
  default = ["ASSESSMENT"]
}

variable "registration_policy_display_name" {
  default = "example-registration-policy"
}

variable "registration_policy_description" {
  default = "Registration policy managed by Terraform"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Create a Data Safe registration policy for the supplied database resource.
resource "oci_data_safe_registration_policy" "example" {
  compartment_id = var.compartment_ocid
  resource_id    = var.registration_policy_resource_id
  features       = var.registration_policy_features
  display_name   = var.registration_policy_display_name
  description    = var.registration_policy_description
}

# Read the policy back by its OCID.
data "oci_data_safe_registration_policy" "example" {
  registration_policy_id = oci_data_safe_registration_policy.example.id
}

# List registration policies in the compartment.
data "oci_data_safe_registration_policies" "example" {
  compartment_id = var.compartment_ocid
}
