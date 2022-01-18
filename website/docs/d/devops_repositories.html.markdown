---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repositories"
sidebar_current: "docs-oci-datasource-devops-repositories"
description: |-
  Provides the list of Repositories in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repositories
This data source provides the list of Repositories in Oracle Cloud Infrastructure Devops service.

Returns a list of repositories given a compartment ID or a project ID.


## Example Usage

```hcl
data "oci_devops_repositories" "test_repositories" {

	#Optional
	compartment_id = var.compartment_id
	name = var.repository_name
	project_id = oci_devops_project.test_project.id
	repository_id = oci_devops_repository.test_repository.id
	state = var.repository_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `project_id` - (Optional) unique project identifier
* `repository_id` - (Optional) Unique repository identifier.
* `state` - (Optional) A filter to return only resources whose lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `repository_collection` - The list of repository_collection.

### Repository Reference

The following attributes are exported:

* `branch_count` - The count of the branches present in the repository.
* `commit_count` - The count of the commits present in the repository.
* `compartment_id` - The OCID of the repository's compartment.
* `default_branch` - The default branch of the repository.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Details of the repository. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `http_url` - HTTP URL that you use to git clone, pull and push.
* `id` - The OCID of the repository. This value is unique and immutable.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `mirror_repository_config` - Configuration information for mirroring the repository.
	* `connector_id` - Upstream git repository connection identifer.
	* `repository_url` - URL of external repository you want to mirror.
	* `trigger_schedule` - Specifies a trigger schedule. Timing information for when to initiate automated syncs.
		* `custom_schedule` - Valid if type is CUSTOM. Following RFC 5545 recurrence rules, we can specify starting time, occurrence frequency, and interval size. Example for frequency could be DAILY/WEEKLY/HOURLY or any RFC 5545 supported frequency, which is followed by start time of this window.  You can control the start time with BYHOUR, BYMINUTE and BYSECONDS. It is followed by the interval size. 
		* `schedule_type` - Different types of trigger schedule: None - No automated synchronization schedule. Default - Trigger schedule is every 30 minutes. Custom - Custom triggering schedule. 
* `name` - Unique name of a repository. This value is mutable.
* `namespace` - Tenancy unique namespace.
* `project_id` - The OCID of the DevOps project containing the repository.
* `project_name` - Unique project name in a namespace.
* `repository_type` - Type of repository: Mirrored - Repository created by mirroring an existing repository. Hosted - Repository created and hosted using Oracle Cloud Infrastructure DevOps code repository. 
* `size_in_bytes` - The size of the repository in bytes.
* `ssh_url` - SSH URL that you use to git clone, pull and push.
* `state` - The current state of the repository.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the repository was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the repository was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `trigger_build_events` - Trigger build events supported for this repository: Push - Build is triggered when a push event occurs. Commit updates - Build is triggered when new commits are mirrored into a repository. 

