# baremetal\_database\_db\_home

Gets information about the specified database home.

## Example Usage

```
data "baremetal_database_db_home" "t" {
  db_home_id = "id"
}
```

## Argument Reference

The following arguments are supported:

* `db_home_id` - (Required) The database home OCID.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_system_id` - The OCID of the DB System.
* `db_version` - The Oracle database version.
* `display_name` - The user-provided name for the database home.
* `id` - The OCID of the database home.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.
