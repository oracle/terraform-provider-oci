---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instances"
sidebar_current: "docs-oci-datasource-core-instances"
description: |-
  Provides the list of Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instances
This data source provides the list of Instances in Oracle Cloud Infrastructure Core service.

Lists the instances in the specified compartment and the specified availability domain.
You can filter the results by specifying an instance name (the list will include all the identically-named
instances in the compartment).

**Note:** To retrieve public and private IP addresses for an instance, use the [ListVnicAttachments](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/VnicAttachment/ListVnicAttachments)
operation to get the VNIC ID for the instance, and then call [GetVnic](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vnic/GetVnic) with the VNIC ID.


## Example Usage

```hcl
data "oci_core_instances" "test_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.instance_availability_domain
	capacity_reservation_id = oci_core_capacity_reservation.test_capacity_reservation.id
	compute_cluster_id = oci_core_compute_cluster.test_compute_cluster.id
	display_name = var.instance_display_name
	state = var.instance_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `capacity_reservation_id` - (Optional) The OCID of the compute capacity reservation.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_cluster_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster. A [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm) is a remote direct memory access (RDMA) network group. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

### Instance Reference

The following attributes are exported:

* `agent_config` - Configuration options for the Oracle Cloud Agent software running on the instance.
	* `are_all_plugins_disabled` - Whether Oracle Cloud Agent can run all of the available plugins. This includes the management and monitoring plugins.

		For more information about the available plugins, see [Managing Plugins with Oracle Cloud Agent](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm). 
	* `is_management_disabled` - Whether Oracle Cloud Agent can run all the available management plugins.

		These are the management plugins: OS Management Service Agent and Compute Instance Run Command.

		The management plugins are controlled by this parameter and by the per-plugin configuration in the `pluginsConfig` object.
		* If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of the per-plugin configuration.
		* If `isManagementDisabled` is false, all of the management plugins are enabled. You can optionally disable individual management plugins by providing a value in the `pluginsConfig` object. 
	* `is_monitoring_disabled` - Whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the monitoring plugins.

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
* `availability_domain` - The availability domain the instance is running in.  Example: `Uocm:PHX-AD-1`
* `cluster_placement_group_id` - The OCID of the cluster placement group of the instance.
* `boot_volume_id` - The OCID of the attached boot volume. If the `source_type` is `bootVolume`, this will be the same OCID as the `source_id`.
* `capacity_reservation_id` - The OCID of the compute capacity reservation this instance is launched under. When this field contains an empty string or is null, the instance is not currently in a capacity reservation. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
* `compartment_id` - The OCID of the compartment that contains the instance.
* `dedicated_vm_host_id` - The OCID of the dedicated virtual machine host that the instance is placed on. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

	They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only). 
	
    If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead. 
* `fault_domain` - The name of the fault domain the instance is running in.

    A fault domain is a grouping of hardware and infrastructure within an availability domain. Each availability domain contains three fault domains. Fault domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or Compute hardware maintenance that affects one fault domain does not affect instances in other fault domains.

	If you do not specify the fault domain, the system selects one for you.

	Example: `FAULT-DOMAIN-1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance.
* `image` - Deprecated. Use `sourceDetails` instead. 
* `instance_configuration_id` - The OCID of the Instance Configuration used to source launch details for this instance. Any other fields supplied in the instance launch request override the details stored in the Instance Configuration for this instance launch.
* `instance_options` - Optional mutable instance options
	* `are_legacy_imds_endpoints_disabled` - Whether to disable the legacy (/v1) instance metadata service endpoints. Customers who have migrated to /v2 should set this to true for added security. Default is false. 
* `ipxe_script` - When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.

	If you want more control over the boot process, you can provide your own custom iPXE script that will run when the instance boots. Be aware that the same iPXE script will run every time an instance boots, not only after the initial LaunchInstance call.

	The default iPXE script connects to the instance's local boot volume over iSCSI and performs a network boot. If you use a custom iPXE script and want to network-boot from the instance's local boot volume over iSCSI the same way as the default iPXE script, use the following iSCSI IP address: 169.254.0.2, and boot volume IQN: iqn.2015-02.oracle.boot.

	If your instance boot volume attachment type is paravirtualized, the boot volume is attached to the instance through virtio-scsi and no iPXE script is used. If your instance boot volume attachment type is paravirtualized and you use custom iPXE to network boot into your instance, the primary boot volume is attached as a data volume through virtio-scsi drive.

	For more information about the Bring Your Own Image feature of Oracle Cloud Infrastructure, see [Bring Your Own Image](https://docs.cloud.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).

	For more information about iPXE, see http://ipxe.org. 
* `is_cross_numa_node` - Whether the instance’s OCPUs and memory are distributed across multiple NUMA nodes. 
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
	* `is_pv_encryption_in_transit_enabled` - Deprecated. Instead use `isPvEncryptionInTransitEnabled` in [LaunchInstanceDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/datatypes/LaunchInstanceDetails). 
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
* `metadata` - Custom metadata that you provide.

* `platform_config` - The platform configuration for the instance. 
	* `are_virtual_instructions_enabled` - Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes or VT-x for Intel shapes. 
	* `config_map` - Instance Platform Configuration Configuration Map for flexible setting input. 
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

* `platform_config` - The platform configuration for the instance.
	* `are_virtual_instructions_enabled` - Whether virtualization instructions are available.
	* `is_access_control_service_enabled` - Whether the Access Control Service is enabled on the instance. When enabled, the platform can enforce PCIe device isolation, required for VFIO device passthrough. 
	* `is_input_output_memory_management_unit_enabled` - Whether the input-output memory management unit is enabled.
	* `is_measured_boot_enabled` - Whether the Measured Boot is to be enabled on the instance.
	* `is_secure_boot_enabled` - Whether the Secure Boot is to be enabled on the instance.
	* `is_symmetric_multi_threading_enabled` - Whether symmetric multi-threading is enabled on the instance.
	* `is_trusted_platform_module_enabled` - Whether the Trusted Platform Module (TPM) is to be enabled on the instance.
	* `numa_nodes_per_socket` - The number of NUMA nodes per socket (NPS).
	* `percentage_of_cores_enabled` - The percentage of cores enabled.
	* `type` - The type of platform being configured. (Supported types=[INTEL_VM, AMD_MILAN_BM, AMD_ROME_BM, AMD_ROME_BM_GPU, INTEL_ICELAKE_BM, INTEL_SKYLAKE_BM])
* `preemptible_instance_config` - (Optional) Configuration options for preemptible instances. 
	* `preemption_action` - (Required) The action to run when the preemptible instance is interrupted for eviction. 
		* `preserve_boot_volume` - (Optional) Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
		* `type` - (Required) The type of action to run when the instance is interrupted for eviction.

* `region` - The region that contains the availability domain the instance is running in.

	For the us-phoenix-1 and us-ashburn-1 regions, `phx` and `iad` are returned, respectively. For all other regions, the full region name is returned.

	Examples: `phx`, `eu-frankfurt-1` 
* `security_attributes` - Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.  Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}` 
* `security_attributes_state` - The lifecycle state of the `securityAttributes`
* `shape` - The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `shape_config` - The shape configuration for an instance. The shape configuration determines the resources allocated to an instance. 
	* `baseline_ocpu_utilization` - The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.

		The following values are supported:
		* `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
		* `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
		* `BASELINE_1_1` - baseline usage is the entire OCPU. This represents a non-burstable instance. 
	* `gpu_description` - A short description of the instance's graphics processing unit (GPU).

		If the instance does not have any GPUs, this field is `null`. 
	* `gpus` - The number of GPUs available to the instance. 
	* `local_disk_description` - A short description of the local disks available to this instance.

		If the instance does not have any local disks, this field is `null`. 
	* `local_disks` - The number of local disks available to the instance. 
	* `local_disks_total_size_in_gbs` - The aggregate size of all local disks, in gigabytes.

		If the instance does not have any local disks, this field is `null`. 
	* `max_vnic_attachments` - The maximum number of VNIC attachments for the instance. 
	* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes. 
	* `networking_bandwidth_in_gbps` - The networking bandwidth available to the instance, in gigabits per second. 
	* `ocpus` - The total number of OCPUs available to the instance. 
	* `processor_description` - A short description of the instance's processor (CPU). 
	* `vcpus` - The total number of VCPUs available to the instance. This can be used instead of OCPUs, in which case the actual number of OCPUs will be calculated based on this value and the actual hardware. This must be a multiple of 2.
* `source_details` - 
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 32,768 GB (32 TB). 
	* `boot_volume_vpus_per_gb` - The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Performance Levels](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.

		Allowed values:
		* `10`: Represents Balanced option.
		* `20`: Represents Higher Performance option.
		* `30`-`120`: Represents the Ultra High Performance option.

		For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB. 
	* `instance_source_image_filter_details` - These are the criteria for selecting an image. This is required if imageId is not specified.
		* `compartment_id` - The OCID of the compartment containing images to search
		* `defined_tags_filter` - Filter based on these defined tags. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
		* `operating_system` - The image's operating system.  Example: `Oracle Linux` 
		* `operating_system_version` - The image's operating system version.  Example: `7.2` 
	* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the boot volume.
	* `source_id` - The OCID of an image or a boot volume to use, depending on the value of `source_type`.
	* `source_type` - The source type for the instance. Use `image` when specifying the image OCID. Use `bootVolume` when specifying the boot volume OCID. 
* `state` - The current state of the instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_maintenance_reboot_due` - The date and time the instance is expected to be stopped / started,  in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time. Regardless of how the instance was stopped, the flag will be reset to empty as soon as instance reaches Stopped state. Example: `2018-05-25T21:10:29.600Z` 

