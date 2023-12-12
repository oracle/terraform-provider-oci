// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cloud_gate_cloud_gate_count" {
  default = 10
}

variable "cloud_gate_cloud_gate_filter" {
  default = ""
}

variable "cloud_gate_active" {
  default = false
}

variable "cloud_gate_authorization" {
  default = "authorization"
}

variable "cloud_gate_description" {
  default = "description"
}

variable "cloud_gate_display_name" {
  default = "cloudGateDisplayName"
}

variable "cloud_gate_last_modified_time" {
  default = "2000-01-01T00:00:00Z"
}

variable "cloud_gate_start_index" {
  default = 1
}

variable "cloud_gate_tags_key" {
  default = "key"
}

variable "cloud_gate_tags_value" {
  default = "value"
}

variable "cloud_gate_tenancy_ocid" {
  default = "tenancyOcid"
}

variable "cloud_gate_type" {
  default = "lbaas"
}

resource "oci_identity_domains_cloud_gate" "test_cloud_gate" {
  #Required
  display_name  = var.cloud_gate_display_name
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGate"]

  #Optional
  active                       = var.cloud_gate_active
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.cloud_gate_authorization
  description                  = var.cloud_gate_description
  last_modified_time           = var.cloud_gate_last_modified_time
  tags {
    #Required
    key   = var.cloud_gate_tags_key
    value = var.cloud_gate_tags_value
  }
  type = var.cloud_gate_type
}

data "oci_identity_domains_cloud_gates" "test_cloud_gates" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  cloud_gate_count             = var.cloud_gate_cloud_gate_count
  cloud_gate_filter            = var.cloud_gate_cloud_gate_filter
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.cloud_gate_authorization
  start_index                  = var.cloud_gate_start_index
}

