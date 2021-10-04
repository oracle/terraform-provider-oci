---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_triggers"
sidebar_current: "docs-oci-datasource-devops-triggers"
description: |-
  Provides the list of Triggers in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_triggers
This data source provides the list of Triggers in Oracle Cloud Infrastructure Devops service.

Returns a list of Triggers.


## Example Usage

```hcl
data "oci_devops_triggers" "test_triggers" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.trigger_display_name
	id = var.trigger_id
	project_id = oci_devops_project.test_project.id
	state = var.trigger_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique Trigger identifier
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only Triggers that matches the given lifecycleState


## Attributes Reference

The following attributes are exported:

* `trigger_collection` - The list of trigger_collection.

### Trigger Reference

The following attributes are exported:

* `actions` - The list of actions that are to be performed for this Trigger
	* `build_pipeline_id` - The id of the build pipeline to be triggered
	* `filter` - The filters for the trigger
		* `events` - The events, example PUSH, PULL_REQUEST_MERGE etc.
		* `include` - Attributes to filter Devops Code Repository events
			* `base_ref` - The target branch for pull requests; not applicable for push
			* `head_ref` - Branch for push event; source branch for pull requests
		* `trigger_source` - Source of the Trigger (allowed values are - GITHUB, GITLAB)
	* `type` - The type of action that will be taken (allowed value - TRIGGER_BUILD_PIPELINE)
* `compartment_id` - Compartment to which the Trigger belongs
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Description about the Trigger
* `display_name` - Name for Trigger.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - Project to which the Trigger belongs
* `repository_id` - The OCID of Oracle Cloud Infrastructure Devops Repository
* `state` - The current state of the Trigger.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the the Trigger was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Trigger was updated. An RFC3339 formatted datetime string
* `trigger_source` - Source of the Trigger (allowed values are - GITHUB, GITLAB)
* `trigger_url` - The endpoint which listens to Trigger events

