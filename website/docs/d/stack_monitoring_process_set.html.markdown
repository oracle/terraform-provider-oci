---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_process_set"
sidebar_current: "docs-oci-datasource-stack_monitoring-process_set"
description: |-
  Provides details about a specific Process Set in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_process_set
This data source provides details about a specific Process Set resource in Oracle Cloud Infrastructure Stack Monitoring service.

API to get the details of a Process Set by identifier.

## Example Usage

```hcl
data "oci_stack_monitoring_process_set" "test_process_set" {
	#Required
	process_set_id = oci_stack_monitoring_process_set.test_process_set.id
}
```

## Argument Reference

The following arguments are supported:

* `process_set_id` - (Required) The Process Set ID


## Attributes Reference

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

