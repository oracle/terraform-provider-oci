---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_policy"
sidebar_current: "docs-oci-resource-database-scheduling_policy"
description: |-
  Provides the Scheduling Policy resource in Oracle Cloud Infrastructure Database service
---

# oci_database_scheduling_policy
This resource provides the Scheduling Policy resource in Oracle Cloud Infrastructure Database service.

Creates a Scheduling Policy resource.


## Example Usage

```hcl
resource "oci_database_scheduling_policy" "test_scheduling_policy" {
	#Required
	cadence = var.scheduling_policy_cadence
	compartment_id = var.compartment_id
	display_name = var.scheduling_policy_display_name

	#Optional
	cadence_start_month {
		#Required
		name = var.scheduling_policy_cadence_start_month_name
	}
	defined_tags = var.scheduling_policy_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cadence` - (Required) (Updatable) The cadence period.
* `cadence_start_month` - (Optional) (Updatable) Start of the month to be followed during the cadence period.
	* `name` - (Required) (Updatable) Name of the month of the year.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the Scheduling Policy. The name does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cadence` - The cadence period.
* `cadence_start_month` - Start of the month to be followed during the cadence period.
	* `name` - Name of the month of the year.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Scheduling Policy. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the Scheduling Policy. Valid states are CREATING, NEEDS_ATTENTION, ACTIVE, UPDATING, FAILED, DELETING and DELETED. 
* `time_created` - The date and time the Scheduling Policy was created.
* `time_next_window_starts` - The date and time of the next scheduling window associated with the schedulingPolicy is planned to start.
* `time_updated` - The last date and time that the Scheduling Policy was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduling Policy
	* `update` - (Defaults to 20 minutes), when updating the Scheduling Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduling Policy


## Import

SchedulingPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_database_scheduling_policy.test_scheduling_policy "id"
```

