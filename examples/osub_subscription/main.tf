// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscribed_service_id" {}
variable "commitment_id" {}
variable "subscription_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osub_subscription_commitments" "test_commitments" {
  #Required
  compartment_id        = var.compartment_id
  subscribed_service_id = var.subscribed_service_id
  x_one_origin_region   = var.region

  #Optional
  #x_one_gateway_subscription_id = var.commitment_x_one_gateway_subscription_id
}

data "oci_osub_subscription_commitment" "test_commitment" {
  #Required
  commitment_id = var.commitment_id

  #Optional
  x_one_origin_region = var.region
}

data "oci_osub_subscription_ratecards" "test_ratecards" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id

  #Optional
  #part_number         = var.ratecard_part_number
  #time_from           = var.ratecard_time_from
  #time_to             = var.ratecard_time_to
  x_one_origin_region = var.region
}

data "oci_osub_subscription_subscriptions" "test_subscriptions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #buyer_email                   = var.subscription_buyer_email
  #is_commit_info_required       = var.subscription_is_commit_info_required
  #plan_number                   = var.subscription_plan_number
  subscription_id                = var.subscription_id
  #x_one_gateway_subscription_id = var.subscription_x_one_gateway_subscription_id
  x_one_origin_region            = var.region
}
