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
	asset_source_ids = var.asset_asset_source_ids
	compute {

		#Optional
		connected_networks = var.asset_compute_connected_networks
		cores_count = var.asset_compute_cores_count
		cpu_model = var.asset_compute_cpu_model
		description = var.asset_compute_description
		disks {

			#Optional
			boot_order = var.asset_compute_disks_boot_order
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

* `asset_source_ids` - (Optional) (Updatable) List of asset source OCID.
* `asset_type` - (Required) (Updatable) The type of asset.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that the asset belongs to.
* `compute` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Compute related properties.
	* `connected_networks` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of connected networks.
	* `cores_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of CPUs.
	* `cpu_model` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) CPU model name.
	* `description` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Information about the asset.
	* `disks` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Lists the set of disks belonging to the virtual machine. This list is unordered.
		* `boot_order` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Order of boot volumes.
		* `location` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Location of the boot/data volume.
		* `name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Disk name.
		* `persistent_mode` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The disk persistent mode.
		* `size_in_mbs` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The size of the volume in MBs.
		* `uuid` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Disk UUID for the virtual disk, if available.
		* `uuid_lun` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Disk UUID LUN for the virtual disk, if available.
	* `disks_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of disks.
	* `dns_name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Fully Qualified DNS Name.
	* `firmware` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Information about firmware type for this virtual machine.
	* `gpu_devices` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) List of GPU devices attached to a virtual machine.
		* `cores_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of GPU cores.
		* `description` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) GPU device description.
		* `manufacturer` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The manufacturer of GPU.
		* `memory_in_mbs` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) GPU memory size in MBs.
		* `name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) GPU device name.
	* `gpu_devices_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of GPU devices.
	* `guest_state` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Guest state.
	* `hardware_version` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Hardware version.
	* `host_name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Host name of the VM.
	* `is_pmem_enabled` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Whether Pmem is enabled. Decides if NVDIMMs are used as a permanent memory.
	* `is_tpm_enabled` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Whether Trusted Platform Module (TPM) is enabled.
	* `latency_sensitivity` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Latency sensitivity.
	* `memory_in_mbs` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Memory size in MBs.
	* `nics` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) List of network ethernet cards attached to a virtual machine.
		* `ip_addresses` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) List of IP addresses.
		* `label` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Provides a label and summary information for the device.
		* `mac_address` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Mac address of the VM.
		* `mac_address_type` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Mac address type.
		* `network_name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Network name.
		* `switch_name` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Switch name.
	* `nics_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of network ethernet cards.
	* `nvdimm_controller` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The asset's NVDIMM configuration.
		* `bus_number` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Bus number.
		* `label` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Provides a label and summary information for the device.
	* `nvdimms` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The properties of the NVDIMMs attached to a virtual machine.
		* `controller_key` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Controller key.
		* `label` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Provides a label and summary information for the device.
		* `unit_number` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The unit number of NVDIMM.
	* `operating_system` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Operating system.
	* `operating_system_version` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Operating system version.
	* `pmem_in_mbs` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Pmem size in MBs.
	* `power_state` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The current power state of the virtual machine.
	* `primary_ip` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Primary IP address of the compute instance.
	* `scsi_controller` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The assets SCSI controller.
		* `label` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Provides a label and summary information for the device.
		* `shared_bus` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Shared bus.
		* `unit_number` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) The unit number of the SCSI controller.
	* `storage_provisioned_in_mbs` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Provision storage size in MBs.
	* `threads_per_core_count` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Number of threads per core.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) Asset display name.
* `external_asset_key` - (Required) The key of the asset from the external environment.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `inventory_id` - (Required) Inventory ID to which an asset belongs.
* `source_key` - (Required) The source key to which the asset belongs.
* `vm` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Virtual machine related properties.
	* `hypervisor_host` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Host name/IP address of VM on which the host is running.
	* `hypervisor_vendor` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Hypervisor vendor.
	* `hypervisor_version` - (Applicable when asset_type=VM | VMWARE_VM) (Updatable) Hypervisor version.
* `vmware_vcenter` - (Optional) (Updatable) VMware vCenter related properties.
	* `data_center` - (Optional) (Updatable) Data center name.
	* `vcenter_key` - (Optional) (Updatable) vCenter unique key.
	* `vcenter_version` - (Optional) (Updatable) Dot-separated version string.
* `vmware_vm` - (Optional) (Updatable) VMware virtual machine related properties.
	* `cluster` - (Optional) (Updatable) Cluster name.
	* `customer_fields` - (Optional) (Updatable) Customer fields.
	* `customer_tags` - (Optional) (Updatable) Customer defined tags.
		* `description` - (Optional) (Updatable) The tag description.
		* `name` - (Optional) (Updatable) The tag name.
	* `fault_tolerance_bandwidth` - (Optional) (Updatable) Fault tolerance bandwidth.
	* `fault_tolerance_secondary_latency` - (Optional) (Updatable) Fault tolerance to secondary latency.
	* `fault_tolerance_state` - (Optional) (Updatable) Fault tolerance state.
	* `instance_uuid` - (Optional) (Updatable) vCenter-specific identifier of the virtual machine.
	* `is_disks_cbt_enabled` - (Optional) (Updatable) Indicates that change tracking is supported for virtual disks of this virtual machine. However, even if change tracking is supported, it might not be available for all disks of the virtual machine. 
	* `is_disks_uuid_enabled` - (Optional) (Updatable) Whether changed block tracking for this VM's disk is active.
	* `path` - (Optional) (Updatable) Path directory of the asset.
	* `vmware_tools_status` - (Optional) (Updatable) VMware tools status.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Asset
	* `update` - (Defaults to 20 minutes), when updating the Asset
	* `delete` - (Defaults to 20 minutes), when destroying the Asset


## Import

Assets can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_bridge_asset.test_asset "id"
```

