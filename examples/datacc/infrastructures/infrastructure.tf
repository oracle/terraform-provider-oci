// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "infrastructure_acfs_file_system_storage_in_gbs" {
  default = 1.0
}

variable "infrastructure_admin_networkcidr" {
  default = "adminNetworkcidr"
}

variable "infrastructure_backup_network_bonding_interface" {
  default = "BTBOND1"
}

variable "infrastructure_backup_network_bonding_mode" {
  default = "ACTIVE_BACKUP"
}

variable "infrastructure_client_network_bonding_interface" {
  default = "BTBOND1"
}

variable "infrastructure_client_network_bonding_mode" {
  default = "ACTIVE_BACKUP"
}

variable "infrastructure_cloud_control_plane_server1" {
  default = "cloudControlPlaneServer1"
}

variable "infrastructure_cloud_control_plane_server2" {
  default = "cloudControlPlaneServer2"
}

variable "infrastructure_contacts_email" {
  default = "email"
}

variable "infrastructure_contacts_is_contact_mos_validated" {
  default = true
}

variable "infrastructure_contacts_is_primary" {
  default = true
}

variable "infrastructure_contacts_name" {
  default = "name"
}

variable "infrastructure_contacts_phone_number" {
  default = "phoneNumber"
}

variable "infrastructure_corporate_proxy" {
  default = "corporateProxy"
}

variable "infrastructure_cps_network_bonding_interface" {
  default = "BTBOND1"
}

variable "infrastructure_cps_network_bonding_mode" {
  default = "ACTIVE_BACKUP"
}

variable "infrastructure_data_disk_percentage" {
  default = 10
}

variable "infrastructure_defined_tags_value" {
  default = "value"
}

variable "infrastructure_description" {
  default = "description"
}

variable "infrastructure_display_name" {
  default = "displayName"
}

variable "infrastructure_dns_servers" {
  default = []
}

variable "infrastructure_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "infrastructure_gateway" {
  default = "gateway"
}

variable "infrastructure_maintenance_window_custom_action_timeout_in_mins" {
  default = 10
}

variable "infrastructure_maintenance_window_days_of_week" {
  default = []
}

variable "infrastructure_maintenance_window_hours_of_day" {
  default = []
}

variable "infrastructure_maintenance_window_is_custom_action_timeout_enabled" {
  default = true
}

variable "infrastructure_maintenance_window_is_monthly_patching_enabled" {
  default = false
}

variable "infrastructure_maintenance_window_lead_time_in_weeks" {
  default = 10
}

variable "infrastructure_maintenance_window_months" {
  default = []
}

variable "infrastructure_maintenance_window_patching_mode" {
  default = "ROLLING"
}

variable "infrastructure_maintenance_window_preference" {
  default = "NO_PREFERENCE"
}

variable "infrastructure_maintenance_window_weeks_of_month" {
  default = []
}

variable "infrastructure_netmask" {
  default = "netmask"
}

variable "infrastructure_ntp_servers" {
  default = []
}

variable "infrastructure_shape" {
  default = "SIX_SSDS"
}

variable "infrastructure_state" {
  default = []
}

variable "infrastructure_system_model" {
  default = "X11_HA_768"
}

variable "vlan_id" {
  default = "vlanId"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_datacc_infrastructure" "test_infrastructure" {
  #Required
  cloud_control_plane_server1 = var.infrastructure_cloud_control_plane_server1
  cloud_control_plane_server2 = var.infrastructure_cloud_control_plane_server2
  compartment_id              = var.compartment_id
  display_name                = var.infrastructure_display_name
  dns_servers                 = var.infrastructure_dns_servers
  gateway                     = var.infrastructure_gateway
  netmask                     = var.infrastructure_netmask
  ntp_servers                 = var.infrastructure_ntp_servers
  shape                       = var.infrastructure_shape
  system_model                = var.infrastructure_system_model

  #Optional
  acfs_file_system_storage_in_gbs  = var.infrastructure_acfs_file_system_storage_in_gbs
  admin_networkcidr                = var.infrastructure_admin_networkcidr
  backup_network_bonding_interface = var.infrastructure_backup_network_bonding_interface
  backup_network_bonding_mode      = var.infrastructure_backup_network_bonding_mode
  client_network_bonding_interface = var.infrastructure_client_network_bonding_interface
  client_network_bonding_mode      = var.infrastructure_client_network_bonding_mode
  contacts {
    #Required
    email      = var.infrastructure_contacts_email
    is_primary = var.infrastructure_contacts_is_primary
    name       = var.infrastructure_contacts_name

    #Optional
    is_contact_mos_validated = var.infrastructure_contacts_is_contact_mos_validated
    phone_number             = var.infrastructure_contacts_phone_number
  }
  corporate_proxy               = var.infrastructure_corporate_proxy
  cps_network_bonding_interface = var.infrastructure_cps_network_bonding_interface
  cps_network_bonding_mode      = var.infrastructure_cps_network_bonding_mode
  data_disk_percentage          = var.infrastructure_data_disk_percentage
  defined_tags                  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.infrastructure_defined_tags_value)
  description                   = var.infrastructure_description
  freeform_tags                 = var.infrastructure_freeform_tags
  maintenance_window {

    #Optional
    custom_action_timeout_in_mins    = var.infrastructure_maintenance_window_custom_action_timeout_in_mins
    days_of_week                     = var.infrastructure_maintenance_window_days_of_week
    hours_of_day                     = var.infrastructure_maintenance_window_hours_of_day
    is_custom_action_timeout_enabled = var.infrastructure_maintenance_window_is_custom_action_timeout_enabled
    is_monthly_patching_enabled      = var.infrastructure_maintenance_window_is_monthly_patching_enabled
    lead_time_in_weeks               = var.infrastructure_maintenance_window_lead_time_in_weeks
    months                           = var.infrastructure_maintenance_window_months
    patching_mode                    = var.infrastructure_maintenance_window_patching_mode
    preference                       = var.infrastructure_maintenance_window_preference
    weeks_of_month                   = var.infrastructure_maintenance_window_weeks_of_month
  }
  vlan_id = var.vlan.id
}

data "oci_datacc_infrastructures" "test_infrastructures" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.infrastructure_display_name
  state        = var.infrastructure_state
}

