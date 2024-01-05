// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "agent_plugin_desired_state" {
  default = "ENABLED"
}

variable "agent_plugin_plugin_name" {
  default = "pluginName"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_agent_plugin" "test_agent_plugin" {
  #Required
  agent_id    = oci_cloud_bridge_agent.test_agent.id
  plugin_name = var.agent_plugin_plugin_name

  #Optional
  desired_state = var.agent_plugin_desired_state
}

data "oci_cloud_bridge_agent_plugins" "test_agent_plugins" {
  #Required
  agent_id    = oci_cloud_bridge_agent.test_agent.id
  plugin_name = var.agent_plugin_plugin_name
}

