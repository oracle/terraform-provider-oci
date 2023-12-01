---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_systems_upgrade_history_entries"
sidebar_current: "docs-oci-datasource-database-db_systems_upgrade_history_entries"
description: |-
  Provides the list of Db Systems Upgrade History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_systems_upgrade_history_entries
This data source provides the list of Db Systems Upgrade History Entries in Oracle Cloud Infrastructure Database service.

Gets the history of the upgrade actions performed on the specified DB system.


## Example Usage

```hcl
data "oci_database_db_systems_upgrade_history_entries" "test_db_systems_upgrade_history_entries" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id

	#Optional
	state = var.db_systems_upgrade_history_entry_state
	upgrade_action = var.db_systems_upgrade_history_entry_upgrade_action
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only upgrade history entries that match the given lifecycle state exactly.
* `upgrade_action` - (Optional) A filter to return only upgradeHistoryEntries that match the specified Upgrade Action.


## Attributes Reference

The following attributes are exported:

* `db_system_upgrade_history_entries` - The list of db_system_upgrade_history_entries.

### DbSystemsUpgradeHistoryEntry Reference

The following attributes are exported:

* `action` - The operating system upgrade action.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the upgrade history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `new_gi_version` - A valid Oracle Grid Infrastructure (GI) software version.
* `new_os_version` - A valid Oracle Software (OS) version eg. Oracle Linux Server release 8
* `old_gi_version` - A valid Oracle Grid Infrastructure (GI) software version.
* `old_os_version` - A valid Oracle Software (OS) version eg. Oracle Linux Server release 8
* `snapshot_retention_period_in_days` - The retention period, in days, for the snapshot that allows you to perform a rollback of the upgrade operation. After this number of days passes, you cannot roll back the upgrade.
* `state` - The current state of the action.
* `time_ended` - The date and time when the upgrade action completed
* `time_started` - The date and time when the upgrade action started.

