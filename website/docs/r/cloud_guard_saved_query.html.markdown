---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_saved_query"
sidebar_current: "docs-oci-resource-cloud_guard-saved_query"
description: |-
  Provides the Saved Query resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_saved_query
This resource provides the Saved Query resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a SavedQuery resource.


## Example Usage

```hcl
resource "oci_cloud_guard_saved_query" "test_saved_query" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.saved_query_display_name
	query = var.saved_query_query

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.saved_query_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment OCID of the saved query
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the saved query
* `display_name` - (Required) (Updatable) Display name of the saved query
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `query` - (Required) (Updatable) The adhoc query expression that is run


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Saved Query
	* `update` - (Defaults to 20 minutes), when updating the Saved Query
	* `delete` - (Defaults to 20 minutes), when destroying the Saved Query


## Import

SavedQueries can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_saved_query.test_saved_query "id"
```

