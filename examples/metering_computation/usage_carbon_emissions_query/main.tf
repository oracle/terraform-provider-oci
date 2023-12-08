// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
}

provider "oci" {
  tenancy_ocid        = var.tenancy_ocid
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

resource "oci_metering_computation_usage_carbon_emissions_query" "test_usage_carbon_emissions_query" {
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
      tenant_id   = var.tenancy_ocid

      #Optional
      compartment_depth = 1
      date_range_name   = "LAST_TWO_MONTHS"
      usage_carbon_emissions_query_filter = <<EOF
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
    }
    version = 1
  }
}

data "oci_metering_computation_usage_carbon_emissions_queries" "test_usage_carbon_emissions_queries" {
  #Required
  compartment_id = var.tenancy_ocid
}
