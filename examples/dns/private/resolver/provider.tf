// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

# Configure the OCI provider and authentication profile for this example.
# See README.md for setup and a walkthrough.

# Version 8.16.0 waits for the VCN resolver association to become available.
terraform {
  required_providers {
    oci = {
      source  = "oracle/oci"
      version = ">= 8.16.0"
    }
  }
}

variable "compartment_ocid" {
  description = "Compartment in which to create the example resources."
  type        = string
}

variable "region" {
  description = "OCI region in which to run the example."
  type        = string
}

variable "auth" {
  description = "Authentication type used by the named OCI configuration profile."
  type        = string
  default     = "APIKey"
}

variable "config_file_profile" {
  description = "Profile in the local OCI configuration file."
  type        = string
  default     = "DEFAULT"
}

# Defaults to API key authentication with the DEFAULT OCI configuration profile.
provider "oci" {
  region              = var.region
  auth                = var.auth
  config_file_profile = var.config_file_profile
}
