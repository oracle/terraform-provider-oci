// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
  default = []
}

variable "external_exadata_infrastructure_display_name" {
  default = "exadata-Terraform-Testing"
}

variable "external_exadata_infrastructure_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "external_exadata_infrastructure_storage_server_names" {
  default = []
}

variable "external_exadata_storage_connector_connection_uri" {
  default = "https://exaInfra01celadm01.us.oracle.com/MS/RESTService/"
}

variable "external_exadata_storage_connector_connector_name" {
  default = "connectorName"
}

variable "external_exadata_storage_connector_credential_info_password" {
  default = "BEstrO0ng_#11"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_location" {
  default = "sslTrustStoreLocation"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_password" {
  default = "sslTrustStorePassword"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_type" {
  default = "JKS"
}

variable "external_exadata_storage_connector_credential_info_username" {
  default = "username"
}

variable "external_exadata_storage_connector_display_name" {
  default = "exaInfra01celadm01-conn"
}

variable "external_exadata_storage_grid_id" {
  default = "id"
}

variable "external_exadata_storage_server_id" {
  default = "id"
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
	agent_id = var.management_agent_id
	connection_uri = var.external_exadata_storage_connector_connection_uri
	connector_name = var.external_exadata_storage_connector_connector_name
	credential_info {
		#Required
		password = var.external_exadata_storage_connector_credential_info_password
		username = var.external_exadata_storage_connector_credential_info_username

		#Optional
		#ssl_trust_store_location = var.external_exadata_storage_connector_credential_info_ssl_trust_store_location
		#ssl_trust_store_password = var.external_exadata_storage_connector_credential_info_ssl_trust_store_password
		#ssl_trust_store_type = var.external_exadata_storage_connector_credential_info_ssl_trust_store_type
	}
	storage_server_id = oci_database_management_storage_server.test_storage_server.id
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

