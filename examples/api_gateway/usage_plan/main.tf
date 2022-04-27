// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apigateway_gateway" "test_gateway" {
  compartment_id = var.compartment_ocid
  endpoint_type  = "PUBLIC"
  subnet_id      = oci_core_subnet.regional_subnet.id
}

resource "oci_core_vcn" "vcn1" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "regional_subnet" {
  cidr_block     = "10.0.1.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn1.id
}

resource "oci_apigateway_deployment" "test_deployment" {
  compartment_id = var.compartment_ocid
  gateway_id     = oci_apigateway_gateway.test_gateway.id
  path_prefix    = "/"

  specification {
    request_policies {
      #Optional
      usage_plans {
        #Required
        token_locations = ["request.headers[apiKeyLocation]"]
      }
    }
    routes {
      backend {
        type = "HTTP_BACKEND"
        url  = "https://api.weather.gov"
      }
      path    = "/hello"
      methods = ["GET"]
    }
  }
}

resource "oci_apigateway_usage_plan" "test_usage_plan" {
  #Required
  compartment_id = var.compartment_ocid
  entitlements {
    #Required
    name = "usagePlanEntitlementsName"

    #Optional
    description = "usage_plan_entitlements_description"
    quota {
      #Required
      operation_on_breach = "REJECT"
      reset_policy        = "CALENDAR"
      unit                = "MINUTE"
      value               = 10
    }
    rate_limit {
      #Required
      unit  = "SECOND"
      value = 10
    }
    targets {
      #Required
      deployment_id = oci_apigateway_deployment.test_deployment.id
    }
  }

  #Optional
  display_name  = "usage_plan_display_name"
  freeform_tags = { "Department" = "Finance" }
}

data "oci_apigateway_usage_plans" "test_usage_plans" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = "usage_plan_display_name"
  state        = "ACTIVE"
}

data "oci_apigateway_usage_plan" "test_usage_plan" {
  #Required
  usage_plan_id = oci_apigateway_usage_plan.test_usage_plan.id
}