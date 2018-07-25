---
layout: "oci"
page_title: "OCI: oci_database_db_system_patch_history_entries"
sidebar_current: "docs-oci-datasource-database-db_system_patch_history_entries"
description: |-
  Provides a list of DbSystemPatchHistoryEntries
---

# Data Source: oci_database_db_system_patch_history_entries
The DbSystemPatchHistoryEntries data source allows access to the list of OCI db_system_patch_history_entries

Gets the history of the patch actions performed on the specified DB System.


## Example Usage

```hcl
data "oci_database_db_system_patch_history_entries" "test_db_system_patch_history_entries" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `patch_history_entries` - The list of patch_history_entries.

### DbSystemPatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The OCID of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The OCID of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed.
* `time_started` - The date and time when the patch action started.

