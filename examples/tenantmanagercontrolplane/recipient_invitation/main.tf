// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "recipient_tenancy_id" {
}

variable "auth" {

}

variable "config_file_profile" {

}

variable "region" {
}

variable "recipient_invitation_id" {

}

variable "recipient_invitation_compartment_id" {

}

provider "oci" {
	tenancy_ocid     = var.recipient_tenancy_id
	auth             = var.auth
	config_file_profile = var.config_file_profile
	region           = var.region
}

data "oci_tenantmanagercontrolplane_recipient_invitation" "recipient_invitation" {
    recipient_invitation_id = var.recipient_invitation_id
}

data "oci_tenantmanagercontrolplane_recipient_invitations" "recipient_invitations" {
    compartment_id = var.recipient_invitation_compartment_id
}