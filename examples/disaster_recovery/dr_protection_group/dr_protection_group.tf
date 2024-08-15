// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
  default = "ocid1.tenancy.oc1..aaaaaaaahowp4zu5z3p3to5mj7vjtlo7zqi2qmbjiij73vfulltlmvtf624a"
}
variable "disassociate_trigger" {
  default = 0
}

variable "dr_protection_group_association_peer_region" {
  default = "us-ashburn-1"
}

variable "dr_protection_group_association_role" {
  default = "STANDBY"
}

variable "dr_protection_group_defined_tags_value" {
  default = "value"
}

variable "dr_protection_group_display_name" {
  default = "displayName"
}

variable "dr_protection_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "dr_protection_group_log_location_bucket" {
  default = "bucket"
}

variable "dr_protection_group_log_location_namespace" {
  default = "namespace"
}

variable "dr_protection_group_members_is_movable" {
  default = false
}

variable "dr_protection_group_members_member_type" {
  default = "VOLUME_GROUP"
}

variable "dr_protection_group_members_member_type_object_storage" {
  default = "OBJECT_STORAGE_BUCKET"
}

variable "dr_protection_group_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_disaster_recovery_dr_protection_group" "test_peer" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.dr_protection_group_display_name
  log_location {
    #Required
    bucket    = data.oci_objectstorage_bucket.test_bucket.name
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  }

  #Optional
  members {
    #Required
    member_id   = data.oci_core_volume_groups.test_volume_groups.volume_groups.0.id
    member_type = var.dr_protection_group_members_member_type

    #Optional
    is_movable                       = var.dr_protection_group_members_is_movable
  }

  members {
    #Required
    member_id   = data.oci_objectstorage_bucket.test_member_bucket.bucket_id
    member_type = var.dr_protection_group_members_member_type_object_storage
    bucket = data.oci_objectstorage_bucket.test_member_bucket.name
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  }
}

resource "oci_disaster_recovery_dr_protection_group" "test_dr_protection_group" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.dr_protection_group_display_name
  log_location {
    #Required
    bucket    = data.oci_objectstorage_bucket.test_bucket.name
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  }

  lifecycle {
    ignore_changes = [defined_tags]
  }

  #Optional
  association {
    #Required
    role = var.dr_protection_group_association_role

    #Optional
    peer_id     = oci_disaster_recovery_dr_protection_group.test_peer.id
    peer_region = var.dr_protection_group_association_peer_region
  }

  #Optional
  disassociate_trigger = var.disassociate_trigger

  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.dr_protection_group_defined_tags_value}")
  freeform_tags = var.dr_protection_group_freeform_tags
}

data "oci_disaster_recovery_dr_protection_groups" "test_dr_protection_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name           = var.dr_protection_group_display_name
  dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
  state                  = var.dr_protection_group_state
}