# Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
  type = string
}

variable "region" {
  type = string
}

variable "compartment_id" {
  type = string
}

variable "ssh_public_key" {
  type    = string
}

variable "admin_password" {
  type    = string
  default = "BEstrO0ng_#12"
}

variable "os_patch_action" {
  default = "PRECHECK"
}

variable "os_patch_trigger" {
  type        = number
  default     = 1
}

variable "db_system_os_patch_history_entry_action" {
  default = "PRECHECK"
}

