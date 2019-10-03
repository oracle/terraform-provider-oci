---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_databases"
sidebar_current: "docs-oci-datasource-database-autonomous_databases"
description: |-
  Provides the list of Autonomous Databases in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_databases
This data source provides the list of Autonomous Databases in Oracle Cloud Infrastructure Database service.

Gets a list of Autonomous Databases.


## Example Usage

```hcl
data "oci_database_autonomous_databases" "test_autonomous_databases" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	autonomous_container_database_id = "${oci_database_autonomous_container_database.test_autonomous_container_database.id}"
	db_workload = "${var.autonomous_database_db_workload}"
	display_name = "${var.autonomous_database_display_name}"
	is_free_tier = "${var.autonomous_database_is_free_tier}"
	state = "${var.autonomous_database_state}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_workload` - (Optional) A filter to return only autonomous database resources that match the specified workload type.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `is_free_tier` - (Optional) Filter on the value of the resource's 'isFreeTier' property. A value of `true` returns only Always Free resources. A value of `false` excludes Always Free resources from the returned results. Omitting this parameter returns both Always Free and paid resources. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_databases` - The list of autonomous_databases.

### AutonomousDatabase Reference

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
	* `machine_learning_user_management_url` - Oracle Machine Learning User Management URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of CPU cores to be made available to the database.
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. Note that auto scaling is available for [serverless deployments](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. 
* `is_dedicated` - True if the database uses the [dedicated deployment](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm) option. 
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. The default for Autonomous Database using the [shared deployment] is BRING_YOUR_OWN_LICENSE. Note that when provisioning an Autonomous Database using the [dedicated deployment](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm) option, this attribute must be null because the attribute is already set on Autonomous Exadata Infrastructure level. 
* `lifecycle_details` - Information about the current lifecycle state.
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `state` - The current state of the database.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the database was created.
* `time_deletion_of_free_autonomous_database` - The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted. 
* `time_reclamation_of_free_autonomous_database` - The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state. 
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for [serverless deployments](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only.  Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet. 

