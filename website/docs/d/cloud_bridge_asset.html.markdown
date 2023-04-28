---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_asset"
sidebar_current: "docs-oci-datasource-cloud_bridge-asset"
description: |-
  Provides details about a specific Asset in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_asset
This data source provides details about a specific Asset resource in Oracle Cloud Infrastructure Cloud Bridge service.

Gets an asset by identifier.

## Example Usage

```hcl
data "oci_cloud_bridge_asset" "test_asset" {
	#Required
	asset_id = oci_cloud_bridge_asset.test_asset.id
}
```

## Argument Reference

The following arguments are supported:

* `asset_id` - (Required) Unique asset identifier.


## Attributes Reference

The following attributes are exported:

* `asset_source_ids` - List of asset source OCID.
* `asset_type` - The type of asset.
* `attached_ebs_volumes_cost` - Cost information for monthly maintenance.
	* `amount` - Monthly costs for maintenance of this asset.
	* `currency_code` - Currency code as defined by ISO-4217.
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
	* `currency_code` - Currency code as defined by ISO-4217.
* `compartment_id` - The OCID of the compartment to which an asset belongs to.
* `compute` - Compute related properties.
	* `connected_networks` - Number of connected networks.
	* `cores_count` - Number of CPUs.
	* `cpu_model` - CPU model name.
	* `description` - Information about the asset.
	* `disks` - Lists the set of disks belonging to the virtual machine. This list is unordered.
		* `boot_order` - Order of boot volumes.
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

