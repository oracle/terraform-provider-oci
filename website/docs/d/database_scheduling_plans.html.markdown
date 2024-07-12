---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_plans"
sidebar_current: "docs-oci-datasource-database-scheduling_plans"
description: |-
  Provides the list of Scheduling Plans in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduling_plans
This data source provides the list of Scheduling Plans in Oracle Cloud Infrastructure Database service.

Lists the Scheduling Plan resources in the specified compartment.


## Example Usage

```hcl
data "oci_database_scheduling_plans" "test_scheduling_plans" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.scheduling_plan_display_name
	id = var.scheduling_plan_id
	resource_id = oci_cloud_guard_resource.test_resource.id
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
	state = var.scheduling_plan_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `id` - (Optional) A filter to return only resources that match the given Schedule Plan id exactly.
* `resource_id` - (Optional) A filter to return only resources that match the given resource id exactly.
* `scheduling_policy_id` - (Optional) A filter to return only resources that match the given scheduling policy id exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `scheduling_plan_collection` - The list of scheduling_plan_collection.

### SchedulingPlan Reference

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

