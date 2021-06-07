---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_heat_wave_cluster"
sidebar_current: "docs-oci-datasource-mysql-heat_wave_cluster"
description: |-
  Provides details about a specific HeatWave cluster in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_heat_wave_cluster
This data source provides details about a specific HeatWave cluster resource in Oracle Cloud Infrastructure MySQL Database service.

Gets information about the HeatWave cluster.

## Example Usage

```hcl
data "oci_mysql_heat_wave_cluster" "test_heat_wave_cluster" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

