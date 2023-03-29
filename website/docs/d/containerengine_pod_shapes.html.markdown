---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_pod_shapes"
sidebar_current: "docs-oci-datasource-containerengine-pod_shapes"
description: |-
  Provides the list of Pod Shapes in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_pod_shapes
This data source provides the list of Pod Shapes in Oracle Cloud Infrastructure Container Engine service.

List all the Pod Shapes in a compartment.

## Example Usage

```hcl
data "oci_containerengine_pod_shapes" "test_pod_shapes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.pod_shape_availability_domain
	name = var.pod_shape_name
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The availability domain of the pod shape.
* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) The name to filter on.


## Attributes Reference

The following attributes are exported:

* `pod_shapes` - The list of pod_shapes.

### PodShape Reference

The following attributes are exported:

* `memory_options` - ShapeMemoryOptions.
	* `default_per_ocpu_in_gbs` - The default amount of memory per OCPU available for this shape, in gigabytes.
	* `max_in_gbs` - The maximum amount of memory, in gigabytes.
	* `max_per_ocpu_in_gbs` - The maximum amount of memory per OCPU available for this shape, in gigabytes.
	* `min_in_gbs` - The minimum amount of memory, in gigabytes.
	* `min_per_ocpu_in_gbs` - The minimum amount of memory per OCPU available for this shape, in gigabytes.
* `name` - The name of the identifying shape.
* `network_bandwidth_options` - ShapeNetworkBandwidthOptions.
	* `default_per_ocpu_in_gbps` - The default amount of networking bandwidth per OCPU, in gigabits per second.
	* `max_in_gbps` - The maximum amount of networking bandwidth, in gigabits per second.
	* `min_in_gbps` - The minimum amount of networking bandwidth, in gigabits per second.
* `ocpu_options` - Options for OCPU shape.
	* `max` - The maximum number of OCPUs.
	* `min` - The minimum number of OCPUs.
* `processor_description` - A short description of the VM's processor (CPU).

