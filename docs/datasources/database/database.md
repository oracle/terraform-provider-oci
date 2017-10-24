# oci\_database\_database

[Database Reference][37332779]

  [37332779]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/Database/ "DatabaseReference"

Gets information about a specific database on a DB System.

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
* `db_name` - The database name. Avoid entering confidential information.
* `db_unique_name` - A system-generated name for the database. This is a unique name that can't be changed.
* `id` - The OCID of the database.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
