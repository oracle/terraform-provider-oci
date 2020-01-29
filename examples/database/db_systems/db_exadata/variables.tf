// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

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
