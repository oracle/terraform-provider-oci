// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}

variable "iot_domain_configure_data_access_db_allow_listed_identity_group_names" {
  default = []
}

variable "iot_domain_configure_data_access_db_allowed_identity_domain_host" {
  default = "dbAllowedIdentityDomainHost"
}

variable "iot_domain_configure_data_access_db_workspace_admin_initial_password" {
  default = "dbWorkspaceAdminInitialPassword"
}

variable "iot_domain_configure_data_access_type" {
  #DIRECT, ORDS or APEX
  default = "DIRECT"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_domain_configure_data_access" "test_iot_domain_configure_data_access" {
  #Required
  iot_domain_id = var.iot_domain_ocid
  type          = var.iot_domain_configure_data_access_type

  #Optional
  #DIRECT
  db_allow_listed_identity_group_names = var.iot_domain_configure_data_access_db_allow_listed_identity_group_names
  #ORDS
  db_allowed_identity_domain_host      = var.iot_domain_configure_data_access_db_allowed_identity_domain_host
  #APEX
  db_workspace_admin_initial_password  = var.iot_domain_configure_data_access_db_workspace_admin_initial_password
}


