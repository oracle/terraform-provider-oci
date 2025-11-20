// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "auth" {}
variable "config_file_profile" {}
variable "region" {}
variable "resource_anchor_id" {}
variable "subscription_id" {}
variable "subscription_service_name" {}


provider "oci" {
  auth                  = var.auth
  config_file_profile   = var.config_file_profile
  region                = var.region
}

data "oci_multicloud_resource_anchor" "test_resource_anchor" {
  #Required
  resource_anchor_id        = var.resource_anchor_id
  subscription_id           = var.subscription_id
  subscription_service_name = var.subscription_service_name
}
