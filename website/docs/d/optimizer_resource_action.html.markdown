---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_resource_action"
sidebar_current: "docs-oci-datasource-optimizer-resource_action"
description: |-
  Provides details about a specific Resource Action in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_resource_action
This data source provides details about a specific Resource Action resource in Oracle Cloud Infrastructure Optimizer service.

Gets the resource action that corresponds to the specified OCID.


## Example Usage

```hcl
data "oci_optimizer_resource_action" "test_resource_action" {
	#Required
	resource_action_id = oci_optimizer_resource_action.test_resource_action.id

	#Optional
	include_resource_metadata = var.resource_action_include_resource_metadata
}
```

## Argument Reference

The following arguments are supported:

* `include_resource_metadata` - (Optional) Supplement additional resource information in extended metadata response.
* `resource_action_id` - (Required) The unique OCID associated with the resource action.


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

