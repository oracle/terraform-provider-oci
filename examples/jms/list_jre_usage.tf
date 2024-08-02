// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_list_jre_usage" "test_list_jre_usage" {

  #Optional
  application_id   = var.application_id
  application_name = var.application_name
  compartment_id   = var.compartment_ocid
  host_id          = var.host_id
  time_end         = var.time_end
  time_start       = var.time_start
}
