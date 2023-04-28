// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" { default = "ocid1.tenancy.oc1..aaaaaaaa5hqxzxt343cvujpjdoexlksvw4753kvyf5dauvlz3r3e5wiuez6q" }
variable "user_ocid" { default = "ocid1.user.oc1..aaaaaaaa2xxpqj2cu5nmzywoaqd2qpyqrpolvrfsixkdiallqxo54hupj6ba" }
variable "fingerprint" { default = "9b:cc:dd:5c:9d:2c:17:1a:52:32:b4:13:ea:0c:d6:73" }
variable "private_key_path" { default = "/Users/ddevidch/.oci/oci_api_key.pem" }
variable "region" { default = "eu-frankfurt-1" }
variable "compartment_id" {
  default = "ocid1.compartment.oc1..aaaaaaaa7ysncgzhx6xx6thdukow3tly67rlkjkpzwdxhr3zqulhvbujicrq"
}

variable "discovery_schedule_display_name" {
  default = "displayName"
}

variable "discovery_schedule_execution_recurrences" {
  default = "FREQ=DAILY;BYHOUR=6"
}

variable "discovery_schedule_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_bridge_discovery_schedule" "test_discovery_schedule" {
  compartment_id        = var.compartment_id
  execution_recurrences = var.discovery_schedule_execution_recurrences
  display_name  = var.discovery_schedule_display_name
}

data "oci_cloud_bridge_discovery_schedules" "test_discovery_schedules" {
  compartment_id = var.compartment_id
  discovery_schedule_id = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
  display_name          = var.discovery_schedule_display_name
  state                 = var.discovery_schedule_state
}

