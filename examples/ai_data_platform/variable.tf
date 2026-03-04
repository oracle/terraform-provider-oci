variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ai_data_platform_ai_data_platform_type" {
default = "aiDataPlatformType"
}

variable "ai_data_platform_display_name" {
default = "displayName"
}

variable "ai_data_platform_workspace_name" {
  default = "workspaceName"
}

variable "ai_data_platform_exclude_lifecycle_state" {
default = "CREATING"
}

variable "ai_data_platform_freeform_tags" {
default = { "Department" = "Finance" }
}

variable "ai_data_platform_include_legacy" {
default = "true"
}

variable "ai_data_platform_state" {
default = "ACTIVE"
}

variable "ai_data_platform_system_tags" {
default = "value"
}