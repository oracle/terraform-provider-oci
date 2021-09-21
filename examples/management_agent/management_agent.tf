// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "managed_agent_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_management_agent_management_agent" "test_management_agent" {
  #Required
  managed_agent_id = var.managed_agent_id

  #Optional
  display_name      = "TF_test_agent"
  deploy_plugins_id = [data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id]
}

data "oci_management_agent_management_agents" "test_management_agents" {
  #Required
  compartment_id = var.compartment_ocid
}

resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  allowed_key_install_count = "200"
  display_name              = "displayName"
  time_expires              = "2022-05-27T17:27:44.398Z"
}

data "oci_management_agent_management_agent_install_keys" "test_management_agent_install_keys" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
  #Required
  management_agent_install_key_id = oci_management_agent_management_agent_install_key.test_management_agent_install_key.id
}

data "oci_management_agent_management_agent_plugins" "test_management_agent_plugins" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_management_agent_management_agent_images" "test_management_agent_images" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_management_agent_management_agent_available_histories" "test_management_agent_available_histories" {
  #Required
  management_agent_id = var.managed_agent_id

  #Optional
  time_availability_status_ended_greater_than      = "2020-01-15T01:01:01.000Z"
  time_availability_status_started_less_than       = "2029-09-28T01:01:01.000Z"

}

data "oci_management_agent_management_agent_get_auto_upgradable_config" "test_management_agent_get_auto_upgradable_config" {
  #Required
  compartment_id = var.tenancy_ocid
}