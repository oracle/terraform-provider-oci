// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}
variable "compartment_id" {}

variable "agent_agent_type" {
  default = "APPLIANCE"
}

variable "agent_agent_version" {
  default = "agentVersion"
}

variable "agent_defined_tags_value" {
  default = "value"
}

variable "agent_display_name" {
  default = "displayName"
}

variable "agent_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "agent_os_version" {
  default = "osVersion"
}

variable "agent_state" {
  default = "ACTIVE"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

variable "environment_display_name" {
  default = "displayName"
}
resource "oci_cloud_bridge_environment" "test_environment" {
  compartment_id = var.compartment_id
  display_name   = var.environment_display_name
}

resource "oci_cloud_bridge_agent" "test_agent" {
  #Required
  agent_type     = var.agent_agent_type
  agent_version  = var.agent_agent_version
  compartment_id = var.compartment_id
  display_name   = var.agent_display_name
  environment_id = oci_cloud_bridge_environment.test_environment.id
  os_version     = var.agent_os_version

  #Optional
  freeform_tags = var.agent_freeform_tags
}

data "oci_cloud_bridge_agents" "test_agents" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  agent_id       = oci_cloud_bridge_agent.test_agent.id
  display_name   = var.agent_display_name
  environment_id = oci_cloud_bridge_environment.test_environment.id
  state          = var.agent_state
}
