---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_automatic_dr_configurations"
sidebar_current: "docs-oci-datasource-disaster_recovery-automatic_dr_configurations"
description: |-
  Provides the list of Automatic Dr Configurations in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_automatic_dr_configurations
This data source provides the list of Automatic Dr Configurations in Oracle Cloud Infrastructure Disaster Recovery service.

Get a summary list of automatic DR configurations for a DR protection group.

## Example Usage

```hcl
data "oci_disaster_recovery_automatic_dr_configurations" "test_automatic_dr_configurations" {
	#Required
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id

	#Optional
	automatic_dr_configuration_id = oci_disaster_recovery_automatic_dr_configuration.test_automatic_dr_configuration.id
	display_name = var.automatic_dr_configuration_display_name
	lifecycle_state_not_equal_to = var.automatic_dr_configuration_lifecycle_state_not_equal_to
	state = var.automatic_dr_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `automatic_dr_configuration_id` - (Optional) The OCID of the automatic DR configuration.  Example: `ocid1.automaticDrConfiguration.oc1..uniqueID` 
* `display_name` - (Optional) A filter to return only resources that match the given display name.  Example: `MyResourceDisplayName` 
* `dr_protection_group_id` - (Required) The OCID of the DR protection group. Mandatory query param.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `lifecycle_state_not_equal_to` - (Optional) A filter to return only Automatic DR configurations that do not match the given lifecycle state. 
* `state` - (Optional) A filter to return only Automatic DR configurations that match the given lifecycle state. 


## Attributes Reference

The following attributes are exported:

* `automatic_dr_configuration_collection` - The list of automatic_dr_configuration_collection.

### AutomaticDrConfiguration Reference

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

