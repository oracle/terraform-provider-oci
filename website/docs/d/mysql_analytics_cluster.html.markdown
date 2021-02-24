---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_analytics_cluster"
sidebar_current: "docs-oci-datasource-mysql-analytics_cluster"
description: |-
  Provides details about a specific Analytics Cluster in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_analytics_cluster
This data source provides details about a specific Analytics Cluster resource in Oracle Cloud Infrastructure MySQL Database service.

Gets information about the Analytics Cluster.

## Example Usage

```hcl
data "oci_mysql_analytics_cluster" "test_analytics_cluster" {
	#Required
	db_system_id = oci_database_db_system.test_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

