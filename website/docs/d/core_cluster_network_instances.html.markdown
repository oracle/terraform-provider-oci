---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cluster_network_instances"
sidebar_current: "docs-oci-datasource-core-cluster_network_instances"
description: |-
  Provides the list of Cluster Network Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cluster_network_instances
This data source provides the list of Cluster Network Instances in Oracle Cloud Infrastructure Core service.

Lists the instances in the specified cluster network.

## Example Usage

```hcl
data "oci_core_cluster_network_instances" "test_cluster_network_instances" {
	#Required
	cluster_network_id = oci_core_cluster_network.test_cluster_network.id
	compartment_id = var.compartment_id

	#Optional
	display_name = var.cluster_network_instance_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster network.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 


## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

### ClusterNetworkInstance Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - The fault domain the instance is running in.
* `id` - The OCID of the instance.
* `instance_configuration_id` - The OCID of the instance confgiuration used to create the instance.
* `load_balancer_backends` - The load balancer backends that are configured for the instance pool instance. 
	* `backend_health_status` - The health of the backend as observed by the load balancer.
	* `backend_name` - The name of the backend in the backend set.
	* `backend_set_name` - The name of the backend set on the load balancer.
	* `load_balancer_id` - The OCID of the load balancer attached to the instance pool.
* `region` - The region that contains the availability domain the instance is running in.
* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `state` - The current state of the instance pool instance.
* `time_created` - The date and time the instance pool instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

