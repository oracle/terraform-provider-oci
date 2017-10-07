# oci\_core\_shape

**API:** [Shape Reference][ba540762]

  [ba540762]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Shape/ "ShapeReference"

Gets the list of shapes that can be used to launch an instance within the specified compartment.

## Example Usage

```
data "oci_core_shape" "s" {
  compartment_id = "compartmentid"
  availability_domain = "availability_domain"
  image_id = "imageid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `availability_domain` - (Required) The name of the Availability Domain.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The page to fetch
* `image_id` - (Optional) The OCID of an image.

## Attributes Reference

The following attributes are exported:

* `shapes` - The list of shapes.

## Shape reference
* `shape` - The name of the shape.
