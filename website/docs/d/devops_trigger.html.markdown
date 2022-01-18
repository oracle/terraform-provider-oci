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

Retrieves a trigger by identifier.

## Example Usage

```hcl
data "oci_devops_trigger" "test_trigger" {
	#Required
	trigger_id = oci_devops_trigger.test_trigger.id
}
```

## Argument Reference

The following arguments are supported:

* `trigger_id` - (Required) Unique trigger identifier.


## Attributes Reference

The following attributes are exported:

* `actions` - The list of actions that are to be performed for this trigger.
	* `build_pipeline_id` - The OCID of the build pipeline to be triggered.
	* `filter` - The filters for the trigger.
		* `events` - The events, for example, PUSH, PULL_REQUEST_MERGE.
		* `include` - Attributes to filter DevOps code repository events.
			* `base_ref` - The target branch for pull requests; not applicable for push requests.
			* `head_ref` - Branch for push event; source branch for pull requests.
		* `trigger_source` - Source of the trigger. Allowed values are, GITHUB and GITLAB.
	* `type` - The type of action that will be taken. Allowed value is TRIGGER_BUILD_PIPELINE.
* `compartment_id` - The OCID of the compartment that contains the trigger.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Description about the trigger.
* `display_name` - Trigger display name. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of the DevOps project to which the trigger belongs to.
* `repository_id` - The OCID of the DevOps code repository.
* `state` - The current state of the trigger.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the trigger was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the trigger was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `trigger_source` - Source of the trigger. Allowed values are, GITHUB, GITLAB and DEVOPS_CODE_REPOSITORY.
* `trigger_url` - The endpoint that listens to trigger events.

