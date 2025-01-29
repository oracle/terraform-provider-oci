// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "sender_tenancy_id" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "sender_invitation_id" {

}

variable "sender_invitation_compartment_id" {

}

provider "oci" {
	tenancy_ocid     = var.sender_tenancy_id
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_sender_invitation" "sender_invitation" {
    sender_invitation_id = var.sender_invitation_id
}

data "oci_tenantmanagercontrolplane_sender_invitations" "sender_invitations" {
    compartment_id = var.sender_invitation_compartment_id
}