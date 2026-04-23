---
subcategory: "Datacc"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacc_maintenance_execution"
sidebar_current: "docs-oci-datasource-datacc-maintenance_execution"
description: |-
  Provides details about a specific Maintenance Execution in Oracle Cloud Infrastructure Datacc service
---

# Data Source: oci_datacc_maintenance_execution
This data source provides details about a specific Maintenance Execution resource in Oracle Cloud Infrastructure Datacc service.

Gets information about the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

## Example Usage

```hcl
data "oci_datacc_maintenance_execution" "test_maintenance_execution" {
	#Required
	maintenance_execution_id = oci_datacc_maintenance_execution.test_maintenance_execution.id
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_execution_id` - (Required) The maintenance execution OCID.


## Attributes Reference

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

