---
layout: "oci"
page_title: "OCI: oci_database_db_home_patch_history_entries"
sidebar_current: "docs-oci-datasource-database-db_home_patch_history_entries"
description: |-
  Provides a list of DbHomePatchHistoryEntries
---

# Data Source: oci_database_db_home_patch_history_entries
The `oci_database_db_home_patch_history_entries` data source allows access to the list of OCI db_home_patch_history_entries

Gets history of the actions taken for patches for the specified database home.


## Example Usage

```hcl
data "oci_database_db_home_patch_history_entries" "test_db_home_patch_history_entries" {
	#Required
	db_home_id = "${oci_database_db_home.test_db_home.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `patch_history_entries` - The list of patch_history_entries.

### DbHomePatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The OCID of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The OCID of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed.
* `time_started` - The date and time when the patch action started.

