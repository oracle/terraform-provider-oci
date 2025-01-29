// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "assigned_subscription_compartment_id" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "assigned_subscription_id" {
    
}

provider "oci" {
	tenancy_ocid     = var.assigned_subscription_compartment_id
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_assigned_subscription_line_items" "test_assigned_subscription_line_items" {
    assigned_subscription_id = var.assigned_subscription_id
}