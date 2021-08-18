---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_projects"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-projects"
description: |-
  Provides the list of Projects in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_projects
This data source provides the list of Projects in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Returns a list of  Projects.


## Example Usage

```hcl
data "oci_ai_anomaly_detection_projects" "test_projects" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.project_display_name
	state = var.project_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `project_collection` - The list of project_collection.

### Project Reference

The following attributes are exported:

* `compartment_id` - The OCID for the project's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the project.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the project that is immutable on creation.
* `state` - The lifecycle state of the Project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

