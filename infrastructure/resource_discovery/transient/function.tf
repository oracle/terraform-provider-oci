// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_functions_application" "function_application_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "functionApplicationRd"
  subnet_ids     = ["${oci_core_subnet.gateway_subnet_rd.id}"]

  #Optional
  config = "${var.config}"
}

data "oci_functions_applications" "function_applications_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "functionApplicationRD"
  id           = "${oci_functions_application.function_application_rd.id}"
  state        = "${var.application_state}"
}

resource "oci_functions_function" "functions_function_rd" {
  #Required
  application_id = "${oci_functions_application.function_application_rd.id}"
  display_name   = "functionsFunctionRD"
  image          = "${var.function_image}"
  memory_in_mbs  = "128"

  #Optional
  config             = "${var.config}"
  image_digest       = "${var.function_image_digest}"
  timeout_in_seconds = "30"
}

data "oci_functions_functions" "test_functions_rd" {
  #Required
  application_id = "${oci_functions_application.function_application_rd.id}"

  #Optional
  display_name = "functionsFunctionRD"
  id           = "${oci_functions_function.functions_function_rd.id}"
  state        = "AVAILABLE"
}
