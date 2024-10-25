---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_assets"
sidebar_current: "docs-oci-datasource-cloud_bridge-assets"
description: |-
  Provides the list of Assets in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_assets
This data source provides the list of Assets in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of assets.


## Example Usage

```hcl
data "oci_cloud_bridge_assets" "test_assets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	asset_id = oci_cloud_bridge_asset.test_asset.id
	asset_type = var.asset_asset_type
	display_name = var.asset_display_name
	external_asset_key = var.asset_external_asset_key
	inventory_id = oci_cloud_bridge_inventory.test_inventory.id
	source_key = var.asset_source_key
	state = var.asset_state
}
```

## Argument Reference

The following arguments are supported:

* `asset_id` - (Optional) Unique asset identifier.
* `asset_type` - (Optional) The type of asset.
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `external_asset_key` - (Optional) External asset key.
* `inventory_id` - (Optional) Unique Inventory identifier.
* `source_key` - (Optional) Source key from where the assets originate.
* `state` - (Optional) A filter to return only assets whose lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `asset_collection` - The list of asset_collection.

### Asset Reference

The following attributes are exported:

* `asset_source_ids` - List of asset source OCID.
* `asset_type` - The type of asset.
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

