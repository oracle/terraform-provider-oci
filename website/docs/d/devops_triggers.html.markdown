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

Returns a list of triggers.


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
* `id` - (Optional) Unique trigger identifier.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only triggers that matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `trigger_collection` - The list of trigger_collection.

### Trigger Reference

The following attributes are exported:

* `actions` - The list of actions that are to be performed for this trigger.
	* `build_pipeline_id` - The OCID of the build pipeline to be triggered.
	* `trigger_filter` - The filters for the trigger.
		* `events` - The events, for example, PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED.
		* `exclude` - Attributes to filter GitLab self-hosted server events. File filter criteria - Changes only affecting excluded files will not invoke a build. if both include and exclude filter are used then exclusion filter will be applied on the result set of inclusion filter.
			* `file_filter` - Attributes to support include/exclude files for triggering build runs.
				* `file_paths` - The file paths/glob pattern for files.
		* `include` - Attributes to filter GitLab self-hosted server events.
			* `base_ref` - The target branch for pull requests; not applicable for push requests.
			* `head_ref` - Branch for push event; source branch for pull requests.
			* `repository_name` - The repository name for trigger events.
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
* `trigger_source` - Source of the trigger. Allowed values are, GITHUB and GITLAB. 
* `trigger_url` - The endpoint that listens to trigger events.

