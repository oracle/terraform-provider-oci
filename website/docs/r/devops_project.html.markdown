---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_project"
sidebar_current: "docs-oci-resource-devops-project"
description: |-
  Provides the Project resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_project
This resource provides the Project resource in Oracle Cloud Infrastructure Devops service.

Creates a new project.

## Example Usage

```hcl
resource "oci_devops_project" "test_project" {
	#Required
	compartment_id = var.compartment_id
	name = var.project_name
	notification_config {
		#Required
		topic_id = oci_ons_notification_topic.test_notification_topic.id
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.project_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the project is created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) Project description.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `name` - (Required) Project name (case-sensitive).
* `notification_config` - (Required) (Updatable) Notification configuration for the project.
	* `topic_id` - (Required) (Updatable) The topic ID for notifications.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where the project is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Project description.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - Project name (case-sensitive).
* `namespace` - Namespace associated with the project.
* `notification_config` - Notification configuration for the project.
	* `topic_id` - The topic ID for notifications.
* `state` - The current state of the project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the project was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the project was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import oci_devops_project.test_project "id"
```

