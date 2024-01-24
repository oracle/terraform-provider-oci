// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_device_my_device_count" {
  default = 10
}

variable "my_device_my_device_filter" {
  default = ""
}

variable "my_device_authorization" {
  default = "authorization"
}

variable "my_device_start_index" {
  default = 1
}


data "oci_identity_domains_my_devices" "test_my_devices" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_device_count  = var.my_device_my_device_count
  my_device_filter = var.my_device_my_device_filter
  attribute_sets   = []
  attributes       = ""
  authorization    = var.my_device_authorization
  start_index      = var.my_device_start_index
}

