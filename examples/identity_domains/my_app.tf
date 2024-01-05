// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "my_app_my_app_count" {
  default = 10
}

variable "my_app_my_app_filter" {
  default = ""
}

variable "my_app_authorization" {
  default = "authorization"
}

variable "my_app_start_index" {
  default = 1
}


data "oci_identity_domains_my_apps" "test_my_apps" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain_for_my_endpoint.url

  #Optional
  my_app_count                 = var.my_app_my_app_count
  my_app_filter                = var.my_app_my_app_filter
  authorization                = var.my_app_authorization
  start_index                  = var.my_app_start_index
}

