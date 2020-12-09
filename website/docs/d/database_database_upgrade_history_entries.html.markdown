---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_database_upgrade_history_entries"
sidebar_current: "docs-oci-datasource-database-database_upgrade_history_entries"
description: |-
  Provides the list of Database Upgrade History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_database_upgrade_history_entries
This data source provides the list of Database Upgrade History Entries in Oracle Cloud Infrastructure Database service.

gets the upgrade history for a specified database.


## Example Usage

```hcl
data "oci_database_database_upgrade_history_entries" "test_database_upgrade_history_entries" {
	#Required
	database_id = oci_database_database.test_database.id

	#Optional
	state = var.database_upgrade_history_entry_state
	upgrade_action = var.database_upgrade_history_entry_upgrade_action
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only upgradeHistoryEntries that match the given lifecycle state exactly.
* `upgrade_action` - (Optional) A filter to return only upgradeHistoryEntries that match the specified Upgrade Action.


## Attributes Reference

The following attributes are exported:

* `database_upgrade_history_entries` - The list of database_upgrade_history_entries.

### DatabaseUpgradeHistoryEntry Reference

The following attributes are exported:

* `action` - action for upgrading database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database upgrade history.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `source` - The source of the database upgrade Use 'DB_HOME' for using existing db home to upgrade the database Use 'DB_VERSION' for using database version to upgrade the database Use 'DB_SOFTWARE_IMAGE' for using database software image to upgrade the database 
* `source_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `state` - Status of database upgrade history SUCCEEDED|IN_PROGRESS|FAILED.
* `target_db_version` - A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions) operation.
* `target_database_software_image_id` - the database software image used for upgrading database.
* `target_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `time_ended` - The date and time when the database upgrade ended.
* `time_started` - The date and time when the database upgrade started.

