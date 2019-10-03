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
Note: `whitelisted_ips` cannot be used during creation.


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
	clone_type = "${var.autonomous_database_clone_type}"
	db_workload = "${var.autonomous_database_db_workload}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.autonomous_database_display_name}"
	freeform_tags = {"Department"= "Finance"}
	is_auto_scaling_enabled = "${var.autonomous_database_is_auto_scaling_enabled}"
	is_dedicated = "${var.autonomous_database_is_dedicated}"
	is_free_tier = "${var.autonomous_database_is_free_tier}"
	is_preview_version_with_service_terms_accepted = "${var.autonomous_database_is_preview_version_with_service_terms_accepted}"
	license_model = "${var.autonomous_database_license_model}"
	source = "${var.autonomous_database_source}"
	source_id = "${oci_database_source.test_source.id}"
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Required) (Updatable) The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `clone_type` - (Required when source=DATABASE) The clone type when cloning an Autonomous Database using a `source`. Supported values:
    * `FULL` - This option creates a new database that includes all source database data.
    * `METADATA` - This option creates a new database that includes the source database schema and select metadata, but not the source database data.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the autonomous database.
* `cpu_core_count` - (Required) (Updatable) The number of CPU Cores to be made available to the database. This input is ignored for Always Free resources.
* `data_storage_size_in_tbs` - (Required) (Updatable) The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. This input is ignored for Always Free resources.
* `db_name` - (Required) The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
* `db_workload` - (Optional) The autonomous database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse. The default is OLTP.
    * `OLTP` - For Autonomous Database workload type.
    * `DW` - For Autonomous Data Warehouse workload type.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_auto_scaling_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database CPU core count. The default value is false. Note that auto scaling is available for [serverless deployments](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI) only. 
* `is_dedicated` - (Optional) True if the database uses the [dedicated deployment](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm) option. 
* `is_free_tier` - (Optional) (Updatable) Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_preview_version_with_service_terms_accepted` - (Optional) If set to true, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for [serverless deployments](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI). 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. The default for Autonomous Database using the [shared deployment] is BRING_YOUR_OWN_LICENSE. Note that when provisioning an Autonomous Database using the [dedicated deployment](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm) option, this attribute must be null because the attribute is already set on Autonomous Exadata Infrastructure level. 
* `source` - (Optional) The source of the database: Use NONE for creating a new Autonomous Database. Use DATABASE for creating a new Autonomous Database by cloning an existing Autonomous Database. 
* `source_id` - (Required when source=DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
	* `machine_learning_user_management_url` - Oracle Machine Learning User Management URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of CPU cores to be made available to the database.
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

## Import

AutonomousDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database.test_autonomous_database "id"
```

