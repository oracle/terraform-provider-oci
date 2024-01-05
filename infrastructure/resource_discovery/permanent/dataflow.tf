// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_dataflow_application" "dataflow_application_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "dataflowApplicationRD"
  driver_shape   = "${var.application_driver_shape}"
  executor_shape = "${var.application_executor_shape}"
  file_uri       = "${var.application_file_uri}"
  language       = "${var.application_language}"
  num_executors  = "${var.application_num_executors}"
  spark_version  = "${var.application_spark_version}"

  #Optional
  arguments  = ["arguments2"]
  class_name = "Hello"

  configuration = {
    "spark.shuffle.io.maxRetries" = "11"
  }

  description     = "description"
  logs_bucket_uri = "${var.application_logs_bucket_uri}"
  archive_uri     = "${var.application_archive_uri}"

  #Required
  parameters {
    name  = "name"
    value = "value"
  }

  warehouse_bucket_uri = "${var.application_warehouse_bucket_uri}"
}

data "oci_dataflow_applications" "dataflow_applications_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "dataflowApplicationsRD"
}
