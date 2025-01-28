// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "container_display_name" {
  default = "displayName"
}

variable "container_name" {
  default = "displayName"
}

variable "container_is_latest" {
  default = false
}

variable "container_state" {
  default = "AVAILABLE"
}

variable "container_tag_query_param" {
  default = "tagQueryParam"
}

variable "container_target_workload" {
  default = "MODEL_DEPLOYMENT"
}

variable "container_usage_query_param" {
  default = "INFERENCE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_datascience_containers" "test_container" {

  #Optional
  container_name    = var.container_name
  display_name      = var.container_display_name
  is_latest         = var.container_is_latest
  state             = var.container_state
  tag_query_param   = var.container_tag_query_param
  target_workload   = var.container_target_workload
  usage_query_param = var.container_usage_query_param
}

data "oci_datascience_containers" "test_containers" {

  #Optional
  container_name    = oci_datascience_containers.test_container.display_name
  display_name      = var.container_display_name
  is_latest         = var.container_is_latest
  state             = var.container_state
  tag_query_param   = var.container_tag_query_param
  target_workload   = var.container_target_workload
  usage_query_param = var.container_usage_query_param
}