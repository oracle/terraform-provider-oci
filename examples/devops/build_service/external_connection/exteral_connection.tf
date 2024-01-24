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

variable "base_url" {
  # dummy vault-secret OCID
  default = "https://test-server.com"
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