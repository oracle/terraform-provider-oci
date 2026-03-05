---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system_os_patch_history_entries"
sidebar_current: "docs-oci-datasource-database-db_system_os_patch_history_entries"
description: |-
  Provides the list of Db System Os Patch History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_system_os_patch_history_entries
This data source provides the list of Db System Os Patch History Entries in Oracle Cloud Infrastructure Database service.

Gets the history of the OS patch actions performed on the specified DB system.


## Example Usage

```hcl
data "oci_database_db_system_os_patch_history_entries" "test_db_system_os_patch_history_entries" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id

	#Optional
	action = var.db_system_os_patch_history_entry_action
	state = var.db_system_os_patch_history_entry_state
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Optional) A filter to return only OS patch history entries that match the specified OS patch action.
* `db_system_id` - (Required) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only OS patch history entries that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `db_system_os_patch_history_entry_collection` - The list of db_system_os_patch_history_entry_collection.

### DbSystemOsPatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OS patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `os_patch_details` - Results of OS patch details for a DB System.
	* `items` - Array of OS patch details for a DB System.
		* `db_node_id` - The OCID of the DB node targeted for this patch action.
		* `is_reboot_required` - Indicates whether a reboot is required after applying the patch.
		* `rpms` - List of OS package identifiers (e.g., RPM strings).
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed
* `time_started` - The date and time when the patch action started.

