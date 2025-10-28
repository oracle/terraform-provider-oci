// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "log_configuration_source_category" {
  default = "vcn"
}

variable "log_configuration_source_parameters" {
  type = map(string)

  default = {
    capture_filter = ""
  }
}

variable "log_configuration_source_resource" {
  default = ""
}

variable "log_configuration_source_service" {
  default = "flowlogs"
}

variable "log_configuration_source_source_type" {
  default = "OCISERVICE"
}

variable "defined_tags_value" {
  default = "tf-value"
}

variable "freeform_tags_value" {
  default = {
    "Department" = "Finance log"
  }
}

variable "log_log_type" {
  type = map(string)

  default = {
    service = "SERVICE"
    custom  = "CUSTOM"
  }
}

variable "log_source_resource" {
  default = ""
}

variable "log_source_service" {
  default = "flowlogs"
}

variable "test_log_group_id" {
  default = ""
}
variable "test_log_name" {
  default = "tf-exampleLog"
}


resource "oci_logging_log" "test_log" {
  #Required
  display_name = var.test_log_name
  log_group_id = var.test_log_group_id
  log_type     = var.log_log_type.service

  #Optional
  configuration {
    #Required
    source {
      #Required
      category    = var.log_configuration_source_category
      resource    = var.log_configuration_source_resource
      service     = var.log_configuration_source_service
      source_type = var.log_configuration_source_source_type

      #Optional
      parameters = var.log_configuration_source_parameters
    }

    #Optional
    compartment_id = ""
  }

  freeform_tags      = var.freeform_tags_value
  retention_duration = "30"

}

data "oci_logging_logs" "test_logs" {
  #Required
  log_group_id = var.test_log_group_id

  #Optional
  display_name    = var.test_log_name
  log_type        = var.log_log_type.service
  source_resource = var.log_source_resource
  source_service  = var.log_source_service
  state           = "ACTIVE"
}

output "test_log_id" {
  value = oci_logging_log.test_log.id
}
