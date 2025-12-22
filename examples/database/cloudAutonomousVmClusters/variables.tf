// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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

variable "autonomous_database_defined_tags_value" {
  default = "value"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_exadata_infrastructure_domain" {
  default = "subnetexadata.examplevcn.oraclevcn.com"
}

variable "autonomous_container_database_backup_config_recovery_window_in_days" {
  default = 10
}

variable "cloud_exadata_infrastructure_shape" {
  default = "Exadata.X8M"
}

variable "cloud_exadata_infrastructure_compute_count" {
  default = "2"
}

variable "cloud_exadata_infrastructure_storage_count" {
  default = "3"
}

variable "cloud_exadata_infrastructure_un_allocated_resource_db_servers" {
  default = []
}

variable "acd_db_version" {
  default = "19.28.0.1.0"
}

// key Store related var
variable "okv_secret" {
}

// key Store related var
variable "kms_vault_ocid" {
}

# Required for imported Autonomous VM Cluster resource - START
variable "cloud_autonomous_vm_cluster_display_name" {
  description = "The display name for the externally created Cloud Autonomous VM Cluster."
  type        = string
  default = "AWS_CAVM"
}

variable "cloud_exadata_infrastructure_id" {
  description = "The OCID of the Exadata Infrastructure hosting the Cloud Autonomous VM Cluster."
  type        = string
}

variable "subnet_id" {
  description = "The OCID of the subnet backing the Cloud Autonomous VM Cluster."
  type        = string
}
# Required for imported Autonomous VM Cluster resource - END
