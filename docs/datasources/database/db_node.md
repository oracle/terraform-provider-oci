# oci\_database\_db\_node

[DbNode Reference][1581cf27]

  [1581cf27]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbNode/ "DbNodeReference"

Gets information about the specified database node.

## Example Usage

```
data "oci_database_db_node" "t" {
  db_node_id = "id"
}
```

## Argument Reference

The following arguments are supported:

* `db_node_id` - (Required) The database node OCID.

## Attributes Reference

The following attributes are exported:

* `db_system_id` - The OCID of the DB System.
* `hostname` - The host name for the DB Node.
* `id` - The OCID of the DB Node.
* `state` - The current state of the database node. Allowed values are: [PROVISIONING, AVAILABLE, UPDATING, STOPPING, STOPPED, STARTING, TERMINATING, TERMINATED, FAILED]
* `time_created` - The date and time that the DB Node was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `vnic_id` - The OCID of the VNIC.
