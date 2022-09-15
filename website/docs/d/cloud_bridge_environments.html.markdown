---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_environments"
sidebar_current: "docs-oci-datasource-cloud_bridge-environments"
description: |-
  Provides the list of Environments in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_environments
This data source provides the list of Environments in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of source environments.


## Example Usage

```hcl
data "oci_cloud_bridge_environments" "test_environments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.environment_display_name
	environment_id = oci_cloud_bridge_environment.test_environment.id
	state = var.environment_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `environment_id` - (Optional) A filter to return only resources that match the given environment ID.
* `state` - (Optional) A filter to return only resources where their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `environment_collection` - The list of environment_collection.

### Environment Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Environment identifier, which can be renamed.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the source environment.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the source environment was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the source environment was updated. An RFC3339 formatted datetime string.

