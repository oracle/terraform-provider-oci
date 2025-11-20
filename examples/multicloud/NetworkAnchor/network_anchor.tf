// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "network_anchor_id" {}
variable "subscription_id" {}
variable "subscription_service_name" {}


provider "oci" {
  auth                  = var.auth
  config_file_profile   = var.config_file_profile
  region                = var.region
}

data "oci_multicloud_network_anchor" "test_network_anchor" {
  #Required
  network_anchor_id         = var.network_anchor_id
  subscription_id           = var.subscription_id
  subscription_service_name = var.subscription_service_name
}