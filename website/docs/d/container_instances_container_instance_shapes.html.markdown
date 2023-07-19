---
subcategory: "Container Instances"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_container_instances_container_instance_shapes"
sidebar_current: "docs-oci-datasource-container_instances-container_instance_shapes"
description: |-
  Provides the list of Container Instance Shapes in Oracle Cloud Infrastructure Container Instances service
---

# Data Source: oci_container_instances_container_instance_shapes
This data source provides the list of Container Instance Shapes in Oracle Cloud Infrastructure Container Instances service.

Lists the shapes that can be used to create container instances.

## Example Usage

```hcl
data "oci_container_instances_container_instance_shapes" "test_container_instance_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.container_instance_shape_availability_domain
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `container_instance_shape_collection` - The list of container_instance_shape_collection.

### ContainerInstanceShape Reference

The following attributes are exported:

* `items` - A list of shapes.
	* `memory_options` - For a flexible shape, the amount of memory available for container instances that use this shape. 
		* `default_per_ocpu_in_gbs` - The default amount of memory per OCPU available for this shape (GB). 
		* `max_in_gbs` - The maximum amount of memory (GB). 
		* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape (GB). 
		* `min_in_gbs` - The minimum amount of memory (GB). 
		* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape (GB). 
	* `name` - The name identifying the shape. 
	* `networking_bandwidth_options` - For a flexible shape, the amount of networking bandwidth available for container instances that use this shape. 
		* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second. 
		* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second. 
		* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second. 
	* `ocpu_options` - For a flexible shape, the number of OCPUs available for container instances that use this shape. 
		* `max` - The maximum number of OCPUs. 
		* `min` - The minimum number of OCPUs. 
	* `processor_description` - A short description of the container instance's processor (CPU). 

