---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_retention_rule"
sidebar_current: "docs-oci-resource-objectstorage-retention_rule"
description: |-
  Provides the Retention Rule resource in Oracle Cloud Infrastructure Object Storage service
---

# oci_objectstorage_retention_rule
This resource provides the Retention Rule resource in Oracle Cloud Infrastructure Object Storage service.

Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds.
Note that a maximum of 100 rules are supported on a bucket.


## Example Usage

```hcl
resource "oci_objectstorage_retention_rule" "test_retention_rule" {
	#Required
	bucket = "${var.retention_rule_bucket}"
	namespace = "${var.retention_rule_namespace}"

	#Optional
	display_name = "${var.retention_rule_display_name}"
	duration {
		#Required
		time_amount = "${var.retention_rule_duration_time_amount}"
		time_unit = "${var.retention_rule_duration_time_unit}"
	}
	time_rule_locked = "${var.retention_rule_time_rule_locked}"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1` 
* `display_name` - (Optional) (Updatable) A user-specified name for the retention rule. Names can be helpful in identifying retention rules.
* `duration` - (Optional) (Updatable) 
	* `time_amount` - (Required) (Updatable) The timeAmount is interpreted in units defined by the timeUnit parameter, and is calculated in relation to each object's Last-Modified timestamp. 
	* `time_unit` - (Required) (Updatable) The unit that should be used to interpret timeAmount.
* `namespace` - (Required) The Object Storage namespace used for the request.
* `time_rule_locked` - (Optional) (Updatable) The date and time as per [RFC 3339](https://tools.ietf.org/html/rfc3339) after which this rule is locked and can only be deleted by deleting the bucket. Once a rule is locked, only increases in the duration are allowed and no other properties can be changed. This property cannot be updated for rules that are in a locked state. Specifying it when a duration is not specified is considered an error. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Import

RetentionRules can be imported using the `id`, e.g.

```
$ terraform import oci_objectstorage_retention_rule.test_retention_rule "id"
```

