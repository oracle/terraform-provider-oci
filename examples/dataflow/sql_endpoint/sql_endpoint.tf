// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_id" {
}

variable "region" {
}

variable "metastore_id" {
}


provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_dataflow_sql_endpoint" "test_sql_endpoint" {
  compartment_id        = var.compartment_id
  display_name          = "test_sql_endpoint"
  driver_shape          = "VM.Standard.E4.Flex"
  driver_shape_config {
    memory_in_gbs = 32
    ocpus = 2
  }
  executor_shape        = "VM.Standard.E4.Flex"
  executor_shape_config {
    memory_in_gbs = 32
    ocpus = 2
  }
  max_executor_count    = 2
  metastore_id          = var.metastore_id
  min_executor_count    = 1
  network_configuration {
    network_type = "SECURE_ACCESS"
  }
  sql_endpoint_version  = "3.2.1"
  #Optional
  #state                = "ACTIVE" OR "INACTIVE" You can use ACTIVE in order to start a stopped sql endpoint and INACTIVE in order to stop a active sql endpoint
  #defined_tags         = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  #description          = var.description
  freeform_tags = {
      "Department" = "Finance"
  }

  lifecycle {
    ignore_changes = [ system_tags ]
  }
}

data "oci_dataflow_sql_endpoints" "tf_sql_endpoints" {
    #Required
    compartment_id = var.compartment_id
}

data "oci_dataflow_sql_endpoint" "tf_sql_endpoint_data" {
    #Required
    sql_endpoint_id = oci_dataflow_sql_endpoint.test_sql_endpoint.id
}