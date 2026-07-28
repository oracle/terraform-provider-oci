// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "region" {}
variable "compartment_ocid" {}
variable "config_file_profile" {}
variable "auth" {
  default = "SecurityToken"
}

variable "ssh_public_key" {
  default = ""
}

variable "db_name" {
  default = "TFDB"
}

variable "test_db_password" {
  default = "BEstrO0ng_#11"
}

variable "cpg_id" {
  default = null
}

variable "subscription_id" {
  default = null
}

variable "autoscale_limit_in_gbs" {
  default = null
}

variable "is_autoscale_enabled" {
  default = null
}
Support for Azure & GCP KMS Integration for ExaDB-XS -1
# Set to "register" or "unregister" only after the ExaDB-XS cluster exists.
# Leave unset during initial provisioning. Increment pkcs_trigger to run the
# selected PKCS SPECIAL_UPDATE action again. Azure and GCP keystore setup is
# performed outside Terraform; the service reports the active type through the
# resource's tde_key_store_type attribute.
variable "pkcs_operation" {
  type    = string
  default = null
}

variable "pkcs_trigger" {
  type    = number
  default = 1
}
