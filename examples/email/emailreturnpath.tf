// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "email_return_path_state" {
	default = "ACTIVE"
}

variable "domain_verification_id" {

}

variable "email_domain" {
	default = "objdomain.email.ap-mumbai-1.oci.oc-test.com"
}

variable "description" {
	default = "description2"
}

variable "department" {
	default = "Accounting"
}


data "oci_email_email_return_paths" "test_email_return_paths" {
	compartment_id = var.compartment_ocid
	filter {
		name = "id"
		values = ["${oci_email_email_return_path.test_email_return_path.id}"]
	}
	id = "${oci_email_email_return_path.test_email_return_path.id}"
	name = "tfrp.${oci_email_email_domain.test_email_domain_rp.name}"
	parent_resource_id = "${oci_email_email_domain.test_email_domain_rp.id}"
	state = var.email_return_path_state
}

resource "oci_email_email_domain" "test_email_domain_rp" {
	compartment_id = var.compartment_ocid
	domain_verification_id = var.domain_verification_id
	name = var.email_domain
}

resource "oci_email_email_return_path" "test_email_return_path" {
	description = var.description
	freeform_tags = {
		"Department" = var.department
	}
	name = "tfrp.${oci_email_email_domain.test_email_domain_rp.name}"
	parent_resource_id = "${oci_email_email_domain.test_email_domain_rp.id}"
}