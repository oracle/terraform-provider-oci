---
subcategory: "Dblm"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dblm_patch_management_databases"
sidebar_current: "docs-oci-datasource-dblm-patch_management_databases"
description: |-
  Provides the list of Patch Management Databases in Oracle Cloud Infrastructure Dblm service
---

# Data Source: oci_dblm_patch_management_databases
This data source provides the list of Patch Management Databases in Oracle Cloud Infrastructure Dblm service.

Gets the list of databases


## Example Usage

```hcl
data "oci_dblm_patch_management_databases" "test_patch_management_databases" {

	#Optional
	compartment_id = var.compartment_id
	database_release = var.patch_management_database_database_release
	database_type = var.patch_management_database_database_type
	display_name = var.patch_management_database_display_name
	drifter_patch_id = oci_fleet_apps_management_patch.test_patch.id
	image_compliance = var.patch_management_database_image_compliance
	image_id = oci_core_image.test_image.id
	severity_type = var.patch_management_database_severity_type
	state = var.patch_management_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `database_release` - (Optional) A filter to return only database that match the given release version.
* `database_type` - (Optional) Filter by database type. Possible values Single Instance or RAC. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `drifter_patch_id` - (Optional) A filter to return only database that have given patchId as additional patch (drifter from image version).
* `image_compliance` - (Optional) Filter databases by image compliance status. 
* `image_id` - (Optional) Subscribed image
* `severity_type` - (Optional) Filter by one or more severity types. Possible values are critical, high, medium, low, info and none. 
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `patch_databases_collection` - The list of patch_databases_collection.

### PatchManagementDatabase Reference

The following attributes are exported:

* `items` - List of patchDatabases.
	* `additional_patches` - List of additional patches on database.
		* `category` - Shows if patch is recommended or is an additional patch from an existing database.
		* `description` - Description of the patch recommendation.
		* `patch_id` - Id for the patch recommendation.
		* `patch_name` - Name for the patch recommendation.
	* `current_patch_watermark` - This is the hashcode representing the list of patches applied.
	* `database_id` - Database ocid.
	* `database_name` - Database name.
	* `database_type` - Database type.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `host_or_cluster` - For SI, hosted on host and for RAC, host on cluster.
	* `image_details` - Image details containing the subscribed image, its status, version, owner and time of creation.
		* `created_by` - Name of the person who created the image.
		* `current_version` - Name of the image version marked as current of the image.
		* `image_id` - Image identifier.
		* `image_owner` - Owner of the image.
		* `image_status` - Image status.
		* `image_version` - Release version of the image.
		* `subscribed_image` - Subscribed image.
		* `time_image_creation` - Date when the image was created.
		* `up_to_date_image_version` - An image version name, that is up to date and has no recommendations.
	* `oracle_home_path` - Path to the Oracle home.
	* `patch_activity_details` - Details of deploy, update and migrate-listener(only for single Instance database) operations for this resource.
		* `deploy_operation_id` - Operation Identifier for deploy operation.
		* `deploy_status` - Status of deploy operation.
		* `deploy_task_id` - Task identifier for deploy operation.
		* `migrate_listener_operation_id` - Operation Identifier for migrate listener operation.
		* `migrate_listener_status` - Status of migrate listener operation.
		* `migrate_listener_task_id` - Task identifier for migrate listener operation.
		* `update_operation_id` - Operation Identifier for update operation.
		* `update_status` - Status of update operation.
		* `update_task_id` - Task identifier for update operation.
	* `patch_compliance_details` - Patch Compliance Status
		* `patch_compliance_status` - Patch compliance status.
		* `patch_compliance_version` - Resource patch compliance version name.
	* `patch_user` - Intermediate user to be used for patching, created and maintained by customers. This user requires sudo access to switch as Oracle home owner and root user
	* `release` - Database release.
	* `release_full_version` - Database release full version.
	* `state` - The current state of the database.
	* `sudo_file_path` - Path to sudo binary (executable) file
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `vulnerabilities_summary` - Summary of vulnerabilities found in registered resources grouped by severity.

