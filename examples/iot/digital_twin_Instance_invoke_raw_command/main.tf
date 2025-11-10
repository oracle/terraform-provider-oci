// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "digital_twin_instance_ocid" {}

variable "digital_twin_instance_invoke_raw_command_request_data" {
  default = "requestData"
}

variable "digital_twin_instance_invoke_raw_command_request_data_content_type" {
  default = "text/plain"
}

variable "digital_twin_instance_invoke_raw_command_request_data_format" {
  #TEXT, JSON or BINARY
  default = "TEXT"
}

variable "digital_twin_instance_invoke_raw_command_request_duration" {
  default = "PT01M"
}

variable "digital_twin_instance_invoke_raw_command_request_endpoint" {
  default = "/requestEndpoint"
}

variable "digital_twin_instance_invoke_raw_command_response_duration" {
  default = "PT01M"
}

variable "digital_twin_instance_invoke_raw_command_response_endpoint" {
  default = "/responseEndpoint"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_digital_twin_instance_invoke_raw_command" "test_digital_twin_instance_invoke_raw_command" {
  #Required
  digital_twin_instance_id = var.digital_twin_instance_ocid
  request_data_format      = var.digital_twin_instance_invoke_raw_command_request_data_format
  request_endpoint         = var.digital_twin_instance_invoke_raw_command_request_endpoint

  #Optional
  request_data              = var.digital_twin_instance_invoke_raw_command_request_data
  request_data_content_type = var.digital_twin_instance_invoke_raw_command_request_data_content_type
  request_duration          = var.digital_twin_instance_invoke_raw_command_request_duration
  response_duration         = var.digital_twin_instance_invoke_raw_command_response_duration
  response_endpoint         = var.digital_twin_instance_invoke_raw_command_response_endpoint
}


