---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_view"
sidebar_current: "docs-oci-datasource-dns-view"
description: |-
  Provides details about a specific View in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_view
This data source provides details about a specific View resource in Oracle Cloud Infrastructure DNS service.

Gets information about a specific view. Note that attempting to get a
view in the DELETED lifecycleState will result in a `404` response to be
consistent with other operations of the API. Requires a `PRIVATE` scope query parameter.


## Example Usage

```hcl
data "oci_dns_view" "test_view" {
	#Required
	view_id = oci_dns_view.test_view.id
	scope = "PRIVATE"
}
```

## Argument Reference

The following arguments are supported:

* `scope` - (Required) Value must be `PRIVATE` when listing views for private zones.
* `view_id` - (Required) The OCID of the target view.


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

