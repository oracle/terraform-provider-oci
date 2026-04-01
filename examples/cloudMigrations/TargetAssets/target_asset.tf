// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default = "compartment_id"
}

variable "target_asset_block_volumes_performance" {
  default = 10
}

variable "target_asset_display_name" {
  default = "displayName"
}

variable "target_asset_is_excluded_from_execution" {
  default = false
}

variable "target_asset_ms_license" {
  default = "msLicense"
}

variable "target_asset_preferred_shape_type" {
  default = "VM"
}

variable "target_asset_state" {
  default = "ACTIVE"
}

variable "target_asset_type" {
  default = "INSTANCE"
}

variable "target_asset_user_spec_agent_config_are_all_plugins_disabled" {
  default = false
}

variable "target_asset_user_spec_agent_config_is_management_disabled" {
  default = false
}

variable "target_asset_user_spec_agent_config_is_monitoring_disabled" {
  default = false
}

variable "target_asset_user_spec_agent_config_plugins_config_desired_state" {
  default = "ENABLED"
}

variable "target_asset_user_spec_agent_config_plugins_config_name" {
  default = "name"
}

variable "target_asset_user_spec_availability_domain" {
  default = "oQNt:US-ASHBURN-AD-1"
}

variable "target_asset_user_spec_create_vnic_details_assign_private_dns_record" {
  default = false
}

variable "target_asset_user_spec_create_vnic_details_assign_public_ip" {
  default = false
}

variable "target_asset_user_spec_create_vnic_details_defined_tags_value" {
  default = "value"
}

variable "target_asset_user_spec_create_vnic_details_display_name" {
  default = "displayName"
}

variable "target_asset_user_spec_create_vnic_details_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "target_asset_user_spec_create_vnic_details_hostname_label" {
  default = "hostnameLabel"
}

variable "target_asset_user_spec_create_vnic_details_nsg_ids" {
  default = []
}

variable "target_asset_user_spec_create_vnic_details_private_ip" {
  default = "privateIp"
}

variable "target_asset_user_spec_create_vnic_details_skip_source_dest_check" {
  default = false
}

variable "target_asset_user_spec_defined_tags_value" {
  default = "value"
}

variable "target_asset_user_spec_display_name" {
  default = "displayName"
}

variable "target_asset_user_spec_fault_domain" {
  default = "faultDomain"
}

variable "target_asset_user_spec_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "target_asset_user_spec_hostname_label" {
  default = "hostnameLabel"
}

variable "target_asset_user_spec_instance_options_are_legacy_imds_endpoints_disabled" {
  default = false
}

variable "target_asset_user_spec_ipxe_script" {
  default = "ipxeScript"
}

variable "target_asset_user_spec_is_pv_encryption_in_transit_enabled" {
  default = false
}

variable "target_asset_user_spec_preemptible_instance_config_preemption_action_preserve_boot_volume" {
  default = false
}

variable "target_asset_user_spec_preemptible_instance_config_preemption_action_type" {
  default = "TERMINATE"
}

variable "target_asset_user_spec_shape" {
  default = "VM.Standard.E4.Flex"
}

variable "target_asset_user_spec_shape_config_baseline_ocpu_utilization" {
  default = "BASELINE_1_8"
}

variable "target_asset_user_spec_shape_config_memory_in_gbs" {
  default = 1.0
}

variable "target_asset_user_spec_shape_config_ocpus" {
  default = 1.0
}

variable "target_asset_user_spec_source_details_boot_volume_size_in_gbs" {
  default = 10
}

variable "target_asset_user_spec_source_details_boot_volume_vpus_per_gb" {
  default = 10
}

variable "target_asset_user_spec_source_details_source_type" {
  default = "image"
}

variable "migration_plan_id" {
  default = "migration_plan_id"
}

variable "subnet_id" {
  default = "subnet_id"
}

variable "image_id" {
  default = "image_id"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

resource "oci_cloud_migrations_target_asset" "test_target_asset" {
  #Required
  is_excluded_from_execution = var.target_asset_is_excluded_from_execution
  migration_plan_id          = var.migration_plan_id
  preferred_shape_type       = var.target_asset_preferred_shape_type
  type                       = var.target_asset_type
  user_spec {

    #Optional
    agent_config {

      #Optional
      are_all_plugins_disabled = var.target_asset_user_spec_agent_config_are_all_plugins_disabled
      is_management_disabled   = var.target_asset_user_spec_agent_config_is_management_disabled
      is_monitoring_disabled   = var.target_asset_user_spec_agent_config_is_monitoring_disabled
      plugins_config {
        #Required
        desired_state = var.target_asset_user_spec_agent_config_plugins_config_desired_state
        name          = var.target_asset_user_spec_agent_config_plugins_config_name
      }
    }
    availability_domain     = var.target_asset_user_spec_availability_domain
    compartment_id          = var.compartment_id
    create_vnic_details {

      #Optional
      assign_private_dns_record = var.target_asset_user_spec_create_vnic_details_assign_private_dns_record
      assign_public_ip          = var.target_asset_user_spec_create_vnic_details_assign_public_ip
      display_name              = var.target_asset_user_spec_create_vnic_details_display_name
      freeform_tags             = var.target_asset_user_spec_create_vnic_details_freeform_tags
      hostname_label            = var.target_asset_user_spec_create_vnic_details_hostname_label
      nsg_ids                   = var.target_asset_user_spec_create_vnic_details_nsg_ids
      private_ip                = var.target_asset_user_spec_create_vnic_details_private_ip
      skip_source_dest_check    = var.target_asset_user_spec_create_vnic_details_skip_source_dest_check
      subnet_id                 = var.subnet_id
    }
    display_name         = var.target_asset_user_spec_display_name
    fault_domain         = var.target_asset_user_spec_fault_domain
    freeform_tags        = var.target_asset_user_spec_freeform_tags
    hostname_label       = var.target_asset_user_spec_hostname_label
    instance_options {

      #Optional
      are_legacy_imds_endpoints_disabled = var.target_asset_user_spec_instance_options_are_legacy_imds_endpoints_disabled
    }
    ipxe_script                         = var.target_asset_user_spec_ipxe_script
    is_pv_encryption_in_transit_enabled = var.target_asset_user_spec_is_pv_encryption_in_transit_enabled
    preemptible_instance_config {
      #Required
      preemption_action {
        #Required
        type = var.target_asset_user_spec_preemptible_instance_config_preemption_action_type

        #Optional
        preserve_boot_volume = var.target_asset_user_spec_preemptible_instance_config_preemption_action_preserve_boot_volume
      }
    }
    shape = var.target_asset_user_spec_shape
    shape_config {

      #Optional
      baseline_ocpu_utilization = var.target_asset_user_spec_shape_config_baseline_ocpu_utilization
      memory_in_gbs             = var.target_asset_user_spec_shape_config_memory_in_gbs
      ocpus                     = var.target_asset_user_spec_shape_config_ocpus
    }
    source_details {
      #Required
      source_type = var.target_asset_user_spec_source_details_source_type

      #Optional
      boot_volume_size_in_gbs = var.target_asset_user_spec_source_details_boot_volume_size_in_gbs
      boot_volume_vpus_per_gb = var.target_asset_user_spec_source_details_boot_volume_vpus_per_gb
      image_id                = var.image_id
    }
  }

  #Optional
  block_volumes_performance = var.target_asset_block_volumes_performance
  ms_license                = var.target_asset_ms_license
}

data "oci_cloud_migrations_target_assets" "test_target_assets" {

  #Optional
  display_name      = var.target_asset_display_name
  migration_plan_id = var.migration_plan_id
  state             = var.target_asset_state
  target_asset_id   = oci_cloud_migrations_target_asset.test_target_asset.id
}

