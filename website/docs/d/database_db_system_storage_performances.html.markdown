---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system_storage_performances"
sidebar_current: "docs-oci-datasource-database-db_system_storage_performances"
description: |-
  Provides the list of Db System Storage Performances in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_system_storage_performances
This data source provides the list of Db System Storage Performances in Oracle Cloud Infrastructure Database service.

Gets a list of possible expected storage performance parameters of a VMDB System based on Configuration.


## Example Usage

```hcl
data "oci_database_db_system_storage_performances" "test_db_system_storage_performances" {
	#Required
	storage_management = var.db_system_storage_performance_storage_management

	#Optional
	shape_type = var.db_system_storage_performance_shape_type
}
```

## Argument Reference

The following arguments are supported:

* `shape_type` - (Optional) Optional. Filters the performance results by shape type.
* `storage_management` - (Required) The DB system storage management option. Used to list database versions available for that storage manager. Valid values are `ASM` and `LVM`.
	* ASM specifies Oracle Automatic Storage Management
	* LVM specifies logical volume manager, sometimes called logical disk manager. 


## Attributes Reference

The following attributes are exported:

* `db_system_storage_performances` - The list of db_system_storage_performances.

### DbSystemStoragePerformance Reference

The following attributes are exported:

* `data_storage_performance_list` - List of storage performance for the DATA disks
	* `balanced_disk_performance` - Representation of disk performance detail parameters. 
		* `disk_iops` - Disk IOPS in thousands.
		* `disk_throughput_in_mbps` - Disk Throughput in Mbps.
	* `high_disk_performance` - Representation of disk performance detail parameters. 
		* `disk_iops` - Disk IOPS in thousands.
		* `disk_throughput_in_mbps` - Disk Throughput in Mbps.
	* `size_in_gbs` - Size in GBs.
* `reco_storage_performance_list` - List of storage performance for the RECO disks
	* `balanced_disk_performance` - Representation of disk performance detail parameters. 
		* `disk_iops` - Disk IOPS in thousands.
		* `disk_throughput_in_mbps` - Disk Throughput in Mbps.
	* `high_disk_performance` - Representation of disk performance detail parameters. 
		* `disk_iops` - Disk IOPS in thousands.
		* `disk_throughput_in_mbps` - Disk Throughput in Mbps.
	* `size_in_gbs` - Size in GBs.
* `shape_type` - ShapeType of the DbSystems,INTEL or AMD

