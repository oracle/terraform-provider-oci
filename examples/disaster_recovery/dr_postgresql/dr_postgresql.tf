// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "disassociate_trigger" {
  default = 0
}

variable "dr_protection_group_association_peer_region" {
  default = "us-ashburn-1"
}

variable "dr_protection_group_association_role" {
  default = "STANDBY"
}

variable "dr_protection_group_display_name" {
  default = "displayName"
}

variable "dr_protection_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "dr_protection_group_members_member_type_postgresql_db_system" {
  default = "POSTGRESQL_DB_SYSTEM"
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
  compartment_id = var.compartment_ocid
  display_name   = var.dr_protection_group_display_name
  log_location {
    #Required
    bucket    = data.oci_objectstorage_bucket.test_bucket.name
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  }

  # Add the primary PostgreSQL DB system as a member
  members {
    #Required
    member_id   = data.oci_psql_db_systems.test_db_systems_primary.db_system_collection[0].items[0].id
    member_type = var.dr_protection_group_members_member_type_postgresql_db_system

    #Optional
    peer_db_system_id = data.oci_psql_db_systems.test_db_systems_standby.db_system_collection[0].items[0].id
  }
}

resource "oci_disaster_recovery_dr_protection_group" "test_dr_protection_group" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.dr_protection_group_display_name
  log_location {
    #Required
    bucket    = data.oci_objectstorage_bucket.test_bucket.name
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  }

  # Add the standby PostgreSQL DB system as a member
  members {
    #Required
    member_id   = data.oci_psql_db_systems.test_db_systems_standby.db_system_collection[0].items[0].id
    member_type = var.dr_protection_group_members_member_type_postgresql_db_system

    #Optional
    peer_db_system_id = data.oci_psql_db_systems.test_db_systems_primary.db_system_collection[0].items[0].id
  }

  association {
    #Required
    role = var.dr_protection_group_association_role

    #Optional
    peer_id     = oci_disaster_recovery_dr_protection_group.test_peer.id
    peer_region = var.dr_protection_group_association_peer_region
  }

  #Optional
  disassociate_trigger = var.disassociate_trigger

  freeform_tags = var.dr_protection_group_freeform_tags
}

data "oci_disaster_recovery_dr_protection_groups" "test_dr_protection_groups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name           = var.dr_protection_group_display_name
  dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
  state                  = var.dr_protection_group_state
}
