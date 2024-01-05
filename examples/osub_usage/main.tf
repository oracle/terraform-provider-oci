// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "subscription_id" {}
variable "parent_product" {}
variable "computed_product" {}
variable "computed_usage_id" {}


variable "computed_usage_aggregated_grouping" {
  default = "MONTHLY"
}

variable "computed_usage_aggregated_time_from" {
  default = "2019-11-20T08:00:00Z"
}

variable "computed_usage_aggregated_time_to" {
  default = "2019-11-20T23:59:59Z"
}

variable "computed_usage_time_from" {
  default = "2019-11-22T08:00:00Z"
}

variable "computed_usage_time_to" {
  default = "2019-11-22T23:59:59Z"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osub_usage_computed_usage_aggregateds" "test_computed_usage_aggregateds" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id
  time_from       = var.computed_usage_aggregated_time_from
  time_to         = var.computed_usage_aggregated_time_to

  #Optional
  grouping = var.computed_usage_aggregated_grouping
  parent_product = var.parent_product
  x_one_origin_region = var.region
}

data "oci_osub_usage_computed_usages" "test_computed_usages" {
  #Required
  compartment_id  = var.compartment_id
  subscription_id = var.subscription_id
  time_from       = var.computed_usage_time_from
  time_to         = var.computed_usage_time_to

  #Optional
  computed_product = var.computed_product
  parent_product = var.parent_product
  x_one_origin_region = var.region
}

data "oci_osub_usage_computed_usage" "test_computed_usage" {
  #Required
  compartment_id  = var.compartment_id
  computed_usage_id = var.computed_usage_id

  #Optional
  x_one_origin_region = var.region
}
	
