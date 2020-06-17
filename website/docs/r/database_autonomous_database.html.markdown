---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database"
sidebar_current: "docs-oci-resource-database-autonomous_database"
description: |-
  Provides the Autonomous Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database
This resource provides the Autonomous Database resource in Oracle Cloud Infrastructure Database service.

Creates a new Autonomous Database.


## Example Usage

```hcl
resource "oci_database_autonomous_database" "test_autonomous_database" {
	#Required
	admin_password = "${var.autonomous_database_admin_password}"
	compartment_id = "${var.compartment_id}"
	cpu_core_count = "${var.autonomous_database_cpu_core_count}"
	data_storage_size_in_tbs = "${var.autonomous_database_data_storage_size_in_tbs}"
	db_name = "${var.autonomous_database_db_name}"

	#Optional
	autonomous_container_database_id = "${oci_database_autonomous_container_database.test_autonomous_container_database.id}"
	autonomous_database_backup_id = "${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}"
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	clone_type = "${var.autonomous_database_clone_type}"
	data_safe_status = "${var.autonomous_database_data_safe_status}"
	db_version = "${var.autonomous_database_db_version}"
	db_workload = "${var.autonomous_database_db_workload}"
	defined_tags = "${var.autonomous_database_defined_tags}"
	display_name = "${var.autonomous_database_display_name}"
	freeform_tags = {"Department"= "Finance"}
	is_auto_scaling_enabled = "${var.autonomous_database_is_auto_scaling_enabled}"
	is_data_guard_enabled = "${var.autonomous_database_is_data_guard_enabled}"
	is_dedicated = "${var.autonomous_database_is_dedicated}"
	is_free_tier = "${var.autonomous_database_is_free_tier}"
	is_preview_version_with_service_terms_accepted = "${var.autonomous_database_is_preview_version_with_service_terms_accepted}"
	license_model = "${var.autonomous_database_license_model}"
	nsg_ids = "${var.autonomous_database_nsg_ids}"
	private_endpoint_label = "${var.autonomous_database_private_endpoint_label}"
	source = "${var.autonomous_database_source}"
	source_id = "${oci_database_source.test_source.id}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
	timestamp = "${var.autonomous_database_timestamp}"
	whitelisted_ips = "${var.autonomous_database_whitelisted_ips}"
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Required) (Updatable) The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_database_backup_id` - (Required when source=BACKUP_FROM_ID) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database Backup that you will clone to create a new Autonomous Database.
* `autonomous_database_id` - (Required when source=BACKUP_FROM_TIMESTAMP) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `clone_type` - (Required when source=BACKUP_FROM_ID | BACKUP_FROM_TIMESTAMP | DATABASE) The Autonomous Database clone type.
	* `FULL` - This option creates a new database that includes all source database data.
	* `METADATA` - This option creates a new database that includes the source database schema and select metadata, but not the source database data.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
* `cpu_core_count` - (Required) (Updatable) The number of OCPU cores to be made available to the database. This input is ignored for Always Free resources.
* `data_safe_status` - (Optional) (Updatable) Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_tbs` - (Required) (Updatable) The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. This input is ignored for Always Free resources.
* `db_name` - (Required) The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
* `db_version` - (Optional) (Updatable) A valid Oracle Database version for Autonomous Database.
* `db_workload` - (Optional) The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_auto_scaling_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database OCPU core count. The default value is `FALSE`. Note that auto scaling is available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. 
* `is_data_guard_enabled` - (Optional) (Updatable) Indicates whether the Autonomous Database has Data Guard enabled.
* `is_dedicated` - (Optional) True if the database is on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm). 
* `is_free_tier` - (Optional) (Updatable) Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_preview_version_with_service_terms_accepted` - (Optional) If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI). 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `nsg_ids` - (Optional) (Updatable) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `private_endpoint_label` - (Optional) (Updatable) The private endpoint label for the resource.
* `source` - (Optional) The source of the database: Use `NONE` for creating a new Autonomous Database. Use `DATABASE` for creating a new Autonomous Database by cloning an existing Autonomous Database.

	For Autonomous Databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), the following cloning options are available: Use `BACKUP_FROM_ID` for creating a new Autonomous Database from a specified backup. Use `BACKUP_FROM_TIMESTAMP` for creating a point-in-time Autonomous Database clone using backups. For more information, see [Cloning an Autonomous Database](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/adbcloning.htm). 
* `source_id` - (Required when source=DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `timestamp` - (Required when source=BACKUP_FROM_TIMESTAMP) The timestamp specified for the point-in-time clone of the source Autonomous Database. The timestamp must be in the past.
* `whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). This feature is available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.0.0/16"]` To remove all whitelisted IPs, set the field to a list with an empty string `[""]`.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
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
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. Note that auto scaling is available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. 
* `is_data_guard_enabled` - Indicates whether the Autonomous Database has Data Guard enabled.
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm). 
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `lifecycle_details` - Information about the current lifecycle state.
* `nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `private_endpoint` - The private endpoint for the resource.
* `private_endpoint_ip` - The private endpoint Ip address for the resource.
* `private_endpoint_label` - The private endpoint label for the resource.
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `standby_db` - 
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `state` - The current state of the Autonomous Database.
* `state` - The current state of the Autonomous Database.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Autonomous Database was created.
* `time_deletion_of_free_autonomous_database` - The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted. 
* `time_maintenance_begin` - The date and time when maintenance will begin.
* `time_maintenance_end` - The date and time when maintenance will end.
* `time_reclamation_of_free_autonomous_database` - The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state. 
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for databases on [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.

	To add the whitelist VCN specific subnet or IP, use a semicoln ';' as a deliminator to add the VCN specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.1.1","ocid1.vcn.oc1.sea.aaaaaaaard2hfx2nn3e5xeo6j6o62jga44xjizkw;1.1.0.0/16"]` 

## Import

AutonomousDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database.test_autonomous_database "id"
```

