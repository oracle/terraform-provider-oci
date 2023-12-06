---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_patches"
sidebar_current: "docs-oci-datasource-database-autonomous_container_patches"
description: |-
  Provides the list of Autonomous Container Patches in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_patches
This data source provides the list of Autonomous Container Patches in Oracle Cloud Infrastructure Database service.

Lists the patches applicable to the requested container database.


## Example Usage

```hcl
data "oci_database_autonomous_container_patches" "test_autonomous_container_patches" {
	#Required
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	compartment_id = var.compartment_id

	#Optional
	autonomous_patch_type = var.autonomous_container_patch_autonomous_patch_type
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_patch_type` - (Optional) Autonomous patch type, either "QUARTERLY" or "TIMEZONE". 
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_patches` - The list of autonomous_patches.

### AutonomousContainerPatch Reference

The following attributes are exported:

* `autonomous_patch_type` - Maintenance run type, either "QUARTERLY" or "TIMEZONE". 
* `description` - The text describing this patch package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `patch_model` - Database patching model preference. See [My Oracle Support note 2285040.1](https://support.oracle.com/rs?type=doc&id=2285040.1) for information on the Release Update (RU) and Release Update Revision (RUR) patching models.
* `quarter` - First month of the quarter in which the patch was released.
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `type` - The type of patch. BUNDLE is one example.
* `version` - The version of this patch package.
* `year` - Year in which the patch was released.

