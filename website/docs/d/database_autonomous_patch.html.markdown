---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_patch"
sidebar_current: "docs-oci-datasource-database-autonomous_patch"
description: |-
  Provides details about a specific Autonomous Patch in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_patch
This data source provides details about a specific Autonomous Patch resource in Oracle Cloud Infrastructure Database service.

Gets information about a specific autonomous patch.

## Example Usage

```hcl
data "oci_database_autonomous_patch" "test_autonomous_patch" {
	#Required
	autonomous_patch_id = oci_database_autonomous_patch.test_autonomous_patch.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_patch_id` - (Required) The autonomous patch [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

