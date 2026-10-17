// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "organization_id" {

}

variable "tenancy_id" {

}

variable "organization_tenancy_admin_email" {

}

variable "organization_tenancy_home_region" {

}

variable "organization_tenancy_name" {

}

variable "organization_tenancy_governance_status" {
	default = "OPTED_IN"
}

variable "organization_tenancy_policy_name" {
	default = "child_tenancy_admin_policy"
}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

resource "oci_tenantmanagercontrolplane_organization_tenancy" "test_organization_tenancy" {
	#Required
	admin_email     = var.organization_tenancy_admin_email
	compartment_id  = var.tenancy_ocid
	home_region     = var.organization_tenancy_home_region
	organization_id = var.organization_id
	tenancy_name    = var.organization_tenancy_name

	#Optional
	governance_status = var.organization_tenancy_governance_status
	policy_name       = var.organization_tenancy_policy_name
}

data "oci_tenantmanagercontrolplane_organization_tenancy" "test_organization_tenancy" {
    organization_id = var.organization_id
    tenancy_id = var.tenancy_id
}

data "oci_tenantmanagercontrolplane_organization_tenancies" "test_organization_tenancies" {
    organization_id = var.organization_id
}
