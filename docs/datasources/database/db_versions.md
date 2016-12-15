# baremetal\_database\_db\_versions

Gets a list of supported Oracle database versions.

## Example Usage

```
data "baremetal_database_db_versions" "t" {
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `limit` - (Required) The maximum number of items to return.
* `page` - (Optional) The pagination token to continue listing from.

## Attributes Reference

The following attributes are exported:

* `db_versions` - A list of supported Oracle database versions.
