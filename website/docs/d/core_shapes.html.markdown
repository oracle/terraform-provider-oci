---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_shapes"
sidebar_current: "docs-oci-datasource-core-shapes"
description: |-
  Provides the list of Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_shapes
This data source provides the list of Shapes in Oracle Cloud Infrastructure Core service.

Lists the shapes that can be used to launch an instance within the specified compartment. You can
filter the list by compatibility with a specific image.


## Example Usage

```hcl
data "oci_core_shapes" "test_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.shape_availability_domain
	image_id = oci_core_image.test_image.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `image_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an image.


## Attributes Reference

The following attributes are exported:

* `shapes` - The list of shapes.

### Shape Reference

The following attributes are exported:

* `baseline_ocpu_utilizations` - For a subcore burstable VM, the supported baseline OCPU utilization for instances that use this shape. 
* `billing_type` - How instances that use this shape are charged. 
* `gpu_description` - A short description of the graphics processing unit (GPU) available for this shape.

	If the shape does not have any GPUs, this field is `null`. 
* `gpus` - The number of GPUs available for this shape. 
* `is_billed_for_stopped_instance` - Whether billing continues when the instances that use this shape are in the stopped state. 
* `is_flexible` - Whether the shape supports creating flexible instances. A [flexible shape](https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible) is a shape that lets you customize the number of OCPUs and the amount of memory when launching or resizing your instance. 
* `is_live_migration_supported` - Whether live migration is supported for this shape. 
* `is_subcore` - Whether the shape supports creating subcore or burstable instances. A [burstable instance](https://docs.cloud.oracle.com/iaas/Content/Compute/References/burstable-instances.htm) is a virtual machine (VM) instance that provides a baseline level of CPU performance with the ability to burst to a higher level to support occasional spikes in usage. 
* `local_disk_description` - A short description of the local disks available for this shape.

	If the shape does not have any local disks, this field is `null`. 
* `local_disks` - The number of local disks available for this shape. 
* `local_disks_total_size_in_gbs` - The aggregate size of the local disks available for this shape, in gigabytes.

	If the shape does not have any local disks, this field is `null`. 
* `max_vnic_attachment_options` - For a flexible shape, the number of VNIC attachments that are available for instances that use this shape.

	If this field is null, then this shape has a fixed maximum number of VNIC attachments equal to `maxVnicAttachments`. 
	* `default_per_ocpu` - The default number of VNIC attachments allowed per OCPU. 
	* `max` - The highest maximum value of VNIC attachments. 
	* `min` - The lowest maximum value of VNIC attachments. 
* `max_vnic_attachments` - The maximum number of VNIC attachments available for this shape. 
* `memory_in_gbs` - The default amount of memory available for this shape, in gigabytes. 
* `memory_options` - For a flexible shape, the amount of memory available for instances that use this shape.

	If this field is null, then this shape has a fixed amount of memory equivalent to `memoryInGBs`. 
	* `default_per_ocpu_in_gbs` - The default amount of memory per OCPU available for this shape, in gigabytes. 
	* `max_in_gbs` - The maximum amount of memory, in gigabytes. 
	* `max_per_numa_node_in_gbs` - The maximum amount of memory per NUMA node, in gigabytes. 
	* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape, in gigabytes. 
	* `min_in_gbs` - The minimum amount of memory, in gigabytes. 
	* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape, in gigabytes. 
* `min_total_baseline_ocpus_required` - For a subcore burstable VM, the minimum total baseline OCPUs required. The total baseline OCPUs is equal to baselineOcpuUtilization chosen multiplied by the number of OCPUs chosen. 
* `name` - The name of the shape. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `network_ports` - The number of physical network interface card (NIC) ports available for this shape. 
* `networking_bandwidth_in_gbps` - The networking bandwidth available for this shape, in gigabits per second. 
* `networking_bandwidth_options` - For a flexible shape, the amount of networking bandwidth available for instances that use this shape.

	If this field is null, then this shape has a fixed amount of bandwidth equivalent to `networkingBandwidthInGbps`. 
	* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second. 
	* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second. 
	* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second. 
* `ocpu_options` - For a flexible shape, the number of OCPUs available for instances that use this shape.

	If this field is null, then this shape has a fixed number of OCPUs equal to `ocpus`. 
	* `max` - The maximum number of OCPUs. 
	* `max_per_numa_node` - The maximum number of cores available per NUMA node. 
	* `min` - The minimum number of OCPUs. 
* `ocpus` - The default number of OCPUs available for this shape. 

* `platform_config_options` - The list of supported platform configuration options for this shape. 
	* `access_control_service_options` - Configuration options for the Access Control Service. 
		* `allowed_values` - Whether the Access Control Service can be enabled. 
		* `is_default_enabled` - Whether the Access Control Service is enabled by default. 
	* `input_output_memory_management_unit_options` - Configuration options for the input-output memory management unit (IOMMU). 
		* `allowed_values` - Whether the input-output memory management unit can be enabled. 
		* `is_default_enabled` - Whether the input-output memory management unit is enabled by default. 
	* `measured_boot_options` - Configuration options for the Measured Boot feature. 
		* `allowed_values` - Boolean values that indicate whether the Measured Boot feature can be enabled or disabled. 
		* `is_default_enabled` - Whether the Measured Boot feature is enabled by default. 
	* `memory_encryption_options` - Configuration options for memory encryption. 
		* `allowed_values` - Whether memory encryption can be enabled. 
		* `is_default_enabled` - Whether memory encryption is enabled by default. 
	* `numa_nodes_per_socket_platform_options` - Configuration options for NUMA nodes per socket. 
		* `allowed_values` - The supported values for this platform configuration property. 
		* `default_value` - The default NUMA nodes per socket configuration. 
	* `percentage_of_cores_enabled_options` - Configuration options for the percentage of cores enabled. 
		* `default_value` - The default percentage of cores enabled. 
		* `max` - The maximum allowed percentage of cores enabled. 
		* `min` - The minimum allowed percentage of cores enabled. 
	* `secure_boot_options` - Configuration options for Secure Boot. 
		* `allowed_values` - Boolean values that indicate whether Secure Boot can be enabled or disabled. 
		* `is_default_enabled` - Whether Secure Boot is enabled by default. 
	* `symmetric_multi_threading_options` - Configuration options for symmetric multithreading (also called simultaneous multithreading or SMT). 
		* `allowed_values` - Whether symmetric multithreading can be enabled. 
		* `is_default_enabled` - Whether symmetric multithreading is enabled by default. 
	* `trusted_platform_module_options` - Configuration options for the Trusted Platform Module (TPM). 
		* `allowed_values` - Boolean values that indicate whether the Trusted Platform Module can be enabled or disabled. 
		* `is_default_enabled` - Whether the Trusted Platform Module is enabled by default. 
	* `type` - The type of platform being configured. 
	* `virtual_instructions_options` - Configuration options for the virtualization instructions. 
		* `allowed_values` - Whether virtualization instructions can be enabled. 
		* `is_default_enabled` - Whether virtualization instructions are enabled by default. 

* `platform_config_options` - The list of supported platform configuration options for this shape.
	* `access_control_service_options` - Configuration options for the Access Control Service.
		* `allowed_values` - Whether the Access Control Service can be enabled.
		* `is_default_enabled` - Whether the Access Control Service is enabled by default.
	* `input_output_memory_management_unit_options` - Configuration options for the input-output memory management unit.
		* `allowed_values` - Whether the input-output memory management unit can be enabled.
		* `is_default_enabled` - Whether the input-output memory management unit is enabled by default.
	* `measured_boot_options` - Configuration options for the Measured Boot feature.
		* `allowed_values` - Boolean values that indicate whether the Measured Boot feature can be enabled or disabled.
		* `is_default_enabled` - Whether the Measured Boot feature is enabled by default.
	* `numa_nodes_per_socket_platform_options` - Configuration options for NUMA nodes per socket.
		* `allowed_values` - The supported values for this platform configuration property.
		* `default_value` - The default NUMA nodes per socket configuration.
	* `percentage_of_cores_enabled_options` - Configuration options for the percentage of cores enabled.
		* `default_value` - The default percentage of cores enabled.
		* `max` - The maximum allowed percentage of cores enabled.
		* `min` - The minimum allowed percentage of cores enabled.
	* `secure_boot_options` - Configuration options for Secure Boot.
		* `allowed_values` - Boolean values that indicate whether Secure Boot can be enabled or disabled.
		* `is_default_enabled` - Whether Secure Boot is enabled by default.
	* `symmetric_multi_threading_options` - Configuration options for symmetric multi-threading.
		* `allowed_values` - Whether symmetric multi-threading can be enabled.
		* `is_default_enabled` - Whether symmetric multi-threading is enabled by default.
	* `trusted_platform_module_options` - Configuration options for the Trusted Platform Module (TPM).
		* `allowed_values` - Boolean values that indicate whether the Trusted Platform Module can be enabled or disabled.
		* `is_default_enabled` - Whether the Trusted Platform Module is enabled by default.
	* `type` - The type of platform being configured. (Supported types=[INTEL_VM, AMD_MILAN_BM, AMD_ROME_BM, AMD_ROME_BM_GPU, INTEL_ICELAKE_BM, INTEL_SKYLAKE_BM])
	* `virtual_instructions_options` - Configuration options for the virtualization instructions.
		* `allowed_values` - Whether virtualization instructions can be enabled.
		* `is_default_enabled` - Whether virtualization instructions are enabled by default.

* `processor_description` - A short description of the shape's processor (CPU). 
* `quota_names` - The list of of compartment quotas for the shape. 
* `rdma_bandwidth_in_gbps` - The networking bandwidth available for the remote direct memory access (RDMA) network for this shape, in gigabits per second. 
* `rdma_ports` - The number of networking ports available for the remote direct memory access (RDMA) network between nodes in a high performance computing (HPC) cluster network. If the shape does not support cluster networks, this value is `0`. 
* `recommended_alternatives` - The list of shapes and shape details (if applicable) that Oracle recommends that you use as an alternative to the current shape. 
	* `shape_name` - The name of the shape. 
* `resize_compatible_shapes` - The list of compatible shapes that this shape can be changed to. For more information, see [Changing the Shape of an Instance](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/resizinginstances.htm). 

