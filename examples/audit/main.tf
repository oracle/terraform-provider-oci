// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to use the audit_configuration Resource to set the event retention period and list events with
 * the audit_events Data Source.
 */

// These variables would commonly be defined as environment variables or sourced in a .env file
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
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_audit_configuration" "audit_configuration" {
  compartment_id        = var.tenancy_ocid
  retention_period_days = "99"
}

data "oci_audit_configuration" "audit_configuration" {
  compartment_id = var.tenancy_ocid
}

output "retention_period_days" {
  value = data.oci_audit_configuration.audit_configuration.retention_period_days
}

data "oci_audit_events" "audit_events" {
  compartment_id = var.compartment_ocid

  # NOTE: These dates should be updated to applicable ranges of events within your tenancy.
  # CAUTION: Specifying wide date ranges may pull excessively large sets of event data from the Audit service.
  start_time = timeadd(timestamp(), "-1m")

  end_time = timestamp()
}

output "audit_events" {
  value = data.oci_audit_events.audit_events.audit_events
}

