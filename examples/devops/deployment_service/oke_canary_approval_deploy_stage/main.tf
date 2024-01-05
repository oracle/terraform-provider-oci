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
  display_name = "deployPipeline"
}

resource "oci_devops_deploy_artifact" "test_deploy_inline_artifact" {
  project_id                 = oci_devops_project.test_project.id
  display_name               = "inlineDeployArtifact"
  deploy_artifact_type       = "KUBERNETES_MANIFEST"
  argument_substitution_mode = "NONE"
  deploy_artifact_source {
    deploy_artifact_source_type = "INLINE"
    base64encoded_content       = "YXBpVmVyc2lvbjogYmF0Y2gvdjEKa2luZDogSm9iCm1ldGFkYXRhOgogIGdlbmVyYXRlTmFtZTogaGVsbG93b3JsZAogIGxhYmVsczoKICAgIGFwcDogaGVsbG93b3JsZApzcGVjOgogIHR0bFNlY29uZHNBZnRlckZpbmlzaGVkOiAxMjAKICB0ZW1wbGF0ZToKICAgIHNwZWM6CiAgICAgIGNvbnRhaW5lcnM6CiAgICAgICAgLSBuYW1lOiBoZWxsb3dvcmxkCiAgICAgICAgICBpbWFnZTogcGh4Lm9jaXIuaW8vYXgwMjJ3dmdtanBxL2hlbGxvd29ybGQtb2tlLXZlcmlmaWVyOmxhdGVzdAogICAgICAgICAgY29tbWFuZDoKICAgICAgICAgICAgLSAiL2Jpbi9iYXNoIgogICAgICAgICAgICAtICItYyIKICAgICAgICAgICAgLSAic2xlZXAgMjsgZWNobyBIZWxsbyBXb3JsZDsiCiAgICAgIHJlc3RhcnRQb2xpY3k6IE5ldmVy"
  }
}

resource "oci_devops_deploy_environment" "test_deploy_oke_environment" {
  #Required
  deploy_environment_type = "OKE_CLUSTER"
  project_id              = oci_devops_project.test_project.id
  cluster_id              = oci_containerengine_cluster.test_cluster.id
  display_name            = "okeDeployEnvironment"
}

resource "oci_devops_deploy_stage" "test_oke_canary_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
    }
  }
  deploy_stage_type = "OKE_CANARY_DEPLOYMENT"

  description  = "description"
  display_name = "okeCanaryDeployStage"

  oke_cluster_deploy_environment_id       = oci_devops_deploy_environment.test_deploy_oke_environment.id
  kubernetes_manifest_deploy_artifact_ids = [oci_devops_deploy_artifact.test_deploy_inline_artifact.id]
  canary_strategy {
    #Required
    strategy_type = "NGINX_CANARY_STRATEGY"
    namespace     = "namespace"
    ingress_name  = "ingressName"
  }
}

resource "oci_devops_deploy_stage" "test_oke_canary_traffic_shift_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_stage.test_oke_canary_deploy_stage.id
    }
  }
  deploy_stage_type = "OKE_CANARY_TRAFFIC_SHIFT"

  oke_canary_deploy_stage_id = oci_devops_deploy_stage.test_oke_canary_deploy_stage.id

  rollout_policy {
    ramp_limit_percent     = 5.0
    batch_delay_in_seconds = 5
    batch_count            = 1
  }

  display_name = "okeCanaryTrafficShiftDeployStage"
}

resource "oci_devops_deploy_stage" "test_oke_canary_approval_deploy_stage" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deploy_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_deploy_stage.test_oke_canary_traffic_shift_deploy_stage.id
    }
  }
  deploy_stage_type = "OKE_CANARY_APPROVAL"

  oke_canary_traffic_shift_deploy_stage_id = oci_devops_deploy_stage.test_oke_canary_traffic_shift_deploy_stage.id
  approval_policy {
    approval_policy_type         = "COUNT_BASED_APPROVAL"
    number_of_approvals_required = 1
  }

  display_name = "okeCanaryApprovalDeployStage"
}