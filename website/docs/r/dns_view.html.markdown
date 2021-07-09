---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_view"
sidebar_current: "docs-oci-resource-dns-view"
description: |-
  Provides the View resource in Oracle Cloud Infrastructure DNS service
---

# oci_dns_view
This resource provides the View resource in Oracle Cloud Infrastructure DNS service.

Creates a new view in the specified compartment. Requires a `PRIVATE` scope query parameter.


## Example Usage

```hcl
resource "oci_dns_view" "test_view" {
	#Required
	compartment_id = var.compartment_id
	scope = "PRIVATE"

	#Optional
	defined_tags = var.view_defined_tags
	display_name = var.view_display_name
	freeform_tags = var.view_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the owning compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - (Optional) (Updatable) The display name of the view. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `scope` - (Required) Value must be `PRIVATE` when creating a view for private zones.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the owning compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - The display name of the view. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the view.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the View
	* `update` - (Defaults to 20 minutes), when updating the View
	* `delete` - (Defaults to 20 minutes), when destroying the View


## Import

For legacy Views that were created without using `scope`, these Views can be imported using the `id`, e.g.

```
$ terraform import oci_dns_view.test_view "id"
```

For Views created using `scope`, these Views can be imported using the `id`, e.g.

```
$ terraform import oci_dns_view.test_view "viewId/{viewId}/scope/{scope}"
```

