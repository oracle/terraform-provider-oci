// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

# Gets the list of file systems in the compartment
data "oci_file_storage_file_systems" "file_systems" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  #Optional fields. Used by the service to filter the results when returning data to the client.
  #display_name = "my_fs_1"
  #id = "ocid1.filesystem.oc1.phx.aaaaaaaaaaaaawynobuhqllqojxwiotqnb4c2ylefuyqaaaa"
  #state = "DELETED"

  #filter {
  #  name = "defined_tags.example-tag-namespace-all.example-tag"
  #  values = ["value"]
  #}
}

# Gets the list of mount targets in the compartment
data "oci_file_storage_mount_targets" "mount_targets" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  #Optional fields. Used by the service to filter the results when returning data to the client.
  #display_name = var.mount_target_display_name
  #export_set_id = var.mount_target_export_set_id
  #id = var.mount_target_id
  #state = var.mount_target_state

  #filter {
  #  name = "freeform_tags.Department"
  #  values = ["Accounting"]
  #}
}

# Gets the list of exports in the compartment
data "oci_file_storage_exports" "exports" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional fields. Used by the service to filter the results when returning data to the client.
  #export_set_id = oci_file_storage_mount_target.my_mount_target_1.export_set_id
  #file_system_id = oci_file_storage_file_system.my_fs.id
  #id = var.export_id
  #state = var.export_state
}

# Gets a list of snapshots for a particular file system
data "oci_file_storage_snapshots" "snapshots" {
  #Required
  file_system_id = oci_file_storage_file_system.my_fs_1.id
  #Optional fields. Used by the service to filter the results when returning data to the client.
  #id = var.snapshot_id
  #state = var.snapshot_state

  #filter {
  #  name = "freeform_tags.Department"
  #  values = ["Accounting"]
  #}
}

# Gets a list of export sets in a compartment and availability domain
data "oci_file_storage_export_sets" "export_sets" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  #Optional fields. Used by the service to filter the results when returning data to the client.
  #display_name = var.export_set_display_name
  #id = var.export_set_id
  #state = var.export_set_state
}

data "oci_core_private_ips" "ip_mount_target1" {
  subnet_id = oci_file_storage_mount_target.my_mount_target_1.subnet_id

  filter {
    name   = "id"
    values = [oci_file_storage_mount_target.my_mount_target_1.private_ip_ids[0]]
  }
}

# Gets a list of replications in a compartment and availability domain
data "oci_file_storage_replications" "test_replications" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  #display_name   = var.replication_display_name
  #file_system_id = oci_file_storage_file_system.test_file_system.id
  #id             = var.replication_id
  #state          = var.replication_state
}

data "oci_file_storage_replication_targets" "test_replication_targets" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  #display_name = var.replication_target_display_name
  #id           = var.replication_target_id
  #state        = var.replication_target_state
}

# Gets a list of filesystem snapshot policies in a compartment and availability domain
data "oci_file_storage_filesystem_snapshot_policies" "filesystem_snapshot_policies" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  #display_name = var.filesystem_snapshot_policy_display_name
  #id           = var.filesystem_snapshot_policy_id
  #state        = var.filesystem_snapshot_policy_state
}


# Gets a list of outbound connectors in a compartment and availability domain
data "oci_file_storage_outbound_connectors" "outbound_connectors" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  #Optional
  #display_name = var.outbound_connector_display_name
  #id           = var.outbound_connector_id
  #state        = var.outbound_connector_state
}
