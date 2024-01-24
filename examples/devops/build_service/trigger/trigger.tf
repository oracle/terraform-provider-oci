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
  # version          = "4.83.0"
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

# Build pipeline
resource "oci_devops_build_pipeline" "test_build_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  description   = "Build pipeline"
  display_name  = "build_pipeline"
}

# Github trigger
resource "oci_devops_trigger" "test_github_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "GITHUB"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "GITHUB"

  #Optional
  description   = "Github Trigger"
  display_name  = "github-trigger"
}

# Gitlab trigger resource
resource "oci_devops_trigger" "test_gitlab_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "GITLAB"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "GITLAB"

  #Optional
  description   = "Gitlab Trigger"
  display_name  = "gitlab-trigger"
  depends_on = [oci_devops_trigger.test_github_trigger]
}

# Gitlab Server trigger resource
resource "oci_devops_trigger" "test_gitlab_server_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "GITLAB_SERVER"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "GITLAB_SERVER"

  #Optional
  description   = "Gitlab server Trigger"
  display_name  = "gitlab-server-trigger"
  depends_on = [oci_devops_trigger.test_gitlab_trigger]
}

# Bitbucket Server trigger resource
resource "oci_devops_trigger" "test_bitbucket_server_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "BITBUCKET_SERVER"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "BITBUCKET_SERVER"

  #Optional
  description   = "Bitbucket server Trigger"
  display_name  = "bitbucket-server-trigger"
  depends_on = [oci_devops_trigger.test_gitlab_server_trigger]
}

# VBS trigger resource
resource "oci_devops_trigger" "test_vbs_trigger" {
  #Required
  actions {
    #Required
    build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
    type              = "TRIGGER_BUILD_PIPELINE"

    #Optional
    filter {
      #Required
      trigger_source = "VBS"

      #Optional
      events = ["PUSH"]
      include {

        #Optional
        base_ref = "baseRef"
        head_ref = "headRef"
      }
    }
  }
  project_id     = oci_devops_project.test_project.id
  trigger_source = "VBS"

  #Optional
  description   = "VBS Trigger"
  display_name  = "vbs-trigger"
  depends_on = [oci_devops_trigger.test_bitbucket_server_trigger]
}