// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "delegation_subscription_defined_tags_value" {
  default = "value"
}

variable "delegation_subscription_description" {
  default = "description"
}

variable "delegation_subscription_display_name" {
  default = "displayName"
}

variable "delegation_subscription_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "delegation_subscription_state" {
  default = "ACTIVE"
}

variable "delegation_subscription_subscribed_service_type" {
  default = "TROUBLESHOOTING"
}

variable "test_service_provider_id" {
}

variable "root_compartment_id" {
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_delegation_management_delegation_subscription" "test_delegation_subscription" {
  #Required
  compartment_id          = var.root_compartment_id
  service_provider_id     = var.test_service_provider_id
  subscribed_service_type = var.delegation_subscription_subscribed_service_type

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.delegation_subscription_defined_tags_value)
  description   = var.delegation_subscription_description
  #freeform_tags = var.delegation_subscription_freeform_tags
}

data "oci_delegation_management_delegation_subscriptions" "test_delegation_subscriptions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.delegation_subscription_display_name
  state        = var.delegation_subscription_state
}