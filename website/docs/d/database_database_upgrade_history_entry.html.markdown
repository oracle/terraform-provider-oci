---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_upgrade_history_entry"
sidebar_current: "docs-oci-datasource-database-database_upgrade_history_entry"
description: |-
  Provides details about a specific Database Upgrade History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_upgrade_history_entry
This data source provides details about a specific Database Upgrade History Entry resource in Oracle Cloud Infrastructure Database service.

gets the upgrade history for a specified database.


## Example Usage

```hcl
data "oci_database_database_upgrade_history_entry" "test_database_upgrade_history_entry" {
	#Required
	database_id = oci_database_database.test_database.id
	upgrade_history_entry_id = oci_database_upgrade_history_entry.test_upgrade_history_entry.id
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `upgrade_history_entry_id` - (Required) The database/db system upgrade History [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `action` - The database upgrade action.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database upgrade history.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `options` - Additional upgrade options supported by DBUA(Database Upgrade Assistant). Example: "-upgradeTimezone false -keepEvents" 
* `source` - The source of the Oracle Database software to be used for the upgrade.
	* Use `DB_VERSION` to specify a generally-available Oracle Database software version to upgrade the database.
	* Use `DB_SOFTWARE_IMAGE` to specify a [database software image](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databasesoftwareimage.htm) to upgrade the database. 
* `source_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `state` - Status of database upgrade history SUCCEEDED|IN_PROGRESS|FAILED.
* `target_db_version` - A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
* `target_database_software_image_id` - the database software image used for upgrading database.
* `target_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `time_ended` - The date and time when the database upgrade ended.
* `time_started` - The date and time when the database upgrade started.

