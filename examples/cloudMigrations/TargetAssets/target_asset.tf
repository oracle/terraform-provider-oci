// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

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
  default = "AVAILABLE"
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
  default = "availabilityDomain"
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
  default = "shape"
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
  default = "sourceType"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_migrations_target_asset" "test_target_asset" {
  #Required
  is_excluded_from_execution = var.target_asset_is_excluded_from_execution
  migration_plan_id          = oci_cloud_migrations_migration_plan.test_migration_plan.id
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
    capacity_reservation_id = oci_cloud_migrations_capacity_reservation.test_capacity_reservation.id
    compartment_id          = var.compartment_id
    create_vnic_details {

      #Optional
      assign_private_dns_record = var.target_asset_user_spec_create_vnic_details_assign_private_dns_record
      assign_public_ip          = var.target_asset_user_spec_create_vnic_details_assign_public_ip
      defined_tags              = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.target_asset_user_spec_create_vnic_details_defined_tags_value)
      display_name              = var.target_asset_user_spec_create_vnic_details_display_name
      freeform_tags             = var.target_asset_user_spec_create_vnic_details_freeform_tags
      hostname_label            = var.target_asset_user_spec_create_vnic_details_hostname_label
      nsg_ids                   = var.target_asset_user_spec_create_vnic_details_nsg_ids
      private_ip                = var.target_asset_user_spec_create_vnic_details_private_ip
      skip_source_dest_check    = var.target_asset_user_spec_create_vnic_details_skip_source_dest_check
      subnet_id                 = oci_core_subnet.test_subnet.id
      vlan_id                   = oci_core_vlan.test_vlan.id
    }
    dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
    defined_tags         = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.target_asset_user_spec_defined_tags_value)
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
      boot_volume_id          = oci_core_boot_volume.test_boot_volume.id
      boot_volume_size_in_gbs = var.target_asset_user_spec_source_details_boot_volume_size_in_gbs
      boot_volume_vpus_per_gb = var.target_asset_user_spec_source_details_boot_volume_vpus_per_gb
      image_id                = oci_core_image.test_image.id
      kms_key_id              = oci_kms_key.test_key.id
    }
  }

  #Optional
  block_volumes_performance = var.target_asset_block_volumes_performance
  ms_license                = var.target_asset_ms_license
}

data "oci_cloud_migrations_target_assets" "test_target_assets" {

  #Optional
  display_name      = var.target_asset_display_name
  migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
  state             = var.target_asset_state
  target_asset_id   = oci_cloud_migrations_target_asset.test_target_asset.id
}

