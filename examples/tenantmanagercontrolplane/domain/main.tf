// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "compartment_id" {

}

variable "domain_id" {

}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_domain" "test_assigned_domain" {
    domain_id = var.domain_id
}

data "oci_tenantmanagercontrolplane_domains" "test_assigned_domains" {
    compartment_id = var.compartment_id
}