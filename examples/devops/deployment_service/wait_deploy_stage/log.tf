// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_logging_log_group" "test_log_group" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "exampleLogGroup"


  description = "description"

}

resource "oci_logging_log" "test_log" {
  #Required
  display_name = "displayName"
  log_group_id = oci_logging_log_group.test_log_group.id
  log_type     = "SERVICE"

  #Optional
  configuration {
    #Required
    source {
      #Required
      category    = "all"
      resource    = oci_devops_project.test_project.id
      service     = "devops"
      source_type = "OCISERVICE"
    }
  }
  is_enabled         = "true"
  retention_duration = "30"
}
