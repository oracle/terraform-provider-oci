---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration"
sidebar_current: "docs-oci-resource-database_migration-migration"
description: |-
Provides the Migration resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_migration
This resource provides the Migration resource in Oracle Cloud Infrastructure Database Migration service.

Create a Migration resource that contains all the details to perform the
database migration operation, such as source and destination database
details, credentials, etc.

Note: If you wish to use the DMS deprecated API version /20210929 it is necessary to pin the Terraform Provider version to v5.47.0. Newer Terraform provider versions will not support the DMS deprecated API version /20210929


## Example Usage

```hcl
resource "oci_database_migration_migration" "test_migration" {
	#Required
	compartment_id = var.compartment_id
	database_combination = var.migration_database_combination
	source_database_connection_id = oci_database_migration_connection.test_connection.id
	target_database_connection_id = oci_database_migration_connection.test_connection.id
	type = var.migration_type

	#Optional
	advanced_parameters {

		#Optional
		data_type = var.migration_advanced_parameters_data_type
		name = var.migration_advanced_parameters_name
		value = var.migration_advanced_parameters_value
	}
	advisor_settings {

		#Optional
		is_ignore_errors = var.migration_advisor_settings_is_ignore_errors
		is_skip_advisor = var.migration_advisor_settings_is_skip_advisor
	}
	bulk_include_exclude_data = var.migration_bulk_include_exclude_data
	data_transfer_medium_details {
		#Required
		type = var.migration_data_transfer_medium_details_type

		#Optional
		access_key_id = oci_kms_key.test_key.id
		name = var.migration_data_transfer_medium_details_name
		object_storage_bucket {

			#Optional
			bucket = var.migration_data_transfer_medium_details_object_storage_bucket_bucket
			namespace = var.migration_data_transfer_medium_details_object_storage_bucket_namespace
		}
		region = var.migration_data_transfer_medium_details_region
		secret_access_key = var.migration_data_transfer_medium_details_secret_access_key
		shared_storage_mount_target_id = oci_file_storage_mount_target.test_mount_target.id
		source {
			#Required
			kind = var.migration_data_transfer_medium_details_source_kind

			#Optional
			oci_home = var.migration_data_transfer_medium_details_source_oci_home
			wallet_location = var.migration_data_transfer_medium_details_source_wallet_location
		}
		target {
			#Required
			kind = var.migration_data_transfer_medium_details_target_kind

			#Optional
			oci_home = var.migration_data_transfer_medium_details_target_oci_home
			wallet_location = var.migration_data_transfer_medium_details_target_wallet_location
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.migration_description
	display_name = var.migration_display_name
	exclude_objects {
		#Required
		object = var.migration_exclude_objects_object

		#Optional
		is_omit_excluded_table_from_replication = var.migration_exclude_objects_is_omit_excluded_table_from_replication
		owner = var.migration_exclude_objects_owner
		schema = var.migration_exclude_objects_schema
		type = var.migration_exclude_objects_type
	}
	freeform_tags = var.migration_freeform_tags
	ggs_details {

		#Optional
		acceptable_lag = var.migration_ggs_details_acceptable_lag
		extract {

			#Optional
			long_trans_duration = var.migration_ggs_details_extract_long_trans_duration
			performance_profile = var.migration_ggs_details_extract_performance_profile
		}
		replicat {

			#Optional
			performance_profile = var.migration_ggs_details_replicat_performance_profile
		}
	}
	hub_details {
		#Required
		key_id = oci_kms_key.test_key.id
		rest_admin_credentials {
			#Required
			password = var.migration_hub_details_rest_admin_credentials_password
			username = var.migration_hub_details_rest_admin_credentials_username
		}
		url = var.migration_hub_details_url
		vault_id = oci_kms_vault.test_vault.id

		#Optional
		acceptable_lag = var.migration_hub_details_acceptable_lag
		compute_id = oci_database_migration_compute.test_compute.id
		extract {

			#Optional
			long_trans_duration = var.migration_hub_details_extract_long_trans_duration
			performance_profile = var.migration_hub_details_extract_performance_profile
		}
		replicat {

			#Optional
			performance_profile = var.migration_hub_details_replicat_performance_profile
		}
	}
	include_objects {
		#Required
		object = var.migration_include_objects_object

		#Optional
		is_omit_excluded_table_from_replication = var.migration_include_objects_is_omit_excluded_table_from_replication
		owner = var.migration_include_objects_owner
		schema = var.migration_include_objects_schema
		type = var.migration_include_objects_type
	}
	initial_load_settings {
		#Required
		job_mode = var.migration_initial_load_settings_job_mode

		#Optional
		compatibility = var.migration_initial_load_settings_compatibility
		data_pump_parameters {

			#Optional
			estimate = var.migration_initial_load_settings_data_pump_parameters_estimate
			exclude_parameters = var.migration_initial_load_settings_data_pump_parameters_exclude_parameters
			export_parallelism_degree = var.migration_initial_load_settings_data_pump_parameters_export_parallelism_degree
			import_parallelism_degree = var.migration_initial_load_settings_data_pump_parameters_import_parallelism_degree
			is_cluster = var.migration_initial_load_settings_data_pump_parameters_is_cluster
			table_exists_action = var.migration_initial_load_settings_data_pump_parameters_table_exists_action
		}
		export_directory_object {

			#Optional
			name = var.migration_initial_load_settings_export_directory_object_name
			path = var.migration_initial_load_settings_export_directory_object_path
		}
		handle_grant_errors = var.migration_initial_load_settings_handle_grant_errors
		import_directory_object {

			#Optional
			name = var.migration_initial_load_settings_import_directory_object_name
			path = var.migration_initial_load_settings_import_directory_object_path
		}
		is_consistent = var.migration_initial_load_settings_is_consistent
		is_ignore_existing_objects = var.migration_initial_load_settings_is_ignore_existing_objects
		is_tz_utc = var.migration_initial_load_settings_is_tz_utc
		metadata_remaps {

			#Optional
			new_value = var.migration_initial_load_settings_metadata_remaps_new_value
			old_value = var.migration_initial_load_settings_metadata_remaps_old_value
			type = var.migration_initial_load_settings_metadata_remaps_type
		}
		primary_key_compatibility = var.migration_initial_load_settings_primary_key_compatibility
		tablespace_details {
			#Required
			target_type = var.migration_initial_load_settings_tablespace_details_target_type

			#Optional
			block_size_in_kbs = var.migration_initial_load_settings_tablespace_details_block_size_in_kbs
			extend_size_in_mbs = var.migration_initial_load_settings_tablespace_details_extend_size_in_mbs
			is_auto_create = var.migration_initial_load_settings_tablespace_details_is_auto_create
			is_big_file = var.migration_initial_load_settings_tablespace_details_is_big_file
			remap_target = var.migration_initial_load_settings_tablespace_details_remap_target
		}
	}
	source_container_database_connection_id = oci_database_migration_connection.test_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `advanced_parameters` - (Applicable when database_combination=ORACLE) (Updatable) List of Migration Parameter objects.
	* `data_type` - (Required when database_combination=ORACLE) (Updatable) Parameter data type.
	* `name` - (Required when database_combination=ORACLE) (Updatable) Parameter name.
	* `value` - (Required when database_combination=ORACLE) (Updatable) If a STRING data type then the value should be an array of characters,  if a INTEGER data type then the value should be an integer value,  if a FLOAT data type then the value should be an float value, if a BOOLEAN data type then the value should be TRUE or FALSE. 
* `advisor_settings` - (Optional) (Updatable) Optional Pre-Migration advisor settings.
  * `is_ignore_errors` - (Optional) (Updatable) True to not interrupt migration execution due to Pre-Migration Advisor errors. Default is false.
  * `is_skip_advisor` - (Optional) (Updatable) True to skip the Pre-Migration Advisor execution. Default is false.
* `bulk_include_exclude_data` - (Optional) Specifies the database objects to be excluded from the migration in bulk. The definition accepts input in a CSV format, newline separated for each entry. More details can be found in the documentation.
* `compartment_id` - (Required) (Updatable) The OCID of the resource being referenced.
* `data_transfer_medium_details` - (Optional) (Updatable) Optional additional properties for data transfer.
  * `access_key_id` - (Applicable when type=AWS_S3) (Updatable) AWS access key credentials identifier Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
  * `name` - (Applicable when type=AWS_S3 | DBLINK) (Updatable) Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link,  if the link does not already exist.
  * `object_storage_bucket` - (Optional) (Updatable) In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium.
    * `bucket` - (Required when type=AWS_S3 | DBLINK | NFS | OBJECT_STORAGE) (Updatable) Bucket name.
    * `namespace` - (Required when type=AWS_S3 | DBLINK | NFS | OBJECT_STORAGE) (Updatable) Namespace name of the object store bucket.
  * `region` - (Applicable when type=AWS_S3) (Updatable) AWS region code where the S3 bucket is located. Region code should match the documented available regions: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions
  * `secret_access_key` - (Applicable when type=AWS_S3) (Updatable) AWS secret access key credentials Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
  * `shared_storage_mount_target_id` - (Applicable when type=NFS) (Updatable) OCID of the shared storage mount target
  * `source` - (Applicable when type=NFS | OBJECT_STORAGE) (Updatable) Optional additional properties for dump transfer in source or target host. Default kind is CURL.
    * `kind` - (Required) (Updatable) Type of dump transfer to use during migration in source or target host. Default kind is CURL
    * `oci_home` - (Applicable when kind=OCI_CLI) (Updatable) Path to the Oracle Cloud Infrastructure CLI installation in the node.
    * `wallet_location` - (Optional) (Updatable) Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node.
  * `target` - (Applicable when type=NFS | OBJECT_STORAGE) (Updatable) Optional additional properties for dump transfer in source or target host. Default kind is CURL.
    * `kind` - (Required) (Updatable) Type of dump transfer to use during migration in source or target host. Default kind is CURL
    * `oci_home` - (Applicable when kind=OCI_CLI) (Updatable) Path to the Oracle Cloud Infrastructure CLI installation in the node.
    * `wallet_location` - (Optional) (Updatable) Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node.
  * `type` - (Required) (Updatable) Type of the data transfer medium to use.
* `database_combination` - (Required) (Updatable) The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `exclude_objects` - (Optional) Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
  * `is_omit_excluded_table_from_replication` - (Applicable when database_combination=ORACLE) Whether an excluded table should be omitted from replication. Only valid for database objects  that have are of type TABLE and object status EXCLUDE.
  * `object` - (Required) Name of the object (regular expression is allowed)
  * `owner` - (Required when database_combination=ORACLE) Owner of the object (regular expression is allowed)
  * `schema` - (Required when database_combination=MYSQL) Schema of the object (regular expression is allowed)
  * `type` - (Optional) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `ggs_details` - (Optional) (Updatable) Optional settings for Oracle GoldenGate processes
  * `acceptable_lag` - (Optional) (Updatable) ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
  * `extract` - (Applicable when database_combination=ORACLE) (Updatable) Parameters for GoldenGate Extract processes.
    * `long_trans_duration` - (Applicable when database_combination=ORACLE) (Updatable) Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions.
    * `performance_profile` - (Applicable when database_combination=ORACLE) (Updatable) Extract performance.
  * `replicat` - (Optional) (Updatable) Parameters for GoldenGate Replicat processes.
    * `performance_profile` - (Optional) (Updatable) Replicat performance.
* `hub_details` - (Optional) (Updatable) Details about Oracle GoldenGate Microservices.
  * `acceptable_lag` - (Optional) (Updatable) ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
  * `compute_id` - (Optional) (Updatable) The OCID of the resource being referenced.
  * `extract` - (Optional) (Updatable) Parameters for GoldenGate Extract processes.
    * `long_trans_duration` - (Optional) (Updatable) Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions.
    * `performance_profile` - (Optional) (Updatable) Extract performance.
  * `key_id` - (Required) (Updatable) The OCID of the resource being referenced.
  * `replicat` - (Optional) (Updatable) Parameters for GoldenGate Replicat processes.
    * `performance_profile` - (Optional) (Updatable) Replicat performance.
  * `rest_admin_credentials` - (Required) (Updatable) Database Administrator Credentials details.
    * `password` - (Required) (Updatable) Administrator password
    * `username` - (Required) (Updatable) Administrator username
  * `url` - (Required) (Updatable) Endpoint URL.
  * `vault_id` - (Required) (Updatable) The OCID of the resource being referenced.
* `include_objects` - (Optional) Database objects to include from migration, cannot be specified alongside 'excludeObjects'
  * `is_omit_excluded_table_from_replication` - (Applicable when database_combination=ORACLE) Whether an excluded table should be omitted from replication. Only valid for database objects  that have are of type TABLE and object status EXCLUDE.
  * `object` - (Required) Name of the object (regular expression is allowed)
  * `owner` - (Required when database_combination=ORACLE) Owner of the object (regular expression is allowed)
  * `schema` - (Required when database_combination=MYSQL) Schema of the object (regular expression is allowed)
  * `type` - (Optional) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded.
* `initial_load_settings` - (Optional) (Updatable) Optional settings for Data Pump Export and Import jobs
  * `compatibility` - (Applicable when database_combination=MYSQL) (Updatable) Apply the specified requirements for compatibility with MySQL Database Service for all tables in the dump  output, altering the dump files as necessary.
  * `data_pump_parameters` - (Applicable when database_combination=ORACLE) (Updatable) Optional parameters for Data Pump Export and Import.
    * `estimate` - (Applicable when database_combination=ORACLE) (Updatable) Estimate size of dumps that will be generated.
    * `exclude_parameters` - (Applicable when database_combination=ORACLE) (Updatable) Exclude paratemers for Export and Import.
    * `export_parallelism_degree` - (Applicable when database_combination=ORACLE) (Updatable) Maximum number of worker processes that can be used for a Data Pump Export job.
    * `import_parallelism_degree` - (Applicable when database_combination=ORACLE) (Updatable) Maximum number of worker processes that can be used for a Data Pump Import job. For an Autonomous Database, ODMS will automatically query its CPU core count and set this property.
    * `is_cluster` - (Applicable when database_combination=ORACLE) (Updatable) Set to false to force Data Pump worker process to run on one instance.
    * `table_exists_action` - (Applicable when database_combination=ORACLE) (Updatable) IMPORT: Specifies the action to be performed when data is loaded into a preexisting table.
  * `export_directory_object` - (Applicable when database_combination=ORACLE) (Updatable) Directory object details, used to define either import or export directory objects in Data Pump Settings. Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error. Export directory will error if there are database link details specified.
    * `name` - (Required when database_combination=ORACLE) (Updatable) Name of directory object in database
    * `path` - (Applicable when database_combination=ORACLE) (Updatable) Absolute path of directory on database server
  * `handle_grant_errors` - (Applicable when database_combination=MYSQL) (Updatable) The action taken in the event of errors related to GRANT or REVOKE errors.
  * `import_directory_object` - (Applicable when database_combination=ORACLE) (Updatable) Directory object details, used to define either import or export directory objects in Data Pump Settings. Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error. Export directory will error if there are database link details specified.
    * `name` - (Required when database_combination=ORACLE) (Updatable) Name of directory object in database
    * `path` - (Applicable when database_combination=ORACLE) (Updatable) Absolute path of directory on database server
  * `is_consistent` - (Applicable when database_combination=MYSQL) (Updatable) Enable (true) or disable (false) consistent data dumps by locking the instance for backup during the dump.
  * `is_ignore_existing_objects` - (Applicable when database_combination=MYSQL) (Updatable) Import the dump even if it contains objects that already exist in the target schema in the MySQL instance.
  * `is_tz_utc` - (Applicable when database_combination=MYSQL) (Updatable) Include a statement at the start of the dump to set the time zone to UTC.
  * `job_mode` - (Required) (Updatable) Oracle Job Mode
  * `metadata_remaps` - (Applicable when database_combination=ORACLE) (Updatable) Defines remapping to be applied to objects as they are processed.
    * `new_value` - (Required when database_combination=ORACLE) (Updatable) Specifies the new value that oldValue should be translated into.
    * `old_value` - (Required when database_combination=ORACLE) (Updatable) Specifies the value which needs to be reset.
    * `type` - (Required when database_combination=ORACLE) (Updatable) Type of remap. Refer to [METADATA_REMAP Procedure ](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D)
  * `primary_key_compatibility` - (Applicable when database_combination=MYSQL) (Updatable) Primary key compatibility option
  * `tablespace_details` - (Applicable when database_combination=ORACLE) (Updatable) Migration tablespace settings.
    * `block_size_in_kbs` - (Applicable when target_type=ADB_D_AUTOCREATE | NON_ADB_AUTOCREATE) (Updatable) Size of Oracle database blocks in KB.
    * `extend_size_in_mbs` - (Applicable when target_type=ADB_D_AUTOCREATE | NON_ADB_AUTOCREATE) (Updatable) Size to extend the tablespace in MB.  Note: Only applicable if 'isBigFile' property is set to true.
    * `is_auto_create` - (Applicable when target_type=ADB_D_AUTOCREATE | NON_ADB_AUTOCREATE) (Updatable) Set this property to true to auto-create tablespaces in the target Database. Note: This is not applicable for Autonomous Database Serverless databases.
    * `is_big_file` - (Applicable when target_type=ADB_D_AUTOCREATE | NON_ADB_AUTOCREATE) (Updatable) Set this property to true to enable tablespace of the type big file.
    * `remap_target` - (Applicable when target_type=ADB_D_REMAP | ADB_S_REMAP | NON_ADB_REMAP) (Updatable) Name of the tablespace on the target database to which the source database tablespace is to be remapped.
    * `target_type` - (Required) (Updatable) Type of Database Base Migration Target.
* `source_container_database_connection_id` - (Applicable when database_combination=ORACLE) (Updatable) The OCID of the resource being referenced.
* `source_database_connection_id` - (Required) (Updatable) The OCID of the resource being referenced.
* `target_database_connection_id` - (Required) (Updatable) The OCID of the resource being referenced.
* `type` - (Required) (Updatable) The type of the migration to be performed. Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advanced_parameters` - List of Migration Parameter objects.
	* `data_type` - Parameter data type.
	* `name` - Parameter name.
	* `value` - If a STRING data type then the value should be an array of characters,  if a INTEGER data type then the value should be an integer value,  if a FLOAT data type then the value should be an float value, if a BOOLEAN data type then the value should be TRUE or FALSE. 
* `advisor_settings` - Details about Oracle Advisor Settings.
  * `is_ignore_errors` - True to not interrupt migration execution due to Pre-Migration Advisor errors. Default is false.
  * `is_skip_advisor` - True to skip the Pre-Migration Advisor execution. Default is false.
* `compartment_id` - The OCID of the resource being referenced.
* `data_transfer_medium_details` - Optional additional properties for data transfer.
  * `access_key_id` - AWS access key credentials identifier Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
  * `name` - Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link,  if the link does not already exist.
  * `object_storage_bucket` - In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium.
    * `bucket` - Bucket name.
    * `namespace` - Namespace name of the object store bucket.
  * `region` - AWS region code where the S3 bucket is located. Region code should match the documented available regions: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions
  * `secret_access_key` - AWS secret access key credentials Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
  * `shared_storage_mount_target_id` - OCID of the shared storage mount target
  * `source` - Optional additional properties for dump transfer in source or target host. Default kind is CURL.
    * `kind` - Type of dump transfer to use during migration in source or target host. Default kind is CURL
    * `oci_home` - Path to the Oracle Cloud Infrastructure CLI installation in the node.
    * `wallet_location` - Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node.
  * `target` - Optional additional properties for dump transfer in source or target host. Default kind is CURL.
    * `kind` - Type of dump transfer to use during migration in source or target host. Default kind is CURL
    * `oci_home` - Path to the Oracle Cloud Infrastructure CLI installation in the node.
    * `wallet_location` - Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node.
  * `type` - Type of the data transfer medium to use.
* `database_combination` - The combination of source and target databases participating in a migration. Example: ORACLE means the migration is meant for migrating Oracle source and target databases.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `executing_job_id` - The OCID of the resource being referenced.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `ggs_details` - Optional settings for Oracle GoldenGate processes
  * `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
  * `extract` - Parameters for Extract processes.
    * `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions.
    * `performance_profile` - Extract performance.
  * `ggs_deployment` - Details about Oracle GoldenGate GGS Deployment.
    * `deployment_id` - The OCID of the resource being referenced.
    * `ggs_admin_credentials_secret_id` - The OCID of the resource being referenced.
  * `replicat` - Parameters for Replicat processes.
    * `performance_profile` - Replicat performance.
* `hub_details` - Details about Oracle GoldenGate Microservices.
  * `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
  * `compute_id` - The OCID of the resource being referenced.
  * `extract` - Parameters for Extract processes.
    * `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions.
    * `performance_profile` - Extract performance.
  * `key_id` - The OCID of the resource being referenced.
  * `replicat` - Parameters for Replicat processes.
    * `performance_profile` - Replicat performance.
  * `rest_admin_credentials` - Database Administrator Credentials details.
    * `username` - Administrator username
  * `url` - Endpoint URL.
  * `vault_id` - The OCID of the resource being referenced.
* `id` - The OCID of the resource being referenced.
* `initial_load_settings` - Optional settings for Data Pump Export and Import jobs
  * `compatibility` - Apply the specified requirements for compatibility with MySQL Database Service for all tables in the dump  output, altering the dump files as necessary.
  * `data_pump_parameters` - Optional parameters for Data Pump Export and Import.
    * `estimate` - Estimate size of dumps that will be generated.
    * `exclude_parameters` - Exclude paratemers for Export and Import.
    * `export_parallelism_degree` - Maximum number of worker processes that can be used for a Data Pump Export job.
    * `import_parallelism_degree` - Maximum number of worker processes that can be used for a Data Pump Import job. For an Autonomous Database, ODMS will automatically query its CPU core count and set this property.
    * `is_cluster` - Set to false to force Data Pump worker process to run on one instance.
    * `table_exists_action` - IMPORT: Specifies the action to be performed when data is loaded into a preexisting table.
  * `export_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings.
    * `name` - Name of directory object in database
    * `path` - Absolute path of directory on database server
  * `handle_grant_errors` - The action taken in the event of errors related to GRANT or REVOKE errors.
  * `import_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings.
    * `name` - Name of directory object in database
    * `path` - Absolute path of directory on database server
  * `is_consistent` - Enable (true) or disable (false) consistent data dumps by locking the instance for backup during the dump.
  * `is_ignore_existing_objects` - Import the dump even if it contains objects that already exist in the target schema in the MySQL instance.
  * `is_tz_utc` - Include a statement at the start of the dump to set the time zone to UTC.
  * `job_mode` - Oracle Job Mode
  * `metadata_remaps` - Defines remapping to be applied to objects as they are processed.
    * `new_value` - Specifies the new value that oldValue should be translated into.
    * `old_value` - Specifies the value which needs to be reset.
    * `type` - Type of remap. Refer to [METADATA_REMAP Procedure ](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D)
  * `primary_key_compatibility` - Primary key compatibility option
  * `tablespace_details` - Migration tablespace settings.
    * `block_size_in_kbs` - Size of Oracle database blocks in KB.
    * `extend_size_in_mbs` - Size to extend the tablespace in MB.  Note: Only applicable if 'isBigFile' property is set to true.
    * `is_auto_create` - Set this property to true to auto-create tablespaces in the target Database. Note: This is not applicable for Autonomous Database Serverless databases.
    * `is_big_file` - Set this property to true to enable tablespace of the type big file.
    * `remap_target` - Name of the tablespace on the target database to which the source database tablespace is to be remapped.
    * `target_type` - Type of Database Base Migration Target.
* `lifecycle_details` - Additional status related to the execution and current state of the Migration.
* `source_container_database_connection_id` - The OCID of the resource being referenced.
* `source_database_connection_id` - The OCID of the resource being referenced.
* `state` - The current state of the Migration resource.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `target_database_connection_id` - The OCID of the resource being referenced.
* `time_created` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
* `time_last_migration` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
* `time_updated` - An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
* `type` - The type of the migration to be performed. Example: ONLINE if no downtime is preferred for a migration. This method uses Oracle GoldenGate for replication.
* `wait_after` - You can optionally pause a migration after a job phase. This property allows you to optionally specify the phase after which you can pause the migration.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Migration
* `update` - (Defaults to 20 minutes), when updating the Migration
* `delete` - (Defaults to 20 minutes), when destroying the Migration


## Import

Migrations can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_migration.test_migration "id"
```
