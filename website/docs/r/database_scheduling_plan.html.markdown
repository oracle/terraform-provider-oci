---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_plan"
sidebar_current: "docs-oci-resource-database-scheduling_plan"
description: |-
  Provides the Scheduling Plan resource in Oracle Cloud Infrastructure Database service
---

# oci_database_scheduling_plan
This resource provides the Scheduling Plan resource in Oracle Cloud Infrastructure Database service.

Creates a Scheduling Plan resource.


## Example Usage

```hcl
resource "oci_database_scheduling_plan" "test_scheduling_plan" {
	#Required
	compartment_id = var.compartment_id
	resource_id = oci_cloud_guard_resource.test_resource.id
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
	service_type = var.scheduling_plan_service_type

	#Optional
	defined_tags = var.scheduling_plan_defined_tags
	freeform_tags = {"Department"= "Finance"}
	is_using_recommended_scheduled_actions = var.scheduling_plan_is_using_recommended_scheduled_actions
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_using_recommended_scheduled_actions` - (Optional) If true, recommended scheduled actions will be generated for the scheduling plan.
* `resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `scheduling_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
* `service_type` - (Required) The service type of the Scheduling Plan.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduling Plan
	* `update` - (Defaults to 20 minutes), when updating the Scheduling Plan
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduling Plan


## Import

SchedulingPlans can be imported using the `id`, e.g.

```
$ terraform import oci_database_scheduling_plan.test_scheduling_plan "id"
```

