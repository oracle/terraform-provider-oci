---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system_patches"
sidebar_current: "docs-oci-datasource-database-db_system_patches"
description: |-
  Provides the list of Db System Patches in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_system_patches
This data source provides the list of Db System Patches in Oracle Cloud Infrastructure Database service.

Lists the patches applicable to the requested DB system.


## Example Usage

```hcl
data "oci_database_db_system_patches" "test_db_system_patches" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `patches` - The list of patches.

### DbSystemPatch Reference

The following attributes are exported:

* `available_actions` - Actions that can possibly be performed using this patch.
* `description` - The text describing this patch package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `last_action` - Action that is currently being performed or was completed last.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `version` - The version of this patch package.

