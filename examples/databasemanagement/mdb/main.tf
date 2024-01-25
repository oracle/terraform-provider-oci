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

variable "compartment_id" {  
  default = "<compartment.ocid>"
}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "managed_database_deployment_type" {
  default = "ONPREMISE"
}
variable "managed_database_management_option" {
  default = ""
}

variable "managed_database_name" {
  default = ""
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