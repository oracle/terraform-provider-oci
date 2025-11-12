---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_installed_patches"
sidebar_current: "docs-oci-datasource-fleet_apps_management-installed_patches"
description: |-
  Provides the list of Installed Patches in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_installed_patches
This data source provides the list of Installed Patches in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of installed patches for the specified target.
CompartmentId should be the compartment OCID of the resource (Containing the target).


## Example Usage

```hcl
data "oci_fleet_apps_management_installed_patches" "test_installed_patches" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	target_id = var.fleet_target_id
	target_name = var.fleet_target_name
	patch_level = var.installed_patch_patch_level
	patch_type = var.installed_patch_patch_type
	severity = var.installed_patch_severity
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `target_id` - (Optional) Target identifier.
* `target_name` - (Optional) Target name.
* `patch_level` - (Optional) Patch level with values like LATEST, LATEST_MINUS_ONE, LATEST_MIUS_TWO etc.,.
* `patch_type` - (Optional) Patch type.
* `severity` - (Optional) Patch severity with values like CRITICAL, HIGH, MEDIUM and LOW.


## Attributes Reference

The following attributes are exported:

* `installed_patch_collection` - The list of installed_patch_collection.

### InstalledPatch Reference

The following attributes are exported:

* `items` - List of installed patches
	* `patch_description` - Description of the patch
	* `patch_id` - The OCID of the patch.
	* `patch_level` - Patch level.
	* `patch_name` - Name of the patch.
	* `patch_type` - Type of the patch.
	* `severity` - Patch severity.
	* `time_applied` - Date on which the patch was applied to the target.
	* `time_released` - Date on which the patch was released.

