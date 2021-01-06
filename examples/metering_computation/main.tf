// Copyright (c) 2017, 2019, 2020, Oracle and/or its affiliates. All rights reserved.

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
}

variable "time_usage_started" {
}

variable "dimensions_value" {
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


  group_by   = ["service"]
  query_type = "COST"
}

data "oci_metering_computation_configuration" "test_configuration" {
  tenant_id = var.tenancy_ocid
}

