// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_id" {
  default = "compartment.ocid"
}

variable "nc_user" {
  default = "SYS"
}

variable "nc_user_role" {
  default = "SYSDBA"
}

variable "key_id" {
  default = "<secret.ocid>"
}

variable "associated_resource_id" {
  default = "database.ocid"
}

variable "associated_resource_updated_id" {
  default = "database.ocid"
}

variable "compartment_id_for_update" {
  default = "compartment.ocid"
}

variable "named_credential_name" {
  default = "namedCredentialName"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Create named credential
resource "oci_database_management_named_credential" "oracle_named_credential" {
  #Required
  compartment_id = var.compartment_id
  name           = var.named_credential_name
  type           = "ORACLE_DB"
  scope          = "GLOBAL"
  content {
    user_name                   = var.nc_user
    credential_type             = "BASIC"
    password_secret_id          = var.key_id
    role                        = var.nc_user_role
    password_secret_access_mode = "USER_PRINCIPAL"
  }

  #Optional
  #display_name = var.external_exadata_storage_server_display_name
}

# Get named credential resource
data "oci_database_management_named_credential" "get_oracle_named_credential" {
  #Required
  named_credential_id = oci_database_management_named_credential.oracle_named_credential.id
}

# List named credential in a compartment
data "oci_database_management_named_credentials" "oracle_named_credentials" {
  #Required
  compartment_id = oci_database_management_named_credential.oracle_named_credential.compartment_id
}