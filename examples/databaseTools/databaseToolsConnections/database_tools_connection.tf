// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# ========================
# Provider Configuration
# ========================
provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}


# ========================
# Variables
# ========================
variable "region" {
 default = "r1"
}

variable "config_file_profile" {
 default = "boat-r1-session"
}

variable "compartment_id" {}
variable "tenancy_ocid" {}
variable "related_resource_id" {}

# Existing resource OCIDs (from environment)
variable "secret_ocid" {
 default = ""
 description = "OCID of existing vault secret (TF_VAR_secret_ocid)"
}

variable "database_tools_private_endpoint_ocid" {
 default = ""
 description = "OCID of existing private endpoint (TF_VAR_database_tools_private_endpoint_ocid)"
}

# Connection variables
variable "database_tools_connection_advanced_properties" {
 default = {
  "oracle.net.ssl_server_dn_match" = "true"
 }
}

variable "database_tools_connection_authentication_type" {
 default = "TOKEN"
}

variable "database_tools_connection_connection_string" {
 default = "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"
}

variable "database_tools_connection_defined_tags_value" {
 default = "value"
}

variable "database_tools_connection_display_name" {
 default = "ADMIN@DB202005191141_low"
}

variable "database_tools_connection_freeform_tags" {
 default = { "bar-key" = "value" }
}

variable "database_tools_connection_key_stores_key_store_content_value_type" {
 default = "SECRETID"
}

variable "database_tools_connection_key_stores_key_store_password_value_type" {
 default = "SECRETID"
}

variable "database_tools_connection_key_stores_key_store_type" {
 default = "PEM"
}

variable "database_tools_connection_locks_message" {
 default = "message"
}

variable "database_tools_connection_locks_type" {
 default = "FULL"
}

variable "database_tools_connection_proxy_client_proxy_authentication_type" {
 default = "USER_NAME"
}

variable "database_tools_connection_proxy_client_roles" {
 default = []
}

variable "database_tools_connection_proxy_client_user_password_value_type" {
 default = "SECRETID"
}

variable "database_tools_connection_related_resource_entity_type" {
 default = "DATABASE"
}

variable "database_tools_connection_related_resource_identifier" {
 default = ""
}

variable "database_tools_connection_runtime_identity" {
 default = "AUTHENTICATED_PRINCIPAL"
}

variable "database_tools_connection_runtime_support" {
 default = "SUPPORTED"
}

variable "database_tools_connection_state" {
 default = "ACTIVE"
}

variable "database_tools_connection_type" {
 default = "ORACLE_DATABASE"
}

variable "database_tools_connection_url" {
 default = "url"
}

variable "database_tools_connection_user_password_value_type" {
 default = "SECRETID"
}

variable "test_user_name" {
 default = "[proxyClient]"
}

# ========================
# Resource: Database Tools Connection (Main)
# ========================
resource "oci_database_tools_database_tools_connection" "test_database_tools_connection" {
 #Required
 compartment_id = var.compartment_id
 display_name = var.database_tools_connection_display_name
 type = var.database_tools_connection_type
 user_name = var.test_user_name

 #Optional
 advanced_properties = var.database_tools_connection_advanced_properties
 authentication_type = var.database_tools_connection_authentication_type
 connection_string = var.database_tools_connection_connection_string
 defined_tags = {
 }
 freeform_tags = var.database_tools_connection_freeform_tags

 key_stores {
 key_store_content {
 value_type = var.database_tools_connection_key_stores_key_store_content_value_type
 secret_id = var.secret_ocid
 }
 key_store_password {
 value_type = var.database_tools_connection_key_stores_key_store_password_value_type
 secret_id = var.secret_ocid
 }
 key_store_type = var.database_tools_connection_key_stores_key_store_type
 }

 private_endpoint_id = var.database_tools_private_endpoint_ocid != "" ? var.database_tools_private_endpoint_ocid : null

 related_resource {
 entity_type = var.database_tools_connection_related_resource_entity_type
 identifier = var.database_tools_connection_related_resource_identifier != "" ? var.database_tools_connection_related_resource_identifier : var.related_resource_id
 }

 runtime_identity = var.database_tools_connection_runtime_identity
 runtime_support = var.database_tools_connection_runtime_support
 url = var.database_tools_connection_url
}

# ========================
# Data Source: Query all connections
# ========================
data "oci_database_tools_database_tools_connections" "test_database_tools_connections" {
 compartment_id = var.compartment_id
 display_name = var.database_tools_connection_display_name
 related_resource_identifier = var.database_tools_connection_related_resource_identifier != "" ? var.database_tools_connection_related_resource_identifier : var.related_resource_id
 runtime_identity = [var.database_tools_connection_runtime_identity]
 runtime_support = [var.database_tools_connection_runtime_support]
 state = var.database_tools_connection_state
 type = [var.database_tools_connection_type]
}

# ========================
# Outputs
# ========================
output "database_tools_connection_id" {
 value = oci_database_tools_database_tools_connection.test_database_tools_connection.id
}

output "database_tools_connection_state" {
 value = oci_database_tools_database_tools_connection.test_database_tools_connection.state
}