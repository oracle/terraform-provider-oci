---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_project"
sidebar_current: "docs-oci-datasource-ai_anomaly_detection-project"
description: |-
  Provides details about a specific Project in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# Data Source: oci_ai_anomaly_detection_project
This data source provides details about a specific Project resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Gets a Project by identifier

## Example Usage

```hcl
data "oci_ai_anomaly_detection_project" "test_project" {
	#Required
	project_id = oci_ai_anomaly_detection_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The OCID of the Project.


## Attributes Reference

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

