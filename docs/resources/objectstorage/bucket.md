# oci\_objectstorage\_bucket

Provides an Objectstorage resource.

## Example Usage

### Bucket w/ Metadata

```
resource "oci_objectstorage_bucket" "t" {
  compartment_id = "compartment_id"
  name = "name"
  access_type = "ObjectRead"
  namespace = "namespace"
  metadata = {
    "foo" = "bar"
  }
}
```

## Argument Reference

The following arguments are supported:

* `compartmentId` - (Required) The compartment ID in which the bucket is authorized.
* `name` - (Required) The name of the bucket.
* `namespace` - (Required) The namespace in which the bucket lives.
* `metadata` - (Optional) Arbitrary string keys and values for user-defined metadata.
* `access_type` - (Optional) Either "ObjectRead" or "NoPublicAccess". If not specified it defaults to "NoPublicAccess"

## Attributes Reference

The following attributes are exported:

* `created_by` - The OCID of the user who created the bucket.
* `time_created` - The date and time at which the bucket was created.
