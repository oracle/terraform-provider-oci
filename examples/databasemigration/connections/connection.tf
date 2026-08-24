// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "connection_additional_attributes_name" {
  default = "name"
}

variable "connection_additional_attributes_value" {
  default = "value"
}

variable "connection_connection_string" {
  default = "connectionString"
}

variable "connection_connection_type" {
  default = "MYSQL"
}

variable "connection_defined_tags_value" {
  default = "value"
}

variable "connection_description" {
  default = "description"
}

variable "connection_display_name" {
  default = "displayName"
}

variable "connection_freeform_tags" {
  default = "freeformTags"
}

variable "connection_host" {
  default = "host"
}

variable "connection_nsg_ids" {
  default = []
}

variable "connection_password" {
  default = "BEstrO0ng_#11"
}

variable "connection_port" {
  default = 10
}

variable "connection_replication_password" {
  default = "replicationPassword"
}

variable "connection_replication_username" {
  default = "replicationUsername"
}

variable "connection_security_protocol" {
  default = "PLAIN"
}

variable "connection_ssh_host" {
  default = "sshHost"
}

variable "connection_ssh_key" {
  default = "sshKey"
}

variable "connection_ssh_sudo_location" {
  default = "sshSudoLocation"
}

variable "connection_ssh_user" {
  default = "sshUser"
}

variable "connection_ssl_ca" {
  default = "sslCa"
}

variable "connection_ssl_cert" {
  default = "sslCert"
}

variable "connection_ssl_crl" {
  default = "sslCrl"
}

variable "connection_ssl_key" {
  default = "sslKey"
}

variable "connection_ssl_mode" {
  default = "DISABLED"
}

variable "connection_state" {
  default = "AVAILABLE"
}

variable "connection_technology_sub_type" {
  default = "technologySubType"
}

variable "connection_technology_type" {
  default = "AMAZON_RDS_ORACLE"
}

variable "connection_username" {
  default = "username"
}

variable "connection_wallet" {
  default = "wallet"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_migration_connection" "test_connection" {
  #Required
  compartment_id  = var.compartment_id
  connection_type = var.connection_connection_type
  display_name    = var.connection_display_name
  key_id          = oci_kms_key.test_key.id
  password        = var.connection_password
  technology_type = var.connection_technology_type
  username        = var.connection_username
  vault_id        = oci_kms_vault.test_vault.id

  #Optional
  additional_attributes {

    #Optional
    name  = var.connection_additional_attributes_name
    value = var.connection_additional_attributes_value
  }
  connection_string    = var.connection_connection_string
  database_id          = oci_database_database.test_database.id
  database_name        = oci_database_database.test_database.name
  db_system_id         = oci_database_db_system.test_db_system.id
  defined_tags         = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.connection_defined_tags_value })
  description          = var.connection_description
  freeform_tags        = var.connection_freeform_tags
  host                 = var.connection_host
  nsg_ids              = var.connection_nsg_ids
  port                 = var.connection_port
  replication_password = var.connection_replication_password
  replication_username = var.connection_replication_username
  security_protocol    = var.connection_security_protocol
  ssh_host             = var.connection_ssh_host
  ssh_key              = var.connection_ssh_key
  ssh_sudo_location    = var.connection_ssh_sudo_location
  ssh_user             = var.connection_ssh_user
  ssl_ca               = var.connection_ssl_ca
  ssl_cert             = var.connection_ssl_cert
  ssl_crl              = var.connection_ssl_crl
  ssl_key              = var.connection_ssl_key
  ssl_mode             = var.connection_ssl_mode
  subnet_id            = oci_core_subnet.test_subnet.id
  wallet               = var.connection_wallet
}

data "oci_database_migration_connections" "test_connections" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  connection_type      = var.connection_connection_type
  display_name         = var.connection_display_name
  source_connection_id = oci_database_migration_connection.test_connection.id
  state                = var.connection_state
  technology_sub_type  = var.connection_technology_sub_type
  technology_type      = var.connection_technology_type
}

