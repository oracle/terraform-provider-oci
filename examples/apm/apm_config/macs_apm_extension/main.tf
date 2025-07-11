// Copyright (c) 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" { }
variable "user_ocid" { }
variable "fingerprint" { }
variable "private_key_path" { }
variable "region" { }
variable "compartment_ocid" { }
variable "management_agent_ocid" { }

variable "config_config_type" {
  default = "MACS_APM_EXTENSION"
}

variable "process_filter" {
  default = [".*org.apache.catalina.startup.Bootstrap.*", ".*jetty.*"]
}

variable "run_as_user" {
  default = "tomcat"
}

variable "service_name" {
  default = "Tomcat"
}

variable "agent_version" {
  default = "1.16.0.585"
}

variable "attach_install_dir" {
  default = "/opt/oracle/apm_attach_process"
}

variable "display_name" {
  default = "Display name"
}

variable "config_defined_tags_value" {
  default = "value"
}

variable "config_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "apm_domain_description" {
  default = "description"
}

variable "apm_domain_display_name" {
  default = "displayName"
}

variable "apm_domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "apm_domain_is_free_tier" {
  default = false
}

variable "apm_domain_state" {
  default = "ACTIVE"
}

variable "data_key_data_key_type" {
  default = "PRIVATE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.apm_domain_display_name

  #Optional
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = true
}


resource "oci_apm_config_config" "test_config" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_type   = var.config_config_type
  management_agent_id = var.management_agent_ocid
  process_filter = var.process_filter
  run_as_user = var.run_as_user
  service_name = var.service_name
  agent_version = var.agent_version
  attach_install_dir = var.attach_install_dir

  #Optional
  display_name = var.display_name
}

data "oci_apm_config_configs" "test_configs" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  config_type  = var.config_config_type
}
