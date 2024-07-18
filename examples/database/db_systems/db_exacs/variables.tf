// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}

variable "ssh_public_key" {
  default = "ssh-rsa"
}
variable "ssh_private_key" {}

variable "cloud_exadata_infrastructure_shape" {
  default = "Exadata.X8M"
}

variable "cloud_exadata_infrastructure_cluster_placement_group_id" {
  default = null
}

variable "tenant_subscription_id" {
  default = null
}

variable "cloud_exadata_infrastructure_compute_count" {
  default = "2"
}

variable "cloud_exadata_infrastructure_storage_count" {
  default = "3"
}

variable "cloud_vm_cluster_cpu_core_count" {
  default = "8"
}

variable "cloud_vm_cluster_ocpu_count" {
  default = "8.0"
}

variable "cloud_vm_cluster_gi_version" {
  default = "19.0.0.0"
}

variable "cloud_vm_cluster_hostname" {
  default = "myoracledb"
}

variable "cloud_vm_cluster_scan_listener_port_tcp" {
  default = "1521"
}

variable "cloud_vm_cluster_scan_listener_port_tcp_ssl" {
  default = "2484"
}

# DBSystem specific
variable "db_system_shape" {
  default = "Exadata.Quarter1.84"
}

variable "cpu_core_count" {
  default = "22"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "db_version" {
  default = "12.1.0.2"
}

variable "db_disk_redundancy" {
  default = "HIGH"
}

variable "sparse_diskgroup" {
  default = true
}

variable "db_system_display_name" {
  default = "MyTFDBSystem"
}

variable "hostname" {
  default = "myoracledb"
}

variable "host_user_name" {
  default = "opc"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

variable "pdb_name" {
  default = "pdbName"
}

variable "data_storage_size_in_gb" {
  default = "256"
}

variable "license_model" {
  default = "LICENSE_INCLUDED"
}

variable "node_count" {
  default = "2"
}

variable "data_storage_percentage" {
  default = "40"
}

variable "time_zone" {
  default = "US/Pacific"
}

variable "cloud_vm_cluster_memory_size_in_gbs" {
  default = 60
}

variable "cloud_vm_cluster_data_storage_size_in_tbs" {
  default = 2.0
}

variable "cloud_vm_cluster_db_node_storage_size_in_gbs" {
  default = 120
}

variable "cloud_vm_cluster_db_servers" {
  default = []
}