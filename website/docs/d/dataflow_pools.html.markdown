---
subcategory: "Data Flow"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataflow_pools"
sidebar_current: "docs-oci-datasource-dataflow-pools"
description: |-
  Provides the list of Pools in Oracle Cloud Infrastructure Data Flow service
---

# Data Source: oci_dataflow_pools
This data source provides the list of Pools in Oracle Cloud Infrastructure Data Flow service.

Lists all pools in the specified compartment. The query must include compartmentId. The query may also include one other parameter. If the query does not include compartmentId, or includes compartmentId, but with two or more other parameters, an error is returned.


## Example Usage

```hcl
data "oci_dataflow_pools" "test_pools" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.pool_display_name
	display_name_starts_with = var.pool_display_name_starts_with
	owner_principal_id = oci_dataflow_owner_principal.test_owner_principal.id
	state = var.pool_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment. 
* `display_name` - (Optional) The query parameter for the Spark application name. 
* `display_name_starts_with` - (Optional) The displayName prefix. 
* `owner_principal_id` - (Optional) The OCID of the user who created the resource. 
* `state` - (Optional) The LifecycleState of the pool. 


## Attributes Reference

The following attributes are exported:

* `pool_collection` - The list of pool_collection.

### Pool Reference

The following attributes are exported:

* `compartment_id` - The OCID of a compartment. 
* `configurations` - List of PoolConfig items. 
	* `max` - Maximum number of compute instances in the pool for a given compute shape. 
	* `min` - Minimum number of compute instances in the pool for a given compute shape. 
	* `shape` - The compute shape of the resources you would like to provision. 
	* `shape_config` - This is used to configure the shape of the driver or executor if a flexible shape is used. 
		* `memory_in_gbs` - The amount of memory used for the driver or executors. 
		* `ocpus` - The total number of OCPUs used for the driver or executors. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user-friendly description. Avoid entering confidential information. 
* `display_name` - A user-friendly name. It does not have to be unique. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of a pool. Unique Id to indentify a dataflow pool resource. 
* `idle_timeout_in_minutes` - Optional timeout value in minutes used to auto stop Pools. A Pool will be auto stopped after inactivity for this amount of time period. If value not set, pool will not be auto stopped auto. 
* `lifecycle_details` - The detailed messages about the lifecycle state. 
* `owner_principal_id` - The OCID of the user who created the resource. 
* `owner_user_name` - The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead. 
* `pool_metrics` - A collection of metrics related to a particular pool. 
	* `active_runs_count` - The number of runs that are currently running that are using this pool. 
	* `actively_used_node_count` - A count of the nodes that are currently being used for each shape in this pool. 
		* `logical_shape` - The compute shape of the nodes that the count is for. 
		* `pool_count` - The node count of this compute shape. 
	* `time_last_metrics_updated` - The last time the mertics were updated for this. 
	* `time_last_started` - The last time this pool was started. 
	* `time_last_stopped` - The last time this pool was stopped. 
	* `time_last_used` - The last time a run used this pool. 
* `schedules` - A list of schedules for pool to auto start and stop. 
	* `day_of_week` - Day of the week SUN-SAT 
	* `start_time` - Hour of the day to start or stop pool.
	* `stop_time` - Hour of the day to stop the pool.
* `state` - The current state of this pool. 
* `time_created` - The date and time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `time_updated` - The date and time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 

