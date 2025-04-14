---
subcategory: "Dblm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dblm_patch_management"
sidebar_current: "docs-oci-datasource-dblm-patch_management"
description: |-
  Provides details about a specific Patch Management in Oracle Cloud Infrastructure Dblm service
---

# Data Source: oci_dblm_patch_management
This data source provides details about a specific Patch Management resource in Oracle Cloud Infrastructure Dblm service.

Overview of Patch Management.


## Example Usage

```hcl
data "oci_dblm_patch_management" "test_patch_management" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	database_release = var.patch_management_database_release
	state = var.patch_management_state
	time_started_greater_than_or_equal_to = var.patch_management_time_started_greater_than_or_equal_to
	time_started_less_than = var.patch_management_time_started_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The required ID of the compartment in which to list resources.
* `database_release` - (Optional) A filter to return only database that match the given release version.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `time_started_greater_than_or_equal_to` - (Optional) A filter to return only resources whose timeStarted is greater than or equal to the given date-time.
* `time_started_less_than` - (Optional) A filter to return only resources whose timeStarted is less than the given date-time.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `images_patch_recommendation_summary` - Summary of image patches recommended to install.
* `message` - A message describing the status of the feature's state
* `patch_operations_summary` - Summary of patch operations.
* `resources` - resources objects
	* `agent_id` - The agent Id of the agent managing the resource.
	* `connector_id` - The connector Id of the resource.
	* `db_platform_type` - The platform type of the resource.
	* `db_version` - The version of the resource.
	* `deployment_type` - The deployment type of the resource.
	* `host_info` - host info objects
		* `host_cores` - Number of host cores.
		* `host_name` - The name of the host.
	* `is_cluster_db` - True if it is a cluster db.
	* `license_type` - The License Type of the resource.
	* `resource_compartment_id` - The compartmentId of the resource.
	* `resource_id` - The Id of the resource.
	* `resource_name` - The name of the resource.
	* `resource_type` - The type of the resource.
* `resources_patch_compliance_summary` - Summary of image patches to be compliant to install.
* `state` - The current state of the feature.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_enabled` - The time the Vulnerability was enabled. An RFC3339 formatted datetime string.

