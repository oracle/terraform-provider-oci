// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// The workflow supplies required_providers_block.tf for its example test.
// For standalone use, copy required_providers_block.tf.example to that filename.

provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  ignore_defined_tags = [
    "Oracle-Tags.CreatedBy",
    "Oracle-Tags.CreatedOn",
  ]
}

variable "region" {
  type        = string
  description = "OCI region with the PSS service enabled."
}

variable "config_file_profile" {
  type        = string
  description = "OCI CLI profile containing a current security token."
  default     = "terraform-federation-test"
}

variable "compartment_ocid" {
  type        = string
  description = "Compartment in which to create the PSS instance."
}

variable "pss_identity_domain_url" {
  type        = string
  description = "URL of the identity domain associated with the PSS instance."
}

variable "pss_admin_identity_domain_user_id" {
  type        = string
  description = "OCID of an existing administrator user in the selected identity domain."
}

variable "display_name" {
  type    = string
  default = "pss-terraform-example"
}

resource "oci_pss_instance" "example" {
  compartment_id                = var.compartment_ocid
  identity_domain_url           = var.pss_identity_domain_url
  admin_identity_domain_user_id = var.pss_admin_identity_domain_user_id
  display_name                  = var.display_name
  freeform_tags                 = { "Example" = "PSS" }
}

data "oci_pss_instance" "example" {
  instance_id = oci_pss_instance.example.id
}

data "oci_pss_instances" "example" {
  compartment_id = var.compartment_ocid
  id             = oci_pss_instance.example.id
  state          = "ACTIVE"
}

output "instance_id" {
  value = oci_pss_instance.example.id
}

output "service_endpoint" {
  value = data.oci_pss_instance.example.service_endpoint
}

output "ui_endpoint" {
  value = data.oci_pss_instance.example.ui_endpoint
}

output "instances" {
  value = data.oci_pss_instances.example.instance_collection
}
