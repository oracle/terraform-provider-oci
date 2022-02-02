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
	compartment_id = var.compartment_id
	db_name = var.autonomous_database_db_name

	#Optional
	admin_password = var.autonomous_database_admin_password
	are_primary_whitelisted_ips_used = var.autonomous_database_are_primary_whitelisted_ips_used
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	autonomous_database_backup_id = oci_database_autonomous_database_backup.test_autonomous_database_backup.id
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	autonomous_maintenance_schedule_type = var.autonomous_database_autonomous_maintenance_schedule_type
	clone_type = var.autonomous_database_clone_type
	cpu_core_count = var.autonomous_database_cpu_core_count
	customer_contacts {

		#Optional
		email = var.autonomous_database_customer_contacts_email
	}
	data_safe_status = var.autonomous_database_data_safe_status
	data_storage_size_in_gb = var.autonomous_database_data_storage_size_in_gb
	data_storage_size_in_tbs = var.autonomous_database_data_storage_size_in_tbs
	db_version = var.autonomous_database_db_version
	db_workload = var.autonomous_database_db_workload
	defined_tags = var.autonomous_database_defined_tags
	display_name = var.autonomous_database_display_name
	freeform_tags = {"Department"= "Finance"}
	is_access_control_enabled = var.autonomous_database_is_access_control_enabled
	is_auto_scaling_enabled = var.autonomous_database_is_auto_scaling_enabled
	is_data_guard_enabled = var.autonomous_database_is_data_guard_enabled
	is_dedicated = var.autonomous_database_is_dedicated
	is_free_tier = var.autonomous_database_is_free_tier
	is_mtls_connection_required = var.autonomous_database_is_mtls_connection_required
	is_preview_version_with_service_terms_accepted = var.autonomous_database_is_preview_version_with_service_terms_accepted
	kms_key_id = oci_kms_key.test_key.id
	license_model = var.autonomous_database_license_model
	nsg_ids = var.autonomous_database_nsg_ids
	ocpu_count = var.autonomous_database_ocpu_count
	private_endpoint_label = var.autonomous_database_private_endpoint_label
	refreshable_mode = var.autonomous_database_refreshable_mode
	scheduled_operations {
		#Required
		day_of_week {
			#Required
			name = var.autonomous_database_scheduled_operations_day_of_week_name
		}

		#Optional
		scheduled_start_time = var.autonomous_database_scheduled_operations_scheduled_start_time
		scheduled_stop_time = var.autonomous_database_scheduled_operations_scheduled_stop_time
	}
	source = var.autonomous_database_source
	source_id = oci_database_source.test_source.id
	standby_whitelisted_ips = var.autonomous_database_standby_whitelisted_ips
	subnet_id = oci_core_subnet.test_subnet.id
	timestamp = var.autonomous_database_timestamp
	vault_id = oci_kms_vault.test_vault.id
	whitelisted_ips = var.autonomous_database_whitelisted_ips
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Optional) (Updatable) The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. The password is mandatory if source value is "BACKUP_FROM_ID", "BACKUP_FROM_TIMESTAMP", "DATABASE" or "NONE".
* `are_primary_whitelisted_ips_used` - (Optional) (Updatable) This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby. It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary. 
* `autonomous_container_database_id` - (Optional) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_database_backup_id` - (Required when source=BACKUP_FROM_ID) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database Backup that you will clone to create a new Autonomous Database.
* `autonomous_database_id` - (Required when source=BACKUP_FROM_TIMESTAMP) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `autonomous_maintenance_schedule_type` - (Optional) The maintenance schedule type of the Autonomous Database on shared Exadata infrastructure. The EARLY maintenance schedule of this Autonomous Database follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle. 
* `clone_type` - (Required when source=BACKUP_FROM_ID | BACKUP_FROM_TIMESTAMP | DATABASE) The Autonomous Database clone type.
	* `FULL` - This option creates a new database that includes all source database data.
	* `METADATA` - This option creates a new database that includes the source database schema and select metadata, but not the source database data.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
* `cpu_core_count` - (Optional) (Updatable) The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details. 

	**Note:** This parameter cannot be used with the `ocpuCount` parameter. This input is ignored for Always Free resources.
* `customer_contacts` - (Optional) (Updatable) Customer Contacts.
	* `email` - (Optional) (Updatable) The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_safe_status` - (Optional) (Updatable) Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_gb` - (Optional) (Updatable) The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

	**Notes** 
	* This parameter is only supported for dedicated Exadata infrastructure.
	* This parameter cannot be used with the `dataStorageSizeInTBs` parameter. 
* `data_storage_size_in_tbs` - (Optional) (Updatable) The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details. 

	**Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter. This input is ignored for Always Free resources.
* `db_name` - (Required) The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
* `db_version` - (Optional) (Updatable) A valid Oracle Database version for Autonomous Database.`db_workload` AJD and APEX are only supported for `db_version` `19c` and above.
* `db_workload` - (Optional) (Updatable) The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type. *Note: `db_workload` can only be updated from AJD to OLTP or from a free OLTP to AJD.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_access_control_enabled` - (Optional) (Updatable) Indicates if the database-level access control is enabled. If disabled, database access is defined by the network security rules. If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional, if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console. When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.

	This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. 
* `is_auto_scaling_enabled` - (Optional) (Updatable) Indicates if auto scaling is enabled for the Autonomous Database OCPU core count. The default value is `FALSE`. 
* `is_data_guard_enabled` - (Optional) (Updatable) Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to  Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dedicated` - (Optional) True if the database is on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm). 
* `is_free_tier` - (Optional) (Updatable) Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. When `db_workload` is `AJD` or `APEX` it cannot be `true`.
* `is_mtls_connection_required` - (Optional) (Updatable) Indicates whether the Autonomous Database requires mTLS connections.
* `is_preview_version_with_service_terms_accepted` - (Optional) If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `kms_key_id` - (Optional) The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `is_refreshable_clone` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) True for creating a refreshable clone and False for detaching the clone from source Autonomous Database. Detaching is one time operation and clone becomes a regular Autonomous Database.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. It is a required field when `db_workload` is AJD and needs to be set to `LICENSE_INCLUDED` as AJD does not support default `license_model` value `BRING_YOUR_OWN_LICENSE`.
* `nsg_ids` - (Optional) (Updatable) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `ocpu_count` - (Optional) (Updatable) The number of OCPU cores to be made available to the database. 

	The following points apply:
	* For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	* To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.

	For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

	**Note:** This parameter cannot be used with the `cpuCoreCount` parameter. 
* `operations_insights_status` - (Optional) (Updatable) Status of Operations Insights for this Autonomous Database. Values supported are `ENABLED` and `NOT_ENABLED`
* `private_endpoint_label` - (Optional) (Updatable) The private endpoint label for the resource.
* `refreshable_mode` - (Applicable when source=CLONE_TO_REFRESHABLE) (Updatable) The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `scheduled_operations` - (Optional) (Updatable) list of scheduled operations
	* `day_of_week` - (Required) (Updatable) Day of the week.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `scheduled_start_time` - (Optional) (Updatable) auto start time. value must be of ISO-8601 format "HH:mm"
	* `scheduled_stop_time` - (Optional) (Updatable) auto stop time. value must be of ISO-8601 format "HH:mm"
* `source` - (Optional) The source of the database: Use `NONE` for creating a new Autonomous Database. Use `DATABASE` for creating a new Autonomous Database by cloning an existing Autonomous Database. Use `CROSS_REGION_DATAGUARD` to create a standby Data Guard database in another region.

	For Autonomous Databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), the following cloning options are available: Use `BACKUP_FROM_ID` for creating a new Autonomous Database from a specified backup. Use `BACKUP_FROM_TIMESTAMP` for creating a point-in-time Autonomous Database clone using backups. For more information, see [Cloning and Moving an Autonomous Database](https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/clone-autonomous-database.html#GUID-D771796F-5081-4CFB-A7FF-0F893EABD7BC). 
* `source_id` - (Required when source=CLONE_TO_REFRESHABLE | DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that you will clone to create a new Autonomous Database.
* `standby_whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. 
* `state` - (Optional) (Updatable) The current state of the Autonomous Database. Could be set to AVAILABLE or STOPPED
* `subnet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `timestamp` - (Required when source=BACKUP_FROM_TIMESTAMP) The timestamp specified for the point-in-time clone of the source Autonomous Database. The timestamp must be in the past.
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `whitelisted_ips` - (Optional) (Updatable) The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. To remove all whitelisted IPs, set the field to a list with an empty string `[""]`.

* `switchover_to` - (Optional) It is applicable only when `is_data_guard_enabled` is true. Could be set to `PRIMARY` or `STANDBY`. Default value is `PRIMARY`.
* `rotate_key_trigger` - (Optional) (Updatable) An optional property when flipped triggers rotation of KMS key. It is only applicable on dedicated databases i.e. where `is_dedicated` is true.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apex_details` - Information about Oracle APEX Application Development.
	* `apex_version` - The Oracle APEX Application Development version.
	* `ords_version` - The Oracle REST Data Services (ORDS) version.
* `are_primary_whitelisted_ips_used` - This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby. It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary. 
* `autonomous_container_database_id` - The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_maintenance_schedule_type` - The maintenance schedule type of the Autonomous Database on shared Exadata infrastructure. The EARLY maintenance schedule of this Autonomous Database follows a schedule that applies patches prior to the REGULAR schedule.The REGULAR maintenance schedule of this Autonomous Database follows the normal cycle. 
* `available_upgrade_versions` - List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.
* `backup_config` - Autonomous Database configuration details for storing [manual backups](https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/backup-restore.html#GUID-9035DFB8-4702-4CEB-8281-C2A303820809) in the [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) service. 
	* `manual_backup_bucket_name` - Name of [Object Storage](https://docs.cloud.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) bucket to use for storing manual backups.
	* `manual_backup_type` - The manual backup destination type.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	* `all_connection_strings` - Returns all connection strings that can be used to connect to the Autonomous Database. For more information, please see [Predefined Database Service Names for Autonomous Transaction Processing](https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE) 
	* `dedicated` - The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `high` - The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	* `low` - The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	* `medium` - The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
	* `profiles` - A list of connection string profiles to allow clients to group, filter and select connection string values based on structured metadata. 
		* `consumer_group` - Consumer group used by the connection.
		* `display_name` - A user-friendly name for the connection.
		* `host_format` - Host format used in connection string.
		* `protocol` - Protocol used by the connection.
		* `session_mode` - Specifies whether the listener performs a direct hand-off of the session, or redirects the session. In RAC deployments where SCAN is used, sessions are redirected to a Node VIP. Use `DIRECT` for direct hand-offs. Use `REDIRECT` to redirect the session.
		* `syntax_format` - Specifies whether the connection string is using the long (`LONG`), Easy Connect (`EZCONNECT`), or Easy Connect Plus (`EZCONNECTPLUS`) format. Autonomous Databases on shared Exadata infrastructure always use the long format. 
		* `tls_authentication` - Specifies whether the TLS handshake is using one-way (`SERVER`) or mutual (`MUTUAL`) authentication.
		* `value` - Connection string value.
* `connection_urls` - The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).  Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}` 
	* `apex_url` - Oracle Application Express (APEX) URL.
	* `graph_studio_url` - The URL of the Graph Studio for the Autonomous Database.
	* `machine_learning_user_management_url` - Oracle Machine Learning user management URL.
	* `sql_dev_web_url` - Oracle SQL Developer Web URL.
* `cpu_core_count` - The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details. 

	**Note:** This parameter cannot be used with the `ocpuCount` parameter. 
* `customer_contacts` - Customer Contacts.
	* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_safe_status` - Status of the Data Safe registration for this Autonomous Database. Could be REGISTERED or NOT_REGISTERED.
* `data_storage_size_in_gb` - The quantity of data in the database, in gigabytes.
* `data_storage_size_in_tbs` - The quantity of data in the database, in terabytes.
* `database_management_status` - Status of Database Management for this Autonomous Database.
* `db_name` - The database name.
* `db_version` - A valid Oracle Database version for Autonomous Database.
* `db_workload` - The Autonomous Database workload type. The following values are valid:
	* OLTP - indicates an Autonomous Transaction Processing database
	* DW - indicates an Autonomous Data Warehouse database
	* AJD - indicates an Autonomous JSON Database
	* APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Database. The name does not have to be unique.
* `failed_data_recovery_in_seconds` - Indicates the number of seconds of data loss for a Data Guard failover.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `infrastructure_type` - The infrastructure type this resource belongs to.
* `is_access_control_enabled` - Indicates if the database-level access control is enabled. If disabled, database access is defined by the network security rules. If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional, if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console. When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.

	This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform. 
* `is_auto_scaling_enabled` - Indicates if auto scaling is enabled for the Autonomous Database CPU core count. 
* `is_data_guard_enabled` - Indicates whether the Autonomous Database has local (in-region) Data Guard enabled. Not applicable to cross-region Autonomous Data Guard associations, or to  Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure. 
* `is_dedicated` - True if the database uses [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html). 
* `is_free_tier` - Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled. 
* `is_mtls_connection_required` - Indicates whether the Autonomous Database requires mTLS connections.
* `is_preview` - Indicates if the Autonomous Database version is a preview version.
* `is_reconnect_clone_enabled` - Indicates if the refreshable clone can be reconnected to its source database.
* `is_refreshable_clone` - Indicates whether the Autonomous Database is a refreshable clone.
* `key_history_entry` - Key History Entry.
	* `id` - The id of the Autonomous Database [Vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts) service key management history entry.
	* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
	* `time_activated` - The date and time the kms key activated.
	* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `key_store_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store.
* `key_store_wallet_name` - The wallet name for Oracle Key Vault.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_lifecycle_details` - KMS key lifecycle details.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Database service. Note that when provisioning an Autonomous Database on [dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the Autonomous Exadata Infrastructure level. When using [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`. 
* `lifecycle_details` - Information about the current lifecycle state.
* `nsg_ids` - A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `ocpu_count` - The number of OCPU cores to be made available to the database. 

	The following points apply:
	* For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	* To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.

	For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See [Characteristics of Infrastructure Shapes](https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.

	**Note:** This parameter cannot be used with the `cpuCoreCount` parameter. 
* `open_mode` - The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
* `operations_insights_status` - Status of Operations Insights for this Autonomous Database.
* `permission_level` - The Autonomous Database permission level. Restricted mode allows access only to admin users.
* `private_endpoint` - The private endpoint for the resource.
* `private_endpoint_ip` - The private endpoint Ip address for the resource.
* `private_endpoint_label` - The private endpoint label for the resource.
* `refreshable_mode` - The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
* `refreshable_status` - The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.
* `role` - The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled. 
* `scheduled_operations` - list of scheduled operations
	* `day_of_week` - Day of the week.
		* `name` - Name of the day of the week.
	* `scheduled_start_time` - auto start time. value must be of ISO-8601 format "HH:mm"
	* `scheduled_stop_time` - auto stop time. value must be of ISO-8601 format "HH:mm"
* `service_console_url` - The URL of the Service Console for the Autonomous Database.
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source Autonomous Database that was cloned to create the current Autonomous Database.
* `standby_db` - Autonomous Data Guard standby database details. 
	* `lag_time_in_seconds` - The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	* `lifecycle_details` - Additional information about the current lifecycle state.
	* `state` - The current state of the Autonomous Database.
* `standby_whitelisted_ips` - The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. 
* `state` - The current state of the Autonomous Database.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	* For Autonomous Database, setting this will disable public secure access to the database.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 
* `supported_regions_to_clone_to` - The list of regions that support the creation of an Autonomous Database clone. 
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
* `time_until_reconnect_clone_enabled` - The time and date as an RFC3339 formatted string, e.g., 2022-01-01T12:00:00.000Z, to set the limit for a refreshable clone to be reconnected to its source database.
* `used_data_storage_size_in_tbs` - The amount of storage that has been used, in terabytes.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `whitelisted_ips` - The client IP access control list (ACL). This feature is available for autonomous databases on [shared Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.

	For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs. Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]` For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations. Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`

	For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 12 hours), when creating the Autonomous Database
	* `update` - (Defaults to 12 hours), when updating the Autonomous Database
	* `delete` - (Defaults to 12 hours), when destroying the Autonomous Database


## Import

AutonomousDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_database.test_autonomous_database "id"
```

