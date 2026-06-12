// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "config_file_profile" {
}

variable "region" {}
variable "compartment_id" {}

variable "hosted_application_display_name" {
  default = "displayName"
}

variable "hosted_application_inbound_auth_config_idcs_config_domain_url" {
  default = "domainUrl"
}

variable "hosted_application_inbound_auth_config_idcs_config_scope" {
  default = "scope"
}

variable "hosted_deployment_active_artifact_artifact_type" {
  default = "SIMPLE_DOCKER_ARTIFACT"
}

variable "hosted_deployment_active_artifact_container_uri" {
  default = ""
}

variable "hosted_deployment_active_artifact_tag" {
  default = "latest"
}

variable "hosted_deployment_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "hosted_deployment_state" {
  default = "ACTIVE"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

locals {
  hosted_deployment_container_uri = var.hosted_deployment_active_artifact_container_uri != "" ? var.hosted_deployment_active_artifact_container_uri : "${var.region}.ocir.io/axk4z7krhqfx/cost-service"
}

resource "oci_generative_ai_hosted_application" "test_hosted_application" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.hosted_application_display_name

  inbound_auth_config {
    #Required
    inbound_auth_config_type = "IDCS_AUTH_CONFIG"

    idcs_config {
      #Required
      domain_url = var.hosted_application_inbound_auth_config_idcs_config_domain_url
      scope      = var.hosted_application_inbound_auth_config_idcs_config_scope
    }
  }

  networking_config {
    #Required
    inbound_networking_config {
      #Required
      endpoint_mode = "PUBLIC"
    }

    outbound_networking_config {
      #Required
      network_mode = "MANAGED"
    }
  }
}

resource "oci_generative_ai_hosted_deployment" "test_hosted_deployment" {
  #Required
  active_artifact {
    #Required
    artifact_type = var.hosted_deployment_active_artifact_artifact_type
    container_uri = local.hosted_deployment_container_uri
    tag           = var.hosted_deployment_active_artifact_tag
  }
  compartment_id        = var.compartment_id
  hosted_application_id = oci_generative_ai_hosted_application.test_hosted_application.id

  #Optional
  freeform_tags = var.hosted_deployment_freeform_tags
}

data "oci_generative_ai_hosted_deployments" "test_hosted_deployments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  application_id = oci_generative_ai_hosted_application.test_hosted_application.id
  id             = oci_generative_ai_hosted_deployment.test_hosted_deployment.id
  state          = var.hosted_deployment_state
}
