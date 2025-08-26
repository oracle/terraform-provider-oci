# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_vm_std_x86
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemVMStdx86
#
#    FILE(S)
#      database_db_system_resource_vm_std_x86_test.go
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created



variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "ssh_public_key" {
  type = string
}

variable "ssh_private_key" {
}

# DBSystem specific
variable "db_system_shape" {
  default = "VM.Standard.x86"
}

variable "compute_model" {
  default = "ECPU"
}

variable "compute_count" {
  default = "4"
}

variable "db_system_storage_volume_performance_mode" {
  default = "BALANCED"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "db_version" {
  default = "19.0.0.0"
}

variable "db_disk_redundancy" {
  default = "NORMAL"
}

variable "sparse_diskgroup" {
  default = true
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
  default = "1"
}

variable "test_database_software_image_ocid" {

}