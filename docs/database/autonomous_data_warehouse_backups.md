# oci_database_autonomous_data_warehouse_backup

## AutonomousDataWarehouseBackup Resource

### AutonomousDataWarehouseBackup Reference

The following attributes are exported:

* `autonomous_data_warehouse_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse.
* `compartment_id` - The OCID of the compartment.
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.



### Create Operation
Creates a new Autonomous Data Warehouse backup for the specified database based on the provided request parameters.


The following arguments are supported:

* `autonomous_data_warehouse_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
* `display_name` - (Required) The user-friendly name for the backup. The name does not have to be unique.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
}
```


## AutonomousDataWarehouseBackup Singular DataSource


### Get Operation
Gets information about the specified Autonomous Data Warehouse backup.

The following arguments are supported:

* `autonomous_data_warehouse_backup_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.


### Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_backup_id = "${oci_database_autonomous_data_warehouse_backup.test_autonomous_data_warehouse_backup.id}"
}
```
# oci_database_autonomous_data_warehouse_backups

## AutonomousDataWarehouseBackup DataSource

Gets a list of autonomous_data_warehouse_backups.

### List Operation
Gets a list of Autonomous Data Warehouse backups based on either the `autonomousDataWarehouseId` or `compartmentId` specified as a query parameter.

The following arguments are supported:

* `autonomous_data_warehouse_id` - (Optional) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


The following attributes are exported:

* `autonomous_data_warehouse_backups` - The list of autonomous_data_warehouse_backups.

### Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_backups" "test_autonomous_data_warehouse_backups" {

	#Optional
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
	state = "${var.autonomous_data_warehouse_backup_state}"
}
```