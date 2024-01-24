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
  #version = "4.98.0"
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

resource "oci_devops_deploy_pipeline" "test_deploy_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  description  = "description"
  display_name = "displayName"
}

resource "oci_devops_deploy_stage" "test_wait_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
    }
  }
  deploy_stage_type = "WAIT"

  description  = "description"
  display_name = "displayName"
  wait_criteria {
    #Required
    wait_duration = "PT5S"
    wait_type     = "ABSOLUTE_WAIT"
  }
}

resource "oci_devops_deployment" "test_deployment" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deployment_type    = "PIPELINE_DEPLOYMENT"

  #Optional
  display_name                  = "test_deployment"
  trigger_new_devops_deployment = false
  deploy_stage_override_arguments {
    #Required
    items {
      #Required
      deploy_stage_id = oci_devops_deploy_stage.test_wait_deploy_stage.id
      name            = "version"
      value           = "1.0"
    }
  }
}