// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "build_pipeline_build_pipeline_parameters_items_default_value" {
  default = "defaultValue"
}

variable "build_pipeline_build_pipeline_parameters_items_description" {
  default = "description"
}

variable "build_pipeline_build_pipeline_parameters_items_name" {
  default = "name"
}

variable "build_pipeline_defined_tags_value" {
  default = "value"
}

variable "build_pipeline_description" {
  default = "description"
}

variable "build_pipeline_display_name" {
  default = "displayName"
}

variable "build_pipeline_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "build_pipeline_id" {
  default = "id"
}

variable "build_pipeline_state" {
  default = "ACTIVE"
}



provider "oci" {
  # version          = "4.110.0"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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

resource "oci_devops_build_pipeline" "test_build_pipeline" {
  #Required
  project_id = oci_devops_project.test_project.id

  #Optional
  build_pipeline_parameters {
    #Required
    items {
      #Required
      default_value = var.build_pipeline_build_pipeline_parameters_items_default_value
      name          = var.build_pipeline_build_pipeline_parameters_items_name

      #Optional
      description = var.build_pipeline_build_pipeline_parameters_items_description
    }
  }
  description   = var.build_pipeline_description
  display_name  = var.build_pipeline_display_name
}

