
# oci_database_patch_history_entries

## PatchHistoryEntry DataSource

Gets a list of patch_history_entries.

### List Operation
Gets history of the actions taken for patches for the specified database home.

The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `patch_history_entries` - The list of patch_history_entries.

### Example Usage

```hcl
data "oci_database_patch_history_entries" "test_patch_history_entries" {
	#Required
	db_home_id = "${var.patch_history_entry_db_home_id}"
}
```
### PatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The OCID of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The OCID of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed.
* `time_started` - The date and time when the patch action started.
