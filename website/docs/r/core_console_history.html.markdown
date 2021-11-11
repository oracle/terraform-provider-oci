---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_console_history"
sidebar_current: "docs-oci-resource-core-console_history"
description: |-
  Provides the Console History resource in Oracle Cloud Infrastructure Core service
---

# oci_core_console_history
This resource provides the Console History resource in Oracle Cloud Infrastructure Core service.

Captures the most recent serial console data (up to a megabyte) for the
specified instance.

The `CaptureConsoleHistory` operation works with the other console history operations
as described below.

1. Use `CaptureConsoleHistory` to request the capture of up to a megabyte of the
most recent console history. This call returns a `ConsoleHistory`
object. The object will have a state of REQUESTED.
2. Wait for the capture operation to succeed by polling `GetConsoleHistory` with
the identifier of the console history metadata. The state of the
`ConsoleHistory` object will go from REQUESTED to GETTING-HISTORY and
then SUCCEEDED (or FAILED).
3. Use `GetConsoleHistoryContent` to get the actual console history data (not the
metadata).
4. Optionally, use `DeleteConsoleHistory` to delete the console history metadata
and the console history data.


## Example Usage

```hcl
resource "oci_core_console_history" "test_console_history" {
	#Required
	instance_id = oci_core_instance.test_instance.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.console_history_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_id` - (Required) The OCID of the instance to get the console history from.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Console History
	* `update` - (Defaults to 20 minutes), when updating the Console History
	* `delete` - (Defaults to 20 minutes), when destroying the Console History


## Import

ConsoleHistories can be imported using the `id`, e.g.

```
$ terraform import oci_core_console_history.test_console_history "id"
```

