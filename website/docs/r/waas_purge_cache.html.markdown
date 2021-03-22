---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_purge_cache"
sidebar_current: "docs-oci-resource-waas-purge_cache"
description: |-
  Provides the Purge Cache resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# oci_waas_purge_cache
This resource provides the Purge Cache resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Performs a purge of the cache for each specified resource. If no resources are passed, the cache for the entire Web Application Firewall will be purged.
For more information, see [Caching Rules](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/cachingrules.htm#purge).

## Example Usage

```hcl
resource "oci_waas_purge_cache" "test_purge_cache" {
	#Required
	waas_policy_id = oci_waas_waas_policy.test_waas_policy.id

	#Optional
	resources = var.purge_cache_resources
}
```

## Argument Reference

The following arguments are supported:

* `resources` - (Optional) A resource to purge, specified by either a hostless absolute path starting with a single slash (Example: `/path/to/resource`) or by a relative path in which the first component will be interpreted as a domain protected by the WAAS policy (Example: `example.com/path/to/resource`).
* `waas_policy_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WAAS policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Purge Cache
	* `update` - (Defaults to 1 hours), when updating the Purge Cache
	* `delete` - (Defaults to 1 hours), when destroying the Purge Cache

## Import

Import is not supported for this resource.

