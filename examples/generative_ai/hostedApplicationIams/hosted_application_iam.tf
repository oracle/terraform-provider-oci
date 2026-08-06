// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "hosted_application_iam_description" {
  default = "description"
}

variable "hosted_application_iam_display_name" {
  default = "qa_plus_agent_test_hosted_application_iam"
}

variable "hosted_application_iam_environment_variables_name" {
  default = "name"
}

variable "hosted_application_iam_environment_variables_type" {
  default = "PLAINTEXT"
}

variable "hosted_application_iam_environment_variables_value" {
  default = "{\"dummyKey\":\"dummyValue\"}"
}

variable "hosted_application_iam_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "hosted_application_iam_networking_config_inbound_networking_config_endpoint_mode" {
  default = "PUBLIC"
}

variable "hosted_application_iam_networking_config_outbound_networking_config_network_mode" {
  default = "MANAGED"
}

variable "hosted_application_iam_scaling_config_max_replica" {
  default = 10
}

variable "hosted_application_iam_scaling_config_min_replica" {
  default = 10
}

variable "hosted_application_iam_scaling_config_scaling_type" {
  default = "CPU"
}

variable "hosted_application_iam_scaling_config_target_concurrency_threshold" {
  default = 10
}

variable "hosted_application_iam_scaling_config_target_cpu_threshold" {
  default = 70
}

variable "hosted_application_iam_scaling_config_target_memory_threshold" {
  default = 80
}

variable "hosted_application_iam_scaling_config_target_rps_threshold" {
  default = 10
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_generative_ai_hosted_application_iam" "test_hosted_application_iam" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.hosted_application_iam_display_name

  #Optional
  description = var.hosted_application_iam_description
  environment_variables {
    #Required
    name  = var.hosted_application_iam_environment_variables_name
    type  = var.hosted_application_iam_environment_variables_type
    value = var.hosted_application_iam_environment_variables_value
  }
  freeform_tags = var.hosted_application_iam_freeform_tags
  networking_config {
    #Required
    inbound_networking_config {
      #Required
      endpoint_mode = var.hosted_application_iam_networking_config_inbound_networking_config_endpoint_mode
    }
    outbound_networking_config {
      #Required
      network_mode = var.hosted_application_iam_networking_config_outbound_networking_config_network_mode
    }
  }
  scaling_config {
    #Required
    scaling_type = var.hosted_application_iam_scaling_config_scaling_type

    #Optional
    max_replica                  = var.hosted_application_iam_scaling_config_max_replica
    min_replica                  = var.hosted_application_iam_scaling_config_min_replica
    target_concurrency_threshold = var.hosted_application_iam_scaling_config_target_concurrency_threshold
    target_cpu_threshold         = var.hosted_application_iam_scaling_config_target_cpu_threshold
    target_memory_threshold      = var.hosted_application_iam_scaling_config_target_memory_threshold
    target_rps_threshold         = var.hosted_application_iam_scaling_config_target_rps_threshold
  }
}

data "oci_generative_ai_hosted_application_iams" "test_hosted_application_iams" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.hosted_application_iam_display_name
  id           = oci_generative_ai_hosted_application_iam.test_hosted_application_iam.id
  state        = "ACTIVE"
}
