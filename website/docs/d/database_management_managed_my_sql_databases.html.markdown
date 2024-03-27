---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_databases"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_databases"
description: |-
  Provides the list of Managed My Sql Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_databases
This data source provides the list of Managed My Sql Databases in Oracle Cloud Infrastructure Database Management service.

Gets the list of Managed MySQL Databases in a specific compartment.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_databases" "test_managed_my_sql_databases" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `managed_my_sql_database_collection` - The list of managed_my_sql_database_collection.

### ManagedMySqlDatabase Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_name` - The name of the MySQL Database.
* `db_version` - The version of the MySQL Database.
* `id` - The OCID of the Managed MySql Database.
* `name` - The name of the Managed MySQL Database.
* `time_created` - The date and time the Managed Database was created.
* `heat_wave_cluster_display_name` - The name of the HeatWave cluster.
* `heat_wave_memory_size` - The total memory belonging to the HeatWave cluster in GBs.
* `heat_wave_node_shape` - Shape of the nodes in the HeatWave cluster.
* `heat_wave_nodes` - The information about an individual HeatWave nodes in the cluster.
	* `id` - The ID associated with the HeatWave node.
	* `status` - The status of the HeatWave node. Indicates whether the status of the node is UP, DOWN, or UNKNOWN at the current time.
	* `time_created` - The date and time the node was created.
* `is_heat_wave_active` - If the HeatWave cluster is active or not.
* `is_heat_wave_enabled` - If HeatWave is enabled for this db system or not.
* `is_lakehouse_enabled` - If HeatWave Lakehouse is enabled for the db system or not.
* `time_created_heat_wave` - The date and time the Managed MySQL Database was created.

