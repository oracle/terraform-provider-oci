// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
  #Required
  compartment_id            = var.compartment_ocid
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id
  cpu_core_count_per_node   = "20"
  autonomous_data_storage_size_in_tbs = "2.0"
  memory_per_oracle_compute_unit_in_gbs = "5"
  total_container_databases             = "2"
  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name              = "autonomousVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  state                     = "AVAILABLE"
}

variable "kms_vault_ocid" {
}

variable "okv_secret" {
}

resource "oci_database_key_store" "test_key_store" {
  compartment_id = var.compartment_ocid
  display_name = "Key Store"
  type_details {
    admin_username = "username1"
    connection_ips = ["192.1.1.1", "192.1.1.2"]
    secret_id = var.okv_secret
    type = "ORACLE_KEY_VAULT"
    vault_id = var.kms_vault_ocid
  }
}

