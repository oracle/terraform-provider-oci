// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
}

variable "time_usage_ended" {
  default = "2023-11-01T00:00:00.000Z"
}

variable "time_usage_started" {
  default = "2023-09-01T00:00:00.000Z"
}

variable "dimensions_value" {
  default = "dimensions_value"
}

resource "oci_metering_computation_usage_carbon_emission" "test_usage_carbon_emission" {
  #Required
  tenant_id          = var.tenancy_ocid
  time_usage_ended   = var.time_usage_ended
  time_usage_started = var.time_usage_started

  #Optional
  compartment_depth = 1
  group_by   = ["service"]
}

data "oci_metering_computation_usage_carbon_emissions_config" "test_usage_carbon_emissions_config" {
  tenant_id = var.tenancy_ocid
}
