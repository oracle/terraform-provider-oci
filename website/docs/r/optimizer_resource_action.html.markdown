---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_resource_action"
sidebar_current: "docs-oci-resource-optimizer-resource_action"
description: |-
  Provides the Resource Action resource in Oracle Cloud Infrastructure Optimizer service
---

# oci_optimizer_resource_action
This resource provides the Resource Action resource in Oracle Cloud Infrastructure Optimizer service.

Updates the resource action that corresponds to the specified OCID. 
Use this operation to implement the following actions:

  * Postpone resource action
  * Ignore resource action
  * Reactivate resource action


## Example Usage

```hcl
resource "oci_optimizer_resource_action" "test_resource_action" {
	#Required
	resource_action_id = oci_optimizer_resource_action.test_resource_action.id
	status = var.resource_action_status

	#Optional
	time_status_end = var.resource_action_time_status_end
}
```

## Argument Reference

The following arguments are supported:

* `resource_action_id` - (Required) The unique OCID associated with the resource action.
* `status` - (Required) (Updatable) The status of the resource action.
* `time_status_end` - (Optional) (Updatable) The date and time the current status will change. The format is defined by RFC3339.

	For example, "The current `postponed` status of the resource action will end and change to `pending` on this date and time." 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - Details about the recommended action. 
	* `description` - Text describing the recommended action.
	* `type` - The status of the resource action.
	* `url` - The URL path to documentation that explains how to perform the action.
* `category_id` - The unique OCID associated with the category.
* `compartment_id` - The OCID of the compartment.
* `compartment_name` - The name associated with the compartment.
* `estimated_cost_saving` - The estimated cost savings, in dollars, for the resource action.
* `extended_metadata` - Additional metadata key/value pairs that you provide. They serve the same purpose and functionality as fields in the `metadata` object.

	They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).

	For example:

	`{"CurrentShape": {"name":"VM.Standard2.16"}, "RecommendedShape": {"name":"VM.Standard2.8"}}` 
* `id` - The unique OCID associated with the resource action.
* `metadata` - Custom metadata key/value pairs for the resource action.

	**Metadata Example**

	"metadata" : { "cpuRecommendedShape": "VM.Standard1.1", "computeMemoryUtilization": "26.05734124418388", "currentShape": "VM.Standard1.2", "instanceRecommendedShape": "VM.Standard1.1", "computeCpuUtilization": "7.930035319720132", "memoryRecommendedShape": "None" } 
* `name` - The name assigned to the resource.
* `recommendation_id` - The unique OCID associated with the recommendation.
* `resource_id` - The unique OCID associated with the resource.
* `resource_type` - The kind of resource.
* `state` - The resource action's current state.
* `status` - The current status of the resource action.
* `time_created` - The date and time the resource action details were created, in the format defined by RFC3339.
* `time_status_begin` - The date and time that the resource action entered its current status. The format is defined by RFC3339.

	For example, "The status of the resource action changed from `pending` to `current(ignored)` on this date and time." 
* `time_status_end` - The date and time the current status will change. The format is defined by RFC3339.

	For example, "The current `postponed` status of the resource action will end and change to `pending` on this date and time." 
* `time_updated` - The date and time the resource action details were last updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Resource Action
	* `update` - (Defaults to 20 minutes), when updating the Resource Action
	* `delete` - (Defaults to 20 minutes), when destroying the Resource Action


## Import

ResourceActions can be imported using the `id`, e.g.

```
$ terraform import oci_optimizer_resource_action.test_resource_action "id"
```

