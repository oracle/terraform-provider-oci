---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database"
description: |-
  Provides details about a specific Managed My Sql Database in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database
This data source provides details about a specific Managed My Sql Database resource in Oracle Cloud Infrastructure Database Management service.

Retrieves General Information for given MySQL Instance.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database" "test_managed_my_sql_database" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of ManagedMySqlDatabase.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_name` - The name of the MySQL Database.
* `db_version` - The version of the MySQL Database.
* `id` - The OCID of the Managed MySQL Database.
* `time_created` - The date and time the Managed MySQL Database was created.
* `heat_wave_cluster_display_name` - The name of the HeatWave cluster.
* `heat_wave_memory_size` - The total memory belonging to the HeatWave cluster in GBs.
* `heat_wave_node_shape` - The shape of the nodes in the HeatWave cluster.
* `heat_wave_nodes` - The information about individual HeatWave nodes in the cluster.
	* `id` - The ID associated with the HeatWave node.
	* `status` - The status of the HeatWave node. Indicates whether the status of the node is UP, DOWN, or UNKNOWN at the current time.
	* `time_created` - The date and time the HeatWave node was created.
* `is_heat_wave_active` - Indicates whether the HeatWave cluster is active or not.
* `is_heat_wave_enabled` - Indicates whether HeatWave is enabled for the MySQL Database System or not.
* `is_lakehouse_enabled` - Indicates whether HeatWave Lakehouse is enabled for the MySQL Database System or not.
* `name` - The name of the Managed MySQL Database.
* `time_created_heat_wave` - The date and time the Managed MySQL Database was created.

