---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_heat_wave_cluster"
sidebar_current: "docs-oci-resource-mysql-heat_wave_cluster"
description: |-
  Provides the Heat Wave Cluster resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_heat_wave_cluster
This resource provides the Heat Wave Cluster resource in Oracle Cloud Infrastructure MySQL Database service.



## Example Usage

```hcl
resource "oci_mysql_heat_wave_cluster" "test_heat_wave_cluster" {
}
```

## Argument Reference

The following arguments are supported:



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

## Import

HeatWaveCluster can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_heat_wave_cluster.test_heat_wave_cluster "dbSystems/{dbSystemId}/heatWaveCluster" 
```

