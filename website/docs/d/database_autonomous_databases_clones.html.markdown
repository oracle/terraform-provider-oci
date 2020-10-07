---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_databases_clones"
sidebar_current: "docs-oci-datasource-database-autonomous_databases_clones"
description: |-
  Provides the list of Autonomous Databases Clones in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_databases_clones
This data source provides the list of Autonomous Databases Clones in Oracle Cloud Infrastructure Database service.

Lists the Autonomous Database clones for the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_databases_clones" "test_autonomous_databases_clones" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	compartment_id = var.compartment_id

	#Optional
	clone_type = var.autonomous_databases_clone_clone_type
	display_name = var.autonomous_databases_clone_display_name
	state = var.autonomous_databases_clone_state
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `clone_type` - (Optional) A filter to return only resources that match the given clone type exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_databases` - The list of autonomous_databases.

### AutonomousDatabasesClone Reference

The following attributes are exported:

* `autonomous_container_database_id` - The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `available_upgrade_versions` - List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Database. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `dedicated` - The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
* `connection_urls` - 
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `machine_learning_user_management_url` - Oracle Machine Learning user management URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of OCPU cores to be made available to the database.
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `failed_data_recovery_in_seconds` - Indicates the number of seconds of data loss for a Data Guard failover.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. 
* `is_data_guard_enabled` - Indicates whether the Autonomous Database has Data Guard enabled.
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm). 
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `is_refreshable_clone` - Indicates whether the Autonomous Database is a refreshable clone.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `lifecycle_details` - Information about the current lifecycle state.
* `nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `open_mode` - The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
* `private_endpoint` - The private endpoint for the resource.
* `private_endpoint_ip` - The private endpoint Ip address for the resource.
* `private_endpoint_label` - The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
* `refreshable_mode` - The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `refreshable_status` - The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
* `standby_db` - 
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
* `state` - The current state of the Autonomous Database.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Autonomous Database was created.
* `time_deletion_of_free_autonomous_database` - The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted. 
* `time_maintenance_begin` - The date and time when maintenance will begin.
* `time_maintenance_end` - The date and time when maintenance will end.
* `time_of_last_failover` - The timestamp of the last failover operation.
* `time_of_last_refresh` - The date and time when last refresh happened.
* `time_of_last_refresh_point` - The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.
* `time_of_last_switchover` - The timestamp of the last switchover operation for the Autonomous Database.
* `time_of_next_refresh` - The date and time of next refresh.
* `time_reclamation_of_free_autonomous_database` - The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state. 
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.

	To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs. For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` 

