// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "notification_topic_id_ocid" {}
variable "delegation_subscription_ids_ocid" { 
  type = list(string)
}

variable "resource_id_ocid" {
  type = list(string)
}

variable "delegation_control_display_name" {
  default = "TersiExample"
}

variable "delegation_control_is_auto_approve_during_maintenance" {
  default = false
}

variable "delegation_control_notification_message_format" {
  default = "JSON"
}



variable "delegation_control_resource_type" {
  default = "VMCLUSTER"
}





provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_delegate_access_control_delegation_control" "test_delegation_control" {
  #Required
  compartment_id              = var.compartment_ocid
  delegation_subscription_ids = var.delegation_subscription_ids_ocid
  display_name                = var.delegation_control_display_name
  notification_message_format = var.delegation_control_notification_message_format
  notification_topic_id       = var.notification_topic_id_ocid
  resource_ids                = var.resource_id_ocid
  resource_type               = var.delegation_control_resource_type


}

data "oci_delegate_access_control_delegation_controls" "test_delegation_controls" {
  compartment_id              = var.compartment_ocid
 }
