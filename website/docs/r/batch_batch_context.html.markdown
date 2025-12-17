---
subcategory: "Batch"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_batch_batch_context"
sidebar_current: "docs-oci-resource-batch-batch_context"
description: |-
  Provides the Batch Context resource in Oracle Cloud Infrastructure Batch service
---

# oci_batch_batch_context
This resource provides the Batch Context resource in Oracle Cloud Infrastructure Batch service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/batch

Creates a batch context.

## Example Usage

```hcl
resource "oci_batch_batch_context" "test_batch_context" {
	#Required
	compartment_id = var.compartment_id
	fleets {
		#Required
		max_concurrent_tasks = var.batch_context_fleets_max_concurrent_tasks
		name = var.batch_context_fleets_name
		shape {
			#Required
			memory_in_gbs = var.batch_context_fleets_shape_memory_in_gbs
			ocpus = var.batch_context_fleets_shape_ocpus
			shape_name = oci_core_shape.test_shape.name
		}
		type = var.batch_context_fleets_type
	}
	network {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		nsg_ids = var.batch_context_network_nsg_ids
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.batch_context_description
	display_name = var.batch_context_display_name
	entitlements = var.batch_context_entitlements
	freeform_tags = {"Department"= "Finance"}
	job_priority_configurations {
		#Required
		tag_key = var.batch_context_job_priority_configurations_tag_key
		tag_namespace = var.batch_context_job_priority_configurations_tag_namespace
		values = var.batch_context_job_priority_configurations_values
		weight = var.batch_context_job_priority_configurations_weight
	}
	logging_configuration {
		#Required
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_apm_traces_log.test_log.id
		type = var.batch_context_logging_configuration_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Summarized information about the batch context.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. If not specified or provided as null or empty string, it will be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchcontext20250914115623. 
* `entitlements` - (Optional) (Updatable) Mapping of concurrent/shared resources used in job tasks to their limits.
* `fleets` - (Required) List of fleet configurations related to the batch context.
	* `max_concurrent_tasks` - (Required) Maximum number of concurrent tasks for the service managed fleet.
	* `name` - (Required) Name of the service managed fleet.
	* `shape` - (Required) Shape of the fleet. Describes hardware resources of each node in the fleet.
		* `memory_in_gbs` - (Required) Amount of memory in GBs required by the shape.
		* `ocpus` - (Required) Number of OCPUs required by the shape.
		* `shape_name` - (Required) The name of the shape.
	* `type` - (Required) Type of the fleet. Also serves as a discriminator for sub-entities.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `job_priority_configurations` - (Optional) (Updatable) List of job priority configurations related to the batch context.
	* `tag_key` - (Required) (Updatable) Name of the tag key.
	* `tag_namespace` - (Required) (Updatable) Name of the corresponding tag namespace.
	* `values` - (Required) (Updatable) Mapping of tag value to its priority.
	* `weight` - (Required) (Updatable) Weight associated with the tag key. Percentage point is the unit of measurement.
* `logging_configuration` - (Optional) Logging configuration for batch context.
	* `log_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
	* `log_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `type` - (Required) Discriminator for sub-entities.
* `network` - (Required) Network configuration of the batch context.
	* `nsg_ids` - (Optional) A list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated network security groups. 
	* `subnet_id` - (Required) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated subnet. 
* `state` - (Optional) (Updatable) The target state for the Batch Context. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Batch Context
	* `update` - (Defaults to 20 minutes), when updating the Batch Context
	* `delete` - (Defaults to 20 minutes), when destroying the Batch Context


## Import

BatchContexts can be imported using the `id`, e.g.

```
$ terraform import oci_batch_batch_context.test_batch_context "id"
```

