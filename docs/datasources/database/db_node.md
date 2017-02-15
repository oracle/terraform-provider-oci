# baremetal\_database\_db\_node

Gets information about the specified database node.

## Example Usage

```
data "baremetal_database_db_node" "t" {
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
* `state` - The current state of the database node.
* `time_created` - The date and time that the DB Node was created.
* `vnic_id` - The OCID of the VNIC.
