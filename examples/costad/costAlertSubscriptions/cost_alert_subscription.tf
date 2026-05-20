// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "cost_alert_subscription_channels" {
  default = "{\"email\":\"test@example.com\"}"
}

variable "cost_alert_subscription_defined_tags_value" {
  default = "value"
}

variable "cost_alert_subscription_description" {
  default = "description"
}

variable "cost_alert_subscription_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cost_alert_subscription_name" {
  default = "Department A email list"
}

variable "cost_alert_subscription_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_costad_cost_alert_subscription" "test_cost_alert_subscription" {
  #Required
  channels       = var.cost_alert_subscription_channels
  compartment_id = var.compartment_id
  name           = var.cost_alert_subscription_name

  #Optional
  description   = var.cost_alert_subscription_description
  freeform_tags = var.cost_alert_subscription_freeform_tags
}

data "oci_costad_cost_alert_subscriptions" "test_cost_alert_subscriptions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name  = var.cost_alert_subscription_name
  state = var.cost_alert_subscription_state
}
