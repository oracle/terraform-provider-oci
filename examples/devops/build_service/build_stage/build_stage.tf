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

variable "access_token_vault_secret_id" {
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

variable "base_url" {
  default = "https://test-server.com"
}

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = join("", ["A", random_string.topicname.result])
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

# GITHUB external connection
resource "oci_devops_connection" "test_github_connection" {
  #Required
  access_token    = var.access_token_vault_secret_id
  connection_type = "GITHUB_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id

  #Optional
  description   = "Github access token"
  display_name  = "github_access_token"
}

# GITLAB external connection
resource "oci_devops_connection" "test_gitlab_connection" {
  #Required
  access_token    = var.access_token_vault_secret_id
  connection_type = "GITLAB_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id

  #Optional
  description   = "Gitlab access token"
  display_name  = "gitlab_access_token"
  depends_on = [oci_devops_connection.test_github_connection]
}

# Gitlab Server external connection
resource "oci_devops_connection" "test_gitlab_server_connection" {
  #Required
  access_token    = var.access_token_vault_secret_id
  connection_type = "GITLAB_SERVER_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id
  base_url        = var.base_url

  #Optional
  description   = "Gitlab server access token"
  display_name  = "gitlab_server_access_token"
  depends_on = [oci_devops_connection.test_gitlab_connection]
}

# Bitbucket server external connection
resource "oci_devops_connection" "test_bitbucket_server_connection" {
  #Required
  access_token    = var.access_token_vault_secret_id
  connection_type = "BITBUCKET_SERVER_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id
  base_url        = var.base_url

  #Optional
  description   = "Bitbucket server access token"
  display_name  = "bitbucket_server_access_token"
  depends_on = [oci_devops_connection.test_gitlab_server_connection]
}

# VBS external connection
resource "oci_devops_connection" "test_vbs_connection" {
  #Required
  access_token    = var.access_token_vault_secret_id
  connection_type = "VBS_ACCESS_TOKEN"
  project_id      = oci_devops_project.test_project.id
  base_url        = var.base_url

  #Optional
  description   = "VBS access token"
  display_name  = "vbs_access_token"
  depends_on = [oci_devops_connection.test_bitbucket_server_connection]
}
# Build pipeline
resource "oci_devops_build_pipeline" "test_build_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  #Optional
  build_pipeline_parameters {
    items {
      #Required
      name = "name"

      #Optional
      default_value = "defaultValue"
      description   = "description"
    }

    items {
      #Required
      name = "name2"

      #Optional
      default_value = "defaultValue2"
      description   = "description2"
    }
  }
  description   = "Build pipeline"
  display_name  = "build_pipeline"
  depends_on = [oci_devops_connection.test_gitlab_connection]
}


# BUILD STAGE
resource "oci_devops_build_pipeline_stage" "test_build_pipeline_build_github_stage" {
  #Required
  build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
  build_pipeline_stage_predecessor_collection {
    #Required
    items {
      #Required
      id = oci_devops_build_pipeline.test_build_pipeline.id
    }
  }
  build_pipeline_stage_type = "BUILD"

  description                        = "Github build stage"
  display_name                       = "github-build-stage"
  build_spec_file                    = "build_spec.yml"
  image                              = "OL7_X86_64_STANDARD_10"
  primary_build_source               = "primaryBuildSource"
  stage_execution_timeout_in_seconds = "10"
  build_runner_shape_config {
    build_runner_type = "CUSTOM"
    memory_in_gbs = 4
    ocpus = 1
  }
  build_source_collection {
    items {
      connection_type   = "GITHUB"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_github_connection.id
      name              = "primaryBuildSource"
      repository_url    = "https://github.com/dlcbld/docktest"
    }
    items {
      connection_type   = "GITLAB"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_gitlab_connection.id
      name              = "source"
      repository_url    = "https://gitlab.com/dlcbld/docktest"
    }
    items {
      connection_type   = "GITLAB_SERVER"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_gitlab_server_connection.id
      name              = "source1"
      repository_url    = "https://gitlabserver.com/dlcbld/docktest"
    }
    items {
      connection_type   = "BITBUCKET_SERVER"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_bitbucket_server_connection.id
      name              = "source2"
      repository_url    = "https://bitbucketserver.com/dlcbld/docktest"
    }
     items {
      connection_type   = "VBS"
      branch            = "branch"
      connection_id     = oci_devops_connection.test_vbs_connection.id
      name              = "source3"
      repository_url    = "https://vbs.com/dlcbld/docktest"
    }
  }
}