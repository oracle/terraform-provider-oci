// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.

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

resource "oci_metering_computation_query" "test_query" {
  #Required
  compartment_id = var.tenancy_ocid
  query_definition {
    #Required
    cost_analysis_ui {

      #Optional
      graph               = "BARS"
      is_cumulative_graph = false
    }
    display_name = "tf_display_name"
    report_query {
      #Required
      granularity = "DAILY"
      tenant_id   = var.tenancy_ocid

      #Optional
      compartment_depth = 1
      date_range_name   = "LAST_TEN_DAYS"
      filter            = <<EOF
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
                    "filters": null,
                    "tags": []
    		    }
            ]
    }
    EOF

      group_by = []
      is_aggregate_by_time = false
      query_type           = "COST"
    }
    version = 1
  }
}

data "oci_metering_computation_queries" "test_queries" {
  #Required
  compartment_id = var.tenancy_ocid
}
