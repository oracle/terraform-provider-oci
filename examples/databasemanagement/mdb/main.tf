// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

variable "managed_database_id" {
   default = "ocid1.test.oc1..<unique_ID>EXAMPLE-managedDatabaseId-Value"
}

variable "managed_database_deployment_type" {
  default = "ONPREMISE"
}

variable "managed_database_management_option" {
  default = "ADVANCED"
}

variable "managed_database_name" {
  default = "EXAMPLE-managedDatabaseName-Value"
}

variable "managed_db_defined_tags_value" {
  default = "managed_db_tag_value"
}

variable "managed_db_freeform_tags" {
  default = { "bar-key" = "value" }
}

# Create a new Tag Namespace.
resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
}

# Create a new Tag definition in the above Tag Namespace.
resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

data "oci_database_management_managed_databases" "test_managed_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_type = var.managed_database_deployment_type
	#external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
	id = var.managed_database_id
	#management_option = var.managed_database_management_option
	#name = var.managed_database_name
}

# Update tags on a Managed Database resource.
resource "oci_database_management_managed_database" "test_managed_database" {
  managed_database_id = var.managed_database_id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.managed_db_defined_tags_value
  }
  freeform_tags = var.managed_db_freeform_tags
}
