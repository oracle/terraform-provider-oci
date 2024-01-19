---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_process_set"
sidebar_current: "docs-oci-resource-stack_monitoring-process_set"
description: |-
  Provides the Process Set resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_process_set
This resource provides the Process Set resource in Oracle Cloud Infrastructure Stack Monitoring service.

API to create Process Set.

## Example Usage

```hcl
resource "oci_stack_monitoring_process_set" "test_process_set" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.process_set_display_name
	specification {
		#Required
		items {

			#Optional
			label = var.process_set_specification_items_label
			process_command = var.process_set_specification_items_process_command
			process_line_regex_pattern = var.process_set_specification_items_process_line_regex_pattern
			process_user = var.process_set_specification_items_process_user
		}
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Name of the Process Set.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `specification` - (Required) (Updatable) Collection of regular expression specifications used to identify the processes to be monitored.
	* `items` - (Required) (Updatable) List of Process Set specification details.
		* `label` - (Optional) (Updatable) Optional label used to identify a single filter.
		* `process_command` - (Optional) (Updatable) String literal used for exact matching on process name.
		* `process_line_regex_pattern` - (Optional) (Updatable) Regex pattern matching on process arguments.
		* `process_user` - (Optional) (Updatable) String literal used for exact matching on process user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Process Set
	* `update` - (Defaults to 20 minutes), when updating the Process Set
	* `delete` - (Defaults to 20 minutes), when destroying the Process Set


## Import

ProcessSets can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_process_set.test_process_set "id"
```

