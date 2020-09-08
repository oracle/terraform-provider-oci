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

* `gpu_description` - A short description of the graphics processing unit (GPU) available for this shape.

	If the shape does not have any GPUs, this field is `null`. 
* `gpus` - The number of GPUs available for this shape. 
* `local_disk_description` - A short description of the local disks available for this shape.

	If the shape does not have any local disks, this field is `null`. 
* `local_disks` - The number of local disks available for this shape. 
* `local_disks_total_size_in_gbs` - The aggregate size of the local disks available for this shape, in gigabytes.

	If the shape does not have any local disks, this field is `null`. 
* `max_vnic_attachment_options` - 
	* `default_per_ocpu` - The default number of VNIC attachments allowed per OCPU. 
	* `max` - The highest maximum value of VNIC attachments. 
	* `min` - The lowest maximum value of VNIC attachments. 
* `max_vnic_attachments` - The maximum number of VNIC attachments available for this shape. 
* `memory_in_gbs` - The default amount of memory available for this shape, in gigabytes. 
* `memory_options` - 
	* `default_per_ocpu_in_gbs` - The default amount of memory per OCPU available for this shape, in gigabytes. 
	* `max_in_gbs` - The maximum amount of memory, in gigabytes. 
	* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape, in gigabytes. 
	* `min_in_gbs` - The minimum amount of memory, in gigabytes. 
	* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape, in gigabytes. 
* `name` - The name of the shape. You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
* `networking_bandwidth_in_gbps` - The networking bandwidth available for this shape, in gigabits per second. 
* `networking_bandwidth_options` - 
	* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second. 
	* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second. 
	* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second. 
* `ocpu_options` - 
	* `max` - The maximum number of OCPUs. 
	* `min` - The minimum number of OCPUs. 
* `ocpus` - The default number of OCPUs available for this shape. 
* `processor_description` - A short description of the shape's processor (CPU). 

