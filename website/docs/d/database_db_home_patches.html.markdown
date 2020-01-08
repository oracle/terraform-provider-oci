---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_home_patches"
sidebar_current: "docs-oci-datasource-database-db_home_patches"
description: |-
  Provides the list of Db Home Patches in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_home_patches
This data source provides the list of Db Home Patches in Oracle Cloud Infrastructure Database service.

Lists patches applicable to the requested Database Home.


## Example Usage

```hcl
data "oci_database_db_home_patches" "test_db_home_patches" {
	#Required
	db_home_id = "${oci_database_db_home.test_db_home.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_home_id` - (Required) The Database Home [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `patches` - The list of patches.

### DbHomePatch Reference

The following attributes are exported:

* `available_actions` - Actions that can possibly be performed using this patch.
* `description` - The text describing this patch package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `last_action` - Action that is currently being performed or was completed last.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `version` - The version of this patch package.

