// Provider OCI variables

variable "tenancy_ocid" {
  description = "The OCID of the tenancy"
  type        = string
}

variable "user_ocid" {
  description = "The OCID of the user"
  type        = string
}

variable "fingerprint" {
  description = "The fingerprint of the user's API key"
  type        = string
}

variable "private_key_path" {
  description = "The path to the private key file"
  type        = string
}

variable "region" {
  description = "The OCI region"
  type        = string
}

variable "compartment_id" {
  description = "The OCID of the compartment"
  type        = string
}

// FSU resources variables for tests

variable "fsu_db_19_target_1" {
  description = "The OCID of the first database target with version 19"
  type        = string
}

variable "fsu_db_23_target_1" {
  description = "The OCID of the first database target with version 23"
  type        = string
}

variable "fsu_db_26_target_1" {
  description = "The OCID of the first database target with version 26"
  type        = string
}

variable "fsu_gi_19_target_1" {
  description = "The OCID of the first GI target with version 19"
  type        = string
}

variable "fsu_gi_23_target_1" {
  description = "The OCID of the first GI target with version 23"
  type        = string
}

variable "fsu_gi_26_target_1" {
  description = "The OCID of the first GI target with version 26"
  type        = string
}

variable "fsu_db_19_software_image_1" {
  description = "The OCID of the first software image for database version 19"
  type        = string
}

variable "fsu_db_19_grid_software_image_1" {
  description = "The OCID of the first software image for GI version 19"
  type        = string
}

