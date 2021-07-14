---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_resolvers"
sidebar_current: "docs-oci-datasource-dns-resolvers"
description: |-
  Provides the list of Resolvers in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_resolvers
This data source provides the list of Resolvers in Oracle Cloud Infrastructure DNS service.

Gets a list of all resolvers within a compartment. The collection can
be filtered by display name, id, or lifecycle state. It can be sorted
on creation time or displayName both in ASC or DESC order. Note that
when no lifecycleState query parameter is provided, the collection
does not include resolvers in the DELETED lifecycleState to be consistent
with other operations of the API. Requires a `PRIVATE` scope query parameter.


## Example Usage

```hcl
data "oci_dns_resolvers" "test_resolvers" {
	#Required
	compartment_id = var.compartment_id
	scope = "PRIVATE"

	#Optional
	display_name = var.resolver_display_name
	id = var.resolver_id
	state = var.resolver_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `display_name` - (Optional) The displayName of a resource.
* `id` - (Optional) The OCID of a resource.
* `scope` - (Required) Value must be `PRIVATE` when listing private name resolvers.
* `state` - (Optional) The state of a resource.


## Attributes Reference

The following attributes are exported:

* `resolvers` - The list of resolvers.

### Resolver Reference

The following attributes are exported:

* `attached_vcn_id` - The OCID of the attached VCN. 
* `compartment_id` - The OCID of the owning compartment.
* `default_view_id` - The OCID of the default view. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Operations": {"CostCenter": "42"}}` 
* `display_name` - The display name of the resolver. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).

	 **Example:** `{"Department": "Finance"}` 
* `id` - The OCID of the resolver.
* `is_protected` - A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `time_updated` - The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 

