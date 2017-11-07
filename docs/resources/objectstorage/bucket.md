# oci\_objectstorage\_bucket

[Bucket Reference][d79680e4]

  [d79680e4]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/ "BucketReference"

Provides a container for object storage.

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
* `name` - (Required) The name of the bucket. Avoid entering confidential information.
* `namespace` - (Required) The namespace in which the bucket lives.
* `metadata` - (Optional) Arbitrary string keys and values for user-defined metadata.
* `access_type` - (Optional) Either `ObjectRead` or `NoPublicAccess`. If not specified, `access_type` defaults to `NoPublicAccess`

## Attributes Reference

The following attributes are exported:

* `created_by` - The OCID of the user who created the bucket.
* `time_created` - The date and time at which the bucket was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
