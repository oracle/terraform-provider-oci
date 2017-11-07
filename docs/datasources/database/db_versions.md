# oci\_database\_db\_versions

[DbVersionSummary Reference][6318c25a]

  [6318c25a]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbVersionSummary/ "DbVersionSummaryReference"

Gets a list of supported Oracle database versions.

## Example Usage

```
data "oci_database_db_versions" "t" {
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
