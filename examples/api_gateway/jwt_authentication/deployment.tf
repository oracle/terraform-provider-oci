// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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

resource "oci_apigateway_gateway" "test_gateway" {
  compartment_id = var.compartment_ocid
  endpoint_type  = "PUBLIC"
  subnet_id      = oci_core_subnet.regional_subnet.id
}

resource "oci_apigateway_deployment" "test_deployment" {
  compartment_id = var.compartment_ocid
  gateway_id     = oci_apigateway_gateway.test_gateway.id
  path_prefix    = "/v1"

  specification {
    logging_policies {
      access_log {
        is_enabled = "false"
      }

      execution_log {
        is_enabled = "false"
        log_level  = "INFO"
      }
    }

    request_policies {
      authentication {
        type                        = "JWT_AUTHENTICATION"
        token_header                = "Authorization"
        token_auth_scheme           = "Bearer"
        is_anonymous_access_allowed = "false"
        issuers                     = ["https://identity.oraclecloud.com/"]
        audiences                   = ["https://www.oracle.com/"]
        max_clock_skew_in_seconds   = "10"

        public_keys {
          type                        = "REMOTE_JWKS"
          max_cache_duration_in_hours = "10"
          uri                         = "https://oracle.com/jwks.json"
        }
      }

      cors {
        allowed_origins = ["*"]
        allowed_methods = ["GET"]
      }

      rate_limiting {
        rate_in_requests_per_second = "10"
        rate_key                    = "CLIENT_IP"
      }
    }

    routes {
      backend {
        type = "HTTP_BACKEND"
        url  = "https://api.weather.gov"
      }

      path    = "/hello"
      methods = ["GET"]

      logging_policies {
        access_log {
          is_enabled = "false"
        }

        execution_log {
          is_enabled = "false"
          log_level  = "INFO"
        }
      }

      request_policies {
        authorization {
          type = "AUTHENTICATION_ONLY"
        }

        cors {
          allowed_headers = ["*"]
          allowed_methods = ["GET"]
          allowed_origins = ["*"]
        }
      }
    }
  }

  display_name = "test_api_gateway_deployment"
}

data "oci_apigateway_gateways" "test_gateways" {
  compartment_id = var.compartment_ocid
}

data "oci_apigateway_deployments" "test_deployments" {
  compartment_id = var.compartment_ocid
}

