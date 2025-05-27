// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

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
variable "compartment_id" {
}

variable "managed_instance_display_name" {
}

variable "managed_instance_id" {
}

variable "managed_instance_plugin_status" {
}

variable "wls_domain_display_name" {
}

variable "wls_domain_id" {
}

variable "wls_domain_middleware_type" {
}

variable "wls_domain_patch_readiness_status" {
}

variable "wls_domain_state" {
}

variable "wls_domain_weblogic_version" {
}

variable "wls_domain_server_name" {
}

variable "wls_server_id" {
}

variable "wls_mw_backup_id" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_wlms_managed_instances" "test_managed_instances" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.managed_instance_display_name
  id             = var.managed_instance_id
  plugin_status  = var.managed_instance_plugin_status
}

data "oci_wlms_managed_instance" "test_managed_instance" {
  #Required
  managed_instance_id = var.managed_instance_id
}

data "oci_wlms_wls_domains" "test_wls_domains" {

  #Optional
  compartment_id         = var.compartment_id
  display_name           = var.wls_domain_display_name
  id                     = var.wls_domain_id
  middleware_type        = var.wls_domain_middleware_type
  patch_readiness_status = var.wls_domain_patch_readiness_status
  state                  = var.wls_domain_state
  weblogic_version       = var.wls_domain_weblogic_version
}

data "oci_wlms_wls_domain" "test_wls_domain" {
  #Required
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_servers" "test_wls_domain_servers" {
  #Required
  wls_domain_id = var.wls_domain_id

  #Optional
  name = var.wls_domain_server_name
}

data "oci_wlms_wls_domain_server" "test_wls_domain_server" {
  #Required
  server_id = var.wls_server_id
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_server_installed_patches" "test_wls_domain_server_installed_patches" {
  #Required
  server_id     = var.wls_server_id
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_server_backups" "test_wls_domain_server_backups" {
  #Required
  server_id     = var.wls_server_id
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_server_backup" "test_wls_domain_server_backup" {
  #Required
  backup_id = var.wls_mw_backup_id
  server_id = var.wls_server_id
  wls_domain_id = var.wls_domain_id
}


data "oci_wlms_wls_domain_server_backup_content" "test_wls_domain_server_backup_content" {
  #Required
  backup_id     = var.wls_mw_backup_id
  server_id     = var.wls_server_id
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_scan_results" "test_wls_domain_scan_results" {
  #Required
  wls_domain_id = var.wls_domain_id

  #Optional
  server_name = var.wls_domain_server_name
}

data "oci_wlms_wls_domain_applicable_patches" "test_wls_domain_applicable_patches" {
  #Required
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_wls_domain_agreement_records" "test_wls_domain_agreement_records" {
  #Required
  wls_domain_id = var.wls_domain_id
}

data "oci_wlms_managed_instance_servers" "test_managed_instance_servers" {
  #Required
  managed_instance_id = var.managed_instance_id

  #Optional
  name = var.wls_domain_server_name
}

data "oci_wlms_managed_instance_server" "test_managed_instance_server" {

  #Required
  managed_instance_id = var.managed_instance_id
  server_id = var.wls_server_id
}


data "oci_wlms_managed_instance_server_installed_patches" "test_managed_instance_server_installed_patches" {
  #Required
  managed_instance_id = var.managed_instance_id
  server_id           = var.wls_server_id
}

data "oci_wlms_managed_instance_scan_results" "test_managed_instance_scan_results" {
  #Required
  managed_instance_id = var.managed_instance_id

  #Optional
  server_name   = var.wls_domain_server_name
  wls_domain_id = var.wls_domain_id
}


