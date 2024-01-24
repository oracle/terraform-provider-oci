// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "association_association_type" {
  default = "CERTIFICATE"
}

variable "association_name" {
  default = "name"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_certificates_management_associations" "test_associations" {

  #Optional
  associated_resource_id   = oci_certificates_management_associated_resource.test_associated_resource.id
  association_id           = oci_certificates_management_association.test_association.id
  association_type         = var.association_association_type
  certificates_resource_id = oci_certificates_management_certificates_resource.test_certificates_resource.id
  compartment_id           = var.compartment_id
  name                     = var.association_name
}

