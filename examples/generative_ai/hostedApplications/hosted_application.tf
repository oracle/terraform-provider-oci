// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "config_file_profile" {
}

variable "region" {}
variable "compartment_id" {}

variable "hosted_application_description" {
  default = "description"
}

variable "hosted_application_display_name" {
  default = "displayName"
}

variable "hosted_application_environment_variables_name" {
  default = "name"
}

variable "hosted_application_environment_variables_type" {
  default = "PLAINTEXT"
}

variable "hosted_application_environment_variables_value" {
  default = "{\"dummyKey\": \"dummyValue\"}"
}

variable "hosted_application_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "hosted_application_inbound_auth_config_idcs_config_audience" {
  default = "audience"
}

variable "hosted_application_inbound_auth_config_idcs_config_domain_url" {
  default = "domainUrl"
}

variable "hosted_application_inbound_auth_config_idcs_config_scope" {
  default = "scope"
}

variable "hosted_application_inbound_auth_config_inbound_auth_config_type" {
  default = "IDCS_AUTH_CONFIG"
}

variable "hosted_application_networking_config_inbound_networking_config_endpoint_mode" {
  default = "PUBLIC"
}

variable "hosted_application_networking_config_outbound_networking_config_network_mode" {
  default = "MANAGED"
}

variable "hosted_application_scaling_config_max_replica" {
  default = 10
}

variable "hosted_application_scaling_config_min_replica" {
  default = 10
}

variable "hosted_application_scaling_config_scaling_type" {
  default = "CPU"
}

variable "hosted_application_scaling_config_target_cpu_threshold" {
  default = 50
}

variable "hosted_application_state" {
  default = "ACTIVE"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_generative_ai_hosted_application" "test_hosted_application" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.hosted_application_display_name

  #Optional
  description = var.hosted_application_description
  environment_variables {
    #Required
    name  = var.hosted_application_environment_variables_name
    type  = var.hosted_application_environment_variables_type
    value = var.hosted_application_environment_variables_value
  }
  freeform_tags = var.hosted_application_freeform_tags
  inbound_auth_config {
    #Required
    inbound_auth_config_type = var.hosted_application_inbound_auth_config_inbound_auth_config_type

    #Optional
    idcs_config {
      #Required
      domain_url = var.hosted_application_inbound_auth_config_idcs_config_domain_url
      scope      = var.hosted_application_inbound_auth_config_idcs_config_scope

      #Optional
      audience = var.hosted_application_inbound_auth_config_idcs_config_audience
    }
  }
  networking_config {
    #Required
    inbound_networking_config {
      #Required
      endpoint_mode = var.hosted_application_networking_config_inbound_networking_config_endpoint_mode

    }
    outbound_networking_config {
      #Required
      network_mode = var.hosted_application_networking_config_outbound_networking_config_network_mode
    }
  }
  scaling_config {
    #Required
    scaling_type = var.hosted_application_scaling_config_scaling_type

    #Optional
    max_replica          = var.hosted_application_scaling_config_max_replica
    min_replica          = var.hosted_application_scaling_config_min_replica
    target_cpu_threshold = var.hosted_application_scaling_config_target_cpu_threshold
  }
}

data "oci_generative_ai_hosted_applications" "test_hosted_applications" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.hosted_application_display_name
  id           = oci_generative_ai_hosted_application.test_hosted_application.id
  state        = var.hosted_application_state
}
