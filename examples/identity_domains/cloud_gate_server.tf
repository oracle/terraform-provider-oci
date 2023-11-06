// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cloud_gate_server_cloud_gate_server_count" {
  default = 10
}

variable "cloud_gate_server_cloud_gate_server_filter" {
  default = ""
}

variable "cloud_gate_server_authorization" {
  default = "authorization"
}

variable "cloud_gate_server_cloud_gate_value" {
  # add ID of CloudGate to use with server
  default = "value"
}

variable "cloud_gate_server_description" {
  default = "description"
}

variable "cloud_gate_server_display_name" {
  default = "displayName"
}

variable "cloud_gate_server_host_name" {
  default = "hostName"
}

variable "cloud_gate_server_nginx_settings" {
  default = "nginxSettings"
}

variable "cloud_gate_server_ocid" {
  default = "ocid"
}

variable "cloud_gate_server_port" {
  default = 10
}

variable "cloud_gate_server_ssl" {
  default = false
}

variable "cloud_gate_server_start_index" {
  default = 1
}

variable "cloud_gate_server_tags_key" {
  default = "key"
}

variable "cloud_gate_server_tags_value" {
  default = "value"
}

variable "cloud_gate_server_tenancy_ocid" {
  default = "tenancyOcid"
}

resource "oci_identity_domains_cloud_gate" "test_cloud_gate_server_cloud_gate" {
  #Required
  display_name  = "name"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGate"]
}

resource "oci_identity_domains_cloud_gate_server" "test_cloud_gate_server" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  cloud_gate {
    #Required
    value = oci_identity_domains_cloud_gate.test_cloud_gate_server_cloud_gate.id
  }
  display_name = var.cloud_gate_server_display_name
  host_name    = var.cloud_gate_server_host_name
  port         = var.cloud_gate_server_port
  schemas      = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGateServer"]
  ssl          = var.cloud_gate_server_ssl

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.cloud_gate_server_authorization
  description    = var.cloud_gate_server_description
  nginx_settings = var.cloud_gate_server_nginx_settings
  tags {
    #Required
    key   = var.cloud_gate_server_tags_key
    value = var.cloud_gate_server_tags_value
  }
}

data "oci_identity_domains_cloud_gate_servers" "test_cloud_gate_servers" {
  #Required  
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  cloud_gate_server_count  = var.cloud_gate_server_cloud_gate_server_count
  cloud_gate_server_filter = var.cloud_gate_server_cloud_gate_server_filter
  attribute_sets           = ["all"]
  attributes               = ""
  authorization            = var.cloud_gate_server_authorization
  start_index              = var.cloud_gate_server_start_index
}
