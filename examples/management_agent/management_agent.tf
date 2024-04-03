// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "root_compartment_ocid" {}

variable "shape" {
  default = "VM.Standard.E4.Flex"
}

variable "subnet" {
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


data "oci_management_agent_management_agents" "find_agent" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_status = "ACTIVE"
  display_name = "terraformTest"
  state = "ACTIVE"
}

resource "oci_management_agent_management_agent" "test_management_agent" {
  #Required
  managed_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id

  #Optional
  deploy_plugins_id = [data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.1.id]
}


data "oci_management_agent_management_agents" "test_management_agents_subtree" {
  #Required
  compartment_id = var.root_compartment_ocid

  #Optional
  access_level = "ACCESSIBLE"
  availability_status = "ACTIVE"
  compartment_id_in_subtree = true
}

data "oci_management_agent_management_agents" "test_management_agents" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  access_level = "ACCESSIBLE"
  availability_status = "ACTIVE"
  compartment_id_in_subtree = false
  display_name = "my agent"
  host_id = "hostid"
  install_type = "AGENT"
  is_customer_deployed = true
  platform_type = ["LINUX"]
  plugin_name = ["Plugin Name"]
  state = "ACTIVE"
  version = ["210101.0101"]
}

resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  allowed_key_install_count = "200"
  display_name              = "terraformTest"
  time_expires              = "2024-09-27T17:27:44.398Z"
}

resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key_unlimited" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name              = "terraformTest"
  is_unlimited              = true
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

  #Optional
  agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id

}

data "oci_management_agent_management_agent_images" "test_management_agent_images" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_management_agent_management_agent_available_histories" "test_management_agent_available_histories" {
  #Required
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id

  #Optional
  time_availability_status_ended_greater_than      = "2020-01-15T01:01:01.000Z"
  time_availability_status_started_less_than       = "2029-09-28T01:01:01.000Z"

}

data "oci_management_agent_management_agent_get_auto_upgradable_config" "test_management_agent_get_auto_upgradable_config" {
  #Required
  compartment_id = var.tenancy_ocid
}


resource "oci_core_instance" "instance" {

  agent_config {
    is_monitoring_disabled = false
    is_management_disabled = false
    plugins_config {
      desired_state = "ENABLED"
      name          = "Management Agent"
    }

  }

  availability_domain = data.oci_identity_availability_domains.ads.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  shape               = var.shape
  shape_config {
    ocpus         = 1
    memory_in_gbs = 8
  }

  source_details {
    source_type = "image"
    source_id   =  data.oci_core_images.compute_images.images[0].id
  }

  create_vnic_details {
    subnet_id              = var.subnet
    display_name           = "example_vnic"
    assign_public_ip       = false
    skip_source_dest_check = false
  }

  display_name = "Terraform example Agent host"

}
data "oci_identity_availability_domains" "ads" {
  compartment_id = var.compartment_ocid
}
data "oci_core_images" "compute_images" {
  #Required
  compartment_id           = var.compartment_ocid
  operating_system         = "Oracle Autonomous Linux"
  operating_system_version = "7.9"
}


data "oci_management_agent_management_agents" "find_compute_agent" {
  compartment_id   = var.compartment_ocid
  host_id          = oci_core_instance.instance.id
  wait_for_host_id = 10
}



resource "oci_management_agent_management_agent" "test_compute_management_agent" {
  freeform_tags     = { "TestingTag" : "TestingValue" }
  managed_agent_id  = data.oci_management_agent_management_agents.find_compute_agent.management_agents[0].id
  deploy_plugins_id = [data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id]
}