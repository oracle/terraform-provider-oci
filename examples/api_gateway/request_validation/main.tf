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

resource "oci_apigateway_deployment" "test_deployment" {
  compartment_id = var.compartment_ocid
  gateway_id = oci_apigateway_gateway.test_gateway.id
  path_prefix = "/"

  specification {
    # An example route making use of a request validation policy
    # requiring a header parameter.
    routes {
      path = "/no-tracking"
      methods = ["GET", "HEAD"]

      backend {
        type = "STOCK_RESPONSE_BACKEND"
        status = 200
        body = "Hello World"
      }

      request_policies {
        header_validations {
          headers {
            name = "DNT"
            required = true
          }

          headers {
            name = "X-Request-ID"
            required = false
          }
        }
      }
    }

    # An example route making use of a request validation policy
    # requiring a query parameter.
    routes {
      path = "/authorize"
      methods = ["GET", "HEAD"]

      backend {
        type = "STOCK_RESPONSE_BACKEND"
        status = 200
        body = "Hello World"
      }

      request_policies {
        query_parameter_validations {
          parameters {
            name = "client_id"
            required = true
          }

          parameters {
            name = "client_secret"
            required = false
          }
        }
      }
    }

    # An example route making use of a request validation policy
    # requiring a JSON or XML request body in permissive mode.
    routes {
      path = "/users"
      methods = ["POST"]

      backend {
        type = "STOCK_RESPONSE_BACKEND"
        status = 201
        body = "User Created"
      }

      request_policies {
        body_validation {
          validation_mode = "PERMISSIVE"

          required = true

          content {
            media_type = "application/json"
            validation_type = "NONE"
          }

          content {
            media_type = "application/xml"
            validation_type = "NONE"
          }
        }
      }
    }
  }
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
