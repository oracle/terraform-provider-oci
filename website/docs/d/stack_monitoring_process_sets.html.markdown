---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_process_sets"
sidebar_current: "docs-oci-datasource-stack_monitoring-process_sets"
description: |-
  Provides the list of Process Sets in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_process_sets
This data source provides the list of Process Sets in Oracle Cloud Infrastructure Stack Monitoring service.

API to get the details of all Process Sets.

## Example Usage

```hcl
data "oci_stack_monitoring_process_sets" "test_process_sets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.process_set_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which data is listed.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `process_set_collection` - The list of process_set_collection.

### ProcessSet Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the Process Set.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Process Set. 
* `revision` - The current revision of the Process Set.
* `specification` - Collection of regular expression specifications used to identify the processes to be monitored.
	* `items` - List of Process Set specification details.
		* `label` - Optional label used to identify a single filter.
		* `process_command` - String literal used for exact matching on process name.
		* `process_line_regex_pattern` - Regex pattern matching on process arguments.
		* `process_user` - String literal used for exact matching on process user.
* `state` - The current state of the Resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the process set was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the process set was last updated. An RFC3339 formatted datetime string.

