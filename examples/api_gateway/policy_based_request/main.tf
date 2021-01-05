// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "image" {
  default = "phx.ocir.io/dxterraformdev/functions/function:0.0.1"
}

variable "image_digest" {
  default = "sha256:73a6de3a706f299f59d2e217c049814e14b42346c12b407996ebbce0c453f1a2"
}

variable "deployment_defined_tags_value" {
  default = "value"
}

variable "deployment_display_name" {
  default = "displayName"
}

variable "deployment_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "config" {
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

variable "deployment_path_prefix" {
  default = "/v1"
}

variable "deployment_specification_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_request_policies_authentication_audiences" {
  default = []
}

variable "deployment_specification_request_policies_authentication_is_anonymous_access_allowed" {
  default = false
}

variable "deployment_specification_request_policies_authentication_issuers" {
  default = []
}

variable "deployment_specification_request_policies_authentication_max_clock_skew_in_seconds" {
  default = 1.0
}

variable "deployment_specification_request_policies_authentication_public_keys_is_ssl_verify_disabled" {
  default = false
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_alg" {
  default = "alg"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_e" {
  default = "e"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_format" {
  default = "PEM"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_key" {
  default = "key"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_key_ops" {
  default = []
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_kid" {
  default = "kid"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_kty" {
  default = "kty"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_n" {
  default = "n"
}

variable "deployment_specification_request_policies_authentication_public_keys_keys_use" {
  default = "use"
}

variable "deployment_specification_request_policies_authentication_public_keys_max_cache_duration_in_hours" {
  default = 10
}

variable "deployment_specification_request_policies_authentication_public_keys_type" {
  default = "REMOTE_JWKS"
}

variable "deployment_specification_request_policies_authentication_public_keys_uri" {
  default = "https://oracle.com/jwks.json"
}

variable "deployment_specification_request_policies_authentication_token_auth_scheme" {
  default = "Bearer"
}

variable "deployment_specification_request_policies_authentication_token_header" {
  default = "Authorization"
}

/*
variable "deployment_specification_request_policies_authentication_token_query_param" {
  default = "tokenQueryParam"
}
*/
variable "deployment_specification_request_policies_authentication_type" {
  default = "CUSTOM_AUTHENTICATION"
}

variable "deployment_specification_request_policies_authentication_verify_claims_is_required" {
  default = false
}

variable "deployment_specification_request_policies_authentication_verify_claims_key" {
  default = "key"
}

variable "deployment_specification_request_policies_authentication_verify_claims_values" {
  default = []
}

variable "deployment_specification_request_policies_cors_allowed_headers" {
  default = ["*", "Content-Type", "X-Foo-Bar"]
}

variable "deployment_specification_request_policies_cors_allowed_methods" {
  default = ["*", "GET", "POST"]
}

variable "deployment_specification_request_policies_cors_allowed_origins" {
  default = ["*", "null", "https://friendly.com:8080"]
}

variable "deployment_specification_request_policies_cors_exposed_headers" {
  default = ["*", "opc-request-id"]
}

variable "deployment_specification_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

variable "deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second" {
  default = 10
}

variable "deployment_specification_request_policies_rate_limiting_rate_key" {
  default = "CLIENT_IP"
}

variable "deployment_specification_routes_backend_body" {
  default = "body"
}

variable "deployment_specification_routes_backend_connect_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_headers_name" {
  default = "name"
}

variable "deployment_specification_routes_backend_headers_value" {
  default = "value"
}

variable "deployment_specification_routes_backend_is_ssl_verify_disabled" {
  default = false
}

variable "deployment_specification_routes_backend_read_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_send_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_status" {
  default = 10
}

variable "deployment_specification_routes_backend_type" {
  default = "HTTP_BACKEND"
}

variable "deployment_specification_routes_backend_url" {
  default = "https://api.weather.gov"
}

variable "deployment_specification_routes_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_routes_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_path" {
  default = "/hello"
}

variable "deployment_specification_routes_request_policies_authorization_allowed_scope" {
  default = []
}

variable "deployment_specification_routes_request_policies_authorization_type" {
  default = "AUTHENTICATION_ONLY"
}

variable "deployment_specification_routes_request_policies_cors_allowed_headers" {
  default = []
}

variable "deployment_specification_routes_request_policies_cors_allowed_methods" {
  default = []
}

variable "deployment_specification_routes_request_policies_cors_allowed_origins" {
  default = ["*", "null", "https://friendly.com:8080"]
}

variable "deployment_specification_routes_request_policies_cors_exposed_headers" {
  default = []
}

variable "deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

variable "deployment_specification_routes_request_policies_header_transformations_filter_headers_items_name" {
  default = "topSecret"
}

variable "deployment_specification_routes_request_policies_header_transformations_filter_headers_type" {
  default = "BLOCK"
}

variable "deployment_specification_routes_request_policies_header_transformations_rename_headers_items_from" {
  default = "from"
}

variable "deployment_specification_routes_request_policies_header_transformations_rename_headers_items_to" {
  default = "to"
}

variable "deployment_specification_routes_request_policies_header_transformations_set_headers_items_if_exists" {
  default = "ifExists"
}

variable "deployment_specification_routes_request_policies_header_transformations_set_headers_items_name" {
  default = "name"
}

variable "deployment_specification_routes_request_policies_header_transformations_set_headers_items_values" {
  default = ["test"]
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_items_name" {
  default = "TOPSECRET"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_type" {
  default = "BLOCK"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_from" {
  default = "from"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_to" {
  default = "to"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_if_exists" {
  default = "ifExists"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_name" {
  default = "name"
}

variable "deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_values" {
  default = ["test"]
}

variable "deployment_specification_routes_response_policies_header_transformations_filter_headers_items_name" {
  default = "Aname"
}

variable "deployment_specification_routes_response_policies_header_transformations_filter_headers_type" {
  default = "BLOCK"
}

variable "deployment_specification_routes_response_policies_header_transformations_rename_headers_items_from" {
  default = "from"
}

variable "deployment_specification_routes_response_policies_header_transformations_rename_headers_items_to" {
  default = "to"
}

variable "deployment_specification_routes_response_policies_header_transformations_set_headers_items_if_exists" {
  default = "ifExists"
}

variable "deployment_specification_routes_response_policies_header_transformations_set_headers_items_name" {
  default = "name"
}

variable "deployment_specification_routes_response_policies_header_transformations_set_headers_items_values" {
  default = ["test"]
}

variable "deployment_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apigateway_deployment" "test_deployment" {
  #Required
  compartment_id = var.compartment_ocid
  gateway_id     = oci_apigateway_gateway.test_gateway.id
  path_prefix    = var.deployment_path_prefix

  specification {
    #Optional
    logging_policies {
      #Optional
      access_log {
        #Optional
        is_enabled = var.deployment_specification_logging_policies_access_log_is_enabled
      }

      execution_log {
        #Optional
        is_enabled = var.deployment_specification_logging_policies_execution_log_is_enabled
        log_level  = var.deployment_specification_logging_policies_execution_log_log_level
      }
    }

    request_policies {
      #Optional
      authentication {
        #Required
        type = var.deployment_specification_request_policies_authentication_type

        #Optional
        audiences                   = var.deployment_specification_request_policies_authentication_audiences
        function_id                 = oci_functions_function.test_function.id
        is_anonymous_access_allowed = var.deployment_specification_request_policies_authentication_is_anonymous_access_allowed
        issuers                     = var.deployment_specification_request_policies_authentication_issuers
        max_clock_skew_in_seconds   = var.deployment_specification_request_policies_authentication_max_clock_skew_in_seconds

        public_keys {
          #Required
          type = var.deployment_specification_request_policies_authentication_public_keys_type

          #Optional
          is_ssl_verify_disabled = var.deployment_specification_request_policies_authentication_public_keys_is_ssl_verify_disabled

          keys {
            #Required
            format = var.deployment_specification_request_policies_authentication_public_keys_keys_format

            #Optional
            alg     = var.deployment_specification_request_policies_authentication_public_keys_keys_alg
            e       = var.deployment_specification_request_policies_authentication_public_keys_keys_e
            key     = var.deployment_specification_request_policies_authentication_public_keys_keys_key
            key_ops = var.deployment_specification_request_policies_authentication_public_keys_keys_key_ops
            kid     = var.deployment_specification_request_policies_authentication_public_keys_keys_kid
            kty     = var.deployment_specification_request_policies_authentication_public_keys_keys_kty
            n       = var.deployment_specification_request_policies_authentication_public_keys_keys_n
            use     = var.deployment_specification_request_policies_authentication_public_keys_keys_use
          }

          max_cache_duration_in_hours = var.deployment_specification_request_policies_authentication_public_keys_max_cache_duration_in_hours
          uri                         = var.deployment_specification_request_policies_authentication_public_keys_uri
        }

        token_auth_scheme = var.deployment_specification_request_policies_authentication_token_auth_scheme
        token_header      = var.deployment_specification_request_policies_authentication_token_header

        verify_claims {
          #Optional
          is_required = var.deployment_specification_request_policies_authentication_verify_claims_is_required
          key         = var.deployment_specification_request_policies_authentication_verify_claims_key
          values      = var.deployment_specification_request_policies_authentication_verify_claims_values
        }
      }

      cors {
        #Required
        allowed_origins = var.deployment_specification_request_policies_cors_allowed_origins

        #Optional
        allowed_headers              = var.deployment_specification_request_policies_cors_allowed_headers
        allowed_methods              = var.deployment_specification_request_policies_cors_allowed_methods
        exposed_headers              = var.deployment_specification_request_policies_cors_exposed_headers
        is_allow_credentials_enabled = var.deployment_specification_request_policies_cors_is_allow_credentials_enabled
        max_age_in_seconds           = var.deployment_specification_request_policies_cors_max_age_in_seconds
      }

      rate_limiting {
        #Required
        rate_in_requests_per_second = var.deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second
        rate_key                    = var.deployment_specification_request_policies_rate_limiting_rate_key
      }
    }

    routes {
      #Required
      backend {
        #Required
        type = var.deployment_specification_routes_backend_type

        #Optional
        body                       = var.deployment_specification_routes_backend_body
        connect_timeout_in_seconds = var.deployment_specification_routes_backend_connect_timeout_in_seconds

        headers {
          #Optional
          name  = var.deployment_specification_routes_backend_headers_name
          value = var.deployment_specification_routes_backend_headers_value
        }

        is_ssl_verify_disabled  = var.deployment_specification_routes_backend_is_ssl_verify_disabled
        read_timeout_in_seconds = var.deployment_specification_routes_backend_read_timeout_in_seconds
        send_timeout_in_seconds = var.deployment_specification_routes_backend_send_timeout_in_seconds
        status                  = var.deployment_specification_routes_backend_status
        url                     = var.deployment_specification_routes_backend_url
      }

      path = var.deployment_specification_routes_path

      #Optional
      logging_policies {
        #Optional
        access_log {
          #Optional
          is_enabled = var.deployment_specification_routes_logging_policies_access_log_is_enabled
        }

        execution_log {
          #Optional
          is_enabled = var.deployment_specification_routes_logging_policies_execution_log_is_enabled
          log_level  = var.deployment_specification_routes_logging_policies_execution_log_log_level
        }
      }

      methods = var.deployment_specification_routes_methods

      request_policies {
        #Optional
        authorization {
          #Optional
          allowed_scope = var.deployment_specification_routes_request_policies_authorization_allowed_scope
          type          = var.deployment_specification_routes_request_policies_authorization_type
        }

        cors {
          #Required
          allowed_origins = var.deployment_specification_routes_request_policies_cors_allowed_origins

          #Optional
          allowed_headers              = var.deployment_specification_routes_request_policies_cors_allowed_headers
          allowed_methods              = var.deployment_specification_routes_request_policies_cors_allowed_methods
          exposed_headers              = var.deployment_specification_routes_request_policies_cors_exposed_headers
          is_allow_credentials_enabled = var.deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled
          max_age_in_seconds           = var.deployment_specification_routes_request_policies_cors_max_age_in_seconds
        }

        header_transformations {
          #Optional
          filter_headers {
            #Required
            items {
              #Required
              name = var.deployment_specification_routes_request_policies_header_transformations_filter_headers_items_name
            }

            type = var.deployment_specification_routes_request_policies_header_transformations_filter_headers_type
          }

          rename_headers {
            #Required
            items {
              #Required
              from = var.deployment_specification_routes_request_policies_header_transformations_rename_headers_items_from
              to   = var.deployment_specification_routes_request_policies_header_transformations_rename_headers_items_to
            }
          }

          set_headers {
            #Required
            items {
              #Required
              name   = var.deployment_specification_routes_request_policies_header_transformations_set_headers_items_name
              values = var.deployment_specification_routes_request_policies_header_transformations_set_headers_items_values

              #Optional
            }
          }
        }

        query_parameter_transformations {
          #Optional
          filter_query_parameters {
            #Required
            items {
              #Required
              name = var.deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_items_name
            }

            type = var.deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_type
          }

          rename_query_parameters {
            #Required
            items {
              #Required
              from = var.deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_from
              to   = var.deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_to
            }
          }

          set_query_parameters {
            #Required
            items {
              #Required
              name   = var.deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_name
              values = var.deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_values

              #Optional
            }
          }
        }
      }

      response_policies {
        #Optional
        header_transformations {
          #Optional
          filter_headers {
            #Required
            items {
              #Required
              name = var.deployment_specification_routes_response_policies_header_transformations_filter_headers_items_name
            }

            type = var.deployment_specification_routes_response_policies_header_transformations_filter_headers_type
          }

          rename_headers {
            #Required
            items {
              #Required
              from = var.deployment_specification_routes_response_policies_header_transformations_rename_headers_items_from
              to   = var.deployment_specification_routes_response_policies_header_transformations_rename_headers_items_to
            }
          }

          set_headers {
            #Required
            items {
              #Required
              name   = var.deployment_specification_routes_response_policies_header_transformations_set_headers_items_name
              values = var.deployment_specification_routes_response_policies_header_transformations_set_headers_items_values

              #Optional
            }
          }
        }
      }
    }
  }

  #Optional
  display_name  = var.deployment_display_name
  freeform_tags = var.deployment_freeform_tags
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

data "oci_apigateway_deployments" "test_deployments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.deployment_display_name
  gateway_id   = oci_apigateway_gateway.test_gateway.id
  state        = var.deployment_state
}

resource "oci_functions_function" "test_function" {
  #Required
  application_id = oci_functions_application.test_application.id
  display_name   = "example-function"
  image          = var.image
  memory_in_mbs  = "128"

  #Optional
  config             = var.config
  image_digest       = var.image_digest
  timeout_in_seconds = "30"
}

resource "oci_functions_application" "test_application" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "example-application"
  subnet_ids     = [oci_core_subnet.regional_subnet.id]

  #Optional
  config = var.config
}
