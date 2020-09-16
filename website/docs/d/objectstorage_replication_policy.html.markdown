---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_replication_policy"
sidebar_current: "docs-oci-datasource-objectstorage-replication_policy"
description: |-
  Provides details about a specific Replication Policy in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_replication_policy
This data source provides details about a specific Replication Policy resource in Oracle Cloud Infrastructure Object Storage service.

Get the replication policy.


## Example Usage

```hcl
data "oci_objectstorage_replication_policy" "test_replication_policy" {
	#Required
	bucket = var.replication_policy_bucket
	namespace = var.replication_policy_namespace
	replication_id = oci_objectstorage_replication.test_replication.id
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `replication_id` - (Required) The ID of the replication policy.


## Attributes Reference

The following attributes are exported:

* `destination_bucket_name` - The bucket to replicate to in the destination region. Replication policy creation does not automatically create a destination bucket. Create the destination bucket before creating the policy. 
* `destination_region_name` - The destination region to replicate to, for example "us-ashburn-1".
* `id` - The id of the replication policy.
* `name` - The name of the policy.
* `status` - The replication status of the policy. If the status is CLIENT_ERROR, once the user fixes the issue described in the status message, the status will become ACTIVE. 
* `status_message` - A human-readable description of the status.
* `time_created` - The date when the replication policy was created as per [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_last_sync` - Changes made to the source bucket before this time has been replicated. 

