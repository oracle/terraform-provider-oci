---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_automatic_dr_configuration"
sidebar_current: "docs-oci-resource-disaster_recovery-automatic_dr_configuration"
description: |-
  Provides the Automatic Dr Configuration resource in Oracle Cloud Infrastructure Disaster Recovery service
---

# oci_disaster_recovery_automatic_dr_configuration
This resource provides the Automatic Dr Configuration resource in Oracle Cloud Infrastructure Disaster Recovery service.

Create a Automatic DR configuration.

## Example Usage

```hcl
resource "oci_disaster_recovery_automatic_dr_configuration" "test_automatic_dr_configuration" {
	#Required
	display_name = var.automatic_dr_configuration_display_name
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id
	members {
		#Required
		member_id = oci_disaster_recovery_member.test_member.id
		member_type = var.automatic_dr_configuration_members_member_type

		#Optional
		is_auto_failover_enabled = var.automatic_dr_configuration_members_is_auto_failover_enabled
		is_auto_switchover_enabled = var.automatic_dr_configuration_members_is_auto_switchover_enabled
	}

	#Optional
	default_failover_dr_plan_id = oci_disaster_recovery_dr_plan.test_dr_plan.id
	default_switchover_dr_plan_id = oci_disaster_recovery_dr_plan.test_dr_plan.id
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `default_failover_dr_plan_id` - (Optional) (Updatable) The unique id of a Failover DR Plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `default_switchover_dr_plan_id` - (Optional) (Updatable) The unique id of a Switchover DR Plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) The display name of the Automatic DR configuration being created.  Example: `Automatic DR Configuration` 
* `dr_protection_group_id` - (Required) The OCID of the DR protection group to which this Automatic DR configuration belongs.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `members` - (Required) (Updatable) A list of members for Automatic DR configuration. 
	* `is_auto_failover_enabled` - (Optional) (Updatable) A flag indicating if the automatic failover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.  Example: `false` 
	* `is_auto_switchover_enabled` - (Optional) (Updatable) A flag indicating if the automatic switchover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.  Example: `false` 
	* `member_id` - (Required) (Updatable) The OCID of the member.  Example: `ocid1.database.oc1..uniqueID` 
	* `member_type` - (Required) (Updatable) The type of the member. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the Automatic DR configuration.  Example: `ocid1.compartment.oc1..uniqueID` 
* `default_failover_dr_plan_id` - The unique id of a Failover DR Plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `default_switchover_dr_plan_id` - The unique id of a Switchover DR Plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the Automatic DR configuration. 
* `dr_protection_group_id` - The OCID of the DR protection group to which this Automatic DR configuration belongs.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Automatic DR configuration.  Example: `ocid1.automaticdrconfiguration.oc1..uniqueID` 
* `last_automatic_dr_execution_submit_details` - A message describing the result of the most recent attempt made to submit an Automatic DR plan execution. 
* `last_automatic_dr_execution_submit_status` - The status of most recent attempt to submit Automatic DR plan execution. 
* `lifecycle_details` - A message describing the Automatic DR configuration's current state in more detail. 
* `members` - The list of members in this Automatic DR configuration. 
	* `is_auto_failover_enabled` - A flag indicating if the automatic failover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.  Example: `false` 
	* `is_auto_switchover_enabled` - A flag indicating if the automatic switchover should be enabled for the Autonomous Database Serverless member in the Automatic DR configuration.  Example: `false` 
	* `member_id` - The unique id of the member. Must not be modified by user.  Example: `ocid1.database.oc1..uniqueID` 
	* `member_type` - The type of the member. 
* `state` - The current state of the Automatic DR configuration. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Automatic DR configuration was created. An RFC3339 formatted datetime string.  Example: `2024-03-29T09:36:42Z` 
* `time_last_automatic_dr_execution_submit_attempt` - The date and time of the most recent attempt made to submit an Automatic DR plan execution. An RFC3339 formatted datetime string.  Example: `2025-06-30T09:36:42Z` 
* `time_updated` - The date and time the Automatic DR configuration was updated. An RFC3339 formatted datetime string.  Example: `2024-03-29T09:36:42Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Automatic Dr Configuration
	* `update` - (Defaults to 20 minutes), when updating the Automatic Dr Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Automatic Dr Configuration


## Import

AutomaticDrConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_disaster_recovery_automatic_dr_configuration.test_automatic_dr_configuration "id"
```

