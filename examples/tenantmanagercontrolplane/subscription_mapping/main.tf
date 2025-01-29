// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "subscription_mapping_compartment_id" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "compartment_ocid" {

}

variable "subscription_mapping_subscription_id" {

}

provider "oci" {
	tenancy_ocid     = var.subscription_mapping_compartment_id
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

resource "oci_tenantmanagercontrolplane_subscription_mapping" "test_subscription_mapping" {
    compartment_id = var.subscription_mapping_compartment_id
    subscription_id = var.subscription_mapping_subscription_id
}

data "oci_tenantmanagercontrolplane_subscription_mappings" "test_subscription_mappings" {
    compartment_id = oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.compartment_id
    subscription_id = var.subscription_mapping_subscription_id
}

data "oci_tenantmanagercontrolplane_subscription_mapping" "test_subscription_mapping" {
    subscription_mapping_id = oci_tenantmanagercontrolplane_subscription_mapping.test_subscription_mapping.id
}