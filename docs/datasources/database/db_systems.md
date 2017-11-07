# oci\_database\_db\_systems

[DbSystemSummary Reference][28cea1db]

  [28cea1db]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystemSummary/ "DbSystemSummaryReference"

Gets a list of the DB Systems in the specified compartment.

## Example Usage

```
data "oci_database_db_systems" "t" {
  compartment_id = "compartmentid"
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

* `db_systems` - A list of the DB Systems in the specified compartment.
