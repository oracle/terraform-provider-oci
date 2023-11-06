// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cloud_gate_mapping_cloud_gate_mapping_count" {
  default = 10
}

variable "cloud_gate_mapping_cloud_gate_mapping_filter" {
  default = ""
}

variable "cloud_gate_mapping_authorization" {
  default = "authorization"
}

variable "cloud_gate_mapping_cloud_gate_value" {
  # add ID of CloudGate
  default = "value"
}

variable "cloud_gate_mapping_compartment_ocid" {
  default = "compartmentOcid"
}

variable "cloud_gate_mapping_description" {
  default = "description"
}

variable "cloud_gate_mapping_domain_ocid" {
  default = "domainOcid"
}

variable "cloud_gate_mapping_gateway_app_name" {
  # add name of App, if different
  default = "name"
}

variable "cloud_gate_mapping_gateway_app_value" {
  # add ID of App
  default = "value"
}

variable "cloud_gate_mapping_nginx_settings" {
  default = "nginxSettings"
}

variable "cloud_gate_mapping_policy_name" {
  default = "default"
}

variable "cloud_gate_mapping_proxy_pass" {
  default = "https://proxyPass.com"
}

variable "cloud_gate_mapping_resource_prefix" {
  default = "resourcePrefix"
}

variable "cloud_gate_mapping_server_value" {
  # add ID of CloudGateServer
  default = "value"
}

variable "cloud_gate_mapping_start_index" {
  default = 10
}

variable "cloud_gate_mapping_tags_key" {
  default = "key"
}

variable "cloud_gate_mapping_tags_value" {
  default = "value"
}

variable "cloud_gate_mapping_tenancy_ocid" {
  default = "tenancyOcid"
}

resource "oci_identity_domains_app" "test_cloud_gate_mapping_app" {
  #Required
  based_on_template {
    #Required
    value = "CustomWebAppTemplateId"
  }
  display_name  = "cloudGateMappingAppDisplayName"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:App"]

  lifecycle {
    ignore_changes = [schemas]
  }
}

resource "oci_identity_domains_cloud_gate" "test_cloud_gate_mapping_cloud_gate" {
  #Required
  display_name  = "cloudGateMappingName"
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGate"]
}

resource "oci_identity_domains_cloud_gate_server" "test_cloud_gate_mapping_server" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  cloud_gate {
    #Required
    value = oci_identity_domains_cloud_gate.test_cloud_gate_mapping_cloud_gate.id
  }
  display_name = "cloudGateMappingServerDisplayName"
  host_name    = "hostName"
  port         = 10
  schemas      = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGateServer"]
  ssl          = false
}

resource "oci_identity_domains_cloud_gate_mapping" "test_cloud_gate_mapping" {
  #Required
  cloud_gate {
    #Required
    value = oci_identity_domains_cloud_gate.test_cloud_gate_mapping_cloud_gate.id
  }
  gateway_app {
    #Required
    name  = oci_identity_domains_app.test_cloud_gate_mapping_app.name
    value = oci_identity_domains_app.test_cloud_gate_mapping_app.id
  }
  idcs_endpoint   = data.oci_identity_domain.test_domain.url
  policy_name     = var.cloud_gate_mapping_policy_name
  resource_prefix = var.cloud_gate_mapping_resource_prefix
  schemas         = ["urn:ietf:params:scim:schemas:oracle:idcs:CloudGateMapping"]
  server {
    #Required
    value = oci_identity_domains_cloud_gate_server.test_cloud_gate_mapping_server.id
  }

  #Optional
  attribute_sets = ["all"]
  attributes     = ""
  authorization  = var.cloud_gate_mapping_authorization
  description    = var.cloud_gate_mapping_description
  nginx_settings = var.cloud_gate_mapping_nginx_settings
  proxy_pass     = var.cloud_gate_mapping_proxy_pass
  tags {
    #Required
    key   = var.cloud_gate_mapping_tags_key
    value = var.cloud_gate_mapping_tags_value
  }
}

data "oci_identity_domains_cloud_gate_mappings" "test_cloud_gate_mappings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  cloud_gate_mapping_count  = var.cloud_gate_mapping_cloud_gate_mapping_count
  cloud_gate_mapping_filter = var.cloud_gate_mapping_cloud_gate_mapping_filter
  attribute_sets            = ["all"]
  attributes                = ""
  authorization             = var.cloud_gate_mapping_authorization
  start_index               = var.cloud_gate_mapping_start_index
}
