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

####################### Exadata Infrastructure Monitoring #########################

variable "compartment_id" {
  default = "<compartment.ocid>"
}

variable "external_exadata_infrastructure_database_managements_management_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "enable_exadata" {
  default = false
}

variable "external_exadata_infrastructure_db_system_ids" {
  default = ["ocid1.test.oc1..<unique_ID>EXAMPLE-externalDbSystemId-Value"]
}

variable "external_exadata_infrastructure_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "external_exadata_infrastructure_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "external_exadata_infrastructure_storage_server_names" {
  default = ["EXAMPLE-storageServerName-Value"]
}

variable "external_exadata_storage_connector_connection_uri" {
  default = "EXAMPLE-connectionUri-Value"
}

variable "external_exadata_storage_connector_connector_name" {
  default = "EXAMPLE-connectorName-Value"
}

variable "external_exadata_storage_connector_credential_info_password" {
  default = "EXAMPLE-password-Value"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_location" {
  default = "EXAMPLE-sslTrustStoreLocation-Value"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_password" {
  default = "EXAMPLE-sslTrustStorePassword-Value"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_type" {
  default = "JKS"
}

variable "external_exadata_storage_connector_credential_info_username" {
  default = "EXAMPLE-username-Value"
}

variable "external_exadata_storage_connector_display_name" {
  default = "EXAMPLE-connectorName-Value"
}

variable "external_exadata_storage_grid_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-externalExadataStorageGridId-Value"
}

variable "external_exadata_storage_server_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-externalExadataStorageServerId-Value"
}

variable "external_exadata_connector_agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "connector_storage_server_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-storageServerId-Value"
}

variable "exadata_infra_defined_tags_value" {
  default = "exadata_infra_tag_value"
}

variable "exadata_infra_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "exadata_connector_defined_tags_value" {
  default = "exadata_connector_tag_value"
}

variable "exadata_connector_freeform_tags" {
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

resource "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
  #Required
  compartment_id = var.compartment_id
  db_system_ids  = var.external_exadata_infrastructure_db_system_ids
  display_name   = var.external_exadata_infrastructure_display_name

  #Optional
  #discovery_key        = var.external_exadata_infrastructure_discovery_key
  #license_model        = var.external_exadata_infrastructure_license_model
  #storage_server_names = var.external_exadata_infrastructure_storage_server_names
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.exadata_infra_defined_tags_value
  }
  freeform_tags = var.exadata_infra_freeform_tags
}

data "oci_database_management_external_exadata_infrastructures" "test_external_exadata_infrastructures" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #display_name = var.external_exadata_infrastructure_display_name
}

data "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
	#Required
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
}

resource "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	agent_id = var.external_exadata_connector_agent_id
	connection_uri = var.external_exadata_storage_connector_connection_uri
	connector_name = var.external_exadata_storage_connector_connector_name
	credential_info {
		#Required
		password = var.external_exadata_storage_connector_credential_info_password
		username = var.external_exadata_storage_connector_credential_info_username

		#Optional
		ssl_trust_store_location = var.external_exadata_storage_connector_credential_info_ssl_trust_store_location
		ssl_trust_store_password = var.external_exadata_storage_connector_credential_info_ssl_trust_store_password
		ssl_trust_store_type = var.external_exadata_storage_connector_credential_info_ssl_trust_store_type
	}
	storage_server_id = var.connector_storage_server_id
    defined_tags  = {
      "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.exadata_connector_defined_tags_value
    }
    freeform_tags = var.exadata_connector_freeform_tags
}

data "oci_database_management_external_exadata_storage_connectors" "test_external_exadata_storage_connectors" {
	#Required
	compartment_id = var.compartment_id
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id

	#Optional
	#display_name = var.external_exadata_storage_connector_display_name
}

data "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	external_exadata_storage_connector_id = oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector.id
}

data "oci_database_management_external_exadata_storage_grid" "test_external_exadata_storage_grid" {
	#Required
	external_exadata_storage_grid_id = var.external_exadata_storage_grid_id
}

data "oci_database_management_external_exadata_storage_server_iorm_plan" "test_external_exadata_storage_server_iorm_plan" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server_open_alert_history" "test_external_exadata_storage_server_open_alert_history" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server_top_sql_cpu_activity" "test_external_exadata_storage_server_top_sql_cpu_activity" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server" "test_external_exadata_storage_server" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_servers" "test_external_exadata_storage_servers" {
	#Required
	compartment_id = var.compartment_id
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id

	#Optional
	#display_name = var.external_exadata_storage_server_display_name
}

resource "oci_database_management_external_exadata_infrastructure_exadata_management" "test_external_exadata_infrastructure_exadata_management" {
  #Required
  external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
  enable_exadata = var.enable_exadata

  #Optional
  #license_model = var.external_exadata_infrastructure_database_managements_management_license_model
}

