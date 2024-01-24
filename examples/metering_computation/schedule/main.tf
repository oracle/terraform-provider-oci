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

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "time_scheduled" {
  default = "2022-10-02T00:00:00.000Z"
}

resource "oci_metering_computation_schedule" "test_schedule" {
  #Required
  compartment_id       = var.tenancy_ocid
  name                 = "name"
  time_scheduled       = var.time_scheduled
  schedule_recurrences = "DAILY"
  result_location {
    #Required
    bucket        = "costCsv"
    location_type = "OBJECT_STORAGE"
    namespace     = "r1uoqjtybbv4"
    region        = "us-seattle-1"
  }
  query_properties {
    #Required
    date_range {
      #Required
      date_range_type         = "DYNAMIC"
      dynamic_date_range_type = "LAST_7_DAYS" 
    }
    granularity = "DAILY"
  }

  #Optional
  description          = "description"
  output_file_format   = "CSV"
}
