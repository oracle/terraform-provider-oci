// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_lustre_file_storage_lustre_file_system" "test_lustre_file_system" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  capacity_in_gbs     = var.lustre_file_system_capacity_in_gbs
  compartment_id      = var.compartment_ocid
  file_system_name    = var.lustre_file_system_name
  performance_tier    = var.lustre_file_system_performance_tier
  root_squash_configuration {

    #Optional
    client_exceptions = var.lustre_file_system_root_squash_configuration_client_exceptions
    identity_squash   = var.lustre_file_system_root_squash_configuration_identity_squash
    squash_gid        = var.lustre_file_system_root_squash_configuration_squash_gid
    squash_uid        = var.lustre_file_system_root_squash_configuration_squash_uid
  }
  subnet_id = oci_core_subnet.my_subnet.id

  #Optional
  cluster_placement_group_id = oci_cluster_placement_groups_cluster_placement_group.test_cpg.id
#   defined_tags               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.lustre_file_system_defined_tags_value)
  display_name               = var.lustre_file_system_display_name
  file_system_description    = var.lustre_file_system_file_system_description

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_lustre_file_storage_lustre_file_system" "lustre_file_system" {
  lustre_file_system_id = oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
  #Optional
#   availability_domain   = var.lustre_file_system_availability_domain
#   compartment_id        = var.compartment_id
#   display_name          = var.lustre_file_system_display_name
#   state                 = var.lustre_file_system_state
}

data "oci_lustre_file_storage_lustre_file_systems" "lustre_file_systems" {
  compartment_id      = var.compartment_ocid
  #Optional
#   availability_domain = var.lustre_file_system_availability_domain
#   display_name        = var.lustre_file_system_display_name
#   id                  = var.lustre_file_system_id
#   state               = var.lustre_file_system_state
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}g