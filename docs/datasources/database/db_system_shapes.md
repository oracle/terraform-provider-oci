# baremetal\_database\_db\_system\_shapes

Gets a list of the shapes that can be used to launch a new DB System.

## Example Usage

```
data "baremetal_database_db_system_shapes" "t" {
  availability_domain = "availability_domain"
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment OCID.
* `limit` - (Required) The maximum number of items to return.
* `page` - (Optional) The pagination token to continue listing from.

## Attributes Reference

The following attributes are exported:

* `db_system_shapes` - A list of the shapes that can be used to launch a new DB System.
