// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "cross_connect_group_display_name" {
  default = "displayName"
}

variable "cross_connect_group_state" {
  default = "AVAILABLE"
}

variable "secret_ocid_ckn" {

}

variable "secret_version_cak" {

}

variable "secret_ocid_cak" {

}

variable "secret_version_ckn" {

}

resource "oci_core_cross_connect_group" "cross_connect_group" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.cross_connect_group_display_name
}

data "oci_core_cross_connect_groups" "cross_connect_groups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.cross_connect_group_display_name
  #state = var.cross_connect_group_state
}

output "cross_connect_groups" {
  value = data.oci_core_cross_connect_groups.cross_connect_groups.cross_connect_groups
}

resource "oci_core_cross_connect_group" "test_cross_connect_group" {
    #Required
    compartment_id = var.compartment_ocid
    #Optional
    customer_reference_name = "customerReferenceName"
    defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")
    display_name = "displayName"
    freeform_tags = {
        "Department" = "Finance"
    }
    macsec_properties {
        #Required
        state = "ENABLED"
        #Optional
        encryption_cipher = "AES256_GCM"
        primary_key {
            #Required
            connectivity_association_key_secret_id = var.secret_ocid_cak
            connectivity_association_name_secret_id = var.secret_ocid_ckn
            #Optional, api will always create with current version, but can use to update
            connectivity_association_key_secret_version = var.secret_version_cak
            connectivity_association_name_secret_version = var.secret_version_ckn
        }
        is_unprotected_traffic_allowed = false

    }
}

resource "oci_core_cross_connect_group" "test_cross_connect_group_2" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  customer_reference_name = "customerReferenceName"
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")
  display_name = "displayName"
  freeform_tags = {
    "Department" = "Finance"
  }
  macsec_properties {
    #Required
    state = "ENABLED"
    #Optional
    encryption_cipher = "AES256_GCM"
    primary_key {
      #Required
      connectivity_association_key_secret_id = var.secret_ocid_cak
      connectivity_association_name_secret_id = var.secret_ocid_ckn
      #secret versions default to current version
    }
    is_unprotected_traffic_allowed = true

  }
}

variable defined_tag_namespace_name { default = "" }
resource "oci_identity_tag_namespace" "tag-namespace1" {
        #Required
        compartment_id = var.tenancy_ocid
        description = "example tag namespace"
        name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

        is_retired = false
}

resource "oci_identity_tag" "tag1" {
        #Required
        description = "example tag"
        name = "example-tag"
        tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

        is_retired = false
}

