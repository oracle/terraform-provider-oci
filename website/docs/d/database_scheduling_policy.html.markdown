---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_policy"
sidebar_current: "docs-oci-datasource-database-scheduling_policy"
description: |-
  Provides details about a specific Scheduling Policy in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduling_policy
This data source provides details about a specific Scheduling Policy resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Scheduling Policy.


## Example Usage

```hcl
data "oci_database_scheduling_policy" "test_scheduling_policy" {
	#Required
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `scheduling_policy_id` - (Required) The Scheduling Policy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

