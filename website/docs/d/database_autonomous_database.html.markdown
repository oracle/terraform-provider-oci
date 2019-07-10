---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database"
sidebar_current: "docs-oci-datasource-database-autonomous_database"
description: |-
  Provides details about a specific Autonomous Database in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database
This data source provides details about a specific Autonomous Database resource in Oracle Cloud Infrastructure Database service.

Gets the details of the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_container_database_id` - The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Database. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `dedicated` - The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
* `connection_urls` - 
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of CPU cores to be made available to the database.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. 
* `is_dedicated` - True if it is dedicated database. 
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. The default is BRING_YOUR_OWN_LICENSE. 
* `lifecycle_details` - Information about the current lifecycle state.
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `state` - The current state of the database.
* `time_created` - The date and time the database was created.
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `whitelisted_ips` - The client IP access control list (ACL). Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet.

