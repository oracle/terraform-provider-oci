// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
/*
 * The following API Gateway and deployment demonstrate using OpenID connect
 * with API Gateway using an OAUTH2 flow.
 */
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "client_id" {
  description = "The OAuth2 Client ID"
}
variable "client_secret_id" {
  description = "An ID to an OCI Secret value containing the OAuth2 client secret"
}
variable "client_secret_version_number" {
  default = 1
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}
resource "oci_apigateway_deployment" "openid_connect_deployment" {
  compartment_id = var.compartment_ocid
  gateway_id = oci_apigateway_gateway.test_gateway.id
  path_prefix = "/"
  specification {
    request_policies {
      authentication {
        type = "TOKEN_AUTHENTICATION"
        token_header = "Authorization"
        token_auth_scheme = "Bearer"
        is_anonymous_access_allowed = false
        validation_policy {
          // Example validation policy using an OAuth2 introspection endpoint
          // (https://datatracker.ietf.org/doc/html/rfc7662) to validate the
          // clients authorization credentials
          type = "REMOTE_DISCOVERY"
          is_ssl_verify_disabled = true
          max_cache_duration_in_hours = 1
          source_uri_details {
            // Discover the OAuth2/OpenID configuration from an RFC8414
            // metadata endpoint (https://www.rfc-editor.org/rfc/rfc8414)
            type = "DISCOVERY_URI"
            uri = "https://auth.example.com/.well-known/oauth-authorization-server"
          }
          client_details {
            // Specify the OAuth client id and secret to use with the
            // introspection endpoint
            type = "CUSTOM"
            client_id = var.client_id
            client_secret_id = var.client_secret_id
            client_secret_version_number = var.client_secret_version_number
          }
          additional_validation_policy {
            issuers                     = ["https://identity.oraclecloud.com/"]
            audiences                   = ["https://www.oracle.com/"]
            verify_claims {
              is_required = true
              key = "key"
              values = ["value"]
            }
          }
        }
        validation_failure_policy {
          // When a client uses the API without auth credentials, or
          // invalid/expired credentials then invoke the OAuth2 flow using
          // the configuration below.
          type = "OAUTH2"
          scopes = ["openid"]
          response_type = "CODE"
          max_expiry_duration_in_hours = 1
          use_cookies_for_intermediate_steps = true
          use_cookies_for_session = true
          use_pkce = true
          fallback_redirect_path = "/fallback"
          source_uri_details {
            // Use the same discovery URI as the validation policy above.
            type = "VALIDATION_BLOCK"
          }
          client_details {
            // Use the same OAuth2 client details as the validation policy above.
            type = "VALIDATION_BLOCK"
          }
        }
      }
    }
    routes {
      path = "/"
      methods = ["GET", "HEAD"]
      backend {
        type = "STOCK_RESPONSE_BACKEND"
        status = 200
        body = "Hello World"
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
