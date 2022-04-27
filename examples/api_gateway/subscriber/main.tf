// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apigateway_usage_plan" "test_usage_plan" {
  compartment_id = var.compartment_ocid
  entitlements {
    name = "name"
  }
}

resource "oci_apigateway_subscriber" "test_subscriber" {
  #Required
  clients {
    #Required
    name  = "subscriber_clients_name"
    token = "subscriber_clients_token"
  }
  compartment_id = var.compartment_ocid
  usage_plans    = [oci_apigateway_usage_plan.test_usage_plan.id]

  #Optional
  display_name  = "subscriber_display_name"
  freeform_tags = { "Department" = "Finance" }
}

data "oci_apigateway_subscribers" "test_subscribers" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = "subscriber_display_name"
  state        = "ACTIVE"
}

data "oci_apigateway_subscriber" "test_subscriber" {
  #Required
  subscriber_id = oci_apigateway_subscriber.test_subscriber.id
}
