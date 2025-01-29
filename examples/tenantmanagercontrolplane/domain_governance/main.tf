// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "compartment_ocid" {

}

variable "domain_id" {

}

variable "domain_governance_id" {

}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_domain_governances" "test_domain_governances" {
    domain_id = var.domain_id
    compartment_id = var.compartment_ocid
}

data "oci_tenantmanagercontrolplane_domain_governance" "test_domain_governance" {
	domain_governance_id = var.domain_governance_id
}