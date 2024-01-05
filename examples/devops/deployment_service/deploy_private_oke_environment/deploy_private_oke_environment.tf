// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "random_string" "topicname" {
  length  = 10
  special = false
}

resource "random_string" "projectname" {
  length  = 10
  special = false
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = random_string.topicname.result
}

resource "oci_devops_project" "test_project" {
  #Required
  compartment_id = var.compartment_ocid
  name           = join("", ["A", random_string.projectname.result])
  notification_config {
    #Required
    topic_id = oci_ons_notification_topic.test_notification_topic.id
  }
}

resource "oci_devops_deploy_environment" "test_deploy_private_oke_environment" {
  #Required
  deploy_environment_type = "OKE_CLUSTER"
  project_id              = oci_devops_project.test_project.id
  cluster_id              = oci_containerengine_cluster.test_cluster.id
  display_name            = "privateOkeDeployEnvironment"
  network_channel {
    #Required
    network_channel_type = "PRIVATE_ENDPOINT_CHANNEL"
    subnet_id            = oci_core_subnet.clusterSubnet_1.id

    #Optional
    nsg_ids = [oci_core_network_security_group.test_nsg.id]
  }
}