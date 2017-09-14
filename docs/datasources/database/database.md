# oci\_database\_database

Gets information about a specific database.

## Example Usage

```
data "oci_database_database" "t" {
  database_id = "id"
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database OCID.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_home_id` - The OCID of the database home.
* `db_name` - The database name.
* `db_unique_name` - A system-generated name for the database.
* `id` - The OCID of the database.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.
