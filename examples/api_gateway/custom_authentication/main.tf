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

variable "gateway_display_name" {
  default = "displayName"
}

variable "gateway_endpoint_type" {
  default = "PUBLIC"
}

variable "gateway_state" {
  default = "ACTIVE"
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

variable "application_state" {
  default = "AVAILABLE"
}

variable "config" {
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

variable "function_image" {
}

variable "function_image_digest" {
}

variable "function_memory_in_mbs" {
  default = 128
}

variable "function_timeout_in_seconds" {
  default = 30
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

variable "deployment_specification_request_policies_authentication_is_anonymous_access_allowed" {
  default = false
}

variable "deployment_specification_request_policies_authentication_token_header" {
  default = "Authorization"
}

variable "deployment_specification_request_policies_authentication_token_query_param" {
  default = "tokenQueryParam"
}

variable "deployment_specification_request_policies_authentication_type" {
  default = "CUSTOM_AUTHENTICATION"
}

variable "deployment_specification_request_policies_cors_allowed_headers" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_methods" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_exposed_headers" {
  default = ["*"]
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
  default = 1
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
  default = 1
}

variable "deployment_specification_routes_backend_send_timeout_in_seconds" {
  default = 1
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
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_exposed_headers" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_cors_max_age_in_seconds" {
  default = "600"
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

# Terraform will take 5 minutes after destroying an application due to a known service issue.
# please refer: https://docs.cloud.oracle.com/iaas/Content/Functions/Tasks/functionsdeleting.htm
resource "oci_functions_application" "test_application" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "example-application"
  subnet_ids     = [oci_core_subnet.regional_subnet.id]

  #Optional
  config = var.config
}

data "oci_functions_applications" "test_applications" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = "example-application"
  id           = oci_functions_application.test_application.id
  state        = var.application_state
}

resource "oci_functions_function" "test_function" {
  #Required
  application_id = oci_functions_application.test_application.id
  display_name   = "example-function"
  image          = var.function_image
  memory_in_mbs  = "128"

  #Optional
  config             = var.config
  image_digest       = var.function_image_digest
  timeout_in_seconds = "30"
}

data "oci_functions_functions" "test_functions" {
  #Required
  application_id = oci_functions_application.test_application.id

  #Optional
  display_name = "example-function"
  id           = oci_functions_function.test_function.id
  state        = "AVAILABLE"
}

resource "oci_apigateway_gateway" "test_gateway" {
  #Required
  compartment_id = var.compartment_ocid
  endpoint_type  = var.gateway_endpoint_type
  subnet_id      = oci_core_subnet.regional_subnet.id

  #Optional
  display_name = var.gateway_display_name
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
        function_id = oci_functions_function.test_function.id
        type        = var.deployment_specification_request_policies_authentication_type

        #Optional
        is_anonymous_access_allowed = var.deployment_specification_request_policies_authentication_is_anonymous_access_allowed
        token_header                = var.deployment_specification_request_policies_authentication_token_header
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
        url  = var.deployment_specification_routes_backend_url
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
          type = var.deployment_specification_routes_request_policies_authorization_type
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
      }
    }
  }

  #Optional
  display_name = var.deployment_display_name
}

data "oci_apigateway_gateways" "test_gateways" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.gateway_display_name
  state        = var.gateway_state
}

data "oci_apigateway_deployments" "test_deployments" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.deployment_display_name
  gateway_id   = oci_apigateway_gateway.test_gateway.id
  state        = var.deployment_state
}

