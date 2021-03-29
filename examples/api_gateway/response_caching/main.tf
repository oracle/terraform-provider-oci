// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "gateway_defined_tags_value" {
  default = "value"
}

variable "gateway_display_name" {
  default = "displayName"
}

variable "gateway_endpoint_type" {
  default = "PUBLIC"
}

variable "gateway_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "gateway_response_cache_details_authentication_secret_id" {}

variable "gateway_response_cache_details_authentication_secret_version_number" {
  default = 1
}

variable "gateway_response_cache_details_connect_timeout_in_ms" {
  default = 1000
}

variable "gateway_response_cache_details_is_ssl_enabled" {
  default = false
}

variable "gateway_response_cache_details_is_ssl_verify_disabled" {
  default = false
}

variable "gateway_response_cache_details_read_timeout_in_ms" {
  default = 1000
}

variable "gateway_response_cache_details_send_timeout_in_ms" {
  default = 1000
}

variable "gateway_response_cache_details_servers_host" {
  default = "host"
}

variable "gateway_response_cache_details_servers_port" {
  default = 6379
}

variable "gateway_response_cache_details_type" {
  default = "EXTERNAL_RESP_CACHE"
}

variable "gateway_state" {
  default = "ACTIVE"
}

variable "deployment_state" {
  default = "ACTIVE"
}

variable "deployment_path_prefix" {
  default = "/v1"
}

variable "deployment_specification_routes_backend_type" {
  default = "HTTP_BACKEND"
}

variable "deployment_specification_routes_backend_url" {
  default = "https://api.weather.gov"
}

variable "deployment_specification_routes_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_path" {
  default = "/hello"
}

variable "deployment_specification_routes_response_policies_response_cache_store_time_to_live_in_seconds" {
  default = 10
}

variable "deployment_specification_routes_response_policies_response_cache_store_type" {
  default = "FIXED_TTL_STORE_POLICY"
}

variable "deployment_specification_routes_request_policies_response_cache_lookup_cache_key_additions" {
  default = ["request.headers[Accept]"]
}

variable "deployment_specification_routes_request_policies_response_cache_lookup_is_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_response_cache_lookup_is_private_caching_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_response_cache_lookup_type" {
  default = "SIMPLE_LOOKUP_POLICY"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_subnet" "regional_subnet" {
  cidr_block        = "10.0.1.0/24"
  display_name      = "regionalSubnet"
  dns_label         = "regionalsubnet"
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.vcn1.id
  security_list_ids = [oci_core_vcn.vcn1.default_security_list_id]
  route_table_id    = oci_core_vcn.vcn1.default_route_table_id
  dhcp_options_id   = oci_core_vcn.vcn1.default_dhcp_options_id
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_apigateway_gateway" "test_gateway" {
  #Required
  compartment_id = var.compartment_ocid
  endpoint_type  = var.gateway_endpoint_type
  subnet_id      = oci_core_subnet.regional_subnet.id

  #Optional
  display_name   = var.gateway_display_name
  freeform_tags  = var.gateway_freeform_tags
  response_cache_details {
    #Required
    type = var.gateway_response_cache_details_type

    #Optional
    authentication_secret_id             = var.gateway_response_cache_details_authentication_secret_id
    authentication_secret_version_number = var.gateway_response_cache_details_authentication_secret_version_number
    connect_timeout_in_ms                = var.gateway_response_cache_details_connect_timeout_in_ms
    is_ssl_enabled                       = var.gateway_response_cache_details_is_ssl_enabled
    is_ssl_verify_disabled               = var.gateway_response_cache_details_is_ssl_verify_disabled
    read_timeout_in_ms                   = var.gateway_response_cache_details_read_timeout_in_ms
    send_timeout_in_ms                   = var.gateway_response_cache_details_send_timeout_in_ms
    servers {

      #Optional
      host = var.gateway_response_cache_details_servers_host
      port = var.gateway_response_cache_details_servers_port
    }
  }
}

resource "oci_apigateway_deployment" "test_deployment" {
  #Required
  compartment_id = var.compartment_ocid
  gateway_id     = oci_apigateway_gateway.test_gateway.id
  path_prefix    = var.deployment_path_prefix

  specification {
    routes {
      #Required
      backend {
        #Required
        type = var.deployment_specification_routes_backend_type
        url  = var.deployment_specification_routes_backend_url
      }
      path = var.deployment_specification_routes_path
      methods = var.deployment_specification_routes_methods
      request_policies {
        response_cache_lookup {
          #Required
          type = var.deployment_specification_routes_request_policies_response_cache_lookup_type

          #Optional
          cache_key_additions        = var.deployment_specification_routes_request_policies_response_cache_lookup_cache_key_additions
          is_enabled                 = var.deployment_specification_routes_request_policies_response_cache_lookup_is_enabled
          is_private_caching_enabled = var.deployment_specification_routes_request_policies_response_cache_lookup_is_private_caching_enabled
        }
      }
      response_policies {
        response_cache_store {
          #Required
          time_to_live_in_seconds = var.deployment_specification_routes_response_policies_response_cache_store_time_to_live_in_seconds
          type                    = var.deployment_specification_routes_response_policies_response_cache_store_type
        }
      }
    }
  }
}

data "oci_apigateway_gateways" "test_gateways" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name   = oci_apigateway_gateway.test_gateway.display_name
  state          = var.gateway_state
}

data "oci_apigateway_deployments" "test_deployments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  gateway_id     = oci_apigateway_gateway.test_gateway.id
  state          = var.deployment_state
}
