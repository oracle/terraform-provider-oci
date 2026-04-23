---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_maintenance_executions"
sidebar_current: "docs-oci-datasource-datacc-maintenance_executions"
description: |-
  Provides the list of Maintenance Executions in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_maintenance_executions
This data source provides the list of Maintenance Executions in Oracle Cloud Infrastructure Datacc service.

Gets a list of the maintenance executions in the specified compartment.


## Example Usage

```hcl
data "oci_datacc_maintenance_executions" "test_maintenance_executions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.maintenance_execution_display_name
	infrastructure_id = oci_datacc_infrastructure.test_infrastructure.id
	maintenance_run_id = oci_database_maintenance_run.test_maintenance_run.id
	maintenance_subtype = var.maintenance_execution_maintenance_subtype
	maintenance_type = var.maintenance_execution_maintenance_type
	state = var.maintenance_execution_state
	target_resource_type = var.maintenance_execution_target_resource_type
	time_accepted_greater_than_or_equal_to = var.maintenance_execution_time_accepted_greater_than_or_equal_to
	time_accepted_less_than_or_equal_to = var.maintenance_execution_time_accepted_less_than_or_equal_to
	type = var.maintenance_execution_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. For list operations, you may provide the tenant [OCID] in this field. When a tenant OCID is provided, it will be validated against the caller's tenant and then treated as tenant scope (compartmentId filtering is not applied). 
* `display_name` - (Optional) A filter to return resources that match the entire display name given. The match is case sensitive.
* `infrastructure_id` - (Optional) The Database Infrastructure ID.
* `maintenance_run_id` - (Optional) The maintenance run OCID.
* `maintenance_subtype` - (Optional) The sub-type of the maintenance run.
* `maintenance_type` - (Optional) The maintenance type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `target_resource_type` - (Optional) The type of the target resource.
* `time_accepted_greater_than_or_equal_to` - (Optional) Filter maintenance run for after given time.
* `time_accepted_less_than_or_equal_to` - (Optional) Filter maintenance run for before given time.
* `type` - (Optional) The maintenance execution type.


## Attributes Reference

The following attributes are exported:

* `maintenance_execution_collection` - The list of maintenance_execution_collection.

### MaintenanceExecution Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each compute server patching operation. Supported values are 15 to 120 minutes. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the maintenance run execution.
* `display_name` - The user-friendly name for the maintenance run execution.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run execution.
* `infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure on which the maintenance run execution occurred.
* `is_custom_action_timeout_enabled` - At the time of execution whether the custom action time out is enabled for the maintenance run that is being executed.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run to which this maintenance execution belongs.
* `maintenance_subtype` - Maintenance run execution sub-type.
* `maintenance_type` - Maintenance type.
* `patching_mode` - The patching mode for the maintenance run that is being executed. 
* `source_version` - The source software version for the Oracle infrastructure.
* `state` - The state of the maintenance run execution. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_resource_type` - The type of the target resource on which the maintenance run execution occurred.
* `target_version` - The target software version for the Database Infrastructure patching operation.
* `time_ended` - The date and time the maintenance run was completed.
* `time_started` - The date and time the maintenance run execution started.
* `total_time_taken_in_mins` - The total time taken by this execution in minutes.
* `workflow_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request executed by this execution.

