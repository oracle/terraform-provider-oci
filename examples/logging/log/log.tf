// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "log_configuration_source_category" {
  type = map(string)

  default = {
    write = "write"
    read  = "read"
  }
}

variable "log_configuration_source_parameters" {
  default = "parameters"
}

variable "log_configuration_source_resource" {
  default = "srishti-bucket"
}

variable "log_configuration_source_service" {
  default = "objectstorage"
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
  default = "srishti-bucket"
}

variable "log_source_service" {
  default = "objectstorage"
}

variable "test_log_group_id" {}
variable "test_log_name" {
  default = "tf-exampleLog"
}
variable "tag_namespace1_name" {}
variable "tag2_name" {}

resource "oci_logging_log" "test_log" {
  #Required
  display_name = var.test_log_name
  log_group_id = var.test_log_group_id
  log_type     = var.log_log_type.custom

  #Optional
  /*configuration {
    #Required
    source {
      #Required
      category    = var.log_configuration_source_category.write}"
      resource    = var.log_configuration_source_resource}"
      service     = var.log_configuration_source_service}"
      source_type = var.log_configuration_source_source_type}"

      #Optional
      //parameters = var.log_configuration_source_parameters}"
    }

    #Optional
    compartment_id = "ocid1.compartment.oc1..aaaaaaaa4rv5j2vzbrwaztnzvtu7kgswtigms4llcbylelylsqt2l3kl7gaa"
  }*/

  defined_tags = {
    "${var.tag_namespace1_name}.${var.tag2_name}" = var.defined_tags_value
  }
  freeform_tags      = var.freeform_tags_value
  is_enabled         = "false"
  retention_duration = "30"

  lifecycle {
    ignore_changes = [ defined_tags ]
  }
}

data "oci_logging_logs" "test_logs" {
  #Required
  log_group_id = var.test_log_group_id

  #Optional
  display_name    = var.test_log_name
  log_type        = var.log_log_type.custom
  source_resource = var.log_source_resource
  source_service  = var.log_source_service
  state           = "ACTIVE"
}

output "test_log_id" {
  value = oci_logging_log.test_log.id
}
