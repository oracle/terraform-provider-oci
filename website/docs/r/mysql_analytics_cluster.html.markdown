---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_analytics_cluster"
sidebar_current: "docs-oci-resource-mysql-analytics_cluster"
description: |-
  Provides the Analytics Cluster resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_analytics_cluster
This resource provides the Analytics Cluster resource in Oracle Cloud Infrastructure MySQL Database service.

DEPRECATED -- please use HeatWave API instead.
Updates the Analytics Cluster.


## Example Usage

```hcl
resource "oci_mysql_analytics_cluster" "test_analytics_cluster" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
	cluster_size = var.analytics_cluster_cluster_size
	shape_name = oci_mysql_shape.test_shape.name
}
```

## Argument Reference

The following arguments are supported:

* `cluster_size` - (Required) (Updatable) A change to the number of nodes in the Analytics Cluster will result in the entire cluster being torn down and re-created with the new cluster of nodes. This may result in a significant downtime for the analytics capability while the Analytics Cluster is re-provisioned. 
* `db_system_id` - (Required) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `shape_name` - (Required) (Updatable) A change to the shape of the nodes in the Analytics Cluster will result in the entire cluster being torn down and re-created with Compute instances of the new Shape. This may result in significant downtime for the analytics capability while the Analytics Cluster is re-provisioned. 
* `state` - (Optional) (Updatable) The target state for the Analytics Cluster. Could be set to `ACTIVE` or `INACTIVE`. 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_nodes` - An Analytics Cluster Node is a compute host that is part of an Analytics Cluster.
	* `node_id` - The ID of the node within MySQL Analytics Cluster.
	* `state` - The current state of the MySQL Analytics Cluster node.
	* `time_created` - The date and time the MySQL Analytics Cluster node was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
	* `time_updated` - The date and time the MySQL Analytics Cluster node was updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the Analytics Cluster. 
* `db_system_id` - The OCID of the parent DB System this Analytics Cluster is attached to.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `shape_name` - The shape determines resources to allocate to the Analytics Cluster nodes - CPU cores, memory. 
* `state` - The current state of the Analytics Cluster.
* `time_created` - The date and time the Analytics Cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The time the Analytics Cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Analytics Cluster
	* `update` - (Defaults to 1 hours), when updating the Analytics Cluster
	* `delete` - (Defaults to 1 hours), when destroying the Analytics Cluster


## Import

AnalyticsCluster can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_analytics_cluster.test_analytics_cluster "dbSystems/{dbSystemId}/analyticsCluster" 
```

