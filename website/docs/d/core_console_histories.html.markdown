---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_console_histories"
sidebar_current: "docs-oci-datasource-core-console_histories"
description: |-
  Provides the list of Console Histories in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_console_histories
This data source provides the list of Console Histories in Oracle Cloud Infrastructure Core service.

Lists the console history metadata for the specified compartment or instance.


## Example Usage

```hcl
data "oci_core_console_histories" "test_console_histories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.console_history_availability_domain
	instance_id = oci_core_instance.test_instance.id
	state = var.console_history_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `console_histories` - The list of console_histories.

### ConsoleHistory Reference

The following attributes are exported:

* `availability_domain` - The availability domain of an instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the console history metadata object.
* `instance_id` - The OCID of the instance this console history was fetched from.
* `state` - The current state of the console history.
* `time_created` - The date and time the history was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

