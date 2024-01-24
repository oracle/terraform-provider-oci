// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

//https://github.com/terraform-providers/terraform-provider-oci/tree/master/examples/compute/image
variable "custom_image_id" {
}

//example reference to schemaData values
variable "schema_data_value" {
  default = "{\"descriptorType\": \"enumstring\",\"source\": \"IMAGE\",\"defaultValue\": \"PARAVIRTUALIZED\", \"values\": [\"PARAVIRTUALIZED\"]}"
}

//example reference to schemaData key
variable "schema_data_key" {
  default = "Network.AttachmentType"
}

data "oci_core_compute_image_capability_schema" "test_compute_image_capability_schema" {
  compute_image_capability_schema_id = oci_core_compute_image_capability_schema.test_compute_image_capability_schema.id
  is_merge_enabled                   = "true"
}

resource "oci_core_compute_image_capability_schema" "test_compute_image_capability_schema" {
  compartment_id                                      = var.compartment_ocid
  compute_global_image_capability_schema_version_name = data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions_datasource.compute_global_image_capability_schema_versions[0].name
  display_name                                        = "displayName"
  image_id                                            = var.custom_image_id

  schema_data = {
    # if using var as key in map, enclose in parenthesis
    (var.schema_data_key) = var.schema_data_value
  }
}

data "oci_core_compute_global_image_capability_schemas_version" "test_compute_global_image_capability_schemas_version_datasource" {
  compute_global_image_capability_schema_id           = data.oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema_datasource.id
  compute_global_image_capability_schema_version_name = data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions_datasource.compute_global_image_capability_schema_versions[0].name
}

data "oci_core_compute_global_image_capability_schemas_versions" "test_compute_global_image_capability_schemas_versions_datasource" {
  compute_global_image_capability_schema_id = data.oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema_datasource.id
}

data "oci_core_compute_global_image_capability_schema" "test_compute_global_image_capability_schema_datasource" {
  compute_global_image_capability_schema_id = data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas_datasource.compute_global_image_capability_schemas[0].id
}

data "oci_core_compute_global_image_capability_schemas" "test_compute_global_image_capability_schemas_datasource" {
}

