---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_heat_wave_cluster"
sidebar_current: "docs-oci-resource-mysql-heat_wave_cluster"
description: |-
  Provides the HeatWave cluster resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_heat_wave_cluster
This resource provides the HeatWave cluster resource in Oracle Cloud Infrastructure MySQL Database service.

Updates the HeatWave cluster.


## Example Usage

```hcl
resource "oci_mysql_heat_wave_cluster" "test_heat_wave_cluster" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
	cluster_size = var.heat_wave_cluster_cluster_size
	shape_name = oci_mysql_shape.test_shape.name
}
```

## Argument Reference

The following arguments are supported:

* `cluster_size` - (Required) (Updatable) A change to the number of nodes in the HeatWave cluster will result in the entire cluster being torn down and re-created with the new cluster of nodes. This may result in a significant downtime for the analytics capability while the HeatWave cluster is re-provisioned.
* `db_system_id` - (Required) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `shape_name` - (Required) (Updatable) A change to the shape of the nodes in the HeatWave cluster will result in the entire cluster being torn down and re-created with Compute instances of the new Shape. This may result in significant downtime for the analytics capability while the HeatWave cluster is re-provisioned.
* `state` - (Optional) (Updatable) The target state for the HeatWave cluster. Could be set to `ACTIVE` or `INACTIVE`.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_nodes` - A HeatWave node is a compute host that is part of a HeatWave cluster.
	* `node_id` - The ID of the node within MySQL HeatWave cluster.
	* `state` - The current state of the MySQL HeatWave node.
	* `time_created` - The date and time the MySQL HeatWave node was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
	* `time_updated` - The date and time the MySQL HeatWave node was updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `cluster_size` - The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster. 
* `db_system_id` - The OCID of the parent DB System this HeatWave cluster is attached to.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `shape_name` - The shape determines resources to allocate to the HeatWave nodes - CPU cores, memory. 
* `state` - The current state of the HeatWave cluster.
* `time_created` - The date and time the HeatWave cluster was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the HeatWave cluster was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Heat Wave Cluster
	* `update` - (Defaults to 1 hours), when updating the Heat Wave Cluster
	* `delete` - (Defaults to 1 hours), when destroying the Heat Wave Cluster


## Import

HeatWaveCluster can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_heat_wave_cluster.test_heat_wave_cluster "dbSystem/{dbSystemId}/heatWaveCluster" 
```

