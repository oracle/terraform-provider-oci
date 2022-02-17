---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system_compute_performances"
sidebar_current: "docs-oci-datasource-database-db_system_compute_performances"
description: |-
  Provides the list of Db System Compute Performances in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_system_compute_performances
This data source provides the list of Db System Compute Performances in Oracle Cloud Infrastructure Database service.

Gets a list of expected compute performance parameters for a virtual machine DB system based on system configuration.


## Example Usage

```hcl
data "oci_database_db_system_compute_performances" "test_db_system_compute_performances" {

	#Optional
	db_system_shape = var.db_system_compute_performance_db_system_shape
}
```

## Argument Reference

The following arguments are supported:

* `db_system_shape` - (Optional) If provided, filters the results to the set of database versions which are supported for the given shape.


## Attributes Reference

The following attributes are exported:

* `db_system_compute_performances` - The list of db_system_compute_performances.

### DbSystemComputePerformance Reference

The following attributes are exported:

* `compute_performance_list` - List of Compute performance details for the specified DB system shape.
	* `cpu_core_count` - The number of OCPU cores available.
	* `memory_in_gbs` - The amount of memory allocated for the VMDB System.
	* `network_bandwidth_in_gbps` - The network bandwidth of the VMDB system in gbps.
	* `network_iops` - IOPS for the VMDB System.
	* `network_throughput_in_mbps` - Network throughput for the VMDB System.
* `shape` - The shape of the DB system.

