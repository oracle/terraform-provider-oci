// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_ocid" {
  default = "example-compartment-id"
}

variable "fleet_ocid" {
  default = "example-fleet-id"
}

variable "fleet_log_group_ocid" {
  default = "example-log-group-id"
}

variable "fleet_inventory_log_ocid" {
  default = "example-inventory-log-id"
}

variable "fleet_operation_log_ocid" {
  default = "example-operation-log-id"
}

variable "fleet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fleet_defined_tags" {
  default  = { "example-tag-namespace-all.example-tag" = "value" }
}

variable "managed_instance_ocid" {
  default="example-managed-instance-id"
}

variable "host_id" {}

variable "host_name" {
  default="example-host-name"
}

variable "application_id" {
  default = "example-application-id"
}

variable "application_name" {
  default = "example-application-name"
}

variable "time_start" {
  default = "2024-01-20T15:15:15.000Z"
}

variable "time_end" {
  default = "2024-01-20T20:20:20.000Z"
}

variable "analytic_bucket_namespace" {
  default = "example-namespace"
}

variable "analytic_bucket_name" {
  default = "example-analytic-bucket-name"
}

variable "crypto_event_log_ocid" {
  default = "example-crypto-event-log-id"
}