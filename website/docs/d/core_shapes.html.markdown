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
* `gpu_description` - A short description of the graphics processing unit (GPU) available for this shape.

	If the shape does not have any GPUs, this field is `null`. 
* `gpus` - The number of GPUs available for this shape. 
* `is_live_migration_supported` - Whether Live Migration is currently supported for this shape. 
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
	* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape, in gigabytes. 
	* `min_in_gbs` - The minimum amount of memory, in gigabytes. 
	* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape, in gigabytes. 
* `min_total_baseline_ocpus_required` - For a subcore burstable VM, the minimum total baseline OCPUs required. The total baseline OCPUs is equal to baselineOcpuUtilization chosen multiplied by the number of OCPUs chosen. 
* `name` - The name of the shape. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `networking_bandwidth_in_gbps` - The networking bandwidth available for this shape, in gigabits per second. 
* `networking_bandwidth_options` - For a flexible shape, the amount of networking bandwidth available for instances that use this shape.

	If this field is null, then this shape has a fixed amount of bandwidth equivalent to `networkingBandwidthInGbps`. 
	* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second. 
	* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second. 
	* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second. 
* `ocpu_options` - For a flexible shape, the number of OCPUs available for instances that use this shape.

	If this field is null, then this shape has a fixed number of OCPUs equal to `ocpus`. 
	* `max` - The maximum number of OCPUs. 
	* `min` - The minimum number of OCPUs. 
* `ocpus` - The default number of OCPUs available for this shape. 
* `platform_config_options` - The list of supported platform configuration options for this shape. 
	* `measured_boot_options` - Available and default options for Measured Boot configuration 
		* `allowed_values` - Possible boolean values indicating whether MeasuredBoot can be enabled or disabled 
		* `is_default_enabled` - Indicates whether Measured Boot is to be enabled by default 
	* `numa_nodes_per_socket_platform_options` - Available and default options for NUMA Nodes Per Socket configuration 
		* `allowed_values` - The supported values for this platform configuration property. 
		* `default_value` - Indicates the default NUMA Nodes Per Socket configuration 
	* `secure_boot_options` - Available and default options for Secure Boot configuration 
		* `allowed_values` - Possible boolean values indicating whether SecureBoot can be enabled or disabled 
		* `is_default_enabled` - Indicates whether Secure Boot is to be enabled by default 
	* `trusted_platform_module_options` - Available and default options for Trusted Platform Module (TPM) configuration 
		* `allowed_values` - Possible boolean values indicating whether Trusted Platform Module (TPM) can be enabled or disabled 
		* `is_default_enabled` - Indicates whether Trusted Platform Module is to be enabled by default 
	* `type` - The type of platform being configured. (Supported types=[INTEL_VM, AMD_MILAN_BM, AMD_ROME_BM, INTEL_SKYLAKE_BM]) 
* `processor_description` - A short description of the shape's processor (CPU). 

