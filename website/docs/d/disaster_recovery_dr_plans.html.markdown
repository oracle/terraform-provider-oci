---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_plans"
sidebar_current: "docs-oci-datasource-disaster_recovery-dr_plans"
description: |-
  Provides the list of Dr Plans in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_dr_plans
This data source provides the list of Dr Plans in Oracle Cloud Infrastructure Disaster Recovery service.

Gets a summary list of all DR Plans for a DR Protection Group.

## Example Usage

```hcl
data "oci_disaster_recovery_dr_plans" "test_dr_plans" {
	#Required
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id

	#Optional
	display_name = var.dr_plan_display_name
	dr_plan_id = oci_disaster_recovery_dr_plan.test_dr_plan.id
	dr_plan_type = var.dr_plan_dr_plan_type
	state = var.dr_plan_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.  Example: `MY UNIQUE DISPLAY NAME` 
* `dr_plan_id` - (Optional) The OCID of the DR Plan.  Example: `ocid1.drplan.oc1.iad.exampleocid` 
* `dr_plan_type` - (Optional) The DR Plan type.
* `dr_protection_group_id` - (Required) The OCID of the DR Protection Group. Mandatory query param.  Example: `ocid1.drprotectiongroup.oc1.phx.exampleocid` 
* `state` - (Optional) A filter to return only DR Plans that match the given lifecycleState. 


## Attributes Reference

The following attributes are exported:

* `dr_plan_collection` - The list of dr_plan_collection.

### DrPlan Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DR Plan.  Example: `ocid1.compartment.oc1..exampleocid1` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of this DR Plan.  Example: `EBS Switchover PHX to IAD` 
* `dr_protection_group_id` - The OCID of the DR Protection Group with which this DR Plan is associated.  Example: `ocid1.drplan.oc1.iad.exampleocid2` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The OCID of this DR Plan.  Example: `ocid1.drplan.oc1.iad.exampleocid2` 
* `life_cycle_details` - A message describing the DR Plan's current state in more detail. 
* `peer_dr_protection_group_id` - The OCID of the peer (remote) DR Protection Group associated with this plan's DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.phx.exampleocid1` 
* `peer_region` - The region of the peer (remote) DR Protection Group associated with this plan's DR Protection Group.  Example: `us-phoenix-1` 
* `plan_groups` - The list of groups in this DR Plan. 
	* `display_name` - The display name of this DR Plan Group.  Example: `DATABASE_SWITCHOVER` 
	* `id` - The unique id of this group. Must not be modified by user.  Example: `sgid1.group..examplegroupsgid` 
	* `steps` - The list of steps in this plan group. 
		* `display_name` - The display name of this DR Plan Group.  Example: `DATABASE_SWITCHOVER` 
		* `error_mode` - The error mode for this step. 
		* `group_id` - The unique id of the group to which this step belongs. Must not be modified by user.  Example: `sgid1.group..examplegroupsgid` 
		* `id` - The unique id of this step. Must not be modified by the user.  Example: `sgid1.step..examplestepsgid` 
		* `is_enabled` - A flag indicating whether this step should be enabled for execution.  Example: `true` 
		* `member_id` - The OCID of the member associated with this step.  Example: `ocid1.database.oc1.phx.exampleocid1` 
		* `timeout` - The timeout in seconds for executing this step.  Example: `600` 
		* `type` - The plan step type. 
		* `user_defined_step` - The details for a user-defined step in a DR Plan.
			* `function_id` - The OCID of function to be invoked.  Example: `ocid1.fnfunc.oc1.iad.exampleocid2` 
			* `function_region` - The region in which the function is deployed.  Example: `us-ashburn-1` 
			* `object_storage_script_location` - Information about an Object Storage script location for a user-defined step in a DR Plan.
				* `bucket` - The bucket name inside the Object Storage namespace.  Example: `custom_dr_scripts` 
				* `namespace` - The namespace in Object Storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
				* `object` - The object name inside the Object Storage bucket.  Example: `validate_app_start.sh` 
			* `request_body` - The request body for the function.  Example: `{ "FnParam1", "FnParam2" }` 
			* `run_as_user` - The userid on the instance to be used for executing the script or command.  Example: `opc` 
			* `run_on_instance_id` - The OCID of the instance where this script or command should be executed.  Example: `ocid1.instance.oc1.phx.exampleocid1` 
			* `run_on_instance_region` - The region of the instance where this script or command should be executed.  Example: `us-phoenix-1` 
			* `script_command` - The script name and arguments.  Example: `/usr/bin/python3 /home/opc/scripts/my_app_script.py arg1 arg2 arg3` 
			* `step_type` - The type of the step.

				RUN_OBJECTSTORE_SCRIPT_PRECHECK - A step which performs a precheck on a script stored in Oracle Object Storage Service

				RUN_LOCAL_SCRIPT_PRECHECK - A step which performs a precheck on a script which resides locally on a compute instance

				INVOKE_FUNCTION_PRECHECK - A step which performs a precheck on an Oracle Function. See https://docs.oracle.com/en-us/iaas/Content/Functions/home.htm.

				RUN_OBJECTSTORE_SCRIPT - A step which runs a script stored in Oracle Object Storage Service

				RUN_LOCAL_SCRIPT - A step which runs a script that resides locally on a compute instance

				INVOKE_FUNCTION - A step which invokes an Oracle Function. See https://docs.oracle.com/en-us/iaas/Content/Functions/home.htm. 
	* `type` - The plan group type. 
* `state` - The current state of the DR Plan. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DR Plan was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The date and time the DR Plan was updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `type` - The type of this DR Plan. 

