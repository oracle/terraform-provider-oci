// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

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

variable "time_usage_ended" {
  default = "2020-01-02T00:00:00.000Z"
}

variable "time_usage_started" {
  default = "2020-01-01T00:00:00.000Z"
}

variable "dimensions_value" {
  default = "dimensions_value"
}

variable "usage_forecast_forecast_type" {
  default = "BASIC"
}

variable "usage_forecast_time_forecast_ended" {
  default = "2020-01-03T00:00:00.000Z"
}

variable "usage_forecast_time_forecast_started" {
  default = "2020-01-02T00:00:00.000Z"
}

resource "oci_metering_computation_usage" "test_usage" {
  #Required
  granularity        = "DAILY"
  tenant_id          = var.tenancy_ocid
  time_usage_ended   = var.time_usage_ended
  time_usage_started = var.time_usage_started

  #Optional
  compartment_depth = 1

  filter = <<EOF
{
        "operator": "AND",
        "dimensions": [],
        "tags": [],
        "filters": [
            {
				"operator": "OR",
                "dimensions": [
				    {
                        "key": "compartmentName",
                        "value": "${var.dimensions_value}"
					}
                ],
                "filters": [],
                "tags": []
		    }
        ]
}
EOF

  forecast {
    #Required
    time_forecast_ended = var.usage_forecast_time_forecast_ended

    #Optional
    forecast_type = var.usage_forecast_forecast_type
    time_forecast_started = var.usage_forecast_time_forecast_started
  }
  group_by   = ["service"]
  query_type = "COST"
}

data "oci_metering_computation_configuration" "test_configuration" {
  tenant_id = var.tenancy_ocid
}
