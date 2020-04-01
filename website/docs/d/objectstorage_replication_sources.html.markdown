---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_replication_sources"
sidebar_current: "docs-oci-datasource-objectstorage-replication_sources"
description: |-
  Provides the list of Replication Sources in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_replication_sources
This data source provides the list of Replication Sources in Oracle Cloud Infrastructure Object Storage service.

List the replication sources of a destination bucket.


## Example Usage

```hcl
data "oci_objectstorage_replication_sources" "test_replication_sources" {
	#Required
	bucket = "${var.replication_source_bucket}"
	namespace = "${var.replication_source_namespace}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.


## Attributes Reference

The following attributes are exported:

* `replication_sources` - The list of replication_sources.

### ReplicationSource Reference

The following attributes are exported:

* `policy_name` - The name of the policy.
* `source_bucket_name` - The source bucket replicating data from.
* `source_region_name` - The source region replicating data from, for example "us-ashburn-1".

