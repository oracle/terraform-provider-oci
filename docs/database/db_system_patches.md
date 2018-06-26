
# oci_database_db_system_patches

## DbSystemPatch DataSource

Gets a list of db_system_patches.

### List Operation
Lists the patches applicable to the requested DB System.

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `patches` - The list of patches.

### Example Usage

```hcl
data "oci_database_db_system_patches" "test_db_system_patches" {
	#Required
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```
### DbSystemPatch Reference

The following attributes are exported:

* `available_actions` - Actions that can possibly be performed using this patch.
* `description` - The text describing this patch package.
* `id` - The OCID of the patch.
* `last_action` - Action that is currently being performed or was completed last.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `version` - The version of this patch package.
