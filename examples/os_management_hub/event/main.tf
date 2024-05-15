// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "event_id" {}
variable "osmh_managed_instance_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


# Event has no public create API. Refer to https://confluence.oci.oraclecorp.com/x/Q5W15Q for creating events.

### Event
# Update tag and compartment
resource "oci_os_management_hub_event" "test_event" {
  compartment_id = "${var.compartment_id}"
  event_id = "${var.event_id}"
  freeform_tags = {
    "Department" = "Finance"
  }
}

# List Event

data "oci_os_management_hub_events" "test_events" {
  compartment_id = "${var.compartment_id}"
  event_summary = "Manually created event 2 for testing caused by <Yijiu>"
  event_summary_contains = "testing"
  filter {
    name = "id"
    values = ["${var.event_id}"]
  }
  id = "${var.event_id}"
  is_managed_by_autonomous_linux = "true"
  resource_id = "${var.osmh_managed_instance_ocid}"
  state = "ACTIVE"
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than = "2088-01-01T00:00:00.000Z"
  type = ["EXPLOIT_ATTEMPT"]
}

# Get Event
data "oci_os_management_hub_event" "test_event" {
  # Required
  event_id = "${var.event_id}"
}
