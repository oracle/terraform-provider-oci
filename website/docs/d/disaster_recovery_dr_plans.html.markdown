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

Get a summary list of all DR plans for a DR protection group.

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

* `display_name` - (Optional) A filter to return only resources that match the given display name.  Example: `MyResourceDisplayName` 
* `dr_plan_id` - (Optional) The OCID of the DR plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `dr_plan_type` - (Optional) The DR plan type.
* `dr_protection_group_id` - (Required) The OCID of the DR protection group. Mandatory query param.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `state` - (Optional) A filter to return only DR plans that match the given lifecycle state. 


## Attributes Reference

The following attributes are exported:

* `dr_plan_collection` - The list of dr_plan_collection.

### DrPlan Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DR plan.  Example: `ocid1.compartment.oc1..uniqueID` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the DR plan.  Example: `EBS Switchover PHX to IAD` 
* `dr_protection_group_id` - The OCID of the DR protection group to which this DR plan belongs.  Example: `ocid1.drplan.oc1..uniqueID` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the DR plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `life_cycle_details` - A message describing the DR plan's current state in more detail. 
* `peer_dr_protection_group_id` - The OCID of the peer DR protection group associated with this plan's DR protection group.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `peer_region` - The region of the peer DR protection group associated with this plan's DR protection group.  Example: `us-ashburn-1` 
* `plan_groups` - The list of groups in this DR plan. 
	* `display_name` - The display name of the group.  Example: `DATABASE_SWITCHOVER` 
	* `id` - The unique id of the group. Must not be modified by user.  Example: `sgid1.group..uniqueID` 
	* `is_pause_enabled` - A flag indicating whether this group should be enabled for execution. This flag is only applicable to the `USER_DEFINED_PAUSE` group. The flag should be null for the remaining group types.  Example: `true` 
	* `steps` - The list of steps in the group. 
		* `display_name` - The display name of the group.  Example: `DATABASE_SWITCHOVER` 
		* `error_mode` - The error mode for this step. 
		* `group_id` - The unique id of the group to which this step belongs. Must not be modified by user.  Example: `sgid1.group..uniqueID` 
		* `id` - The unique id of the step. Must not be modified by the user.  Example: `sgid1.step..uniqueID` 
		* `is_enabled` - A flag indicating whether this step should be enabled for execution.  Example: `true` 
		* `member_id` - The OCID of the member associated with this step.  Example: `ocid1.database.oc1..uniqueID` 
		* `timeout` - The timeout in seconds for executing this step.  Example: `600` 
		* `type` - The plan step type. 
		* `user_defined_step` - The details for a user-defined step in a DR plan.
			* `function_id` - The OCID of function to be invoked.  Example: `ocid1.fnfunc.oc1..uniqueID` 
			* `function_region` - The region in which the function is deployed.  Example: `us-ashburn-1` 
			* `object_storage_script_location` - The details of an object storage script location for a user-defined step in a DR plan.
				* `bucket` - The bucket name inside the object storage namespace.  Example: `custom_dr_scripts` 
				* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
				* `object` - The object name inside the object storage bucket.  Example: `validate_app_start.sh` 
			* `request_body` - The request body for the function.  Example: `{ "FnParam1", "FnParam2" }` 
			* `run_as_user` - The userid on the instance to be used for executing the script or command.  Example: `opc` 
			* `run_on_instance_id` - The OCID of the instance on which this script or command should be executed.  

				**For moving instances:** *runOnInstanceId* must be the OCID of the instance in the region where the  instance is currently present.  

				**For non-moving instances:** *runOnInstanceId* must be the OCID of the non-moving instance.  

				Example: `ocid1.instance.oc1..uniqueID` 
			* `run_on_instance_region` - The region of the instance where this script or command should be executed.  Example: `us-ashburn-1` 
			* `script_command` - The script name and arguments.  Example: `/usr/bin/python3 /home/opc/scripts/my_app_script.py arg1 arg2 arg3` 
			* `step_type` - The type of the user-defined step.

				**RUN_OBJECTSTORE_SCRIPT_PRECHECK** - A step which performs a precheck on a script stored in Oracle Cloud Infrastructure object storage.

				**RUN_LOCAL_SCRIPT_PRECHECK** - A step which performs a precheck on a script which resides locally on a compute instance.

				**INVOKE_FUNCTION_PRECHECK** - A step which performs a precheck on an Oracle Cloud Infrastructure function. See https://docs.oracle.com/en-us/iaas/Content/Functions/home.htm.

				**RUN_OBJECTSTORE_SCRIPT** - A step which runs a script stored in Oracle Cloud Infrastructure object storage.

				**RUN_LOCAL_SCRIPT** - A step which runs a script that resides locally on a compute instance.

				**INVOKE_FUNCTION** - A step which invokes an Oracle Cloud Infrastructure function. See https://docs.oracle.com/en-us/iaas/Content/Functions/home.htm. 
	* `type` - The group type.  Example: `BUILT_IN` 
* `state` - The current state of the DR plan. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the DR plan was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The date and time the DR plan was updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `type` - The type of the DR plan. 

