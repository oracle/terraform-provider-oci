---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_saved_query"
sidebar_current: "docs-oci-datasource-cloud_guard-saved_query"
description: |-
  Provides details about a specific Saved Query in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_saved_query
This data source provides details about a specific Saved Query resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a SavedQuery resource identified by savedQueryId.

## Example Usage

```hcl
data "oci_cloud_guard_saved_query" "test_saved_query" {
	#Required
	saved_query_id = oci_cloud_guard_saved_query.test_saved_query.id
}
```

## Argument Reference

The following arguments are supported:

* `saved_query_id` - (Required) Saved query OCID


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment OCID of the saved query
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the saved query
* `display_name` - Display name of the saved query
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - OCID for the saved query
* `query` - The saved query expression
* `state` - The current lifecycle state of the resource
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the saved query was created. Format defined by RFC3339.
* `time_updated` - The date and time the saved query was updated. Format defined by RFC3339.

