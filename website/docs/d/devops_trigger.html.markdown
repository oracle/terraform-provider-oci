---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_trigger"
sidebar_current: "docs-oci-datasource-devops-trigger"
description: |-
  Provides details about a specific Trigger in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_trigger
This data source provides details about a specific Trigger resource in Oracle Cloud Infrastructure Devops service.

Gets a Trigger by identifier

## Example Usage

```hcl
data "oci_devops_trigger" "test_trigger" {
	#Required
	trigger_id = oci_devops_trigger.test_trigger.id
}
```

## Argument Reference

The following arguments are supported:

* `trigger_id` - (Required) unique Trigger identifier


## Attributes Reference

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

