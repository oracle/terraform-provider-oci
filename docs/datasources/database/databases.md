# oci\_database\_databases

Gets a list of the databases in the specified database home.

## Example Usage

```
data "oci_database_databases" "t" {
  compartment_id = "compartment_id"
  db_home_id = "db_home_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `db_home_id` - (Required) A database home OCID.
* `limit` - (Required) The maximum number of items to return.
* `page` - (Optional) The pagination token to continue listing from.

## Attributes Reference

The following attributes are exported:

* `databases` - A list of the databases in the specified database home.
