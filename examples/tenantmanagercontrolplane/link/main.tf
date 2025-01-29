// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "link_id" {

}

variable "child_tenancy_id" {

}

variable "parent_tenancy_id" {

}

variable "state" {
    
}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_link" "test_link" {
    link_id = var.link_id
}

data "oci_tenantmanagercontrolplane_links" "test_links" {
    child_tenancy_id = var.child_tenancy_id
    parent_tenancy_id = var.parent_tenancy_id
    state = var.state
}