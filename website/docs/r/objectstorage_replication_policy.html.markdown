---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_replication_policy"
sidebar_current: "docs-oci-resource-objectstorage-replication_policy"
description: |-
  Provides the Replication Policy resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_replication_policy
This resource provides the Replication Policy resource in Oracle Cloud Infrastructure Object Storage service.

Creates a replication policy for the specified bucket.


## Example Usage

```hcl
resource "oci_objectstorage_replication_policy" "test_replication_policy" {
	#Required
	bucket = var.replication_policy_bucket
	destination_bucket_name = oci_objectstorage_bucket.test_bucket.name
	destination_region_name = oci_identity_region.test_region.name
	name = var.replication_policy_name
	namespace = var.replication_policy_namespace
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `destination_bucket_name` - (Required) The bucket to replicate to in the destination region. Replication policy creation does not automatically create a destination bucket. Create the destination bucket before creating the policy. 
* `destination_region_name` - (Required) The destination region to replicate to, for example "us-ashburn-1".
* `name` - (Required) The name of the policy. Avoid entering confidential information.
* `namespace` - (Required) The Object Storage namespace used for the request.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Replication Policy
	* `update` - (Defaults to 20 minutes), when updating the Replication Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Replication Policy


## Import

ReplicationPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_replication_policy.test_replication_policy "n/{namespaceName}/b/{bucketName}/replicationPolicies/{replicationId}" 
```

