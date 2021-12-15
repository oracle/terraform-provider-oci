// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "gateway_display_name" {
  default = "displayName"
}

variable "gateway_endpoint_type" {
  default = "PUBLIC"
}

variable "gateway_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "ca_bundle_ocid" {}

variable "gateway_ca_bundles_type" {
  default = "CA_BUNDLE"
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

variable "deployment_specification_request_policies_mutual_tls_allowed_sans" {
  default = ["*.abc.com"]
}

variable "deployment_specification_request_policies_mutual_tls_is_verified_certificate_required" {
  default = true
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
  ca_bundles {
    #Required
    type = var.gateway_ca_bundles_type

    #Optional
    ca_bundle_id             = var.ca_bundle_ocid
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
    }
    request_policies {
      #Optional
      mutual_tls {
        #Optional
        allowed_sans                      = var.deployment_specification_request_policies_mutual_tls_allowed_sans
        is_verified_certificate_required  = var.deployment_specification_request_policies_mutual_tls_is_verified_certificate_required
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
