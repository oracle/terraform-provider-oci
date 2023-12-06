---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_systems_upgrade_history_entry"
sidebar_current: "docs-oci-datasource-database-db_systems_upgrade_history_entry"
description: |-
  Provides details about a specific Db Systems Upgrade History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_systems_upgrade_history_entry
This data source provides details about a specific Db Systems Upgrade History Entry resource in Oracle Cloud Infrastructure Database service.

Gets the details of the specified operating system upgrade operation for the specified DB system.


## Example Usage

```hcl
data "oci_database_db_systems_upgrade_history_entry" "test_db_systems_upgrade_history_entry" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
	upgrade_history_entry_id = oci_database_upgrade_history_entry.test_upgrade_history_entry.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `upgrade_history_entry_id` - (Required) The database/db system upgrade History [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

