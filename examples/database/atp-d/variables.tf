// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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