// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

variable "github_access_token_vault_id" {
}

variable "gitlab_access_token_vault_id" {
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
  name           = random_string.projectname.result
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

resource "oci_devops_deploy_artifact" "test_deploy_ocir_artifact" {
  project_id              = oci_devops_project.test_project.id
  display_name = "Display_name"
  deploy_artifact_type = "DOCKER_IMAGE"
  argument_substitution_mode = "NONE"
  deploy_artifact_source {
    deploy_artifact_source_type = "OCIR"
    image_uri = "iad.ocir.io/ax022wvgmjpq/fake/hello-java:0.0.2"
    image_digest = "38598585.fakedigest1"
  }
}

resource "oci_devops_deploy_artifact" "test_deploy_generic_artifact" {
  argument_substitution_mode = "NONE"
  deploy_artifact_source {
    deploy_artifact_path = "helloworld-oke.yaml"
    deploy_artifact_source_type = "GENERIC_ARTIFACT"
    deploy_artifact_version = "v1"
    repository_id = "ocid1.artifactrepository.oc1.iad.0.amaaaaaansx72maa7nbce5ebmsqkan3msgyosvxe5d5a6jghn5su6ykgw7vq"
  }
  deploy_artifact_type = "KUBERNETES_MANIFEST"
  project_id = "${oci_devops_project.test_project.id}"
}

resource "oci_devops_deploy_environment" "test_deploy_environment" {
  #Required
  deploy_environment_type = "OKE_CLUSTER"
  project_id              = oci_devops_project.test_project.id
  cluster_id = "ocid1.cluster.oc1.us-ashburn-1.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstocluster1"
  display_name = "Display_name"
}

resource "oci_devops_deploy_environment" "test_deploy_function_environment" {
  #Required
  deploy_environment_type = "FUNCTION"
  project_id              = oci_devops_project.test_project.id
  function_id = "ocid1.fnfunc.oc1.us-ashburn-1.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstofunction1"
  display_name = "Display_name"
}

resource "oci_devops_deploy_environment" "test_deploy_instance_group_environment" {
  compute_instance_group_selectors {
    items {
      compute_instance_ids = ["ocid1.instance.oc1.iad.anuwcljtnsx72macffe5fbkzbj4eerle5ot56g2cexj3jvfsr242pye44ghq"]
      selector_type = "INSTANCE_IDS"
    }
  }
  deploy_environment_type = "COMPUTE_INSTANCE_GROUP"
  project_id = "${oci_devops_project.test_project.id}"
}

resource "oci_devops_deploy_stage" "test_deploy_stage" {
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

  description                                  = "description"
  display_name                                 = "displayName"
  wait_criteria {
    #Required
    wait_duration = "PT5S"
    wait_type     = "ABSOLUTE_WAIT"
  }
}

resource "oci_devops_deploy_stage" "test_deploy_stage" {
  compute_instance_group_deploy_environment_id = "${oci_devops_deploy_environment.test_deploy_instance_group_environment.id}"
  defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  deploy_pipeline_id = "${oci_devops_deploy_pipeline.test_deploy_pipeline.id}"
  deploy_stage_predecessor_collection {
    items {
      id = "${oci_devops_deploy_pipeline.test_deploy_pipeline.id}"
    }
  }
  deploy_stage_type = "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"
  deployment_spec_deploy_artifact_id = "${oci_devops_deploy_artifact.test_deploy_generic_artifact.id}"
  description = "description"
  display_name = "displayName"
  failure_policy {
    failure_count = "1"
    policy_type = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"
  }
  freeform_tags = {
    "bar-key" = "value"
  }
  load_balancer_config {
    backend_port = "8080"
    listener_name = "LoadBalancerListener"
    load_balancer_id = "ocid1.loadbalancer.oc1.phx.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstofake1"
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

resource "oci_devops_deployment" "test_deployment" {
  #Required
  deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
  deployment_type    = "PIPELINE_DEPLOYMENT"

  #Optional
  display_name           = "test_deployment"

}