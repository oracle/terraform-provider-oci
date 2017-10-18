# oci\_objectstorage\_bucket\_summary

[BucketSummary Reference][cb1f26ec]

  [cb1f26ec]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/BucketSummary/ "BucketSummaryReference"

Get a list of all the `BucketSummary`'s in a namespace. A `BucketSummary` contains only summary fields for the bucket and does not contain fields like the user-defined metadata.


## Example Usage

```
    data "oci_objectstorage_bucket_summaries" "t" {
      compartment_id = "compartmentid"
      namespace = "namespace"
      limit = 1
      page = "page"
    }
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment ID in which to create the bucket.
* `namespace` - (Required) The top-level namespace used for the request.
* `limit` - (Optional) The maximum number of items to return.
* `page` - (Optional) The page at which to start retrieving results.
