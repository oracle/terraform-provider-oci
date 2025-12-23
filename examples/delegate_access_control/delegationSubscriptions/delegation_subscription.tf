// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "delegation_serviceprovider_ocid" {}


variable "delegation_subscription_display_name" {
  default = "TersiTest"
}

variable "delegation_subscription_state" {
  default = "ACTIVE"
}

variable "delegation_subscription_subscribed_service_type" {
  default = "TROUBLESHOOTING"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_delegate_access_control_delegation_subscription" "test_delegation_subscription" {
  #Required
  compartment_id          = var.compartment_ocid
  service_provider_id     = var.delegation_serviceprovider_ocid
  subscribed_service_type = var.delegation_subscription_subscribed_service_type
}

data "oci_delegate_access_control_delegation_subscriptions" "test_delegation_subscriptions" {
  #Required
  compartment_id = var.compartment_ocid
}

