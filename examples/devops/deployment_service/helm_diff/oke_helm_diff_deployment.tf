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

resource "oci_devops_deploy_pipeline" "test_deploy_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  description   = "description"
  display_name  = "displayName"
}

resource "oci_devops_deploy_artifact" "test_deploy_helm_artifact" {
  project_id              = oci_devops_project.test_project.id
  display_name = "Display_name"
  deploy_artifact_type = "HELM_CHART"
  argument_substitution_mode = "NONE"
  deploy_artifact_source {
    deploy_artifact_source_type = "HELM_CHART"
    chart_url = "iad.ocir.io/ax022wvgmjpq/fake"
    deploy_artifact_version = "0.1"
  }
}

resource "oci_devops_deploy_environment" "test_deploy_oke_environment" {
  #Required
  deploy_environment_type = "OKE_CLUSTER"
  project_id              = oci_devops_project.test_project.id
  cluster_id              = oci_containerengine_cluster.test_cluster.id
  display_name            = "okeDeployEnvironment"
}

resource "oci_devops_deploy_stage" "test_helm_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_stage.test_oke_canary_traffic_shift_deploy_stage.id
    }
  }
  deploy_stage_type = "OKE_HELM_CHART_DEPLOYMENT"
  release_name = "release-name"
  helm_chart_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_helm_artifact.id
}

resource "oci_devops_deployment" "test_deployment" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deployment_type = "PIPELINE_DEPLOYMENT"

  deploy_stage_id = oci_devops_deploy_stage.test_helm_deploy_stage.id
  display_name = "HelmDiffDeployment"
  deployment_arguments {

    items {
      name = "PLAN_DRY_RUN"
      value = "true"
    }

  }
}