// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ccc_upgrade_schedule_access_level" {
  default = "RESTRICTED"
}

variable "ccc_upgrade_schedule_compartment_id_in_subtree" {
  default = false
}

variable "ccc_upgrade_schedule_defined_tags_value" {
  default = "value"
}

variable "ccc_upgrade_schedule_description" {
  default = "Month-start upgrade window"
}

variable "ccc_upgrade_schedule_display_name" {
  default = "example_cccUpgradeSchedule"
}

variable "ccc_upgrade_schedule_display_name_contains" {
  default = "displayNameContains"
}

variable "ccc_upgrade_schedule_events_description" {
  default = "Month-start upgrade window"
}

variable "ccc_upgrade_schedule_events_schedule_event_duration" {
  default = "P2DT6H"
}

variable "ccc_upgrade_schedule_events_schedule_event_recurrences" {
  default = "FREQ=MONTHLY;INTERVAL=1;BYMONTHDAY=1"
}

variable "ccc_upgrade_schedule_events_time_start" {
  default = "2024-01-25T22:00:00Z"
}

variable "ccc_upgrade_schedule_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "ccc_upgrade_schedule_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
  		#Required
		compartment_id = var.compartment_id
  		description = "example tag namespace"
  		name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

		is_retired = false
}

resource "oci_identity_tag" "tag1" {
  		#Required
  		description = "example tag"
  		name = "example-tag"
        tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

		is_retired = false
}

resource "oci_compute_cloud_at_customer_ccc_upgrade_schedule" "test_ccc_upgrade_schedule" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.ccc_upgrade_schedule_display_name
  events {
    #Required
    description             = var.ccc_upgrade_schedule_events_description
    schedule_event_duration = var.ccc_upgrade_schedule_events_schedule_event_duration
    time_start              = var.ccc_upgrade_schedule_events_time_start

    #Optional
    schedule_event_recurrences = var.ccc_upgrade_schedule_events_schedule_event_recurrences
  }

  #Optional
  defined_tags  = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", var.ccc_upgrade_schedule_defined_tags_value)
  description   = var.ccc_upgrade_schedule_description
  freeform_tags = var.ccc_upgrade_schedule_freeform_tags
}

data "oci_compute_cloud_at_customer_ccc_upgrade_schedules" "test_ccc_upgrade_schedules" {

  #Optional
  access_level              = var.ccc_upgrade_schedule_access_level
  ccc_upgrade_schedule_id   = oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id
  compartment_id            = var.compartment_id
  compartment_id_in_subtree = var.ccc_upgrade_schedule_compartment_id_in_subtree
  display_name              = var.ccc_upgrade_schedule_display_name
  display_name_contains     = var.ccc_upgrade_schedule_display_name_contains
  state                     = var.ccc_upgrade_schedule_state
}

