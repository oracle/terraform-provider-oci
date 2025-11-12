---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_recommended_patches"
sidebar_current: "docs-oci-datasource-fleet_apps_management-recommended_patches"
description: |-
  Provides the list of Recommended Patches in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_recommended_patches
This data source provides the list of Recommended Patches in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of recommended patches for the specified target.


## Example Usage

```hcl
data "oci_fleet_apps_management_recommended_patches" "test_recommended_patches" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	target_id = var.fleet_target_id
	target_name = var.fleet_target_name
	patch_id = oci_fleet_apps_management_patch.test_patch.id
	patch_level = var.recommended_patch_patch_level
	patch_type = var.recommended_patch_patch_type
	severity = var.recommended_patch_severity
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `target_id` - (Optional) Fleet target identifier.
* `target_name` - (Optional) Fleet target name.
* `patch_id` - (Optional) Patch identifier.
* `patch_level` - (Optional) Patch level with values like LATEST, LATEST_MINUS_ONE, LATEST_MIUS_TWO etc.,.
* `patch_type` - (Optional) Patch type.
* `severity` - (Optional) Patch severity with values like CRITICAL, HIGH, MEDIUM and LOW.


## Attributes Reference

The following attributes are exported:

* `recommended_patch_collection` - The list of recommended_patch_collection.

### RecommendedPatch Reference

The following attributes are exported:

* `items` - List of recommended patches
	* `patch_description` - Description of the patch
	* `patch_id` - The OCID of the patch.
	* `patch_level` - Patch level
	* `patch_name` - Name of the patch.
	* `patch_type` - Type of the patch.
	* `severity` - Patch severity.
	* `time_released` - Date on which the patch was released.

