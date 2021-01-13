// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

variable "log_defined_tags_value" {
  default = "value"
}

variable "log_freeform_tags" {
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

resource "oci_logging_log" "test_log" {
  #Required
  display_name = "displayName"
  log_group_id = oci_logging_log_group.test_log_group.id
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
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag2.name}" = var.log_defined_tags_value
  }
  freeform_tags      = var.log_freeform_tags
  is_enabled         = "false"
  retention_duration = "30"
}

data "oci_logging_logs" "test_logs" {
  #Required
  log_group_id = oci_logging_log_group.test_log_group.id

  #Optional
  display_name    = "displayName"
  log_type        = var.log_log_type.custom
  source_resource = var.log_source_resource
  source_service  = var.log_source_service
  state           = "ACTIVE"
}

