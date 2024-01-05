// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_subnet" "gateway_subnet_rd" {
  cidr_block        = "10.0.1.0/24"
  display_name      = "gatewaySubnetRD"
  dns_label         = "gatewaySubnetRD"
  compartment_id    = "${var.compartment_ocid}"
  vcn_id            = "${oci_core_vcn.vcn3_rd.id}"
  security_list_ids = ["${oci_core_vcn.vcn3_rd.default_security_list_id}"]
  route_table_id    = "${oci_core_vcn.vcn3_rd.default_route_table_id}"
  dhcp_options_id   = "${oci_core_vcn.vcn3_rd.default_dhcp_options_id}"
}

resource "oci_apigateway_gateway" "apigateway_gateway_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  endpoint_type  = "${var.gateway_endpoint_type}"
  subnet_id      = "${oci_core_subnet.gateway_subnet_rd.id}"

  #Optional
  display_name = "${var.gateway_display_name}"
}

resource "oci_apigateway_deployment" "apigateway_deployment_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  gateway_id     = "${oci_apigateway_gateway.apigateway_gateway_rd.id}"
  path_prefix    = "${var.deployment_path_prefix}"

  specification {
    #Optional
    logging_policies {
      #Optional
      access_log {
        #Optional
        is_enabled = "${var.deployment_specification_logging_policies_access_log_is_enabled}"
      }

      execution_log {
        #Optional
        is_enabled = "${var.deployment_specification_logging_policies_execution_log_is_enabled}"
        log_level  = "${var.deployment_specification_logging_policies_execution_log_log_level}"
      }
    }

    request_policies {
      #Optional
      authentication {
        #Required
        function_id = "${oci_functions_function.functions_function_rd.id}"
        type        = "${var.deployment_specification_request_policies_authentication_type}"

        #Optional
        is_anonymous_access_allowed = "${var.deployment_specification_request_policies_authentication_is_anonymous_access_allowed}"
        token_header                = "${var.deployment_specification_request_policies_authentication_token_header}"
      }

      cors {
        #Required
        allowed_origins = "${var.deployment_specification_request_policies_cors_allowed_origins}"

        #Optional
        allowed_headers              = "${var.deployment_specification_request_policies_cors_allowed_headers}"
        allowed_methods              = "${var.deployment_specification_request_policies_cors_allowed_methods}"
        exposed_headers              = "${var.deployment_specification_request_policies_cors_exposed_headers}"
        is_allow_credentials_enabled = "${var.deployment_specification_request_policies_cors_is_allow_credentials_enabled}"
        max_age_in_seconds           = "${var.deployment_specification_request_policies_cors_max_age_in_seconds}"
      }

      rate_limiting {
        #Required
        rate_in_requests_per_second = "${var.deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second}"
        rate_key                    = "${var.deployment_specification_request_policies_rate_limiting_rate_key}"
      }
    }

    routes {
      #Required
      backend {
        #Required
        type = "${var.deployment_specification_routes_backend_type}"
        url  = "${var.deployment_specification_routes_backend_url}"
      }

      path = "${var.deployment_specification_routes_path}"

      #Optional
      logging_policies {
        #Optional
        access_log {
          #Optional
          is_enabled = "${var.deployment_specification_routes_logging_policies_access_log_is_enabled}"
        }

        execution_log {
          #Optional
          is_enabled = "${var.deployment_specification_routes_logging_policies_execution_log_is_enabled}"
          log_level  = "${var.deployment_specification_routes_logging_policies_execution_log_log_level}"
        }
      }

      methods = "${var.deployment_specification_routes_methods}"

      request_policies {
        #Optional
        authorization {
          #Optional
          type = "${var.deployment_specification_routes_request_policies_authorization_type}"
        }

        cors {
          #Required
          allowed_origins = "${var.deployment_specification_routes_request_policies_cors_allowed_origins}"

          #Optional
          allowed_headers              = "${var.deployment_specification_routes_request_policies_cors_allowed_headers}"
          allowed_methods              = "${var.deployment_specification_routes_request_policies_cors_allowed_methods}"
          exposed_headers              = "${var.deployment_specification_routes_request_policies_cors_exposed_headers}"
          is_allow_credentials_enabled = "${var.deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled}"
          max_age_in_seconds           = "${var.deployment_specification_routes_request_policies_cors_max_age_in_seconds}"
        }
      }
    }
  }

  #Optional
  display_name = "apigatewayDeploymentRD"
}

data "oci_apigateway_gateways" "apigateway_gateways_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.gateway_display_name}"
  state        = "${var.gateway_state}"
}

data "oci_apigateway_deployments" "apigateway_deployments_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "apigatewayDeploymentsRD"
  gateway_id   = "${oci_apigateway_gateway.apigateway_gateway_rd.id}"
  state        = "${var.deployment_state}"
}
