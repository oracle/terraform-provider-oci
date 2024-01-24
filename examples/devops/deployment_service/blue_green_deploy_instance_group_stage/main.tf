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
  #version = "4.62.0"
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
  name = join("", [
    "A",
    random_string.projectname.result])
  notification_config {
    #Required
    topic_id = oci_ons_notification_topic.test_notification_topic.id
  }
}

resource "oci_devops_deploy_pipeline" "test_deploy_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  description  = "description"
  display_name = "testDeployPipeline"
  deploy_pipeline_parameters {
    items {
      #Required
      name = "test"
      description = "test pipeline"
    }
  }
}

resource "oci_devops_deploy_artifact" "test_deploy_artifact" {
  argument_substitution_mode = "SUBSTITUTE_PLACEHOLDERS"
  deploy_artifact_type       = "DEPLOYMENT_SPEC"
  project_id                 = oci_devops_project.test_project.id

  deploy_artifact_source {
    deploy_artifact_source_type = "INLINE"
    base64encoded_content       = file("${path.module}/manifest/spec.yaml")
  }
}

resource "oci_devops_deploy_environment" "test_deploy_instance_group_environment_a" {
  compute_instance_group_selectors {
    items {
      compute_instance_ids = [oci_core_instance.example_instance_a.id]
      selector_type = "INSTANCE_IDS"
    }
  }
  deploy_environment_type = "COMPUTE_INSTANCE_GROUP"
  project_id = oci_devops_project.test_project.id
}

resource "oci_devops_deploy_environment" "test_deploy_instance_group_environment_b" {
  compute_instance_group_selectors {
    items {
      compute_instance_ids = [oci_core_instance.example_instance_b.id]
      selector_type = "INSTANCE_IDS"
    }
  }
  deploy_environment_type = "COMPUTE_INSTANCE_GROUP"
  project_id = oci_devops_project.test_project.id
}

resource "oci_devops_deploy_stage" "test_blue_green_deploy_instance_group_stage" {
  deploy_environment_id_a = oci_devops_deploy_environment.test_deploy_instance_group_environment_a.id
  deploy_environment_id_b = oci_devops_deploy_environment.test_deploy_instance_group_environment_b.id
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    items {
      id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
    }
  }
  deploy_stage_type = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT"
  deployment_spec_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
  description = "description"
  display_name = "testBlueGreenDeployInstanceGroupStage"
  failure_policy {
    failure_count = "1"
    policy_type = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"
  }
  freeform_tags = {
    "bar-key" = "value"
  }
  test_load_balancer_config {
    backend_port = "8080"
    listener_name = oci_load_balancer_listener.test_load_balancer_listener.name
    load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
  }
  production_load_balancer_config {
    backend_port = "8081"
    listener_name = oci_load_balancer_listener.prod_load_balancer_listener.name
    load_balancer_id = oci_load_balancer_load_balancer.prod_load_balancer.id
  }
  rollback_policy {
    policy_type = "AUTOMATED_STAGE_ROLLBACK_POLICY"
  }
  rollout_policy {
    batch_count = "5"
    batch_delay_in_seconds = "10"
    policy_type = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"
  }
}

resource "oci_devops_deploy_stage" "test_blue_green_traffic_shift_instance_group_stage" {
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    items {
      id = oci_devops_deploy_stage.test_blue_green_deploy_instance_group_stage.id
    }
  }
  deploy_stage_type = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT"
  compute_instance_group_blue_green_deployment_deploy_stage_id = oci_devops_deploy_stage.test_blue_green_deploy_instance_group_stage.id
  description = "description"
  display_name = "testBlueGreenTrafficShiftInstanceGroupStage"
  freeform_tags = {
    "bar-key" = "value"
  }
}