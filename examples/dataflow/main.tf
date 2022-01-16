// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

variable "application_display_name" {
  default = "tf_app"
}

variable "application_driver_shape" {
  default = "VM.Standard2.1"
}

variable "application_executor_shape" {
  default = "VM.Standard2.1"
}

variable "application_file_uri" {
}

variable "application_archive_uri" {
}

variable "metastore_id" {
}

variable "dataflow_logs_bucket_uri" {
}

variable "dataflow_warehouse_bucket_uri" {
}

variable "application_language" {
  default = "PYTHON"
}

variable "application_num_executors" {
  default = 1
}

variable "application_spark_version" {
  default = "2.4"
}

variable "invoke_run_display_name" {
  default = "tf_run"
}

resource "oci_dataflow_application" "tf_application" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.application_display_name
  driver_shape   = var.application_driver_shape
  executor_shape = var.application_executor_shape
  file_uri       = var.application_file_uri
  language       = var.application_language
  num_executors  = var.application_num_executors
  spark_version  = var.application_spark_version

  #Optional
  #arguments       = var.application_arguments
  #class_name      = var.application_class_name
  #configuration   = var.application_configuration
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  #description     = var.application_description
  #freeform_tags   = var.application_freeform_tags
  #logs_bucket_uri = var.application_logs_bucket_uri}"
  type             = "BATCH"
  archive_uri = var.application_archive_uri
  #parameters {
  #Required
  #name  = var.application_parameters_name}"
  #value = var.application_parameters_value}"
  #}

  #warehouse_bucket_uri = var.application_warehouse_bucket_uri}"
  metastore_id = var.metastore_id
}

data "oci_dataflow_applications" "tf_applications" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.application_display_name
}

resource "oci_dataflow_invoke_run" "tf_invoke_run" {
  #Required
  application_id = oci_dataflow_application.tf_application.id
  compartment_id = var.compartment_id
  display_name   = var.invoke_run_display_name
  #Optional
  #arguments       = var.invoke_run_arguments}"
  #configuration   = var.invoke_run_configuration}"
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  #driver_shape    = var.invoke_run_driver_shape}"
  #executor_shape  = var.invoke_run_executor_shape}"
  #freeform_tags   = var.invoke_run_freeform_tags}"
  #logs_bucket_uri = var.invoke_run_logs_bucket_uri}"
  #num_executors   = var.invoke_run_num_executors}"

  #parameters {
  #Required
  #name  = var.invoke_run_parameters_name}"
  #value = var.invoke_run_parameters_value}"
  #}

  #warehouse_bucket_uri = var.invoke_run_warehouse_bucket_uri}"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_dataflow_private_endpoint" "test_private_endpoint" {
  compartment_id = var.compartment_id
  description    = "description"
  display_name   = "pe_name"
  dns_zones      = ["custpvtsubnet.oraclevcn.com"]

  freeform_tags = {
    "Department" = "Finance"
  }

  max_host_count = "256"
  nsg_ids        = [oci_core_network_security_group.test_network_security_group.id]
  subnet_id      = oci_core_subnet.test_subnet.id
}

resource "oci_dataflow_application" "test_application" {
  archive_uri    = var.application_archive_uri
  arguments      = ["arguments"]
  compartment_id = var.compartment_id

  configuration = {
    "spark.shuffle.io.maxRetries" = "10"
  }

  description    = "description"
  display_name   = "test_wordcount_app"
  driver_shape   = "VM.Standard2.1"
  executor_shape = "VM.Standard2.1"
  file_uri       = var.application_file_uri

  freeform_tags = {
    "Department" = "Finance"
  }

  language        = "PYTHON"
  logs_bucket_uri = var.dataflow_logs_bucket_uri
  num_executors   = "1"

  parameters {
    name  = "name"
    value = "value"
  }

  private_endpoint_id  = oci_dataflow_private_endpoint.test_private_endpoint.id
  spark_version        = "2.4"
  warehouse_bucket_uri = var.dataflow_warehouse_bucket_uri
  metastore_id = var.metastore_id
}

resource "oci_dataflow_invoke_run" "test_invoke_run" {
  application_id = oci_dataflow_application.test_application.id
  compartment_id = var.compartment_id
  display_name   = "test_run_name"
}

resource "oci_dataflow_application" "test_application_submit" {
  #Required
  compartment_id = var.compartment_id
  execute        = "--conf spark.shuffle.io.maxRetries=10 ${var.application_file_uri} arguments"
  display_name   = "test_wordcount_app_submit"
  driver_shape   = "VM.Standard2.1"
  executor_shape = "VM.Standard2.1"
  file_uri       = var.application_file_uri
  language       = "PYTHON"
  num_executors  = "1"
  spark_version  = "2.4"
  #Optional
  archive_uri    = var.application_archive_uri
  private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
  metastore_id = var.metastore_id
}

resource "oci_dataflow_invoke_run" "test_invokey_run_submit" {
  #Required
  compartment_id = var.compartment_id
  execute        = "--conf spark.shuffle.io.maxRetries=10 ${var.application_file_uri} arguments"
  #Optional
  application_id = oci_dataflow_application.test_application_submit.id
  archive_uri    = var.application_archive_uri
  display_name   = "test_wordcount_run_submit"
  spark_version  = "2.4"
  metastore_id = var.metastore_id
}

data "oci_dataflow_private_endpoints" "test_private_endpoints" {
  compartment_id = var.compartment_id

  // service supports using only one filter at a time for LIST API call
  display_name = "pe_name"

  filter {
    name   = "id"
    values = [oci_dataflow_private_endpoint.test_private_endpoint.id]
  }
}

data "oci_dataflow_invoke_runs" "tf_invoke_runs" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  application_id = oci_dataflow_application.tf_application.id
}

data "oci_dataflow_run_logs" "tf_run_logs" {
  #Required
  run_id = oci_dataflow_invoke_run.tf_invoke_run.id
}

