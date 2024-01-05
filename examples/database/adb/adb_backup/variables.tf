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

# The Autonomous Database needs to be pre-configured for backups. Refer:
# https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/atpbackingup.htm
variable "autonomous_database_id" {
  default = ""
}

variable "autonomous_database_backup_display_name" {
  default = "Monthly Backup"
}

