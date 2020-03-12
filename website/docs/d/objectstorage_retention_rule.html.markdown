---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_retention_rule"
sidebar_current: "docs-oci-datasource-objectstorage-retention_rule"
description: |-
  Provides details about a specific Retention Rule in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_retention_rule
This data source provides details about a specific Retention Rule resource in Oracle Cloud Infrastructure Object Storage service.

Get the specified retention rule.

## Example Usage

```hcl
data "oci_objectstorage_retention_rule" "test_retention_rule" {
	#Required
	bucket = "${var.retention_rule_bucket}"
	namespace = "${var.retention_rule_namespace}"
	retention_rule_id = "${oci_objectstorage_retention_rule.test_retention_rule.id}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `namespace` - (Required) The Object Storage namespace used for the request.
* `retention_rule_id` - (Required) The ID of the retention rule.


## Attributes Reference

The following attributes are exported:

* `display_name` - User specified name for the retention rule.
* `duration` - 
	* `time_amount` - The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp. 
	* `time_unit` - The unit that should be used to interpret timeAmount.
* `etag` - The entity tag (ETag) for the retention rule.
* `id` - Unique identifier for the retention rule.
* `time_created` - The date and time that the retention rule was created as per [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_modified` - The date and time that the retention rule was modified as per [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_rule_locked` - The date and time as per [RFC 3339](https://tools.ietf.org/html/rfc3339) after which this rule becomes locked. and can only be deleted by deleting the bucket. 

