---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_configuration"
sidebar_current: "docs-oci-resource-core-instance_configuration"
description: |-
  Provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_configuration
This resource provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service.

Creates an instance configuration. An instance configuration is a template that defines the
settings to use when creating Compute instances.


## Example Usage

```hcl
resource "oci_core_instance_configuration" "test_instance_configuration" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_configuration_display_name
	freeform_tags = {"Department"= "Finance"}
	instance_details {
		#Required
		instance_type = var.instance_configuration_instance_details_instance_type

		#Optional
		block_volumes {

			#Optional
			attach_details {
				#Required
				type = var.instance_configuration_instance_details_block_volumes_attach_details_type

				#Optional
				device = var.instance_configuration_instance_details_block_volumes_attach_details_device
				display_name = var.instance_configuration_instance_details_block_volumes_attach_details_display_name
				is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_block_volumes_attach_details_is_pv_encryption_in_transit_enabled
				is_read_only = var.instance_configuration_instance_details_block_volumes_attach_details_is_read_only
				is_shareable = var.instance_configuration_instance_details_block_volumes_attach_details_is_shareable
				use_chap = var.instance_configuration_instance_details_block_volumes_attach_details_use_chap
			}
			create_details {

				#Optional
				autotune_policies {
					#Required
					autotune_type = var.instance_configuration_instance_details_block_volumes_create_details_autotune_policies_autotune_type

					#Optional
					max_vpus_per_gb = var.instance_configuration_instance_details_block_volumes_create_details_autotune_policies_max_vpus_per_gb
				}
				availability_domain = var.instance_configuration_instance_details_block_volumes_create_details_availability_domain
				backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
				block_volume_replicas {
					#Required
					availability_domain = var.instance_configuration_instance_details_block_volumes_create_details_block_volume_replicas_availability_domain

					#Optional
					display_name = var.instance_configuration_instance_details_block_volumes_create_details_block_volume_replicas_display_name
				}
				cluster_placement_group_id = var.cluster_placement_group_id
				compartment_id = var.compartment_id
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_block_volumes_create_details_display_name
				freeform_tags = {"Department"= "Finance"}
				is_auto_tune_enabled = var.instance_configuration_instance_details_block_volumes_create_details_is_auto_tune_enabled
				kms_key_id = oci_kms_key.test_key.id
				size_in_gbs = var.instance_configuration_instance_details_block_volumes_create_details_size_in_gbs
				source_details {
					#Required
					type = var.instance_configuration_instance_details_block_volumes_create_details_source_details_type

					#Optional
					id = var.instance_configuration_instance_details_block_volumes_create_details_source_details_id
				}
				vpus_per_gb = var.instance_configuration_instance_details_block_volumes_create_details_vpus_per_gb
				xrc_kms_key_id = oci_kms_key.test_key.id
			}
			volume_id = oci_core_volume.test_volume.id
		}
		launch_details {

			#Optional
			agent_config {

				#Optional
				are_all_plugins_disabled = var.instance_configuration_instance_details_launch_details_agent_config_are_all_plugins_disabled
				is_management_disabled = var.instance_configuration_instance_details_launch_details_agent_config_is_management_disabled
				is_monitoring_disabled = var.instance_configuration_instance_details_launch_details_agent_config_is_monitoring_disabled
				plugins_config {

					#Optional
					desired_state = var.instance_configuration_instance_details_launch_details_agent_config_plugins_config_desired_state
					name = var.instance_configuration_instance_details_launch_details_agent_config_plugins_config_name
				}
			}
			availability_config {

				#Optional
				is_live_migration_preferred = var.instance_configuration_instance_details_launch_details_availability_config_is_live_migration_preferred
				recovery_action = var.instance_configuration_instance_details_launch_details_availability_config_recovery_action
			}
			availability_domain = var.instance_configuration_instance_details_launch_details_availability_domain
			capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id
			cluster_placement_group_id = oci_identity_group.test_group.id
			compartment_id = var.compartment_id
			create_vnic_details {

				#Optional
				assign_ipv6ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_ipv6ip
				assign_private_dns_record = var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_private_dns_record
				assign_public_ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_public_ip
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_launch_details_create_vnic_details_display_name
				freeform_tags = {"Department"= "Finance"}
				hostname_label = var.instance_configuration_instance_details_launch_details_create_vnic_details_hostname_label
				ipv6address_ipv6subnet_cidr_pair_details {

					#Optional
					ipv6address = var.instance_configuration_instance_details_launch_details_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6address
					ipv6subnet_cidr = var.instance_configuration_instance_details_launch_details_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
				}				
				nsg_ids = var.instance_configuration_instance_details_launch_details_create_vnic_details_nsg_ids
				private_ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_private_ip
				security_attributes = var.instance_configuration_instance_details_launch_details_create_vnic_details_security_attributes
				skip_source_dest_check = var.instance_configuration_instance_details_launch_details_create_vnic_details_skip_source_dest_check
				subnet_id = oci_core_subnet.test_subnet.id
			}
			dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
			defined_tags = {"Operations.CostCenter"= "42"}
			display_name = var.instance_configuration_instance_details_launch_details_display_name
			extended_metadata = var.instance_configuration_instance_details_launch_details_extended_metadata
			fault_domain = var.instance_configuration_instance_details_launch_details_fault_domain
			freeform_tags = {"Department"= "Finance"}
			instance_options {

				#Optional
				are_legacy_imds_endpoints_disabled = var.instance_configuration_instance_details_launch_details_instance_options_are_legacy_imds_endpoints_disabled
			}
			ipxe_script = var.instance_configuration_instance_details_launch_details_ipxe_script
			is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_launch_details_is_pv_encryption_in_transit_enabled
			launch_mode = var.instance_configuration_instance_details_launch_details_launch_mode
			launch_options {

				#Optional
				boot_volume_type = var.instance_configuration_instance_details_launch_details_launch_options_boot_volume_type
				firmware = var.instance_configuration_instance_details_launch_details_launch_options_firmware
				is_consistent_volume_naming_enabled = var.instance_configuration_instance_details_launch_details_launch_options_is_consistent_volume_naming_enabled
				is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_launch_details_launch_options_is_pv_encryption_in_transit_enabled
				network_type = var.instance_configuration_instance_details_launch_details_launch_options_network_type
				remote_data_volume_type = var.instance_configuration_instance_details_launch_details_launch_options_remote_data_volume_type
			}
			licensing_configs {
				#Required
				type = var.instance_configuration_instance_details_launch_details_licensing_configs_type

				#Optional
				license_type = var.instance_configuration_instance_details_launch_details_licensing_configs_license_type
			}
			metadata = var.instance_configuration_instance_details_launch_details_metadata
			platform_config {
				#Required
				type = var.instance_configuration_instance_details_launch_details_platform_config_type

				#Optional
				are_virtual_instructions_enabled = var.instance_configuration_instance_details_launch_details_platform_config_are_virtual_instructions_enabled
				config_map = var.instance_configuration_instance_details_launch_details_platform_config_config_map
				is_access_control_service_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_access_control_service_enabled
				is_input_output_memory_management_unit_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_input_output_memory_management_unit_enabled
				is_measured_boot_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_measured_boot_enabled
				is_memory_encryption_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_memory_encryption_enabled
				is_secure_boot_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_secure_boot_enabled
				is_symmetric_multi_threading_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_symmetric_multi_threading_enabled
				is_trusted_platform_module_enabled = var.instance_configuration_instance_details_launch_details_platform_config_is_trusted_platform_module_enabled
				numa_nodes_per_socket = var.instance_configuration_instance_details_launch_details_platform_config_numa_nodes_per_socket
				percentage_of_cores_enabled = var.instance_configuration_instance_details_launch_details_platform_config_percentage_of_cores_enabled
			}
			preemptible_instance_config {

				#Optional
				preemption_action {
					#Required
					type = var.instance_configuration_instance_details_launch_details_preemptible_instance_config_preemption_action_type

					#Optional
					preserve_boot_volume = var.instance_configuration_instance_details_launch_details_preemptible_instance_config_preemption_action_preserve_boot_volume
				}
			}
			preferred_maintenance_action = var.instance_configuration_instance_details_launch_details_preferred_maintenance_action
			security_attributes = var.instance_configuration_instance_details_launch_details_security_attributes
			shape = var.instance_configuration_instance_details_launch_details_shape
			shape_config {

				#Optional
				baseline_ocpu_utilization = var.instance_configuration_instance_details_launch_details_shape_config_baseline_ocpu_utilization
				memory_in_gbs = var.instance_configuration_instance_details_launch_details_shape_config_memory_in_gbs
				nvmes = var.instance_configuration_instance_details_launch_details_shape_config_nvmes
				ocpus = var.instance_configuration_instance_details_launch_details_shape_config_ocpus
				vcpus = var.instance_configuration_instance_details_launch_details_shape_config_vcpus
			}
			source_details {
				#Required
				source_type = var.instance_configuration_instance_details_launch_details_source_details_source_type

				#Optional
				boot_volume_id = oci_core_boot_volume.test_boot_volume.id
				boot_volume_size_in_gbs = var.instance_configuration_instance_details_launch_details_source_details_boot_volume_size_in_gbs
				boot_volume_vpus_per_gb = var.instance_configuration_instance_details_launch_details_source_details_boot_volume_vpus_per_gb
				image_id = oci_core_image.test_image.id
				kms_key_id = oci_kms_key.test_key.id
				instance_source_image_filter_details {

					#Optional
					compartment_id = var.compartment_id
					defined_tags_filter = var.instance_configuration_instance_details_launch_details_source_details_instance_source_image_filter_details_defined_tags_filter
					operating_system = var.instance_configuration_instance_details_launch_details_source_details_instance_source_image_filter_details_operating_system
					operating_system_version = var.instance_configuration_instance_details_launch_details_source_details_instance_source_image_filter_details_operating_system_version
				}
			}
		}
		options {

			#Optional
			block_volumes {

				#Optional
				attach_details {
					#Required
					type = var.instance_configuration_instance_details_options_block_volumes_attach_details_type

					#Optional
					device = var.instance_configuration_instance_details_options_block_volumes_attach_details_device
					display_name = var.instance_configuration_instance_details_options_block_volumes_attach_details_display_name
					is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_options_block_volumes_attach_details_is_pv_encryption_in_transit_enabled
					is_read_only = var.instance_configuration_instance_details_options_block_volumes_attach_details_is_read_only
					is_shareable = var.instance_configuration_instance_details_options_block_volumes_attach_details_is_shareable
					use_chap = var.instance_configuration_instance_details_options_block_volumes_attach_details_use_chap
				}
				create_details {

					#Optional
					autotune_policies {
						#Required
						autotune_type = var.instance_configuration_instance_details_options_block_volumes_create_details_autotune_policies_autotune_type

						#Optional
						max_vpus_per_gb = var.instance_configuration_instance_details_options_block_volumes_create_details_autotune_policies_max_vpus_per_gb
					}
					availability_domain = var.instance_configuration_instance_details_options_block_volumes_create_details_availability_domain
					backup_policy_id = data.oci_core_volume_backup_policies.test_volume_backup_policies.volume_backup_policies.0.id
					cluster_placement_group_id = var.cluster_placement_group_id
					compartment_id = var.compartment_id
					defined_tags = {"Operations.CostCenter"= "42"}
					display_name = var.instance_configuration_instance_details_options_block_volumes_create_details_display_name
					freeform_tags = {"Department"= "Finance"}
					kms_key_id = oci_kms_key.test_key.id
					size_in_gbs = var.instance_configuration_instance_details_options_block_volumes_create_details_size_in_gbs
					source_details {
						#Required
						type = var.instance_configuration_instance_details_options_block_volumes_create_details_source_details_type

						#Optional
						id = var.instance_configuration_instance_details_options_block_volumes_create_details_source_details_id
					}
					vpus_per_gb = var.instance_configuration_instance_details_options_block_volumes_create_details_vpus_per_gb
					xrc_kms_key_id = oci_kms_key.test_key.id
				}
				volume_id = oci_core_volume.test_volume.id
			}
			launch_details {

				#Optional
				agent_config {

					#Optional
					are_all_plugins_disabled = var.instance_configuration_instance_details_options_launch_details_agent_config_are_all_plugins_disabled
					is_management_disabled = var.instance_configuration_instance_details_options_launch_details_agent_config_is_management_disabled
					is_monitoring_disabled = var.instance_configuration_instance_details_options_launch_details_agent_config_is_monitoring_disabled
					plugins_config {

						#Optional
						desired_state = var.instance_configuration_instance_details_options_launch_details_agent_config_plugins_config_desired_state
						name = var.instance_configuration_instance_details_options_launch_details_agent_config_plugins_config_name
					}
				}
				availability_config {

					#Optional
					recovery_action = var.instance_configuration_instance_details_options_launch_details_availability_config_recovery_action
				}
				availability_domain = var.instance_configuration_instance_details_options_launch_details_availability_domain
				capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id
				cluster_placement_group_id = oci_identity_group.test_group.id
				compartment_id = var.compartment_id
				create_vnic_details {

					#Optional
					assign_ipv6ip = var.instance_configuration_instance_details_launch_details_create_vnic_details_assign_ipv6ip
					assign_private_dns_record = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_assign_private_dns_record
					assign_public_ip = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_assign_public_ip
					defined_tags = {"Operations.CostCenter"= "42"}
					display_name = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_display_name
					freeform_tags = {"Department"= "Finance"}
					hostname_label = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_hostname_label
					ipv6address_ipv6subnet_cidr_pair_details {

						#Optional
						ipv6address = var.instance_configuration_instance_details_launch_details_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6address
						ipv6subnet_cidr = var.instance_configuration_instance_details_launch_details_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
					}
					nsg_ids = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_nsg_ids
					private_ip = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_private_ip
					security_attributes = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_security_attributes
					skip_source_dest_check = var.instance_configuration_instance_details_options_launch_details_create_vnic_details_skip_source_dest_check
					subnet_id = oci_core_subnet.test_subnet.id
				}
				dedicated_vm_host_id = oci_core_dedicated_vm_host.test_dedicated_vm_host.id
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_options_launch_details_display_name
				extended_metadata = var.instance_configuration_instance_details_options_launch_details_extended_metadata
				fault_domain = var.instance_configuration_instance_details_options_launch_details_fault_domain
				freeform_tags = {"Department"= "Finance"}
				instance_options {

					#Optional
					are_legacy_imds_endpoints_disabled = var.instance_configuration_instance_details_options_launch_details_instance_options_are_legacy_imds_endpoints_disabled
				}
				ipxe_script = var.instance_configuration_instance_details_options_launch_details_ipxe_script
				is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_options_launch_details_is_pv_encryption_in_transit_enabled
				launch_mode = var.instance_configuration_instance_details_options_launch_details_launch_mode
				launch_options {

					#Optional
					boot_volume_type = var.instance_configuration_instance_details_options_launch_details_launch_options_boot_volume_type
					firmware = var.instance_configuration_instance_details_options_launch_details_launch_options_firmware
					is_consistent_volume_naming_enabled = var.instance_configuration_instance_details_options_launch_details_launch_options_is_consistent_volume_naming_enabled
					is_pv_encryption_in_transit_enabled = var.instance_configuration_instance_details_options_launch_details_launch_options_is_pv_encryption_in_transit_enabled
					network_type = var.instance_configuration_instance_details_options_launch_details_launch_options_network_type
					remote_data_volume_type = var.instance_configuration_instance_details_options_launch_details_launch_options_remote_data_volume_type
				}
				licensing_configs {
					#Required
					type = var.instance_configuration_instance_details_options_launch_details_licensing_configs_type

					#Optional
					license_type = var.instance_configuration_instance_details_options_launch_details_licensing_configs_license_type
				}
				metadata = var.instance_configuration_instance_details_options_launch_details_metadata
				platform_config {
					#Required
					type = var.instance_configuration_instance_details_options_launch_details_platform_config_type

					#Optional
					are_virtual_instructions_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_are_virtual_instructions_enabled
					is_access_control_service_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_access_control_service_enabled
					is_input_output_memory_management_unit_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_input_output_memory_management_unit_enabled
					is_measured_boot_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_measured_boot_enabled
					is_memory_encryption_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_memory_encryption_enabled
					is_secure_boot_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_secure_boot_enabled
					is_symmetric_multi_threading_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_symmetric_multi_threading_enabled
					is_trusted_platform_module_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_is_trusted_platform_module_enabled
					numa_nodes_per_socket = var.instance_configuration_instance_details_options_launch_details_platform_config_numa_nodes_per_socket
					percentage_of_cores_enabled = var.instance_configuration_instance_details_options_launch_details_platform_config_percentage_of_cores_enabled
				}
				preemptible_instance_config {

					#Optional
					preemption_action {
						#Required
						type = var.instance_configuration_instance_details_options_launch_details_preemptible_instance_config_preemption_action_type

						#Optional
						preserve_boot_volume = var.instance_configuration_instance_details_options_launch_details_preemptible_instance_config_preemption_action_preserve_boot_volume
					}
				}
				preferred_maintenance_action = var.instance_configuration_instance_details_options_launch_details_preferred_maintenance_action
				security_attributes = var.instance_configuration_instance_details_options_launch_details_security_attributes
				shape = var.instance_configuration_instance_details_options_launch_details_shape
				shape_config {

					#Optional
					baseline_ocpu_utilization = var.instance_configuration_instance_details_options_launch_details_shape_config_baseline_ocpu_utilization
					memory_in_gbs = var.instance_configuration_instance_details_options_launch_details_shape_config_memory_in_gbs
					nvmes = var.instance_configuration_instance_details_options_launch_details_shape_config_nvmes
					ocpus = var.instance_configuration_instance_details_options_launch_details_shape_config_ocpus
					vcpus = var.instance_configuration_instance_details_options_launch_details_shape_config_vcpus
				}
				source_details {
					#Required
					source_type = var.instance_configuration_instance_details_options_launch_details_source_details_source_type

					#Optional
					boot_volume_id = oci_core_boot_volume.test_boot_volume.id
					boot_volume_size_in_gbs = var.instance_configuration_instance_details_options_launch_details_source_details_boot_volume_size_in_gbs
					boot_volume_vpus_per_gb = var.instance_configuration_instance_details_options_launch_details_source_details_boot_volume_vpus_per_gb
					image_id = oci_core_image.test_image.id
					instance_source_image_filter_details {

						#Optional
						compartment_id = var.compartment_id
						defined_tags_filter = var.instance_configuration_instance_details_options_launch_details_source_details_instance_source_image_filter_details_defined_tags_filter
						operating_system = var.instance_configuration_instance_details_options_launch_details_source_details_instance_source_image_filter_details_operating_system
						operating_system_version = var.instance_configuration_instance_details_options_launch_details_source_details_instance_source_image_filter_details_operating_system_version
					}
				}
			}
			secondary_vnics {

				#Optional
				create_vnic_details {

					#Optional
					assign_ipv6ip = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_ipv6ip
					assign_private_dns_record = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_assign_private_dns_record
					assign_public_ip = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_assign_public_ip
					defined_tags = {"Operations.CostCenter"= "42"}
					display_name = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_display_name
					freeform_tags = {"Department"= "Finance"}
					hostname_label = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_hostname_label
					ipv6address_ipv6subnet_cidr_pair_details {

						#Optional
						ipv6address = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6address
						ipv6subnet_cidr = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_ipv6address_ipv6subnet_cidr_pair_details_ipv6subnet_cidr
					}
					nsg_ids = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_nsg_ids
					private_ip = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_private_ip
					security_attributes = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_security_attributes
					skip_source_dest_check = var.instance_configuration_instance_details_options_secondary_vnics_create_vnic_details_skip_source_dest_check
					subnet_id = oci_core_subnet.test_subnet.id
				}
				display_name = var.instance_configuration_instance_details_options_secondary_vnics_display_name
				nic_index = var.instance_configuration_instance_details_options_secondary_vnics_nic_index
			}
		}
		secondary_vnics {

			#Optional
			create_vnic_details {

				#Optional
				assign_private_dns_record = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_private_dns_record
				assign_public_ip = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_public_ip
				defined_tags = {"Operations.CostCenter"= "42"}
				display_name = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_display_name
				freeform_tags = {"Department"= "Finance"}
				hostname_label = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_hostname_label
				nsg_ids = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_nsg_ids
				private_ip = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_private_ip
				security_attributes = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_security_attributes
				skip_source_dest_check = var.instance_configuration_instance_details_secondary_vnics_create_vnic_details_skip_source_dest_check
				subnet_id = oci_core_subnet.test_subnet.id
			}
			display_name = var.instance_configuration_instance_details_secondary_vnics_display_name
			nic_index = var.instance_configuration_instance_details_secondary_vnics_nic_index
		}
	}
	instance_id = oci_core_instance.test_instance.id
	source = var.instance_configuration_source
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance configuration.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `instance_details` - (Required when source=NONE)
	* `block_volumes` - (Applicable when instance_type=compute) Block volume parameters.
		* `attach_details` - (Applicable when instance_type=compute) Volume attachmentDetails. Please see [AttachVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AttachVolumeDetails/)
			* `device` - (Applicable when instance_type=compute) The device name.
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
			* `is_pv_encryption_in_transit_enabled` - (Applicable when type=paravirtualized) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `is_read_only` - (Applicable when instance_type=compute) Whether the attachment should be created in read-only mode.
			* `is_shareable` - (Applicable when instance_type=compute) Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
			* `type` - (Required) The type of volume. The only supported values are "iscsi" and "paravirtualized"
			* `use_chap` - (Applicable when type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false. 
		* `create_details` - (Applicable when instance_type=compute) Creates a new block volume. Please see [CreateVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVolumeDetails/) 
			* `autotune_policies` - (Applicable when instance_type=compute) The list of autotune policies enabled for this volume.
				* `autotune_type` - (Required) This specifies the type of autotunes supported by OCI.
				* `max_vpus_per_gb` - (Required when autotune_type=PERFORMANCE_BASED) This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
			* `availability_domain` - (Applicable when instance_type=compute) The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - (Applicable when instance_type=compute) If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `block_volume_replicas` - (Applicable when instance_type=compute) The list of block volume replicas to be enabled for this volume in the specified destination availability domains. 
				* `availability_domain` - (Required when instance_type=compute) The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1` 
				* `display_name` - (Applicable when instance_type=compute) The display name of the block volume replica. You may optionally specify a *display name* for the block volume replica, otherwise a default is provided. 
			* `compartment_id` - (Applicable when instance_type=compute) The OCID of the compartment that contains the volume.
			* `defined_tags` - (Applicable when instance_type=compute) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - (Applicable when instance_type=compute) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `is_auto_tune_enabled` - (Applicable when instance_type=compute) Specifies whether the auto-tune performance is enabled for this boot volume. This field is deprecated. Use the `InstanceConfigurationDetachedVolumeAutotunePolicy` instead to enable the volume for detached autotune. 
			* `kms_key_id` - (Applicable when instance_type=compute) The OCID of the Vault service key to assign as the master encryption key for the volume. 
			* `size_in_gbs` - (Applicable when instance_type=compute) The size of the volume in GBs.
			* `source_details` - (Applicable when instance_type=compute) 
				* `id` - (Optional) The OCID of the volume backup.
				* `type` - (Required) The type can be one of these values: `volume`, `volumeBackup`
			* `vpus_per_gb` - (Applicable when instance_type=compute) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				Allowed values:
				* `0`: Represents Lower Cost option.
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option.
				* `30`-`120`: Represents the Ultra High Performance option.

				For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 
			* `xrc_kms_key_id` - (Applicable when instance_type=compute) The OCID of the Vault service key which is the master encryption key for the block volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
		* `volume_id` - (Applicable when instance_type=compute) The OCID of the volume.
	* `instance_type` - (Required) The type of instance details. Supported instanceType is compute
	* `launch_details` - (Applicable when instance_type=compute) Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.
    See [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) for more information.
        * `agent_config` - (Applicable when instance_type=compute) Configuration options for the Oracle Cloud Agent software running on the instance.
            * `are_all_plugins_disabled` - (Applicable when instance_type=compute) Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

				To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
			* `is_management_disabled` - (Applicable when instance_type=compute) Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

				These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

				The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
				* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
			* `is_monitoring_disabled` - (Applicable when instance_type=compute) Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

				These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

				The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
				* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
			* `plugins_config` - (Applicable when instance_type=compute) The configuration of plugins associated with this instance.
				* `desired_state` - (Required when instance_type=compute) Whether the plugin should be enabled or disabled.

					To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
				* `name` - (Required when instance_type=compute) The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `availability_config` - (Applicable when instance_type=compute) Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware. 
			* `is_live_migration_preferred` - (Optional) Whether to live migrate supported VM instances to a healthy physical VM host without disrupting running instances during infrastructure maintenance events. If null, Oracle chooses the best option for migrating the VM during infrastructure maintenance events.
			* `recovery_action` - (Applicable when instance_type=compute) The lifecycle state for an instance when it is recovered after infrastructure maintenance.
				* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
				* `STOP_INSTANCE` - The instance is recovered in the stopped state. 
		* `availability_domain` - (Applicable when instance_type=compute) The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - (Applicable when instance_type=compute) The OCID of the compute capacity reservation this instance is launched under.
		* `cluster_placement_group_id` - (Applicable when instance_type=compute) The OCID of the cluster placement group of the instance.
		* `compartment_id` - (Applicable when instance_type=compute) The OCID of the compartment containing the instance. Instances created from instance configurations are placed in the same compartment as the instance that was used to create the instance configuration. 
		* `create_vnic_details` - (Applicable when instance_type=compute) Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
			* `assign_ipv6ip` - (Applicable when instance_type=compute) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
			* `assign_private_dns_record` - (Applicable when instance_type=compute) Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `assign_ipv6ip` - (Optional) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you.
			* `assign_public_ip` - (Applicable when instance_type=compute) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - (Applicable when instance_type=compute) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - (Applicable when instance_type=compute) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - (Applicable when instance_type=compute) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure selects an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range.
				* `ipv6address` - (Optional) Optional. An available IPv6 address of your subnet from a valid IPv6 prefix on the subnet (otherwise the IP address is automatically assigned).
				* `ipv6subnet_cidr` - (Optional) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation.
			* `nsg_ids` - (Applicable when instance_type=compute) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - (Applicable when instance_type=compute) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `security_attributes` - (Applicable when instance_type=compute) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
			* `skip_source_dest_check` - (Applicable when instance_type=compute) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - (Applicable when instance_type=compute) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `dedicated_vm_host_id` - (Applicable when instance_type=compute) The OCID of the dedicated virtual machine host to place the instance on.

			Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools. 
		* `defined_tags` - (Applicable when instance_type=compute) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `extended_metadata` - (Applicable when instance_type=compute) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `fault_domain` - (Applicable when instance_type=compute) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			If you do not specify the fault domain, the system selects one for you.

			 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `FAULT-DOMAIN-1` 
		* `freeform_tags` - (Applicable when instance_type=compute) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `instance_options` - (Applicable when instance_type=compute) Optional mutable instance options. As a part of Instance Metadata Service Security Header, This allows user to disable the legacy imds endpoints.
			* `are_legacy_imds_endpoints_disabled` - (Applicable when instance_type=compute) Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
		* `ipxe_script` - (Applicable when instance_type=compute) This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `is_pv_encryption_in_transit_enabled` - (Applicable when instance_type=compute) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
		* `launch_mode` - (Applicable when instance_type=compute) Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
			* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
			* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
			* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
			* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
		* `launch_options` - (Applicable when instance_type=compute) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
			* `boot_volume_type` - (Applicable when instance_type=compute) Emulation type for the boot volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
			* `firmware` - (Applicable when instance_type=compute) Firmware used to boot VM. Select the option that matches your operating system.
				* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
				* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images. 
			* `is_consistent_volume_naming_enabled` - (Applicable when instance_type=compute) Whether to enable consistent volume naming feature. Defaults to false.
			* `is_pv_encryption_in_transit_enabled` - (Applicable when instance_type=compute) Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationLaunchInstanceDetails). 
			* `network_type` - (Applicable when instance_type=compute) Emulation type for the physical network interface card (NIC).
				* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
				* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
			* `remote_data_volume_type` - (Applicable when instance_type=compute) Emulation type for volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
		* `metadata` - (Applicable when instance_type=compute) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
			* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
			* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			**Providing Cloud-Init Metadata**

			You can use the following metadata key names to provide information to Cloud-Init:

			**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			**Metadata Example**

			"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `platform_config` - (Applicable when instance_type=compute) (Optional) (Updatable only for VM's) The platform configuration requested for the instance.

			If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `are_virtual_instructions_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes. 
			* `config_map` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Instance Platform Configuration Configuration Map for flexible setting input. 
			* `is_access_control_service_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM) Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through. 
			* `is_input_output_memory_management_unit_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) Whether the input-output memory management unit is enabled. 
			* `is_measured_boot_enabled` - (Applicable when instance_type=compute) Whether the Measured Boot feature is enabled on the instance. 
			* `is_memory_encryption_enabled` - (Applicable when instance_type=compute) Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`. 
			* `is_secure_boot_enabled` - (Applicable when instance_type=compute) Whether Secure Boot is enabled on the instance. 
			* `is_symmetric_multi_threading_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) (Updatable only for AMD_VM and INTEL_VM) Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.

				Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads. 
			* `is_trusted_platform_module_enabled` - (Applicable when instance_type=compute) Whether the Trusted Platform Module (TPM) is enabled on the instance. 
			* `numa_nodes_per_socket` - (Applicable when type=AMD_MILAN_BM | AMD_MILAN_BM_GPU | AMD_ROME_BM | AMD_ROME_BM_GPU | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The number of NUMA nodes per socket (NPS). 
			* `percentage_of_cores_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | GENERIC_BM | INTEL_ICELAKE_BM | INTEL_SKYLAKE_BM) The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage results in a fractional number of cores, the system rounds up the number of cores across processors and provisions an instance with a whole number of cores.

				If the applications that you run on the instance use a core-based licensing model and need fewer cores than the full size of the shape, you can disable cores to reduce your licensing costs. The instance itself is billed for the full shape, regardless of whether all cores are enabled. 
			* `type` - (Required) The type of platform being configured. 
		* `preemptible_instance_config` - (Applicable when instance_type=compute) Configuration options for preemptible instances. 
			* `preemption_action` - (Required when instance_type=compute) The action to run when the preemptible instance is interrupted for eviction. 
				* `preserve_boot_volume` - (Optional) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - (Required) The type of action to run when the instance is interrupted for eviction.
		* `preferred_maintenance_action` - (Applicable when instance_type=compute) The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
			* `LIVE_MIGRATE` - Run maintenance using a live migration.
			* `REBOOT` - Run maintenance using a reboot. 
		* `security_attributes` - (Applicable when instance_type=compute) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
		* `shape` - (Applicable when instance_type=compute) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
		* `shape_config` - (Applicable when instance_type=compute) The shape configuration requested for the instance.

			If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `baseline_ocpu_utilization` - (Applicable when instance_type=compute) The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

				The following values are supported:
				* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
				* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
				* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - (Applicable when instance_type=compute) The total amount of memory available to the instance, in gigabytes. 
			* `nvmes` - (Applicable when instance_type=compute) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. 
			* `ocpus` - (Applicable when instance_type=compute) The total number of OCPUs available to the instance. 
			* `vcpus` - (Applicable when instance_type=compute) The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2. 
		* `source_details` - (Applicable when instance_type=compute) 
			* `boot_volume_id` - (Applicable when source_type=bootVolume) The OCID of the boot volume used to boot the instance.
			* `boot_volume_size_in_gbs` - (Applicable when source_type=image) The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
			* `boot_volume_vpus_per_gb` - (Applicable when source_type=image) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				Allowed values:
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option.
				* `30`-`120`: Represents the Ultra High Performance option.

			* `image_id` - (Applicable when source_type=image) The OCID of the image used to boot the instance.
			* `kms_key_id` - (Applicable when source_type=image) The OCID of the Vault service key to assign as the master encryption key for the boot volume.          
			* `instance_source_image_filter_details` - (Applicable when source_type=image) These are the criteria for selecting an image. This is required if imageId is not specified. 
				* `compartment_id` - (Applicable when source_type=image) (Updatable) The OCID of the compartment containing images to search
				* `defined_tags_filter` - (Applicable when source_type=image) Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
				* `operating_system` - (Applicable when source_type=image) The image's operating system.  Example: `Oracle Linux` 
				* `operating_system_version` - (Applicable when source_type=image) The image's operating system version.  Example: `7.2` 
			* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `options` - (Applicable when instance_type=instance_options) Multiple Compute Instance Configuration instance details.
		* `block_volumes` - (Applicable when instance_type=instance_options) Block volume parameters.
			* `attach_details` - (Applicable when instance_type=instance_options) Volume attachmentDetails. Please see [AttachVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AttachVolumeDetails/)
				* `device` - (Applicable when instance_type=instance_options) The device name.
				* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
				* `is_pv_encryption_in_transit_enabled` - (Applicable when type=paravirtualized) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
				* `is_read_only` - (Applicable when instance_type=instance_options) Whether the attachment should be created in read-only mode.
				* `is_shareable` - (Applicable when instance_type=instance_options) Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
				* `type` - (Required) The type of volume. The only supported values are "iscsi" and "paravirtualized". 
				* `use_chap` - (Applicable when type=iscsi) Whether to use CHAP authentication for the volume attachment. Defaults to false. 
			* `create_details` - (Applicable when instance_type=instance_options) Creates a new block volume. Please see [CreateVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVolumeDetails/) 
				* `autotune_policies` - (Applicable when instance_type=instance_options) The list of autotune policies enabled for this volume.
					* `autotune_type` - (Required) This specifies the type of autotunes supported by OCI.
					* `max_vpus_per_gb` - (Required when autotune_type=PERFORMANCE_BASED) This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
				* `availability_domain` - (Applicable when instance_type=instance_options) The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
				* `backup_policy_id` - (Applicable when instance_type=instance_options) If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
				* `cluster_placement_group_id` - (Applicable when instance_type=instance_options) The clusterPlacementGroup Id of the volume for volume placement.
				* `compartment_id` - (Applicable when instance_type=instance_options) (Updatable) The OCID of the compartment that contains the volume.
				* `defined_tags` - (Applicable when instance_type=instance_options) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
				* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
				* `freeform_tags` - (Applicable when instance_type=instance_options) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
				* `kms_key_id` - (Applicable when instance_type=instance_options) The OCID of the Vault service key to assign as the master encryption key for the volume.
				* `size_in_gbs` - (Applicable when instance_type=instance_options) The size of the volume in GBs.
				* `source_details` - (Applicable when instance_type=instance_options)
					* `id` - (Optional) The OCID of the volume backup.
					* `type` - (Required) The type can be one of these values: `volume`, `volumeBackup`
				* `vpus_per_gb` - (Applicable when instance_type=instance_options) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				  Allowed values:
					* `0`: Represents Lower Cost option.
					* `10`: Represents Balanced option.
					* `20`: Represents Higher Performance option.
					* `30`-`120`: Represents the Ultra High Performance option.
					For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 
				* `xrc_kms_key_id` - (Applicable when instance_type=instance_options) The OCID of the Vault service key which is the master encryption key for the block volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
			* `volume_id` - (Applicable when instance_type=instance_options) The OCID of the volume.
		* `launch_details` - (Applicable when instance_type=instance_options) Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.

		  See [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) for more information.
			* `agent_config` - (Applicable when instance_type=instance_options) Configuration options for the Oracle Cloud Agent software running on the instance.
				* `are_all_plugins_disabled` - (Applicable when instance_type=instance_options) Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

				  To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
				* `is_management_disabled` - (Applicable when instance_type=instance_options) Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

				  These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

				  The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
					* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
					* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object.
				* `is_monitoring_disabled` - (Applicable when instance_type=instance_options) Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

				  These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

				  The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
					* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
					* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object.
				* `plugins_config` - (Applicable when instance_type=instance_options) The configuration of plugins associated with this instance.
					* `desired_state` - (Required when instance_type=instance_options) Whether the plugin should be enabled or disabled.

					  To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false.
					* `name` - (Required when instance_type=instance_options) The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
			* `availability_config` - (Applicable when instance_type=instance_options) Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware.
				* `recovery_action` - (Applicable when instance_type=instance_options) The lifecycle state for an instance when it is recovered after infrastructure maintenance.
					* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
					* `STOP_INSTANCE` - The instance is recovered in the stopped state.
			* `availability_domain` - (Applicable when instance_type=instance_options) The availability domain of the instance.  Example: `Uocm:PHX-AD-1`
			* `capacity_reservation_id` - (Applicable when instance_type=instance_options) The OCID of the compute capacity reservation this instance is launched under.
			* `compartment_id` - (Applicable when instance_type=instance_options) (Updatable) The OCID of the compartment containing the instance. Instances created from instance configurations are placed in the same compartment as the instance that was used to create the instance configuration. 
			* `create_vnic_details` - (Applicable when instance_type=instance_options) Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
		* `assign_ipv6ip` - (Optional) Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you.
		* `ipv6address_ipv6subnet_cidr_pair_details` - (Optional) A list of IPv6 prefix ranges from which the VNIC should be assigned an IPv6 address. You can provide only the prefix ranges and Oracle Cloud Infrastructure selects an available address from the range. You can optionally choose to leave the prefix range empty and instead provide the specific IPv6 address that should be used from within that range.
			* `ipv6address` - (Optional) Optional. An available IPv6 address of your subnet from a valid IPv6 prefix on the subnet (otherwise the IP address is automatically assigned).
			* `ipv6subnet_cidr` - (Optional) Optional. Used to disambiguate which subnet prefix should be used to create an IPv6 allocation.
				* `assign_private_dns_record` - (Applicable when instance_type=instance_options) Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information.
				* `assign_public_ip` - (Applicable when instance_type=instance_options) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `defined_tags` - (Applicable when instance_type=instance_options) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
				* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
				* `freeform_tags` - (Applicable when instance_type=instance_options) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
				* `hostname_label` - (Applicable when instance_type=instance_options) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `nsg_ids` - (Applicable when instance_type=instance_options) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
				* `private_ip` - (Applicable when instance_type=instance_options) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `security_attributes` - (Applicable when instance_type=instance_options) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
				* `skip_source_dest_check` - (Applicable when instance_type=instance_options) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `subnet_id` - (Applicable when instance_type=instance_options) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `dedicated_vm_host_id` - (Applicable when instance_type=instance_options) The OCID of the dedicated virtual machine host to place the instance on.

			  Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools.
			* `defined_tags` - (Applicable when instance_type=instance_options) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
			* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
			* `extended_metadata` - (Applicable when instance_type=instance_options) Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			  They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			  The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
			* `fault_domain` - (Applicable when instance_type=instance_options) A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			  If you do not specify the fault domain, the system selects one for you.

			  To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			  Example: `FAULT-DOMAIN-1`
			* `freeform_tags` - (Applicable when instance_type=instance_options) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
			* `instance_options` - (Applicable when instance_type=instance_options) Optional mutable instance options. As a part of Instance Metadata Service Security Header, This allows user to disable the legacy imds endpoints.
				* `are_legacy_imds_endpoints_disabled` - (Applicable when instance_type=instance_options) Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false.
			* `ipxe_script` - (Applicable when instance_type=instance_options) This is an advanced option.

			  When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			  If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			  The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			  For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			  For more information about iPXE, see http://ipxe.org.
			* `is_pv_encryption_in_transit_enabled` - (Applicable when instance_type=instance_options) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `launch_mode` - (Applicable when instance_type=instance_options) Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
				* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
				* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
				* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter.
			* `launch_options` - (Applicable when instance_type=instance_options) Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values.
				* `boot_volume_type` - (Applicable when instance_type=instance_options) Emulation type for the boot volume.
					* `ISCSI` - ISCSI attached block storage device.
					* `SCSI` - Emulated SCSI disk.
					* `IDE` - Emulated IDE disk.
					* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
					* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images.
				* `firmware` - (Applicable when instance_type=instance_options) Firmware used to boot VM. Select the option that matches your operating system.
					* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
					* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images.
				* `is_consistent_volume_naming_enabled` - (Applicable when instance_type=instance_options) Whether to enable consistent volume naming feature. Defaults to false.
				* `is_pv_encryption_in_transit_enabled` - (Applicable when instance_type=instance_options) Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationLaunchInstanceDetails).
				* `network_type` - (Applicable when instance_type=instance_options) Emulation type for the physical network interface card (NIC).
					* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
					* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
					* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
				* `remote_data_volume_type` - (Applicable when instance_type=instance_options) Emulation type for volume.
					* `ISCSI` - ISCSI attached block storage device.
					* `SCSI` - Emulated SCSI disk.
					* `IDE` - Emulated IDE disk.
					* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
					* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
			* `licensing_configs` - (Applicable when instance_type=instance_options) List of licensing configurations associated with target launch values.
				* `license_type` - (Optional) License Type for the OS license.
					* `OCI_PROVIDED` - Oracle Cloud Infrastructure provided license (e.g. metered $/OCPU-hour).
					* `BRING_YOUR_OWN_LICENSE` - Bring your own license. 
				* `type` - (Required) Operating System type of the Configuration.
			* `metadata` - (Applicable when instance_type=instance_options) Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			  A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
				* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
				* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			  **Providing Cloud-Init Metadata**

			  You can use the following metadata key names to provide information to Cloud-Init:

			  **"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			  **"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			  **Metadata Example**

			  "metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			  To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			  curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			  You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			  The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
			* `platform_config` - (Applicable when instance_type=instance_options) The platform configuration requested for the instance.

			  If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

			  Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.
				* `are_virtual_instructions_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | AMD_ROME_BM_GPU) Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes.
				* `is_access_control_service_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | AMD_ROME_BM_GPU) Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through.
				* `is_input_output_memory_management_unit_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | AMD_ROME_BM_GPU | INTEL_ICELAKE_BM) Whether the input-output memory management unit is enabled.
				* `is_measured_boot_enabled` - (Applicable when instance_type=instance_options) Whether the Measured Boot feature is enabled on the instance.
				* `is_memory_encryption_enabled` - (Applicable when instance_type=instance_options) Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
				* `is_secure_boot_enabled` - (Applicable when instance_type=instance_options) Whether Secure Boot is enabled on the instance.
				* `is_symmetric_multi_threading_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | AMD_ROME_BM_GPU | INTEL_ICELAKE_BM) (Updatable only for AMD_VM and INTEL_VM) Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.

				  Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads.
				* `is_trusted_platform_module_enabled` - (Applicable when instance_type=instance_options) Whether the Trusted Platform Module (TPM) is enabled on the instance.
				* `numa_nodes_per_socket` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | AMD_ROME_BM_GPU | INTEL_ICELAKE_BM) The number of NUMA nodes per socket (NPS).
				* `percentage_of_cores_enabled` - (Applicable when type=AMD_MILAN_BM | AMD_ROME_BM | INTEL_ICELAKE_BM) The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage results in a fractional number of cores, the system rounds up the number of cores across processors and provisions an instance with a whole number of cores.

				  If the applications that you run on the instance use a core-based licensing model and need fewer cores than the full size of the shape, you can disable cores to reduce your licensing costs. The instance itself is billed for the full shape, regardless of whether all cores are enabled.
				* `type` - (Required) The type of platform being configured.
			* `preemptible_instance_config` - (Applicable when instance_type=instance_options) Configuration options for preemptible instances.
				* `preemption_action` - (Required when instance_type=instance_options) The action to run when the preemptible instance is interrupted for eviction.
					* `preserve_boot_volume` - (Optional) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
					* `type` - (Required) The type of action to run when the instance is interrupted for eviction.
			* `preferred_maintenance_action` - (Applicable when instance_type=instance_options) The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
				* `LIVE_MIGRATE` - Run maintenance using a live migration.
				* `REBOOT` - Run maintenance using a reboot. 
			* `security_attributes` - (Applicable when instance_type=instance_options) [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
			* `shape` - (Applicable when instance_type=instance_options) The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			  You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes).
			* `shape_config` - (Applicable when instance_type=instance_options) The shape configuration requested for the instance.

			  If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

			  Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.
				* `baseline_ocpu_utilization` - (Applicable when instance_type=instance_options) The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

				  The following values are supported:
					* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
					* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
					* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance.
				* `memory_in_gbs` - (Applicable when instance_type=instance_options) The total amount of memory available to the instance, in gigabytes.
				* `nvmes` - (Applicable when instance_type=instance_options) The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
				* `ocpus` - (Applicable when instance_type=instance_options) The total number of OCPUs available to the instance.
				* `vcpus` - (Applicable when instance_type=instance_options) The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
			* `source_details` - (Applicable when instance_type=instance_options)
				* `boot_volume_id` - (Applicable when source_type=bootVolume) The OCID of the boot volume used to boot the instance.
				* `boot_volume_size_in_gbs` - (Applicable when source_type=image) The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB).
				* `boot_volume_vpus_per_gb` - (Applicable when source_type=image) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				  Allowed values:
					* `10`: Represents Balanced option.
					* `20`: Represents Higher Performance option.
					* `30`-`120`: Represents the Ultra High Performance option.

				  For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB.
				* `image_id` - (Applicable when source_type=image) The OCID of the image used to boot the instance.
				* `instance_source_image_filter_details` - (Applicable when source_type=image) These are the criteria for selecting an image. This is required if imageId is not specified.
					* `compartment_id` - (Applicable when source_type=image) (Updatable) The OCID of the compartment containing images to search
					* `defined_tags_filter` - (Applicable when source_type=image) Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
					* `operating_system` - (Applicable when source_type=image) The image's operating system.  Example: `Oracle Linux`
					* `operating_system_version` - (Applicable when source_type=image) The image's operating system version.  Example: `7.2`
				* `source_type` - (Required) The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID.
		* `secondary_vnics` - (Applicable when instance_type=instance_options) Secondary VNIC parameters.
			* `create_vnic_details` - (Applicable when instance_type=instance_options) Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
				* `assign_private_dns_record` - (Applicable when instance_type=instance_options) Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `assign_public_ip` - (Applicable when instance_type=instance_options) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `defined_tags` - (Applicable when instance_type=instance_options) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
				* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
				* `freeform_tags` - (Applicable when instance_type=instance_options) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
				* `hostname_label` - (Applicable when instance_type=instance_options) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `nsg_ids` - (Applicable when instance_type=instance_options) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
				* `private_ip` - (Applicable when instance_type=instance_options) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information.
				* `security_attributes` - (Applicable when instance_type=instance_options) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
				* `skip_source_dest_check` - (Applicable when instance_type=instance_options) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `subnet_id` - (Applicable when instance_type=instance_options) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `display_name` - (Applicable when instance_type=instance_options) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `nic_index` - (Applicable when instance_type=instance_options) Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
	* `secondary_vnics` - (Applicable when instance_type=compute) Secondary VNIC parameters.
		* `create_vnic_details` - (Applicable when instance_type=compute) Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
			* `assign_private_dns_record` - (Applicable when instance_type=compute) Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `assign_public_ip` - (Applicable when instance_type=compute) Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - (Applicable when instance_type=compute) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - (Applicable when instance_type=compute) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - (Applicable when instance_type=compute) The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `nsg_ids` - (Applicable when instance_type=compute) A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - (Applicable when instance_type=compute) A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information.
			* `security_attributes` - (Applicable when instance_type=compute) Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
			* `skip_source_dest_check` - (Applicable when instance_type=compute) Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - (Applicable when instance_type=compute) The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `display_name` - (Applicable when instance_type=compute) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `nic_index` - (Applicable when instance_type=compute) Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm).
* `instance_id` - (Required when source=INSTANCE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
* `source` - (Optional) The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.

  When you create an instance configuration using an existing instance as a template, the instance configuration does not include any information from the source instance's boot volume, such as installed applications, binaries, and files on the instance. It also does not include the contents of any block volumes that are attached to the instance.

  To create an instance configuration that includes the custom setup from an instance's boot volume, you must first create a custom image from the instance (see [CreateImage](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Image/CreateImage)). Then, use the custom image to launch a new instance (see [LaunchInstance](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstance)). Finally, create the instance configuration based on the instance that you created from the custom image.

  To include block volume contents with an instance configuration, first create a backup of the attached block volumes (see [CreateVolumeBackup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/VolumeBackup/CreateVolumeBackup)). Then, create the instance configuration by specifying the list of settings, using [InstanceConfigurationVolumeSourceFromVolumeBackupDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationVolumeSourceFromVolumeBackupDetails) to include the block volume backups in the list of settings.

  The following values are supported:
	* `NONE`: Creates an instance configuration using the list of settings that you specify.
	* `INSTANCE`: Creates an instance configuration using an existing instance as a template.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance configuration.
* `deferred_fields` - Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstanceConfiguration) operation.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration.
* `instance_details` -
	* `block_volumes` - Block volume parameters.
		* `attach_details` - Volume attachmentDetails. Please see [AttachVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AttachVolumeDetails/)
			* `device` - The device name.
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
			* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `is_read_only` - Whether the attachment should be created in read-only mode.
			* `is_shareable` - Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
			* `type` - The type of volume. The only supported values are "iscsi" and "paravirtualized". 
			* `use_chap` - Whether to use CHAP authentication for the volume attachment. Defaults to false. 
		* `create_details` - Creates a new block volume. Please see [CreateVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVolumeDetails/) 
			* `autotune_policies` - The list of autotune policies enabled for this volume.
				* `autotune_type` - This specifies the type of autotunes supported by OCI.
				* `max_vpus_per_gb` - This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
			* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
			* `backup_policy_id` - If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
			* `block_volume_replicas` - The list of block volume replicas to be enabled for this volume in the specified destination availability domains. 
				* `availability_domain` - The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1` 
				* `display_name` - The display name of the block volume replica. You may optionally specify a *display name* for the block volume replica, otherwise a default is provided. 
			* `compartment_id` - The OCID of the compartment that contains the volume.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `is_auto_tune_enabled` - Specifies whether the auto-tune performance is enabled for this boot volume. This field is deprecated. Use the `InstanceConfigurationDetachedVolumeAutotunePolicy` instead to enable the volume for detached autotune. 
			* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the volume. 
			* `size_in_gbs` - The size of the volume in GBs.
			* `source_details` - 
				* `id` - The OCID of the volume backup.
				* `type` - The type can be one of these values: `volume`, `volumeBackup`
			* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				Allowed values:
				* `0`: Represents Lower Cost option.
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option.
				* `30`-`120`: Represents the Ultra High Performance option.

				For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 
			* `xrc_kms_key_id` - The OCID of the Vault service key which is the master encryption key for the block volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
		* `volume_id` - The OCID of the volume.
	* `instance_type` - The type of instance details. Supported instanceType is compute
	* `launch_details` - Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.
		See [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) for more information. 
		* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
			* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

				To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
			* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

				These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

				The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
				* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
			* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

				These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

				The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
				* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
				* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object. 
			* `plugins_config` - The configuration of plugins associated with this instance.
				* `desired_state` - Whether the plugin should be enabled or disabled.

					To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false. 
				* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
		* `availability_config` - Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware. 
			* `is_live_migration_preferred` - Whether to live migrate supported VM instances to a healthy physical VM host without disrupting running instances during infrastructure maintenance events. If null, Oracle chooses the best option for migrating the VM during infrastructure maintenance events. 
			* `recovery_action` - The lifecycle state for an instance when it is recovered after infrastructure maintenance.
				* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
				* `STOP_INSTANCE` - The instance is recovered in the stopped state. 
		* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under.
		* `compartment_id` - The OCID of the compartment containing the instance. Instances created from instance configurations are placed in the same compartment as the instance that was used to create the instance configuration. 
		* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
			* `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information.
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `dedicated_vm_host_id` - The OCID of the dedicated virtual machine host to place the instance on.

			Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools. 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			If you do not specify the fault domain, the system selects one for you.

			 To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			Example: `FAULT-DOMAIN-1` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `instance_options` - Optional mutable instance options. As a part of Instance Metadata Service Security Header, This allows user to disable the legacy imds endpoints.
			* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
		* `ipxe_script` - This is an advanced option.

			When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			For more information about iPXE, see http://ipxe.org. 
		* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
		* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
			* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
			* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
			* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
			* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter. 
		* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values. 
			* `boot_volume_type` - Emulation type for the boot volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
			* `firmware` - Firmware used to boot VM. Select the option that matches your operating system.
				* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
				* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images. 
			* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
			* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationLaunchInstanceDetails). 
			* `network_type` - Emulation type for the physical network interface card (NIC).
				* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
				* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers. 
			* `remote_data_volume_type` - Emulation type for volume.
				* `ISCSI` - ISCSI attached block storage device.
				* `SCSI` - Emulated SCSI disk.
				* `IDE` - Emulated IDE disk.
				* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
				* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
		* `metadata` - Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
			* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
			* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			**Providing Cloud-Init Metadata**

			You can use the following metadata key names to provide information to Cloud-Init:

			**"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			**"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			**Metadata Example**

			"metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes. 
		* `platform_config` - The platform configuration requested for the instance.

			If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `are_virtual_instructions_enabled` - Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes. 
			* `config_map` - Instance Platform Configuration Configuration Map for flexible setting input. 
			* `is_access_control_service_enabled` - Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through. 
			* `is_input_output_memory_management_unit_enabled` - Whether the input-output memory management unit is enabled. 
			* `is_measured_boot_enabled` - Whether the Measured Boot feature is enabled on the instance. 
			* `is_memory_encryption_enabled` - Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`. 
			* `is_secure_boot_enabled` - Whether Secure Boot is enabled on the instance. 
			* `is_symmetric_multi_threading_enabled` - Whether symmetric multi-threading is enabled on the instance.

				Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads. 
			* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is enabled on the instance. 
			* `numa_nodes_per_socket` - The number of NUMA nodes per socket (NPS). 
			* `percentage_of_cores_enabled` - The percentage of cores enabled.
			* `type` - The type of platform being configured. 
		* `preemptible_instance_config` - Configuration options for preemptible instances. 
			* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction. 
				* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - The type of action to run when the instance is interrupted for eviction.
		* `preferred_maintenance_action` - The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
			* `LIVE_MIGRATE` - Run maintenance using a live migration.
			* `REBOOT` - Run maintenance using a reboot. 
		* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
		* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
		* `shape_config` - The shape configuration requested for the instance.

			If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

			Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned. 
			* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

				The following values are supported:
				* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
				* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
				* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance. 
			* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
			* `nvmes` - The number of NVMe drives to be used for storage. A single drive has 6.8 TB available. 
			* `ocpus` - The total number of OCPUs available to the instance. 
			* `vcpus` - The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2. 
		* `source_details` - 
			* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
			* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB). 
			* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				Allowed values:
				* `10`: Represents Balanced option.
				* `20`: Represents Higher Performance option.
				* `30`-`120`: Represents the Ultra High Performance option.
			* `image_id` - The OCID of the image used to boot the instance.
			* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the boot volume.
			* `instance_source_image_filter_details` - These are the criteria for selecting an image. This is required if imageId is not specified. 
				* `compartment_id` - The OCID of the compartment containing images to search
				* `defined_tags_filter` - Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
				* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
				* `operating_system_version` - The image's operating system version.  Example: `7.2` 
			* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the boot volume.
			* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
	* `options` - Multiple Compute Instance Configuration instance details.
		* `block_volumes` - Block volume parameters.
			* `attach_details` - Volume attachmentDetails. Please see [AttachVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/AttachVolumeDetails/)
				* `device` - The device name.
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
				* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
				* `is_read_only` - Whether the attachment should be created in read-only mode.
				* `is_shareable` - Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified. 
				* `type` - The type of volume. The only supported values are "iscsi" and "paravirtualized". 
				* `use_chap` - Whether to use CHAP authentication for the volume attachment. Defaults to false. 
			* `create_details` - Creates a new block volume. Please see [CreateVolumeDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVolumeDetails/) 
				* `autotune_policies` - The list of autotune policies enabled for this volume.
					* `autotune_type` - This specifies the type of autotunes supported by OCI.
					* `max_vpus_per_gb` - This will be the maximum VPUs/GB performance level that the volume will be auto-tuned temporarily based on performance monitoring. 
				* `availability_domain` - The availability domain of the volume.  Example: `Uocm:PHX-AD-1` 
				* `backup_policy_id` - If provided, specifies the ID of the volume backup policy to assign to the newly created volume. If omitted, no policy will be assigned. 
				* `cluster_placement_group_id` - The clusterPlacementGroup Id of the volume for volume placement.
				* `compartment_id` - The OCID of the compartment that contains the volume.
				* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
				* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
				* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the volume.
				* `size_in_gbs` - The size of the volume in GBs.
				* `source_details` -
					* `id` - The OCID of the volume backup.
					* `type` - The type can be one of these values: `volume`, `volumeBackup`
				* `vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				  Allowed values:
					* `0`: Represents Lower Cost option.
					* `10`: Represents Balanced option.
					* `20`: Represents Higher Performance option.
					* `30`-`120`: Represents the Ultra High Performance option.
					For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB. 
				* `xrc_kms_key_id` - The OCID of the Vault service key which is the master encryption key for the block volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
			* `volume_id` - The OCID of the volume.
		* `launch_details` - Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails` parameter to specify whether a boot volume or an image should be used to launch a new instance.

		  See [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/LaunchInstanceDetails) for more information.
			* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
				* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all the available plugins. This includes the management and monitoring plugins.

				  To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
				* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins. Default value is false (management plugins are enabled).

				  These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

				  The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
					* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
					* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object.
				* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins. Default value is false (monitoring plugins are enabled).

				  These are the monitoring plugins: Compute Instance Monitoring and Custom Logs Monitoring.

				  The monitoring plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
					* If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of the per-plugin configuration.
					* If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig` object.
				* `plugins_config` - The configuration of plugins associated with this instance.
					* `desired_state` - Whether the plugin should be enabled or disabled.

					  To enable the monitoring and management plugins, the `isMonitoringDisabled` and `isManagementDisabled` attributes must also be set to false.
					* `name` - The plugin name. To get a list of available plugins, use the [ListInstanceagentAvailablePlugins](https://docs.cloud.oracle.com/iaas/api/#/en/instanceagent/20180530/Plugin/ListInstanceagentAvailablePlugins) operation in the Oracle Cloud Agent API. For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
			* `availability_config` - Options for defining the availabiity of a VM instance after a maintenance event that impacts the underlying hardware.
				* `recovery_action` - The lifecycle state for an instance when it is recovered after infrastructure maintenance.
					* `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event. If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
					* `STOP_INSTANCE` - The instance is recovered in the stopped state.
			* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1`
			* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under.
			* `compartment_id` - The OCID of the compartment containing the instance. Instances created from instance configurations are placed in the same compartment as the instance that was used to create the instance configuration. 
			* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
				* `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
				* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
				* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
				* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
				* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `dedicated_vm_host_id` - The OCID of the dedicated virtual machine host to place the instance on.

			  Dedicated VM hosts can be used when launching individual instances from an instance configuration. They cannot be used to launch instance pools.
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
			* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

			  They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

			  The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
			* `fault_domain` - A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

			  If you do not specify the fault domain, the system selects one for you.

			  To get a list of fault domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

			  Example: `FAULT-DOMAIN-1`
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
			* `instance_options` - Optional mutable instance options. As a part of Instance Metadata Service Security Header, This allows user to disable the legacy imds endpoints.
				* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false.
			* `ipxe_script` - This is an advanced option.

			  When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

			  If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots; however, you should be aware that the same iPXE script will run every time an instance boots; not only after the initial LaunchInstance call.

			  The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, you should use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

			  For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

			  For more information about iPXE, see http://ipxe.org.
			* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
			* `launch_mode` - Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
				* `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
				* `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
				* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
				* `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter.
			* `launch_options` - Options for tuning the compatibility and performance of VM shapes. The values that you specify override any default values.
				* `boot_volume_type` - Emulation type for the boot volume.
					* `ISCSI` - ISCSI attached block storage device.
					* `SCSI` - Emulated SCSI disk.
					* `IDE` - Emulated IDE disk.
					* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
					* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images.
				* `firmware` - Firmware used to boot VM. Select the option that matches your operating system.
					* `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating systems that boot using MBR style bootloaders.
					* `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the default for platform images.
				* `is_consistent_volume_naming_enabled` - Whether to enable consistent volume naming feature. Defaults to false.
				* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [InstanceConfigurationLaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/InstanceConfigurationLaunchInstanceDetails).
				* `network_type` - Emulation type for the physical network interface card (NIC).
					* `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
					* `VFIO` - Direct attached Virtual Function network controller. This is the networking type when you launch an instance using hardware-assisted (SR-IOV) networking.
					* `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
				* `remote_data_volume_type` - Emulation type for volume.
					* `ISCSI` - ISCSI attached block storage device.
					* `SCSI` - Emulated SCSI disk.
					* `IDE` - Emulated IDE disk.
					* `VFIO` - Direct attached Virtual Function storage. This is the default option for local data volumes on platform images.
					* `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block storage volumes on platform images. 
			* `licensing_configs` - List of licensing configurations associated with target launch values.
				* `license_type` - License Type for the OS license.
					* `OCI_PROVIDED` - Oracle Cloud Infrastructure provided license (e.g. metered $/OCPU-hour).
					* `BRING_YOUR_OWN_LICENSE` - Bring your own license. 
				* `type` - Operating System type of the Configuration.
			* `metadata` - Custom metadata key/value pairs that you provide, such as the SSH public key required to connect to the instance.

			  A metadata service runs on every launched instance. The service is an HTTP endpoint listening on 169.254.169.254. You can use the service to:
				* Provide information to [Cloud-Init](https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
				* Get information about the instance, including the custom metadata that you provide when you launch the instance.

			  **Providing Cloud-Init Metadata**

			  You can use the following metadata key names to provide information to Cloud-Init:

			  **"ssh_authorized_keys"** - Provide one or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on the instance. Use a newline character to separate multiple keys. The SSH keys must be in the format necessary for the `authorized_keys` file, as shown in the example below.

			  **"user_data"** - Provide your own base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For information about how to take advantage of user data, see the [Cloud-Init Documentation](http://cloudinit.readthedocs.org/en/latest/topics/format.html).

			  **Metadata Example**

			  "metadata" : { "quake_bot_level" : "Severe", "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227", "user_data" : "<your_public_SSH_key>==" } **Getting Metadata on the Instance**

			  To get information about your instance, connect to the instance using SSH and issue any of the following GET requests:

			  curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/ curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>

			  You'll get back a response that includes all the instance information; only the metadata information; or the metadata information for the specified key name, respectively.

			  The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
			* `platform_config` - The platform configuration requested for the instance.

			  If you provide the parameter, the instance is created with the platform configuration that you specify. For any values that you omit, the instance uses the default configuration values for the `shape` that you specify. If you don't provide the parameter, the default values for the `shape` are used.

			  Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.
				* `are_virtual_instructions_enabled` - Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes.
				* `is_access_control_service_enabled` - Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device pass-through.
				* `is_input_output_memory_management_unit_enabled` - Whether the input-output memory management unit is enabled.
				* `is_measured_boot_enabled` - Whether the Measured Boot feature is enabled on the instance.
				* `is_memory_encryption_enabled` - Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
				* `is_secure_boot_enabled` - Whether Secure Boot is enabled on the instance.
				* `is_symmetric_multi_threading_enabled` - Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also called simultaneous multithreading (SMT) or Intel Hyper-Threading.

				  Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple independent threads of execution, to better use the resources and increase the efficiency of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which can provide higher or more predictable performance for some workloads.
				* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is enabled on the instance.
				* `numa_nodes_per_socket` - The number of NUMA nodes per socket (NPS).
				* `percentage_of_cores_enabled` - The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage results in a fractional number of cores, the system rounds up the number of cores across processors and provisions an instance with a whole number of cores.

				  If the applications that you run on the instance use a core-based licensing model and need fewer cores than the full size of the shape, you can disable cores to reduce your licensing costs. The instance itself is billed for the full shape, regardless of whether all cores are enabled.
				* `type` - The type of platform being configured.
			* `preemptible_instance_config` - Configuration options for preemptible instances.
				* `preemption_action` - The action to run when the preemptible instance is interrupted for eviction.
					* `preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
					* `type` - The type of action to run when the instance is interrupted for eviction.
			* `preferred_maintenance_action` - The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
				* `LIVE_MIGRATE` - Run maintenance using a live migration.
				* `REBOOT` - Run maintenance using a reboot.
			* `security_attributes` - Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
			* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

			  You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes).
			* `shape_config` - The shape configuration requested for the instance.

			  If the parameter is provided, the instance is created with the resources that you specify. If some properties are missing or the entire parameter is not provided, the instance is created with the default configuration values for the `shape` that you specify.

			  Each shape only supports certain configurable values. If the values that you provide are not valid for the specified `shape`, an error is returned.
				* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

				  The following values are supported:
					* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
					* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
					* `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance.
				* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes.
				* `nvmes` - The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
				* `ocpus` - The total number of OCPUs available to the instance.
				* `vcpus` - The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
			* `source_details` -
				* `boot_volume_id` - The OCID of the boot volume used to boot the instance.
				* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. The minimum value is 50 GB and the maximum value is 32,768 GB (32 TB).
				* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

				  Allowed values:
					* `10`: Represents Balanced option.
					* `20`: Represents Higher Performance option.
					* `30`-`120`: Represents the Ultra High Performance option.

				  For performance autotune enabled volumes, it would be the Default(Minimum) VPUs/GB.
				* `image_id` - The OCID of the image used to boot the instance.
				* `instance_source_image_filter_details` - These are the criteria for selecting an image. This is required if imageId is not specified.
					* `compartment_id` - The OCID of the compartment containing images to search
					* `defined_tags_filter` - Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
					* `operating_system` - The image's operating system.  Example: `Oracle Linux`
					* `operating_system_version` - The image's operating system version.  Example: `7.2`
				* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID.
		* `secondary_vnics` - Secondary VNIC parameters.
			* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
				* `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
				* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
				* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
				* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
				* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
				* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `nic_index` - Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
	* `secondary_vnics` - Secondary VNIC parameters.
		* `create_vnic_details` - Contains the properties of the VNIC for an instance configuration. See [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) and [Instance Configurations](https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm#config) for more information. 
			* `assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet. Default: False. When provided you may optionally provide an IPv6 prefix (`ipv6SubnetCidr`) of your choice to assign the IPv6 address from. If `ipv6SubnetCidr` is not provided then an IPv6 prefix is chosen for you. 
			* `assign_private_dns_record` - Whether the VNIC should be assigned a private DNS record. Defaults to true. See the `assignPrivateDnsRecord` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/) for more information.
			* `assign_public_ip` - Whether the VNIC should be assigned a public IP address. See the `assignPublicIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
			* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
			* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
			* `hostname_label` - The hostname for the VNIC's primary private IP. See the `hostnameLabel` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
			* `private_ip` - A private IP address of your choice to assign to the VNIC. See the `privateIp` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `security_attributes` - [Security attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels for a resource that can be referenced in a [Zero Trust Packet Routing](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm) (ZPR) policy to control access to ZPR-supported resources.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
			* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC. See the `skipSourceDestCheck` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
			* `subnet_id` - The OCID of the subnet to create the VNIC in. See the `subnetId` attribute of [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/) for more information. 
		* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
		* `nic_index` - Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `time_created` - The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Instance Configuration
* `update` - (Defaults to 20 minutes), when updating the Instance Configuration
* `delete` - (Defaults to 20 minutes), when destroying the Instance Configuration


## Import

InstanceConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_configuration.test_instance_configuration "id"
```
