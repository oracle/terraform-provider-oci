---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_asset"
sidebar_current: "docs-oci-resource-cloud_bridge-asset"
description: |-
  Provides the Asset resource in Oracle Cloud Infrastructure Cloud Bridge service
---

# oci_cloud_bridge_asset
This resource provides the Asset resource in Oracle Cloud Infrastructure Cloud Bridge service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/OCB/latest/Asset

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/cloudBridge

Creates an asset.

## Example Usage

```hcl
resource "oci_cloud_bridge_asset" "test_asset" {
	#Required
	asset_type = var.asset_asset_type
	compartment_id = var.compartment_id
	external_asset_key = var.asset_external_asset_key
	inventory_id = oci_cloud_bridge_inventory.test_inventory.id
	source_key = var.asset_source_key

	#Optional
	asset_class_name = var.asset_asset_class_name
	asset_class_version = var.asset_asset_class_version
	asset_details = var.asset_asset_details
	asset_source_ids = var.asset_asset_source_ids
	attached_ebs_volumes_cost {

		#Optional
		amount = var.asset_attached_ebs_volumes_cost_amount
		currency_code = var.asset_attached_ebs_volumes_cost_currency_code
	}
	aws_ebs {

		#Optional
		attachments {

			#Optional
			device = var.asset_aws_ebs_attachments_device
			instance_key = var.asset_aws_ebs_attachments_instance_key
			is_delete_on_termination = var.asset_aws_ebs_attachments_is_delete_on_termination
			status = var.asset_aws_ebs_attachments_status
			volume_key = var.asset_aws_ebs_attachments_volume_key
		}
		availability_zone = var.asset_aws_ebs_availability_zone
		iops = var.asset_aws_ebs_iops
		is_encrypted = var.asset_aws_ebs_is_encrypted
		is_multi_attach_enabled = var.asset_aws_ebs_is_multi_attach_enabled
		size_in_gi_bs = var.asset_aws_ebs_size_in_gi_bs
		status = var.asset_aws_ebs_status
		tags {

			#Optional
			key = var.asset_aws_ebs_tags_key
			value = var.asset_aws_ebs_tags_value
		}
		throughput = var.asset_aws_ebs_throughput
		volume_key = var.asset_aws_ebs_volume_key
		volume_type = var.asset_aws_ebs_volume_type
	}
	aws_ec2 {

		#Optional
		architecture = var.asset_aws_ec2_architecture
		are_elastic_inference_accelerators_present = var.asset_aws_ec2_are_elastic_inference_accelerators_present
		boot_mode = var.asset_aws_ec2_boot_mode
		capacity_reservation_key = var.asset_aws_ec2_capacity_reservation_key
		image_key = var.asset_aws_ec2_image_key
		instance_key = var.asset_aws_ec2_instance_key
		instance_lifecycle = var.asset_aws_ec2_instance_lifecycle
		instance_type = var.asset_aws_ec2_instance_type
		ip_address = var.asset_aws_ec2_ip_address
		ipv6address = var.asset_aws_ec2_ipv6address
		is_enclave_options = var.asset_aws_ec2_is_enclave_options
		is_hibernation_options = var.asset_aws_ec2_is_hibernation_options
		is_source_dest_check = var.asset_aws_ec2_is_source_dest_check
		is_spot_instance = var.asset_aws_ec2_is_spot_instance
		kernel_key = var.asset_aws_ec2_kernel_key
		licenses = var.asset_aws_ec2_licenses
		maintenance_options = var.asset_aws_ec2_maintenance_options
		monitoring = var.asset_aws_ec2_monitoring
		network_interfaces {

			#Optional
			association {

				#Optional
				carrier_ip = var.asset_aws_ec2_network_interfaces_association_carrier_ip
				customer_owned_ip = var.asset_aws_ec2_network_interfaces_association_customer_owned_ip
				ip_owner_key = var.asset_aws_ec2_network_interfaces_association_ip_owner_key
				public_dns_name = var.asset_aws_ec2_network_interfaces_association_public_dns_name
				public_ip = var.asset_aws_ec2_network_interfaces_association_public_ip
			}
			attachment {

				#Optional
				attachment_key = var.asset_aws_ec2_network_interfaces_attachment_attachment_key
				device_index = var.asset_aws_ec2_network_interfaces_attachment_device_index
				is_delete_on_termination = var.asset_aws_ec2_network_interfaces_attachment_is_delete_on_termination
				network_card_index = var.asset_aws_ec2_network_interfaces_attachment_network_card_index
				status = var.asset_aws_ec2_network_interfaces_attachment_status
				time_attach = var.asset_aws_ec2_network_interfaces_attachment_time_attach
			}
			description = var.asset_aws_ec2_network_interfaces_description
			interface_type = var.asset_aws_ec2_network_interfaces_interface_type
			ipv4prefixes = var.asset_aws_ec2_network_interfaces_ipv4prefixes
			ipv6addresses = var.asset_aws_ec2_network_interfaces_ipv6addresses
			ipv6prefixes = var.asset_aws_ec2_network_interfaces_ipv6prefixes
			is_source_dest_check = var.asset_aws_ec2_network_interfaces_is_source_dest_check
			mac_address = var.asset_aws_ec2_network_interfaces_mac_address
			network_interface_key = var.asset_aws_ec2_network_interfaces_network_interface_key
			owner_key = var.asset_aws_ec2_network_interfaces_owner_key
			private_ip_addresses {

				#Optional
				association {

					#Optional
					carrier_ip = var.asset_aws_ec2_network_interfaces_private_ip_addresses_association_carrier_ip
					customer_owned_ip = var.asset_aws_ec2_network_interfaces_private_ip_addresses_association_customer_owned_ip
					ip_owner_key = var.asset_aws_ec2_network_interfaces_private_ip_addresses_association_ip_owner_key
					public_dns_name = var.asset_aws_ec2_network_interfaces_private_ip_addresses_association_public_dns_name
					public_ip = var.asset_aws_ec2_network_interfaces_private_ip_addresses_association_public_ip
				}
				is_primary = var.asset_aws_ec2_network_interfaces_private_ip_addresses_is_primary
				private_dns_name = var.asset_aws_ec2_network_interfaces_private_ip_addresses_private_dns_name
				private_ip_address = var.asset_aws_ec2_network_interfaces_private_ip_addresses_private_ip_address
			}
			security_groups {

				#Optional
				group_key = var.asset_aws_ec2_network_interfaces_security_groups_group_key
				group_name = oci_identity_group.test_group.name
			}
			status = var.asset_aws_ec2_network_interfaces_status
			subnet_key = var.asset_aws_ec2_network_interfaces_subnet_key
		}
		placement {

			#Optional
			affinity = var.asset_aws_ec2_placement_affinity
			availability_zone = var.asset_aws_ec2_placement_availability_zone
			group_name = oci_identity_group.test_group.name
			host_key = var.asset_aws_ec2_placement_host_key
			host_resource_group_arn = var.asset_aws_ec2_placement_host_resource_group_arn
			partition_number = var.asset_aws_ec2_placement_partition_number
			spread_domain = var.asset_aws_ec2_placement_spread_domain
			tenancy = var.asset_aws_ec2_placement_tenancy
		}
		private_dns_name = var.asset_aws_ec2_private_dns_name
		private_ip_address = var.asset_aws_ec2_private_ip_address
		root_device_name = var.asset_aws_ec2_root_device_name
		root_device_type = var.asset_aws_ec2_root_device_type
		security_groups {

			#Optional
			group_key = var.asset_aws_ec2_security_groups_group_key
			group_name = oci_identity_group.test_group.name
		}
		sriov_net_support = var.asset_aws_ec2_sriov_net_support
		state {

			#Optional
			code = var.asset_aws_ec2_state_code
			name = var.asset_aws_ec2_state_name
		}
		subnet_key = var.asset_aws_ec2_subnet_key
		tags {

			#Optional
			key = var.asset_aws_ec2_tags_key
			value = var.asset_aws_ec2_tags_value
		}
		time_launch = var.asset_aws_ec2_time_launch
		tpm_support = var.asset_aws_ec2_tpm_support
		virtualization_type = var.asset_aws_ec2_virtualization_type
		vpc_key = var.asset_aws_ec2_vpc_key
	}
	aws_ec2cost {

		#Optional
		amount = var.asset_aws_ec2cost_amount
		currency_code = var.asset_aws_ec2cost_currency_code
	}
	compute {

		#Optional
		connected_networks = var.asset_compute_connected_networks
		cores_count = var.asset_compute_cores_count
		cpu_model = var.asset_compute_cpu_model
		description = var.asset_compute_description
		disks {

			#Optional
			boot_order = var.asset_compute_disks_boot_order
			is_cbt_enabled = var.asset_compute_disks_is_cbt_enabled
			location = var.asset_compute_disks_location
			name = var.asset_compute_disks_name
			persistent_mode = var.asset_compute_disks_persistent_mode
			size_in_mbs = var.asset_compute_disks_size_in_mbs
			uuid = var.asset_compute_disks_uuid
			uuid_lun = var.asset_compute_disks_uuid_lun
		}
		disks_count = var.asset_compute_disks_count
		dns_name = var.asset_compute_dns_name
		firmware = var.asset_compute_firmware
		gpu_devices {

			#Optional
			cores_count = var.asset_compute_gpu_devices_cores_count
			description = var.asset_compute_gpu_devices_description
			manufacturer = var.asset_compute_gpu_devices_manufacturer
			memory_in_mbs = var.asset_compute_gpu_devices_memory_in_mbs
			name = var.asset_compute_gpu_devices_name
		}
		gpu_devices_count = var.asset_compute_gpu_devices_count
		guest_state = var.asset_compute_guest_state
		hardware_version = var.asset_compute_hardware_version
		host_name = var.asset_compute_host_name
		is_pmem_enabled = var.asset_compute_is_pmem_enabled
		is_tpm_enabled = var.asset_compute_is_tpm_enabled
		latency_sensitivity = var.asset_compute_latency_sensitivity
		memory_in_mbs = var.asset_compute_memory_in_mbs
		nics {

			#Optional
			ip_addresses = var.asset_compute_nics_ip_addresses
			label = var.asset_compute_nics_label
			mac_address = var.asset_compute_nics_mac_address
			mac_address_type = var.asset_compute_nics_mac_address_type
			network_name = var.asset_compute_nics_network_name
			switch_name = var.asset_compute_nics_switch_name
		}
		nics_count = var.asset_compute_nics_count
		nvdimm_controller {

			#Optional
			bus_number = var.asset_compute_nvdimm_controller_bus_number
			label = var.asset_compute_nvdimm_controller_label
		}
		nvdimms {

			#Optional
			controller_key = var.asset_compute_nvdimms_controller_key
			label = var.asset_compute_nvdimms_label
			unit_number = var.asset_compute_nvdimms_unit_number
		}
		operating_system = var.asset_compute_operating_system
		operating_system_version = var.asset_compute_operating_system_version
		pmem_in_mbs = var.asset_compute_pmem_in_mbs
		power_state = var.asset_compute_power_state
		primary_ip = var.asset_compute_primary_ip
		scsi_controller {

			#Optional
			label = var.asset_compute_scsi_controller_label
			shared_bus = var.asset_compute_scsi_controller_shared_bus
			unit_number = var.asset_compute_scsi_controller_unit_number
		}
		storage_provisioned_in_mbs = var.asset_compute_storage_provisioned_in_mbs
		threads_per_core_count = var.asset_compute_threads_per_core_count
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.asset_display_name
	environment_type = var.asset_environment_type
	freeform_tags = {"Department"= "Finance"}
	vm {

		#Optional
		hypervisor_host = var.asset_vm_hypervisor_host
		hypervisor_vendor = var.asset_vm_hypervisor_vendor
		hypervisor_version = var.asset_vm_hypervisor_version
	}
	vmware_vcenter {

		#Optional
		data_center = var.asset_vmware_vcenter_data_center
		vcenter_key = var.asset_vmware_vcenter_vcenter_key
		vcenter_version = var.asset_vmware_vcenter_vcenter_version
	}
	vmware_vm {

		#Optional
		cluster = var.asset_vmware_vm_cluster
		customer_fields = var.asset_vmware_vm_customer_fields
		customer_tags {

			#Optional
			description = var.asset_vmware_vm_customer_tags_description
			name = var.asset_vmware_vm_customer_tags_name
		}
		fault_tolerance_bandwidth = var.asset_vmware_vm_fault_tolerance_bandwidth
		fault_tolerance_secondary_latency = var.asset_vmware_vm_fault_tolerance_secondary_latency
		fault_tolerance_state = var.asset_vmware_vm_fault_tolerance_state
		instance_uuid = var.asset_vmware_vm_instance_uuid
		is_disks_cbt_enabled = var.asset_vmware_vm_is_disks_cbt_enabled
		is_disks_uuid_enabled = var.asset_vmware_vm_is_disks_uuid_enabled
		path = var.asset_vmware_vm_path
		vmware_tools_status = var.asset_vmware_vm_vmware_tools_status
	}
}
```

## Argument Reference

The following arguments are supported:

* `asset_class_name` - (Required when asset_type=INVENTORY_ASSET) (Updatable) The class name of the asset.
* `asset_class_version` - (Required when asset_type=INVENTORY_ASSET) (Updatable) The version of the asset class.
* `asset_details` - (Required when asset_type=INVENTORY_ASSET) (Updatable) The details of the asset.
* `asset_source_ids` - (Optional) (Updatable) List of asset source OCID.
* `asset_type` - (Required) (Updatable) The type of asset.
* `attached_ebs_volumes_cost` - (Applicable when asset_type=AWS_EC2) (Updatable) Cost information for monthly maintenance.
	* `amount` - (Required when asset_type=AWS_EC2) (Updatable) Monthly costs for maintenance of this asset.
	* `currency_code` - (Required when asset_type=AWS_EC2) (Updatable) Currency code
* `aws_ebs` - (Required when asset_type=AWS_EBS) (Updatable) AWS EBS volume related properties.
	* `attachments` - (Applicable when asset_type=AWS_EBS) (Updatable) Information about the volume attachments.
		* `device` - (Applicable when asset_type=AWS_EBS) (Updatable) The device name.
		* `instance_key` - (Applicable when asset_type=AWS_EBS) (Updatable) The ID of the instance.
		* `is_delete_on_termination` - (Applicable when asset_type=AWS_EBS) (Updatable) Indicates whether the EBS volume is deleted on instance termination.
		* `status` - (Applicable when asset_type=AWS_EBS) (Updatable) The attachment state of the volume.
		* `volume_key` - (Applicable when asset_type=AWS_EBS) (Updatable) The ID of the volume.
	* `availability_zone` - (Applicable when asset_type=AWS_EBS) (Updatable) The Availability Zone for the volume.
	* `iops` - (Applicable when asset_type=AWS_EBS) (Updatable) The number of I/O operations per second.
	* `is_encrypted` - (Required when asset_type=AWS_EBS) (Updatable) Indicates whether the volume is encrypted.
	* `is_multi_attach_enabled` - (Required when asset_type=AWS_EBS) (Updatable) Indicates whether Amazon EBS Multi-Attach is enabled.
	* `size_in_gi_bs` - (Required when asset_type=AWS_EBS) (Updatable) The size of the volume, in GiBs.
	* `status` - (Applicable when asset_type=AWS_EBS) (Updatable) The volume state.
	* `tags` - (Applicable when asset_type=AWS_EBS) (Updatable) Any tags assigned to the volume.
		* `key` - (Applicable when asset_type=AWS_EBS) (Updatable) The key of the tag.
		* `value` - (Applicable when asset_type=AWS_EBS) (Updatable) The value of the tag.
	* `throughput` - (Applicable when asset_type=AWS_EBS) (Updatable) The throughput that the volume supports, in MiB/s.
	* `volume_key` - (Required when asset_type=AWS_EBS) (Updatable) The ID of the volume.
	* `volume_type` - (Required when asset_type=AWS_EBS) (Updatable) The volume type.
* `aws_ec2` - (Required when asset_type=AWS_EC2) (Updatable) AWS virtual machine related properties.
	* `architecture` - (Required when asset_type=AWS_EC2) (Updatable) The architecture of the image.
	* `are_elastic_inference_accelerators_present` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates if the elastic inference accelerators attached to an instance
	* `boot_mode` - (Applicable when asset_type=AWS_EC2) (Updatable) The boot mode of the instance.
	* `capacity_reservation_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the Capacity Reservation.
	* `image_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the AMI used to launch the instance.
	* `instance_key` - (Required when asset_type=AWS_EC2) (Updatable) The ID of the instance.
	* `instance_lifecycle` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether this is a Spot Instance or a Scheduled Instance.
	* `instance_type` - (Required when asset_type=AWS_EC2) (Updatable) The instance type.
	* `ip_address` - (Applicable when asset_type=AWS_EC2) (Updatable) The public IPv4 address, or the Carrier IP address assigned to the instance.
	* `ipv6address` - (Applicable when asset_type=AWS_EC2) (Updatable) The IPv6 address assigned to the instance.
	* `is_enclave_options` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether the instance is enabled for AWS Nitro Enclaves.
	* `is_hibernation_options` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether the instance is enabled for hibernation.
	* `is_source_dest_check` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether source/destination checking is enabled.
	* `is_spot_instance` - (Applicable when asset_type=AWS_EC2) (Updatable) If the request is a Spot Instance request, this value will be true.
	* `kernel_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The kernel associated with this instance, if applicable.
	* `licenses` - (Applicable when asset_type=AWS_EC2) (Updatable) The license configurations for the instance.
	* `maintenance_options` - (Applicable when asset_type=AWS_EC2) (Updatable) Provides information on the recovery and maintenance options of your instance.
	* `monitoring` - (Applicable when asset_type=AWS_EC2) (Updatable) The monitoring for the instance.
	* `network_interfaces` - (Applicable when asset_type=AWS_EC2) (Updatable) The network interfaces for the instance.
		* `association` - (Applicable when asset_type=AWS_EC2) (Updatable) Describes association information for an Elastic IP address (IPv4).
			* `carrier_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The carrier IP address associated with the network interface.
			* `customer_owned_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The customer-owned IP address associated with the network interface.
			* `ip_owner_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the owner of the Elastic IP address.
			* `public_dns_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The public DNS name.
			* `public_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The public IP address or Elastic IP address bound to the network interface.
		* `attachment` - (Applicable when asset_type=AWS_EC2) (Updatable) Describes a network interface attachment.
			* `attachment_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the network interface attachment.
			* `device_index` - (Applicable when asset_type=AWS_EC2) (Updatable) The index of the device on the instance for the network interface attachment.
			* `is_delete_on_termination` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether the network interface is deleted when the instance is terminated.
			* `network_card_index` - (Applicable when asset_type=AWS_EC2) (Updatable) The index of the network card.
			* `status` - (Applicable when asset_type=AWS_EC2) (Updatable) The attachment state.
			* `time_attach` - (Applicable when asset_type=AWS_EC2) (Updatable) The timestamp when the attachment initiated.
		* `description` - (Applicable when asset_type=AWS_EC2) (Updatable) The description.
		* `interface_type` - (Applicable when asset_type=AWS_EC2) (Updatable) The type of network interface.
		* `ipv4prefixes` - (Applicable when asset_type=AWS_EC2) (Updatable) The IPv4 delegated prefixes that are assigned to the network interface.
		* `ipv6addresses` - (Applicable when asset_type=AWS_EC2) (Updatable) The IPv6 addresses associated with the network interface.
		* `ipv6prefixes` - (Applicable when asset_type=AWS_EC2) (Updatable) The IPv6 delegated prefixes that are assigned to the network interface.
		* `is_source_dest_check` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether source/destination checking is enabled.
		* `mac_address` - (Applicable when asset_type=AWS_EC2) (Updatable) The MAC address.
		* `network_interface_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the network interface.
		* `owner_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the AWS account that created the network interface.
		* `private_ip_addresses` - (Applicable when asset_type=AWS_EC2) (Updatable) The private IPv4 addresses associated with the network interface.
			* `association` - (Applicable when asset_type=AWS_EC2) (Updatable) Describes association information for an Elastic IP address (IPv4).
				* `carrier_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The carrier IP address associated with the network interface.
				* `customer_owned_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The customer-owned IP address associated with the network interface.
				* `ip_owner_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the owner of the Elastic IP address.
				* `public_dns_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The public DNS name.
				* `public_ip` - (Applicable when asset_type=AWS_EC2) (Updatable) The public IP address or Elastic IP address bound to the network interface.
			* `is_primary` - (Applicable when asset_type=AWS_EC2) (Updatable) Indicates whether this IPv4 address is the primary private IP address of the network interface.
			* `private_dns_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The private IPv4 DNS name.
			* `private_ip_address` - (Applicable when asset_type=AWS_EC2) (Updatable) The private IPv4 address of the network interface.
		* `security_groups` - (Applicable when asset_type=AWS_EC2) (Updatable) The security groups.
			* `group_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the security group.
			* `group_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The name of the security group.
		* `status` - (Applicable when asset_type=AWS_EC2) (Updatable) The status of the network interface.
		* `subnet_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the subnet.
	* `placement` - (Applicable when asset_type=AWS_EC2) (Updatable) Describes the placement of an instance.
		* `affinity` - (Applicable when asset_type=AWS_EC2) (Updatable) The affinity setting for the instance on the Dedicated Host.
		* `availability_zone` - (Applicable when asset_type=AWS_EC2) (Updatable) The Availability Zone of the instance.
		* `group_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The name of the placement group the instance is in.
		* `host_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the Dedicated Host on which the instance resides.
		* `host_resource_group_arn` - (Applicable when asset_type=AWS_EC2) (Updatable) The ARN of the host resource group in which to launch the instances.
		* `partition_number` - (Applicable when asset_type=AWS_EC2) (Updatable) The number of the partition that the instance is in.
		* `spread_domain` - (Applicable when asset_type=AWS_EC2) (Updatable) Reserved for future use.
		* `tenancy` - (Applicable when asset_type=AWS_EC2) (Updatable) The tenancy of the instance (if the instance is running in a VPC).
	* `private_dns_name` - (Applicable when asset_type=AWS_EC2) (Updatable) (IPv4 only) The private DNS hostname name assigned to the instance.
	* `private_ip_address` - (Applicable when asset_type=AWS_EC2) (Updatable) The private IPv4 address assigned to the instance.
	* `root_device_name` - (Required when asset_type=AWS_EC2) (Updatable) The device name of the root device volume.
	* `root_device_type` - (Applicable when asset_type=AWS_EC2) (Updatable) The root device type used by the AMI. The AMI can use an EBS volume or an instance store volume.
	* `security_groups` - (Applicable when asset_type=AWS_EC2) (Updatable) The security groups for the instance.
		* `group_key` - (Applicable when asset_type=AWS_EC2) (Updatable) The ID of the security group.
		* `group_name` - (Applicable when asset_type=AWS_EC2) (Updatable) The name of the security group.
	* `sriov_net_support` - (Applicable when asset_type=AWS_EC2) (Updatable) Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.
	* `state` - (Required when asset_type=AWS_EC2) (Updatable) Describes the current state of an instance.
		* `code` - (Applicable when asset_type=AWS_EC2) (Updatable) The state of the instance as a 16-bit unsigned integer.
		* `name` - (Applicable when asset_type=AWS_EC2) (Updatable) The current state of the instance.
	* `subnet_key` - (Applicable when asset_type=AWS_EC2) (Updatable) EC2-VPC The ID of the subnet in which the instance is running.
	* `tags` - (Applicable when asset_type=AWS_EC2) (Updatable) Any tags assigned to the instance.
		* `key` - (Applicable when asset_type=AWS_EC2) (Updatable) The key of the tag.
		* `value` - (Applicable when asset_type=AWS_EC2) (Updatable) The value of the tag.
	* `time_launch` - (Applicable when asset_type=AWS_EC2) (Updatable) The time the instance was launched.
	* `tpm_support` - (Applicable when asset_type=AWS_EC2) (Updatable) If the instance is configured for NitroTPM support, the value is v2.0.
	* `virtualization_type` - (Applicable when asset_type=AWS_EC2) (Updatable) The virtualization type of the instance.
	* `vpc_key` - (Applicable when asset_type=AWS_EC2) (Updatable) EC2-VPC The ID of the VPC in which the instance is running.
* `aws_ec2cost` - (Applicable when asset_type=AWS_EC2) (Updatable) Cost information for monthly maintenance.
	* `amount` - (Required when asset_type=AWS_EC2) (Updatable) Monthly costs for maintenance of this asset.
	* `currency_code` - (Required when asset_type=AWS_EC2) (Updatable) Currency code
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that the asset belongs to.
* `compute` - (Required) (Updatable) Compute related properties.
	* `connected_networks` - (Optional) (Updatable) Number of connected networks.
	* `cores_count` - (Optional) (Updatable) Number of CPUs.
	* `cpu_model` - (Optional) (Updatable) CPU model name.
	* `description` - (Optional) (Updatable) Information about the asset.
	* `disks` - (Optional) (Updatable) Lists the set of disks belonging to the virtual machine. This list is unordered.
		* `boot_order` - (Optional) (Updatable) Order of boot volumes.
		* `is_cbt_enabled` - (Optional) (Updatable) Indicates that CBT (change disk tracking) is enabled for this virtual disk.
		* `location` - (Optional) (Updatable) Location of the boot/data volume.
		* `name` - (Optional) (Updatable) Disk name.
		* `persistent_mode` - (Optional) (Updatable) The disk persistent mode.
		* `size_in_mbs` - (Optional) (Updatable) The size of the volume in MBs.
		* `uuid` - (Optional) (Updatable) Disk UUID for the virtual disk, if available.
		* `uuid_lun` - (Optional) (Updatable) Disk UUID LUN for the virtual disk, if available.
	* `disks_count` - (Optional) (Updatable) Number of disks.
	* `dns_name` - (Optional) (Updatable) Fully Qualified DNS Name.
	* `firmware` - (Optional) (Updatable) Information about firmware type for this virtual machine.
	* `gpu_devices` - (Optional) (Updatable) List of GPU devices attached to a virtual machine.
		* `cores_count` - (Optional) (Updatable) Number of GPU cores.
		* `description` - (Optional) (Updatable) GPU device description.
		* `manufacturer` - (Optional) (Updatable) The manufacturer of GPU.
		* `memory_in_mbs` - (Optional) (Updatable) GPU memory size in MBs.
		* `name` - (Optional) (Updatable) GPU device name.
	* `gpu_devices_count` - (Optional) (Updatable) Number of GPU devices.
	* `guest_state` - (Optional) (Updatable) Guest state.
	* `hardware_version` - (Optional) (Updatable) Hardware version.
	* `host_name` - (Optional) (Updatable) Host name of the VM.
	* `is_pmem_enabled` - (Optional) (Updatable) Whether Pmem is enabled. Decides if NVDIMMs are used as a permanent memory.
	* `is_tpm_enabled` - (Optional) (Updatable) Whether Trusted Platform Module (TPM) is enabled.
	* `latency_sensitivity` - (Optional) (Updatable) Latency sensitivity.
	* `memory_in_mbs` - (Optional) (Updatable) Memory size in MBs.
	* `nics` - (Optional) (Updatable) List of network ethernet cards attached to a virtual machine.
		* `ip_addresses` - (Optional) (Updatable) List of IP addresses.
		* `label` - (Optional) (Updatable) Provides a label and summary information for the device.
		* `mac_address` - (Optional) (Updatable) Mac address of the VM.
		* `mac_address_type` - (Optional) (Updatable) Mac address type.
		* `network_name` - (Optional) (Updatable) Network name.
		* `switch_name` - (Optional) (Updatable) Switch name.
	* `nics_count` - (Optional) (Updatable) Number of network ethernet cards.
	* `nvdimm_controller` - (Optional) (Updatable) The asset's NVDIMM configuration.
		* `bus_number` - (Optional) (Updatable) Bus number.
		* `label` - (Optional) (Updatable) Provides a label and summary information for the device.
	* `nvdimms` - (Optional) (Updatable) The properties of the NVDIMMs attached to a virtual machine.
		* `controller_key` - (Optional) (Updatable) Controller key.
		* `label` - (Optional) (Updatable) Provides a label and summary information for the device.
		* `unit_number` - (Optional) (Updatable) The unit number of NVDIMM.
	* `operating_system` - (Optional) (Updatable) Operating system.
	* `operating_system_version` - (Optional) (Updatable) Operating system version.
	* `pmem_in_mbs` - (Optional) (Updatable) Pmem size in MBs.
	* `power_state` - (Optional) (Updatable) The current power state of the virtual machine.
	* `primary_ip` - (Optional) (Updatable) Primary IP address of the compute instance.
	* `scsi_controller` - (Optional) (Updatable) The assets SCSI controller.
		* `label` - (Optional) (Updatable) Provides a label and summary information for the device.
		* `shared_bus` - (Optional) (Updatable) Shared bus.
		* `unit_number` - (Optional) (Updatable) The unit number of the SCSI controller.
	* `storage_provisioned_in_mbs` - (Optional) (Updatable) Provision storage size in MBs.
	* `threads_per_core_count` - (Optional) (Updatable) Number of threads per core.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) Asset display name.
* `external_asset_key` - (Required) The key of the asset from the external environment.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `inventory_id` - (Required) Inventory ID to which an asset belongs.
* `source_key` - (Required) The source key to which the asset belongs.
* `vm` - (Required) (Updatable) Virtual machine related properties.
	* `hypervisor_host` - (Optional) (Updatable) Host name/IP address of VM on which the host is running.
	* `hypervisor_vendor` - (Optional) (Updatable) Hypervisor vendor.
	* `hypervisor_version` - (Optional) (Updatable) Hypervisor version.
* `vmware_vcenter` - (Required when asset_type=VMWARE_VM) (Updatable) VMware vCenter related properties.
	* `data_center` - (Applicable when asset_type=VMWARE_VM) (Updatable) Data center name.
	* `vcenter_key` - (Applicable when asset_type=VMWARE_VM) (Updatable) vCenter unique key.
	* `vcenter_version` - (Applicable when asset_type=VMWARE_VM) (Updatable) Dot-separated version string.
* `vmware_vm` - (Required when asset_type=VMWARE_VM) (Updatable) VMware virtual machine related properties.
	* `cluster` - (Applicable when asset_type=VMWARE_VM) (Updatable) Cluster name.
	* `customer_fields` - (Applicable when asset_type=VMWARE_VM) (Updatable) Customer fields.
	* `customer_tags` - (Applicable when asset_type=VMWARE_VM) (Updatable) Customer defined tags.
		* `description` - (Applicable when asset_type=VMWARE_VM) (Updatable) The tag description.
		* `name` - (Applicable when asset_type=VMWARE_VM) (Updatable) The tag name.
	* `fault_tolerance_bandwidth` - (Applicable when asset_type=VMWARE_VM) (Updatable) Fault tolerance bandwidth.
	* `fault_tolerance_secondary_latency` - (Applicable when asset_type=VMWARE_VM) (Updatable) Fault tolerance to secondary latency.
	* `fault_tolerance_state` - (Applicable when asset_type=VMWARE_VM) (Updatable) Fault tolerance state.
	* `instance_uuid` - (Applicable when asset_type=VMWARE_VM) (Updatable) vCenter-specific identifier of the virtual machine.
	* `is_disks_cbt_enabled` - (Applicable when asset_type=VMWARE_VM) (Updatable) Indicates that change tracking is supported for virtual disks of this virtual machine. However, even if change tracking is supported, it might not be available for all disks of the virtual machine. 
	* `is_disks_uuid_enabled` - (Applicable when asset_type=VMWARE_VM) (Updatable) Whether changed block tracking for this VM's disk is active.
	* `path` - (Applicable when asset_type=VMWARE_VM) (Updatable) Path directory of the asset.
	* `vmware_tools_status` - (Applicable when asset_type=VMWARE_VM) (Updatable) VMware tools status.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `asset_class_name` - The class name of the asset.
* `asset_class_version` - The version of the asset class.
* `asset_details` - The details of the asset.
* `asset_source_ids` - List of asset source OCID.
* `asset_type` - The type of asset.
* `attached_ebs_volumes_cost` - Cost information for monthly maintenance.
	* `amount` - Monthly costs for maintenance of this asset.
	* `currency_code` - Currency code
* `aws_ebs` - AWS EBS volume related properties.
	* `attachments` - Information about the volume attachments.
		* `device` - The device name.
		* `instance_key` - The ID of the instance.
		* `is_delete_on_termination` - Indicates whether the EBS volume is deleted on instance termination.
		* `status` - The attachment state of the volume.
		* `volume_key` - The ID of the volume.
	* `availability_zone` - The Availability Zone for the volume.
	* `iops` - The number of I/O operations per second.
	* `is_encrypted` - Indicates whether the volume is encrypted.
	* `is_multi_attach_enabled` - Indicates whether Amazon EBS Multi-Attach is enabled.
	* `size_in_gi_bs` - The size of the volume, in GiBs.
	* `status` - The volume state.
	* `tags` - Any tags assigned to the volume.
		* `key` - The key of the tag.
		* `value` - The value of the tag.
	* `throughput` - The throughput that the volume supports, in MiB/s.
	* `volume_key` - The ID of the volume.
	* `volume_type` - The volume type.
* `aws_ec2` - AWS virtual machine related properties.
	* `architecture` - The architecture of the image.
	* `are_elastic_inference_accelerators_present` - Indicates if the elastic inference accelerators attached to an instance
	* `boot_mode` - The boot mode of the instance.
	* `capacity_reservation_key` - The ID of the Capacity Reservation.
	* `image_key` - The ID of the AMI used to launch the instance.
	* `instance_key` - The ID of the instance.
	* `instance_lifecycle` - Indicates whether this is a Spot Instance or a Scheduled Instance.
	* `instance_type` - The instance type.
	* `ip_address` - The public IPv4 address, or the Carrier IP address assigned to the instance.
	* `ipv6address` - The IPv6 address assigned to the instance.
	* `is_enclave_options` - Indicates whether the instance is enabled for AWS Nitro Enclaves.
	* `is_hibernation_options` - Indicates whether the instance is enabled for hibernation.
	* `is_source_dest_check` - Indicates whether source/destination checking is enabled.
	* `is_spot_instance` - If the request is a Spot Instance request, this value will be true.
	* `kernel_key` - The kernel associated with this instance, if applicable.
	* `licenses` - The license configurations for the instance.
	* `maintenance_options` - Provides information on the recovery and maintenance options of your instance.
	* `monitoring` - The monitoring for the instance.
	* `network_interfaces` - The network interfaces for the instance.
		* `association` - Describes association information for an Elastic IP address (IPv4).
			* `carrier_ip` - The carrier IP address associated with the network interface.
			* `customer_owned_ip` - The customer-owned IP address associated with the network interface.
			* `ip_owner_key` - The ID of the owner of the Elastic IP address.
			* `public_dns_name` - The public DNS name.
			* `public_ip` - The public IP address or Elastic IP address bound to the network interface.
		* `attachment` - Describes a network interface attachment.
			* `attachment_key` - The ID of the network interface attachment.
			* `device_index` - The index of the device on the instance for the network interface attachment.
			* `is_delete_on_termination` - Indicates whether the network interface is deleted when the instance is terminated.
			* `network_card_index` - The index of the network card.
			* `status` - The attachment state.
			* `time_attach` - The timestamp when the attachment initiated.
		* `description` - The description.
		* `interface_type` - The type of network interface.
		* `ipv4prefixes` - The IPv4 delegated prefixes that are assigned to the network interface.
		* `ipv6addresses` - The IPv6 addresses associated with the network interface.
		* `ipv6prefixes` - The IPv6 delegated prefixes that are assigned to the network interface.
		* `is_source_dest_check` - Indicates whether source/destination checking is enabled.
		* `mac_address` - The MAC address.
		* `network_interface_key` - The ID of the network interface.
		* `owner_key` - The ID of the AWS account that created the network interface.
		* `private_ip_addresses` - The private IPv4 addresses associated with the network interface.
			* `association` - Describes association information for an Elastic IP address (IPv4).
				* `carrier_ip` - The carrier IP address associated with the network interface.
				* `customer_owned_ip` - The customer-owned IP address associated with the network interface.
				* `ip_owner_key` - The ID of the owner of the Elastic IP address.
				* `public_dns_name` - The public DNS name.
				* `public_ip` - The public IP address or Elastic IP address bound to the network interface.
			* `is_primary` - Indicates whether this IPv4 address is the primary private IP address of the network interface.
			* `private_dns_name` - The private IPv4 DNS name.
			* `private_ip_address` - The private IPv4 address of the network interface.
		* `security_groups` - The security groups.
			* `group_key` - The ID of the security group.
			* `group_name` - The name of the security group.
		* `status` - The status of the network interface.
		* `subnet_key` - The ID of the subnet.
	* `placement` - Describes the placement of an instance.
		* `affinity` - The affinity setting for the instance on the Dedicated Host.
		* `availability_zone` - The Availability Zone of the instance.
		* `group_name` - The name of the placement group the instance is in.
		* `host_key` - The ID of the Dedicated Host on which the instance resides.
		* `host_resource_group_arn` - The ARN of the host resource group in which to launch the instances.
		* `partition_number` - The number of the partition that the instance is in.
		* `spread_domain` - Reserved for future use.
		* `tenancy` - The tenancy of the instance (if the instance is running in a VPC).
	* `private_dns_name` - (IPv4 only) The private DNS hostname name assigned to the instance.
	* `private_ip_address` - The private IPv4 address assigned to the instance.
	* `root_device_name` - The device name of the root device volume.
	* `root_device_type` - The root device type used by the AMI. The AMI can use an EBS volume or an instance store volume.
	* `security_groups` - The security groups for the instance.
		* `group_key` - The ID of the security group.
		* `group_name` - The name of the security group.
	* `sriov_net_support` - Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.
	* `state` - Describes the current state of an instance.
		* `code` - The state of the instance as a 16-bit unsigned integer.
		* `name` - The current state of the instance.
	* `subnet_key` - EC2-VPC The ID of the subnet in which the instance is running.
	* `tags` - Any tags assigned to the instance.
		* `key` - The key of the tag.
		* `value` - The value of the tag.
	* `time_launch` - The time the instance was launched.
	* `tpm_support` - If the instance is configured for NitroTPM support, the value is v2.0.
	* `virtualization_type` - The virtualization type of the instance.
	* `vpc_key` - EC2-VPC The ID of the VPC in which the instance is running.
* `aws_ec2cost` - Cost information for monthly maintenance.
	* `amount` - Monthly costs for maintenance of this asset.
	* `currency_code` - Currency code
* `compartment_id` - The OCID of the compartment to which an asset belongs to.
* `compute` - Compute related properties.
	* `connected_networks` - Number of connected networks.
	* `cores_count` - Number of CPUs.
	* `cpu_model` - CPU model name.
	* `description` - Information about the asset.
	* `disks` - Lists the set of disks belonging to the virtual machine. This list is unordered.
		* `boot_order` - Order of boot volumes.
		* `is_cbt_enabled` - Indicates that CBT (change disk tracking) is enabled for this virtual disk.
		* `location` - Location of the boot/data volume.
		* `name` - Disk name.
		* `persistent_mode` - The disk persistent mode.
		* `size_in_mbs` - The size of the volume in MBs.
		* `uuid` - Disk UUID for the virtual disk, if available.
		* `uuid_lun` - Disk UUID LUN for the virtual disk, if available.
	* `disks_count` - Number of disks.
	* `dns_name` - Fully Qualified DNS Name.
	* `firmware` - Information about firmware type for this virtual machine.
	* `gpu_devices` - List of GPU devices attached to a virtual machine.
		* `cores_count` - Number of GPU cores.
		* `description` - GPU device description.
		* `manufacturer` - The manufacturer of GPU.
		* `memory_in_mbs` - GPU memory size in MBs.
		* `name` - GPU device name.
	* `gpu_devices_count` - Number of GPU devices.
	* `guest_state` - Guest state.
	* `hardware_version` - Hardware version.
	* `host_name` - Host name of the VM.
	* `is_pmem_enabled` - Whether Pmem is enabled. Decides if NVDIMMs are used as a permanent memory.
	* `is_tpm_enabled` - Whether Trusted Platform Module (TPM) is enabled.
	* `latency_sensitivity` - Latency sensitivity.
	* `memory_in_mbs` - Memory size in MBs.
	* `nics` - List of network ethernet cards attached to a virtual machine.
		* `ip_addresses` - List of IP addresses.
		* `label` - Provides a label and summary information for the device.
		* `mac_address` - Mac address of the VM.
		* `mac_address_type` - Mac address type.
		* `network_name` - Network name.
		* `switch_name` - Switch name.
	* `nics_count` - Number of network ethernet cards.
	* `nvdimm_controller` - The asset's NVDIMM configuration.
		* `bus_number` - Bus number.
		* `label` - Provides a label and summary information for the device.
	* `nvdimms` - The properties of the NVDIMMs attached to a virtual machine.
		* `controller_key` - Controller key.
		* `label` - Provides a label and summary information for the device.
		* `unit_number` - The unit number of NVDIMM.
	* `operating_system` - Operating system.
	* `operating_system_version` - Operating system version.
	* `pmem_in_mbs` - Pmem size in MBs.
	* `power_state` - The current power state of the virtual machine.
	* `primary_ip` - Primary IP address of the compute instance.
	* `scsi_controller` - The assets SCSI controller.
		* `label` - Provides a label and summary information for the device.
		* `shared_bus` - Shared bus.
		* `unit_number` - The unit number of the SCSI controller.
	* `storage_provisioned_in_mbs` - Provision storage size in MBs.
	* `threads_per_core_count` - Number of threads per core.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Asset display name.
* `environment_type` - Specifies if this is the Source or Destination point for migration - different assets may be discovered depending on setting.
* `external_asset_key` - The key of the asset from the external environment.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Asset OCID that is immutable on creation.
* `inventory_id` - Inventory ID to which an asset belongs to.
* `source_key` - The source key that the asset belongs to.
* `state` - The current state of the asset.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the asset was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the asset was updated. An RFC3339 formatted datetime string.
* `vm` - Virtual machine related properties.
	* `hypervisor_host` - Host name/IP address of VM on which the host is running.
	* `hypervisor_vendor` - Hypervisor vendor.
	* `hypervisor_version` - Hypervisor version.
* `vmware_vcenter` - VMware vCenter related properties.
	* `data_center` - Data center name.
	* `vcenter_key` - vCenter unique key.
	* `vcenter_version` - Dot-separated version string.
* `vmware_vm` - VMware virtual machine related properties.
	* `cluster` - Cluster name.
	* `customer_fields` - Customer fields.
	* `customer_tags` - Customer defined tags.
		* `description` - The tag description.
		* `name` - The tag name.
	* `fault_tolerance_bandwidth` - Fault tolerance bandwidth.
	* `fault_tolerance_secondary_latency` - Fault tolerance to secondary latency.
	* `fault_tolerance_state` - Fault tolerance state.
	* `instance_uuid` - vCenter-specific identifier of the virtual machine.
	* `is_disks_cbt_enabled` - Indicates that change tracking is supported for virtual disks of this virtual machine. However, even if change tracking is supported, it might not be available for all disks of the virtual machine. 
	* `is_disks_uuid_enabled` - Whether changed block tracking for this VM's disk is active.
	* `path` - Path directory of the asset.
	* `vmware_tools_status` - VMware tools status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Asset
	* `update` - (Defaults to 20 minutes), when updating the Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Asset


## Import

Assets can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_bridge_asset.test_asset "id"
```
