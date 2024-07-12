---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_plan"
sidebar_current: "docs-oci-datasource-database-scheduling_plan"
description: |-
  Provides details about a specific Scheduling Plan in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduling_plan
This data source provides details about a specific Scheduling Plan resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Scheduling Plan.


## Example Usage

```hcl
data "oci_database_scheduling_plan" "test_scheduling_plan" {
	#Required
	scheduling_plan_id = oci_database_scheduling_plan.test_scheduling_plan.id
}
```

## Argument Reference

The following arguments are supported:

* `scheduling_plan_id` - (Required) The Schedule Plan [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The display name of the Scheduling Plan.
* `estimated_time_in_mins` - The estimated time for the Scheduling Plan.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
* `is_using_recommended_scheduled_actions` - If true, recommended scheduled actions will be generated for the scheduling plan.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `plan_intent` - The current intent the Scheduling Plan. Valid states is EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE. 
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `scheduling_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
* `service_type` - The service type of the Scheduling Plan.
* `state` - The current state of the Scheduling Plan. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Scheduling Plan Resource was created.
* `time_updated` - The date and time the Scheduling Plan Resource was updated.

