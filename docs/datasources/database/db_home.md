# oci\_database\_db\_home

[DbHome Reference][922d9bef]

  [922d9bef]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbHome/ "DbHomeReference"

Gets information about the specified database home. 

## Example Usage

```
data "oci_database_db_home" "t" {
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
* `display_name` - The user-provided name for the database home. It does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the database home.
* `state` - The current state of the database home. Allowed values are: [PROVISIONING, AVAILABLE, UPDATING, TERMINATING, TERMINATED, FAILED]
* `time_created` - The date and time the database home was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
