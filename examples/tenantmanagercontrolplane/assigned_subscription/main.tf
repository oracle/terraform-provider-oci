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

variable "entity_version" {

}

provider "oci" {
	tenancy_ocid     = var.assigned_subscription_compartment_id
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_assigned_subscription" "test_assigned_subscription" {
    assigned_subscription_id = var.assigned_subscription_id
}

data "oci_tenantmanagercontrolplane_assigned_subscriptions" "test_assigned_subscriptions" {
    compartment_id = var.assigned_subscription_compartment_id
    entity_version = var.entity_version
    subscription_id = var.assigned_subscription_id
}