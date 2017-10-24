# oci\_database\_db\_homes

[DbHome Summary][40f7e502]

  [40f7e502]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbHomeSummary/ "DbHomeSummary"

Gets a list of database homes in the specified DB System and compartment.

## Example Usage

```
data "oci_database_db_homes" "t" {
  compartment_id = "compartment_id"
  db_system_id = "db_system_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `db_system_id` - (Required) The OCID of the DB System.
* `limit` - (Required) The maximum number of items to return.
* `page` - (Optional) The pagination token to continue listing from.

## Attributes Reference

The following attributes are exported:

* `db_homes` - A list of database homes in the specified DB System and compartment.
