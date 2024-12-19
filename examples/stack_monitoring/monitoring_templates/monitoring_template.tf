// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
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

resource "oci_ons_notification_topic" "test_notification_topic" {
  #Required
  compartment_id = var.compartment_ocid
  name           = random_string.topicname.result
}

resource "oci_stack_monitoring_monitoring_template" "test_monitoring_template_example" {
  # Required
  compartment_id = var.compartment_ocid
  display_name = "MT_MonitoringTemplateTerraformExample"
  destinations =  [oci_ons_notification_topic.test_notification_topic.id]
  message_format = "ONS_OPTIMIZED"
  members {
    id = "ocid1.stackmonitoringresourcetype.apache_tomcat"
    type = "RESOURCE_TYPE"
  }

  # Optional
  description = "Example MT for resource type Apache Tomcat"
}

data "oci_stack_monitoring_monitoring_template" "test_monitoring_template_example" {
  # Required
  monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template_example.id
}
