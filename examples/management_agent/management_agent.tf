// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0



///
///
//  If running tests against devtest-ashburn,  set is_production_test = false
///
locals {
  is_production_test = false

  oca_devtest_ashburn = base64encode(file("${path.module}/oca_to_devtest-ashburn.txt"))
  oca_devtest_london = base64encode(file("${path.module}/oca_to_devtest-london.txt"))

  cloud_init_script = local.is_production_test ? null : local.oca_devtest_london
  plugin_name = local.is_production_test ? "Logging Analytics" : "Test Plugin Ric"
}



variable "tenancy_ocid" {}
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
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region           = var.region
  # version = "7.2.0"
}

//  Agent simulator must be running against the compartment, see https://confluence.oraclecorp.com/confluence/display/MGMTAGENT/How+to+complete+TERSI+tickets
//  Otherwise this will fail with
# Error: Invalid index
#
# on management_agent.tf line 114, in data "oci_management_agent_management_agent_plugins" "test_management_agent_plugins":
# 114:   agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
# |----------------
# | data.oci_management_agent_management_agents.find_agent.management_agents is empty list of object
data "oci_management_agent_management_agents" "find_agent" {
  compartment_id = var.compartment_ocid
  availability_status = "ACTIVE"
  display_name = "terraformTest"
  state = "ACTIVE"
}


// Find a plugin in the MACS environment
data "oci_management_agent_management_agent_plugins" "test_management_agent_plugins" {
  compartment_id = var.compartment_ocid
  display_name = local.plugin_name
  agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id

}

// Using one of the simulator agents,  add it to terraform and deploy plugin to it
// also modify the freeform tags
resource "oci_management_agent_management_agent" "test_management_agent" {
  managed_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  deploy_plugins_id = [data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id]
  freeform_tags = {"tagKey":"tagValue"}
}

// Test data load of all agents with subtree
data "oci_management_agent_management_agents" "test_management_agents_subtree" {
  compartment_id = var.root_compartment_ocid
  access_level = "ACCESSIBLE"
  availability_status = "ACTIVE"
  compartment_id_in_subtree = true
}

// Test data load for specific agent
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

// Create an install key
//  NOTE this can fail if the time_expires is older than today.. you may have to change this value
resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
  compartment_id = var.compartment_ocid
  allowed_key_install_count = "200"
  display_name              = "terraformTest"
  time_expires              = "2026-02-19T17:27:44.398Z"
}

// Create unlimited install key
resource "oci_management_agent_management_agent_install_key" "test_management_agent_install_key_unlimited" {
  compartment_id = var.compartment_ocid
  display_name              = "terraformTest"
  is_unlimited              = true
}

// Test data load for install keys
data "oci_management_agent_management_agent_install_keys" "test_management_agent_install_keys" {
  compartment_id = var.compartment_ocid
}

// Test data load for specific install key
data "oci_management_agent_management_agent_install_key" "test_management_agent_install_key" {
  management_agent_install_key_id = oci_management_agent_management_agent_install_key.test_management_agent_install_key.id
}

// Test data load for agent images
data "oci_management_agent_management_agent_images" "test_management_agent_images" {
  compartment_id = var.compartment_ocid
}

// Load history agent
data "oci_management_agent_management_agent_available_histories" "test_management_agent_available_histories" {
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  time_availability_status_ended_greater_than      = "2020-01-15T01:01:01.000Z"
  time_availability_status_started_less_than       = "2029-09-28T01:01:01.000Z"

}

// load auto upgrade config for tenancy
data "oci_management_agent_management_agent_get_auto_upgradable_config" "test_management_agent_get_auto_upgradable_config" {
  compartment_id = var.tenancy_ocid
}


// Create a compute instance, with OCA plugin enabled for management agent
// If in devtest-ashburn, add cloud-init script to change endpoint of OCA
resource "oci_core_instance" "instance" {

  agent_config {
    is_monitoring_disabled = false
    is_management_disabled = false
    plugins_config {
      desired_state = "ENABLED"
      name          = "Management Agent"
    }

  }
  metadata = {
    user_data = local.cloud_init_script
    #ssh_authorized_keys = file("/Users/rambridg/.oci/macs_test_host1.pub")

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


// Find the management agent created by the compute instance, it will wait for 10 minutes for the agent to appear
data "oci_management_agent_management_agents" "find_compute_agent" {
  compartment_id   = var.compartment_ocid
  host_id          = oci_core_instance.instance.id
  wait_for_host_id = 10
}


// Update the OCA management agent, deploy plugin and update freeform tags
resource "oci_management_agent_management_agent" "test_compute_management_agent" {
  freeform_tags     = { "TestingTag" : "TestingValue" }
  managed_agent_id  = data.oci_management_agent_management_agents.find_compute_agent.management_agents[0].id
  deploy_plugins_id = [data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id]
}

output "updated_agent" {
  value = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
}