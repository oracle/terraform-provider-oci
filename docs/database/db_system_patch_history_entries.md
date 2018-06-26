
# oci_database_db_system_patch_history_entries

## DbSystemPatchHistoryEntry DataSource

Gets a list of db_system_patch_history_entries.

### List Operation
Gets the history of the patch actions performed on the specified DB System.

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `patch_history_entries` - The list of patch_history_entries.

### Example Usage

```hcl
data "oci_database_db_system_patch_history_entries" "test_db_system_patch_history_entries" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```
### DbSystemPatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The OCID of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The OCID of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed.
* `time_started` - The date and time when the patch action started.
