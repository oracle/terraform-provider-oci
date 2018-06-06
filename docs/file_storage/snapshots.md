# oci_file_storage_snapshot

## Snapshot Resource

### Snapshot Reference

The following attributes are exported:

* `file_system_id` - The OCID of the file system from which the snapshot was created. 
* `id` - The OCID of the snapshot.
* `name` - Name of the snapshot. This value is immutable.  Avoid entering confidential information.  Example: `Sunday` 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new snapshot of the specified file system. You
can access the snapshot at `.snapshot/<name>`.


The following arguments are supported:

* `file_system_id` - (Required) The OCID of this export's file system.
* `name` - (Required) Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.  Avoid entering confidential information.  Example: `Sunday` 


### Update Operation
Updates the specified snapshot's information.

The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_file_storage_snapshot" "test_snapshot" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	name = "${var.snapshot_name}"
}
```

# oci_file_storage_snapshots

## Snapshot DataSource

Gets a list of snapshots.

### List Operation
Lists snapshots of the specified file system.

The following arguments are supported:

* `file_system_id` - (Required) The OCID of the file system.
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


The following attributes are exported:

* `snapshots` - The list of snapshots.

### Example Usage

```hcl
data "oci_file_storage_snapshots" "test_snapshots" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"

	#Optional
	id = "${var.snapshot_id}"
	state = "${var.snapshot_state}"
}
```