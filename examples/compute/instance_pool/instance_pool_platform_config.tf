// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "instance_configuration_platform_config_instance_shape" {
  default = "BM.DenseIO.E4.128"
}

variable "instance_configuration_platform_config_numa_nodes_per_socket" {
  default = "NPS1"
}

variable "instance_configuration_platform_config_type" {
  default = "AMD_MILAN_BM"
}

resource "oci_core_instance_configuration" "test_instance_configuration_platform_config" {
  count = length(data.oci_core_images.instance_config_supported_platform_config_shape_images.images) == 0 ? 0 : 1
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfigurationPlatformConfig"

  instance_details {
    instance_type = "compute"

    block_volumes {
      create_details {
        compartment_id      = var.compartment_ocid
        display_name        = "TestCreateVolumeDetails"
        availability_domain = data.oci_identity_availability_domain.ad.name
        size_in_gbs         = 50
        vpus_per_gb         = 2
      }

      attach_details {
        type                                = "paravirtualized"
        display_name                        = "TestAttachVolumeDetails"
        is_read_only                        = true
        device                              = "TestDeviceName"
        is_pv_encryption_in_transit_enabled = true
        is_shareable                        = true
      }
    }

    launch_details {
      compartment_id                      = var.compartment_ocid
      ipxe_script                         = "ipxeScript"
      shape                               = var.instance_configuration_platform_config_instance_shape
      display_name                        = "TestInstanceConfigurationPlatformConfigLaunchDetails"
      is_pv_encryption_in_transit_enabled = false
      preferred_maintenance_action        = "LIVE_MIGRATE"
      launch_mode                         = "NATIVE"

      agent_config {
        is_management_disabled = false
        is_monitoring_disabled = false
      }

      launch_options {
        network_type = "PARAVIRTUALIZED"
      }

      instance_options {
        are_legacy_imds_endpoints_disabled = false
      }

      platform_config {
        type = var.instance_configuration_platform_config_type
        numa_nodes_per_socket = var.instance_configuration_platform_config_numa_nodes_per_socket
      }

      create_vnic_details {
        assign_public_ip       = true
        display_name           = "TestInstanceConfigurationPlatformConfigVNIC"
        skip_source_dest_check = false
      }

      extended_metadata = {
        some_string   = "stringA"
        nested_object = "{\"some_string\": \"stringB\", \"object\": {\"some_string\": \"stringC\"}}"
      }

      source_details {
        source_type = "image"
        image_id    = data.oci_core_images.instance_config_supported_platform_config_shape_images.images[0]["id"]
      }
    }
  }
}

# Gets a list of all images that support a given Instance shape
data "oci_core_images" "instance_config_supported_platform_config_shape_images" {
  compartment_id   = var.tenancy_ocid
  shape            = var.instance_configuration_platform_config_instance_shape
  operating_system = "Oracle Linux"

}
