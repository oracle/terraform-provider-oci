// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "tenancy_name" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "admin_email" {
}

variable "idcs_access_token" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_oce_oce_instance" "test_oce_instance" {
  admin_email              = var.admin_email
  compartment_id           = var.compartment_ocid
  idcs_access_token        = var.idcs_access_token
  name                     = "testoceinstance"
  object_storage_namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  tenancy_id               = var.tenancy_ocid
  tenancy_name             = var.tenancy_name

  #optional
  freeform_tags = {
    "freeformkey" = "freeformvalue"
  }
  description           = "description"
  instance_access_type  = "PUBLIC"
  instance_usage_type   = "NONPRIMARY"
  instance_license_type = "BYOL"
  upgrade_schedule      = "UPGRADE_IMMEDIATELY"
  waf_primary_domain    = "oracle.com"
  
  timeouts {
    create = "2h"
    update = "60m"
    delete = "1h"
  }
}

data "oci_oce_oce_instances" "test_oce_instances" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_oce_oce_instance.test_oce_instance.id]
  }

  state = "Active"
}

data "oci_oce_oce_instance" "test_oce_instance" {
  oce_instance_id = oci_oce_oce_instance.test_oce_instance.id
}

output "active_oce_instances" {
  value = [data.oci_oce_oce_instances.test_oce_instances.oce_instances]
}

output "output_nested_service_data" {
  value = jsondecode(data.oci_oce_oce_instance.test_oce_instance.service.dns).A.domain
}

data "oci_objectstorage_namespace" "test_namespace" {
    #Optional
    compartment_id = var.compartment_ocid
}
