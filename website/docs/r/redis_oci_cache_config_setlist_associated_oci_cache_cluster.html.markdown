---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster"
sidebar_current: "docs-oci-resource-redis-oci_cache_config_setlist_associated_oci_cache_cluster"
description: |-
  Provides the Oci Cache Config Setlist Associated Oci Cache Cluster resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster
This resource provides the Oci Cache Config Setlist Associated Oci Cache Cluster resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/OciCacheConfigSetlistAssociatedOciCacheCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Gets a list of associated Oracle Cloud Infrastructure Cache clusters for an Oracle Cloud Infrastructure Cache Config Set.


## Example Usage

```hcl
resource "oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster" "test_oci_cache_config_setlist_associated_oci_cache_cluster" {
	#Required
	oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set.id
}
```

## Argument Reference

The following arguments are supported:

* `oci_cache_config_set_id` - (Required) Unique Oracle Cloud Infrastructure Cache Config Set identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `items` - List of clusters with the same Oracle Cloud Infrastructure Cache Config Set ID.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache Config Setlist Associated Oci Cache Cluster
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache Config Setlist Associated Oci Cache Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache Config Setlist Associated Oci Cache Cluster


## Import

Import is not supported for this resource.

