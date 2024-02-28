// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_availability_catalog_base64encoded_catalog_details" {
  default = "U2VxIE5vLEZpbmFsIEN1c3RvbWVyIE9yZGVFkIFR5cGUsTmFtZSxBdmFpbGFibGUgUXVhbnRpdHksVW5pdA0KMSwyMDIzLTAyLTAyLDIwMjMtMDItMTYsU0VSVkVSIEhXLEdlbmVyaWMsQk0uU3RhbmRhcmQyLjUyMSwyMzUsU2VydmVycw0KMiwyMDIzLTAyLTAyLDIwMjMtMDItMTYsU0VSVkVSIEhXLEdlbmVyaWMsQk0uU3RhbmRhcmQyLjUyMywzNjcsU2VydmVycw0KMywyMDIzLTAyLTAyLDIwMjMtMDItMTYsU0VSVkVSIEhXLEdlbmVyaWMsQk0uU3RhbmRhcmQyLjUyLDEyLFNlcnZlcnMNCjQsMjAyMy0wMi0wMiwyMDIzLTAyLTE2LFNFUlZFUiBIVyxST1csQk0uRGVuc2VPLkU0LjEzMSwzMTQsU2VydmVycw0KNSwyMDIzLTAyLTE2LDIwMjMtMDItMDMsU0VSVkVSIEhXLEdlbmVyaWMsQk0uRGVuc2VPLkU0LjEyOCwyMzUsU2VydmVycw0KNiwyMDIzLTAyLTE2LDIwMjMtMDItMDMsU0VSVkVSIEhXLEdlbmVyaWMsQk0uQmlnRGF0YTMuRTQuMTI4LDM2NyxTZXJ2ZXJzDQo3LDIwMjMtMDItMTYsMjAyMy0wMi0wMyxTRVJWRVIgSFcsR2VuZXJpYyxCTS5TdGFuZGFyZDIuNTIsMTIsU2VydmVycw0KOCwyMDIzLTAyLTE2LDIwMjMtMDItMDMsU0VSVkVSIEhXLFJPVyxCTS5EZW5zZU8uRTQuMTMxLDMxNCxTZXJ2ZXJzDQo5LDIwMjMtMDItMTYsMjAyMy0wMi0wMyxTRVJWRVIgSFcsUk9XLEU0MlQgV2hpdGVib3gsNzIyLFNlcnZlcnMNCjEwLDIwMjMtMDItMDIsMjAyMy0wMi0xNixDQVBBQ0lUWSBDT05TVFJBSU5ULEdlbmVyaWMsVVMtQVNIQlVSTi0xLUFELTEsNDg2NCxTZXJ2ZXJzDQoxMSwyMDIzLTAyLTAyLDIwMjMtMDItMTYsQ0FQQUNJVFkgQ09OU1RSQUlOVCxST1csVVMtQVNIQlVSTi0xLUFELTEsNDg2NCxTZXJ2ZXJzDQoxMiwyMDIzLTAyLTAyLDIwMjMtMDItMTYsQ0FQQUNJVFkgQ09OU1RSQUlOVCxHZW5lcmljLFVTLUFTSEJVUk4tMS1BRC0yLDQ4NjQsU2VydmVycw0KMTMsMjAyMy0wMi0wMiwyMDIzLTAyLTE2LENBUEFDSVRZIENPTlNUUkFJTlQsUk9XLFVTLUFTSEJVUk4tMS1BRC0yLDQ4NjQsU2VydmVycw0KMTQsMjAyMy0wMi0xNiwyMDIzLTAyLTAzLENBUEFDSVRZIENPTlNUUkFJTlQsR2VuZXJpYyxVUy1BU0hCVVJOLTEtQUQtMSw2MDIyLFNlcnZlcnMNCjE1LDIwMjMtMDItDT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiw2MDIyLFNlcnZlcnM="
}

variable "occ_availability_catalog_catalog_state" {
  default = "NOT_UPLOADED"
}

variable "occ_availability_catalog_defined_tags_value" {
  default = "value"
}

variable "occ_availability_catalog_description" {
  default = "description"
}

variable "occ_availability_catalog_display_name" {
  default = "displayName"
}

variable "occ_availability_catalog_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occ_availability_catalog_id" {
  default = "id"
}

variable "occ_availability_catalog_metadata_details_format_version" {
  default = "V1"
}

variable "occ_availability_catalog_namespace" {
  default = "COMPUTE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occ_availability_catalog" "test_occ_availability_catalog" {
  #Required
  base64encoded_catalog_details = var.occ_availability_catalog_base64encoded_catalog_details
  compartment_id                = var.compartment_id
  display_name                  = var.occ_availability_catalog_display_name
  namespace                     = var.occ_availability_catalog_namespace
  occ_customer_group_id         = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_availability_catalog_defined_tags_value)
  description   = var.occ_availability_catalog_description
  freeform_tags = var.occ_availability_catalog_freeform_tags
  metadata_details {
    #Required
    format_version = var.occ_availability_catalog_metadata_details_format_version
  }
}

data "oci_capacity_management_occ_availability_catalogs" "test_occ_availability_catalogs" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  catalog_state = var.occ_availability_catalog_catalog_state
  display_name  = var.occ_availability_catalog_display_name
  id            = var.occ_availability_catalog_id
  namespace     = var.occ_availability_catalog_namespace
}