---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_contexts"
sidebar_current: "docs-oci-datasource-batch-batch_contexts"
description: |-
  Provides the list of Batch Contexts in Oracle Cloud Infrastructure Batch service
---

# Data Source: oci_batch_batch_contexts
This data source provides the list of Batch Contexts in Oracle Cloud Infrastructure Batch service.

Lists the batch contexts by compartment or context [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). You can filter and sort them by various properties like lifecycle state, name and also ocid. All properties require an exact match. List operation only provides a summary information, use GetBatchContext to get the full details on a specific context

## Example Usage

```hcl
data "oci_batch_batch_contexts" "test_batch_contexts" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.batch_context_display_name
	id = var.batch_context_id
	state = var.batch_context_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch context.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `batch_context_collection` - The list of batch_context_collection.

### BatchContext Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Summarized information about the batch context.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `entitlements` - Mapping of concurrent/shared resources used in job tasks to their limits.
* `fleets` - List of fleet configurations related to the batch context.
	* `details` - A message that describes the current state of the service manage fleet configuration in more detail.
	* `max_concurrent_tasks` - Maximum number of concurrent tasks for the service managed fleet.
	* `name` - Name of the service managed fleet.
	* `shape` - Shape of the fleet. Describes hardware resources of each node in the fleet.
		* `memory_in_gbs` - Amount of memory in GBs required by the shape.
		* `ocpus` - Number of OCPUs required by the shape.
		* `shape_name` - The name of the shape.
	* `state` - Current state of the service manage fleet configuration.
	* `type` - Type of the fleet. Also serves as a discriminator for sub-entities.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch context.
* `job_priority_configurations` - List of job priority configurations related to the batch context.
	* `tag_key` - Name of the tag key.
	* `tag_namespace` - Name of the corresponding tag namespace.
	* `values` - Mapping of tag value to its priority.
	* `weight` - Weight associated with the tag key. Percentage point is the unit of measurement.
* `lifecycle_details` - A message that describes the current state in more detail. For example,   can be used to provide actionable information for a resource in the Failed state. 
* `logging_configuration` - Logging configuration for batch context.
	* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `type` - Discriminator for sub-entities.
* `network` - Network configuration of the batch context.
	* `nsg_ids` - A list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated network security groups. 
	* `subnet_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated subnet. 
	* `vnics` - A list of private endpoint's VNICs. 
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint's VNIC, which resides in the customer's VCN. 
		* `source_ips` - A list of private IP addresses (in the customer's VCN) that represent access points for the service. 
* `state` - The current state of the batch context. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the batch context was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the batch context was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

