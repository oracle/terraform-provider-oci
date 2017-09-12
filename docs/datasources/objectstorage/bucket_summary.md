# oci\_objectstorage\_bucket\_summary

Get a list of all Bucket summaries in a namespace. A BucketSummary contains only summary fields for the bucket and does not contain fields like the user-defined metadata.

To use this and other API operations, you must be authorized in an IAM policy.
See [List Buckets API Docs](https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/BucketSummary/ListBuckets) for details.


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
