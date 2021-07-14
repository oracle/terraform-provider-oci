// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "monitor_configuration_config_type" {
  default = "BROWSER_CONFIG"
}

variable "monitor_configuration_is_certificate_validation_enabled" {
  default = false
}

variable "monitor_configuration_is_failure_retried" {
  default = false
}

variable "monitor_configuration_is_redirection_enabled" {
  default = false
}

variable "monitor_configuration_req_authentication_details_auth_headers_header_name" {
  default = "headerName"
}

variable "monitor_configuration_req_authentication_details_auth_headers_header_value" {
  default = "headerValue"
}

variable "monitor_configuration_req_authentication_details_auth_request_method" {
  default = "GET"
}

variable "monitor_configuration_req_authentication_details_auth_request_post_body" {
  default = "authRequestPostBody"
}

variable "monitor_configuration_req_authentication_details_auth_token" {
  default = "authToken"
}

variable "monitor_configuration_req_authentication_details_auth_url" {
  default = "authUrl"
}

variable "monitor_configuration_req_authentication_details_auth_user_password" {
  default = "authUserPassword"
}

variable "monitor_configuration_req_authentication_details_oauth_scheme" {
  default = "NONE"
}

variable "monitor_configuration_req_authentication_scheme" {
  default = "OAUTH"
}

variable "monitor_configuration_request_headers_header_name" {
  default = "headerName"
}

variable "monitor_configuration_request_headers_header_value" {
  default = "headerValue"
}

variable "monitor_configuration_request_method" {
  default = "GET"
}

variable "monitor_configuration_request_post_body" {
  default = "requestPostBody"
}

variable "monitor_configuration_request_query_params_param_name" {
  default = "paramName"
}

variable "monitor_configuration_request_query_params_param_value" {
  default = "paramValue"
}

variable "monitor_configuration_verify_response_codes" {
  default = []
}

variable "monitor_configuration_verify_response_content" {
  default = "verifyResponseContent"
}

variable "monitor_configuration_verify_texts_text" {
  default = "text"
}

variable "monitor_defined_tags_value" {
  default = "value"
}

variable "monitor_display_name" {
  default = "displayName"
}

variable "monitor_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "monitor_monitor_type" {
  default = "SCRIPTED_BROWSER"
}

variable "monitor_repeat_interval_in_seconds" {
  default = 600
}

variable "monitor_script_parameters_param_name" {
  default = "paramName"
}

variable "monitor_script_parameters_param_value" {
  default = "paramValue"
}

variable "monitor_status" {
  default = "ENABLED"
}

variable "monitor_target" {
  default = "target"
}

variable "monitor_timeout_in_seconds" {
  default = 60
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_synthetics_monitor" "test_monitor" {
  #Required
  apm_domain_id              = oci_apm_apm_domain.test_apm_domain.id
  display_name               = var.monitor_display_name
  monitor_type               = var.monitor_monitor_type
  repeat_interval_in_seconds = var.monitor_repeat_interval_in_seconds
  vantage_points {
  }

  #Optional
  configuration {

    #Optional
    config_type                       = var.monitor_configuration_config_type
    is_certificate_validation_enabled = var.monitor_configuration_is_certificate_validation_enabled
    is_failure_retried                = var.monitor_configuration_is_failure_retried
    is_redirection_enabled            = var.monitor_configuration_is_redirection_enabled
    req_authentication_details {

      #Optional
      auth_headers {

        #Optional
        header_name  = var.monitor_configuration_req_authentication_details_auth_headers_header_name
        header_value = var.monitor_configuration_req_authentication_details_auth_headers_header_value
      }
      auth_request_method    = var.monitor_configuration_req_authentication_details_auth_request_method
      auth_request_post_body = var.monitor_configuration_req_authentication_details_auth_request_post_body
      auth_token             = var.monitor_configuration_req_authentication_details_auth_token
      auth_url               = var.monitor_configuration_req_authentication_details_auth_url
      auth_user_name         = oci_identity_user.test_user.name
      auth_user_password     = var.monitor_configuration_req_authentication_details_auth_user_password
      oauth_scheme           = var.monitor_configuration_req_authentication_details_oauth_scheme
    }
    req_authentication_scheme = var.monitor_configuration_req_authentication_scheme
    request_headers {

      #Optional
      header_name  = var.monitor_configuration_request_headers_header_name
      header_value = var.monitor_configuration_request_headers_header_value
    }
    request_method    = var.monitor_configuration_request_method
    request_post_body = var.monitor_configuration_request_post_body
    request_query_params {

      #Optional
      param_name  = var.monitor_configuration_request_query_params_param_name
      param_value = var.monitor_configuration_request_query_params_param_value
    }
    verify_response_codes   = var.monitor_configuration_verify_response_codes
    verify_response_content = var.monitor_configuration_verify_response_content
    verify_texts {

      #Optional
      text = var.monitor_configuration_verify_texts_text
    }
  }
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.monitor_defined_tags_value)
  freeform_tags = var.monitor_freeform_tags
  script_id     = oci_apm_synthetics_script.test_script.id
  script_parameters {
    #Required
    param_name  = var.monitor_script_parameters_param_name
    param_value = var.monitor_script_parameters_param_value
  }
  status             = var.monitor_status
  target             = var.monitor_target
  timeout_in_seconds = var.monitor_timeout_in_seconds
}

data "oci_apm_synthetics_monitors" "test_monitors" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  display_name = var.monitor_display_name
  monitor_type = var.monitor_monitor_type
  script_id    = oci_apm_synthetics_script.test_script.id
  status       = var.monitor_status
}

