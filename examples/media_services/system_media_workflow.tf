// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "system_media_workflow_name" {
  default = "name"
}

data "oci_media_services_system_media_workflow" "test_system_media_workflow" {

  #Optional
  compartment_id = var.compartment_id
  name           = var.system_media_workflow_name
}

