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


## Example Usage

```hcl
resource "oci_database_migration_migration" "test_migration" {
	#Required
	compartment_id = var.compartment_id
	source_database_connection_id = oci_database_migration_connection.test_connection.id
	target_database_connection_id = oci_database_migration_connection.test_connection.id
	type = var.migration_type

	#Optional
	advisor_settings {

		#Optional
		is_ignore_errors = var.migration_advisor_settings_is_ignore_errors
		is_skip_advisor = var.migration_advisor_settings_is_skip_advisor
	}
	agent_id = oci_database_migration_agent.test_agent.id
	csv_text = var.migration_csv_text
	data_transfer_medium_details {

		#Optional
		database_link_details {

			#Optional
			name = var.migration_data_transfer_medium_details_database_link_details_name
			wallet_bucket {
				#Required
				bucket = var.migration_data_transfer_medium_details_database_link_details_wallet_bucket_bucket
				namespace = var.migration_data_transfer_medium_details_database_link_details_wallet_bucket_namespace
			}
		}
		object_storage_details {
			#Required
			bucket = var.migration_data_transfer_medium_details_object_storage_details_bucket
			namespace = var.migration_data_transfer_medium_details_object_storage_details_namespace
		}
	}
	datapump_settings {

		#Optional
		data_pump_parameters {

			#Optional
			estimate = var.migration_datapump_settings_data_pump_parameters_estimate
			exclude_parameters = var.migration_datapump_settings_data_pump_parameters_exclude_parameters
			export_parallelism_degree = var.migration_datapump_settings_data_pump_parameters_export_parallelism_degree
			import_parallelism_degree = var.migration_datapump_settings_data_pump_parameters_import_parallelism_degree
			is_cluster = var.migration_datapump_settings_data_pump_parameters_is_cluster
			table_exists_action = var.migration_datapump_settings_data_pump_parameters_table_exists_action
		}
		export_directory_object {
			#Required
			name = var.migration_datapump_settings_export_directory_object_name

			#Optional
			path = var.migration_datapump_settings_export_directory_object_path
		}
		import_directory_object {
			#Required
			name = var.migration_datapump_settings_import_directory_object_name

			#Optional
			path = var.migration_datapump_settings_import_directory_object_path
		}
		job_mode = var.migration_datapump_settings_job_mode
		metadata_remaps {
			#Required
			new_value = var.migration_datapump_settings_metadata_remaps_new_value
			old_value = var.migration_datapump_settings_metadata_remaps_old_value
			type = var.migration_datapump_settings_metadata_remaps_type
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.migration_display_name
	dump_transfer_details {

		#Optional
		source {
			#Required
			kind = var.migration_dump_transfer_details_source_kind

			#Optional
			oci_home = var.migration_dump_transfer_details_source_oci_home
			wallet_location = var.migration_dump_transfer_details_source_wallet_location
		}
		target {
			#Required
			kind = var.migration_dump_transfer_details_target_kind

			#Optional
			oci_home = var.migration_dump_transfer_details_target_oci_home
			wallet_location = var.migration_dump_transfer_details_target_wallet_location
		}
	}
	exclude_objects {
		#Required
		object = var.migration_exclude_objects_object
		owner = var.migration_exclude_objects_owner

		#Optional
		is_omit_excluded_table_from_replication = var.migration_exclude_objects_is_omit_excluded_table_from_replication
		type = var.migration_exclude_objects_type
	}
	freeform_tags = {"bar-key"= "value"}
	golden_gate_details {
		#Required
		hub {
			#Required
			rest_admin_credentials {
				#Required
				password = var.migration_golden_gate_details_hub_rest_admin_credentials_password
				username = var.migration_golden_gate_details_hub_rest_admin_credentials_username
			}
			url = var.migration_golden_gate_details_hub_url

			#Optional
			compute_id = oci_database_migration_compute.test_compute.id
			source_container_db_admin_credentials {
				#Required
				password = var.migration_golden_gate_details_hub_source_container_db_admin_credentials_password
				username = var.migration_golden_gate_details_hub_source_container_db_admin_credentials_username
			}
			source_db_admin_credentials {
				#Required
				password = var.migration_golden_gate_details_hub_source_db_admin_credentials_password
				username = var.migration_golden_gate_details_hub_source_db_admin_credentials_username
			}
			source_microservices_deployment_name = oci_apigateway_deployment.test_deployment.name
			target_db_admin_credentials {
				#Required
				password = var.migration_golden_gate_details_hub_target_db_admin_credentials_password
				username = var.migration_golden_gate_details_hub_target_db_admin_credentials_username
			}
			target_microservices_deployment_name = oci_apigateway_deployment.test_deployment.name
		}

		#Optional
		settings {

			#Optional
			acceptable_lag = var.migration_golden_gate_details_settings_acceptable_lag
			extract {

				#Optional
				long_trans_duration = var.migration_golden_gate_details_settings_extract_long_trans_duration
				performance_profile = var.migration_golden_gate_details_settings_extract_performance_profile
			}
			replicat {

				#Optional
				map_parallelism = var.migration_golden_gate_details_settings_replicat_map_parallelism
				max_apply_parallelism = var.migration_golden_gate_details_settings_replicat_max_apply_parallelism
				min_apply_parallelism = var.migration_golden_gate_details_settings_replicat_min_apply_parallelism
			}
		}
	}
	golden_gate_service_details {

		#Optional
		settings {

			#Optional
			acceptable_lag = var.migration_golden_gate_service_details_settings_acceptable_lag
			extract {

				#Optional
				long_trans_duration = var.migration_golden_gate_service_details_settings_extract_long_trans_duration
				performance_profile = var.migration_golden_gate_service_details_settings_extract_performance_profile
			}
			replicat {

				#Optional
				map_parallelism = var.migration_golden_gate_service_details_settings_replicat_map_parallelism
				max_apply_parallelism = var.migration_golden_gate_service_details_settings_replicat_max_apply_parallelism
				min_apply_parallelism = var.migration_golden_gate_service_details_settings_replicat_min_apply_parallelism
			}
		}
		source_container_db_credentials {
			#Required
			password = var.migration_golden_gate_service_details_source_container_db_credentials_password
			username = var.migration_golden_gate_service_details_source_container_db_credentials_username
		}
		source_db_credentials {
			#Required
			password = var.migration_golden_gate_service_details_source_db_credentials_password
			username = var.migration_golden_gate_service_details_source_db_credentials_username
		}
		target_db_credentials {
			#Required
			password = var.migration_golden_gate_service_details_target_db_credentials_password
			username = var.migration_golden_gate_service_details_target_db_credentials_username
		}
	}
	include_objects {
		#Required
		object = var.migration_include_objects_object
		owner = var.migration_include_objects_owner

		#Optional
		is_omit_excluded_table_from_replication = var.migration_include_objects_is_omit_excluded_table_from_replication
		type = var.migration_include_objects_type
	}
	source_container_database_connection_id = oci_database_migration_connection.test_connection.id
	vault_details {
		#Required
		compartment_id = var.compartment_id
		key_id = oci_kms_key.test_key.id
		vault_id = oci_kms_vault.test_vault.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `advisor_settings` - (Optional) (Updatable) Optional Pre-Migration advisor settings. 
	* `is_ignore_errors` - (Optional) (Updatable) True to not interrupt migration execution due to Pre-Migration Advisor errors. Default is false. 
	* `is_skip_advisor` - (Optional) (Updatable) True to skip the Pre-Migration Advisor execution. Default is false. 
* `agent_id` - (Optional) (Updatable) The OCID of the registered ODMS Agent. Only valid for Offline Logical Migrations. 
* `compartment_id` - (Required) (Updatable) OCID of the compartment 
* `csv_text` - (Optional) Database objects to exclude/include from migration in CSV format. The excludeObjects and includeObjects fields will be ignored if this field is not null. 
* `data_transfer_medium_details` - (Optional) (Updatable) Data Transfer Medium details for the Migration. If not specified, it will default to Database Link. Only one type of data transfer medium can be specified. 
	* `database_link_details` - (Optional) (Updatable) Optional details for creating a network database link from Oracle Cloud Infrastructure database to on-premise database. 
		* `name` - (Optional) (Updatable) Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link, if the link does not already exist. 
		* `wallet_bucket` - (Optional) (Updatable) In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
			* `bucket` - (Required) (Updatable) Bucket name. 
			* `namespace` - (Required) (Updatable) Namespace name of the object store bucket. 
	* `object_storage_details` - (Optional) (Updatable) In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
		* `bucket` - (Required) (Updatable) Bucket name. 
		* `namespace` - (Required) (Updatable) Namespace name of the object store bucket.
* `data_transfer_medium_details_v2` - (Optional) (Updatable) Optional additional properties for dump transfer in source or target host. 
	* `access_key_id` - (Applicable when type=AWS_S3) (Updatable) AWS access key credentials identifier Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys 
	* `name` - (Applicable when type=AWS_S3 | DBLINK) (Updatable) Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link, if the link does not already exist. 
	* `object_storage_bucket` - (Optional) (Updatable) In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
		* `bucket` - (Required when type=AWS_S3 | DBLINK | NFS | OBJECT_STORAGE) (Updatable) Bucket name. 
		* `namespace` - (Required when type=AWS_S3 | DBLINK | NFS | OBJECT_STORAGE) (Updatable) Namespace name of the object store bucket. 
	* `region` - (Applicable when type=AWS_S3) (Updatable) AWS region code where the S3 bucket is located. Region code should match the documented available regions: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions 
	* `secret_access_key` - (Applicable when type=AWS_S3) (Updatable) AWS secret access key credentials Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys 
	* `type` - (Required) (Updatable) Type of the data transfer medium to use for the datapump
* `datapump_settings` - (Optional) (Updatable) Optional settings for Data Pump Export and Import jobs 
	* `data_pump_parameters` - (Optional) (Updatable) Optional parameters for Data Pump Export and Import. Refer to [Configuring Optional Initial Load Advanced Settings](https://docs.us.oracle.com/en/cloud/paas/database-migration/dmsus/working-migration-resources.html#GUID-24BD3054-FDF8-48FF-8492-636C1D4B71ED) 
		* `estimate` - (Optional) (Updatable) Estimate size of dumps that will be generated. 
		* `exclude_parameters` - (Optional) (Updatable) Exclude paratemers for Export and Import. 
		* `export_parallelism_degree` - (Optional) (Updatable) Maximum number of worker processes that can be used for a Data Pump Export job. 
		* `import_parallelism_degree` - (Optional) (Updatable) Maximum number of worker processes that can be used for a Data Pump Import job. For an Autonomous Database, ODMS will automatically query its CPU core count and set this property. 
		* `is_cluster` - (Optional) (Updatable) Set to false to force Data Pump worker process to run on one instance. 
		* `table_exists_action` - (Optional) (Updatable) IMPORT: Specifies the action to be performed when data is loaded into a preexisting table. 
	* `export_directory_object` - (Optional) (Updatable) Directory object details, used to define either import or export directory objects in Data Pump Settings. Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error. Export directory will error if there are database link details specified. 
		* `name` - (Required) (Updatable) Name of directory object in database 
		* `path` - (Optional) (Updatable) Absolute path of directory on database server 
	* `import_directory_object` - (Optional) (Updatable) Directory object details, used to define either import or export directory objects in Data Pump Settings. Import directory is required for Non-Autonomous target connections. If specified for an autonomous target, it will show an error. Export directory will error if there are database link details specified. 
		* `name` - (Required) (Updatable) Name of directory object in database 
		* `path` - (Optional) (Updatable) Absolute path of directory on database server 
	* `job_mode` - (Optional) (Updatable) Data Pump job mode. Refer to [link text](https://docs.oracle.com/en/database/oracle/oracle-database/19/sutil/oracle-data-pump-export-utility.html#GUID-8E497131-6B9B-4CC8-AA50-35F480CAC2C4) 
	* `metadata_remaps` - (Optional) (Updatable) Defines remapping to be applied to objects as they are processed. Refer to [DATA_REMAP](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-E75AAE6F-4EA6-4737-A752-6B62F5E9D460) 
		* `new_value` - (Required) (Updatable) Specifies the new value that oldValue should be translated into. 
		* `old_value` - (Required) (Updatable) Specifies the value which needs to be reset. 
		* `type` - (Required) (Updatable) Type of remap. Refer to [METADATA_REMAP Procedure ](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Migration Display Name 
* `dump_transfer_details` - (Optional) (Updatable) Optional additional properties for dump transfer. 
	* `source` - (Optional) (Updatable) Optional additional properties for dump transfer in source or target host. Default kind is CURL 
		* `kind` - (Required) (Updatable) Type of dump transfer to use during migration in source or target host. Default kind is CURL 
		* `oci_home` - (Required when kind=OCI_CLI) (Updatable) Path to the Oracle Cloud Infrastructure CLI installation in the node. 
		* `wallet_location` - (Optional) (Updatable) Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node. 
	* `target` - (Optional) (Updatable) Optional additional properties for dump transfer in source or target host. Default kind is CURL 
		* `kind` - (Required) (Updatable) Type of dump transfer to use during migration in source or target host. Default kind is CURL 
		* `oci_home` - (Required when kind=OCI_CLI) (Updatable) Path to the Oracle Cloud Infrastructure CLI installation in the node. 
		* `wallet_location` - (Optional) (Updatable) Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node. 
* `exclude_objects` - (Optional) (Updatable) Database objects to exclude from migration, cannot be specified alongside 'includeObjects' 
	* `is_omit_excluded_table_from_replication` - (Optional) (Updatable) Whether an excluded table should be omitted from replication. Only valid for database objects that have are of type TABLE and that are included in the exludeObjects. 
	* `object` - (Required) (Updatable) Name of the object (regular expression is allowed) 
	* `owner` - (Required) (Updatable) Owner of the object (regular expression is allowed) 
	* `type` - (Optional) (Updatable) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `golden_gate_details` - (Optional) (Updatable) Details about Oracle GoldenGate Microservices. Required for online logical migration. 
	* `hub` - (Required) (Updatable) Details about Oracle GoldenGate Microservices. Required for online logical migration. 
		* `compute_id` - (Optional) (Updatable) OCID of GoldenGate Microservices compute instance. 
		* `rest_admin_credentials` - (Required) (Updatable) Database Administrator Credentials details. 
			* `password` - (Required) (Updatable) Administrator password 
			* `username` - (Required) (Updatable) Administrator username 
		* `source_container_db_admin_credentials` - (Optional) (Updatable) Database Administrator Credentials details. 
			* `password` - (Required) (Updatable) Administrator password 
			* `username` - (Required) (Updatable) Administrator username 
		* `source_db_admin_credentials` - (Optional) (Updatable) Database Administrator Credentials details. 
			* `password` - (Required) (Updatable) Administrator password 
			* `username` - (Required) (Updatable) Administrator username 
		* `source_microservices_deployment_name` - (Optional) (Updatable) Name of GoldenGate Microservices deployment to operate on source database 
		* `target_db_admin_credentials` - (Optional) (Updatable) Database Administrator Credentials details. 
			* `password` - (Required) (Updatable) Administrator password 
			* `username` - (Required) (Updatable) Administrator username 
		* `target_microservices_deployment_name` - (Optional) (Updatable) Name of GoldenGate Microservices deployment to operate on target database 
		* `url` - (Required) (Updatable) Oracle GoldenGate Microservices hub's REST endpoint. Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA 
	* `settings` - (Optional) (Updatable) Optional settings for GoldenGate Microservices processes 
		* `acceptable_lag` - (Optional) (Updatable) ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds. 
		* `extract` - (Optional) (Updatable) Parameters for GoldenGate Extract processes. 
			* `long_trans_duration` - (Optional) (Updatable) Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions. 
			* `performance_profile` - (Optional) (Updatable) Extract performance. 
		* `replicat` - (Optional) (Updatable) Parameters for GoldenGate Replicat processes. 
			* `map_parallelism` - (Optional) (Updatable) Number of threads used to read trail files (valid for Parallel Replicat) 
			* `max_apply_parallelism` - (Optional) (Updatable) Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
			* `min_apply_parallelism` - (Optional) (Updatable) Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
* `golden_gate_service_details` - (Optional) (Updatable) Details about Oracle GoldenGate GGS Deployment. 
	* `settings` - (Optional) (Updatable) Optional settings for GoldenGate Microservices processes 
		* `acceptable_lag` - (Optional) (Updatable) ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds. 
		* `extract` - (Optional) (Updatable) Parameters for GoldenGate Extract processes. 
			* `long_trans_duration` - (Optional) (Updatable) Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions. 
			* `performance_profile` - (Optional) (Updatable) Extract performance. 
		* `replicat` - (Optional) (Updatable) Parameters for GoldenGate Replicat processes. 
			* `map_parallelism` - (Optional) (Updatable) Number of threads used to read trail files (valid for Parallel Replicat) 
			* `max_apply_parallelism` - (Optional) (Updatable) Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
			* `min_apply_parallelism` - (Optional) (Updatable) Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
	* `source_container_db_credentials` - (Optional) (Updatable) Database Credentials details. 
		* `password` - (Required) (Updatable) Database  password 
		* `username` - (Required) (Updatable) Database username 
	* `source_db_credentials` - (Optional) (Updatable) Database Credentials details. 
		* `password` - (Required) (Updatable) Database  password 
		* `username` - (Required) (Updatable) Database username 
	* `target_db_credentials` - (Optional) (Updatable) Database Credentials details. 
		* `password` - (Required) (Updatable) Database  password 
		* `username` - (Required) (Updatable) Database username 
* `include_objects` - (Optional) (Updatable) Database objects to include from migration, cannot be specified alongside 'excludeObjects' 
	* `is_omit_excluded_table_from_replication` - (Optional) (Updatable) Whether an excluded table should be omitted from replication. Only valid for database objects that have are of type TABLE and that are included in the exludeObjects. 
	* `object` - (Required) (Updatable) Name of the object (regular expression is allowed) 
	* `owner` - (Required) (Updatable) Owner of the object (regular expression is allowed) 
	* `type` - (Optional) (Updatable) Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `source_container_database_connection_id` - (Optional) (Updatable) The OCID of the Source Container Database Connection. Only used for Online migrations. Only Connections of type Non-Autonomous can be used as source container databases. 
* `source_database_connection_id` - (Required) (Updatable) The OCID of the Source Database Connection. 
* `target_database_connection_id` - (Required) (Updatable) The OCID of the Target Database Connection. 
* `type` - (Required) (Updatable) Migration type. 
* `vault_details` - (Optional) (Updatable) Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets 
	* `compartment_id` - (Required) (Updatable) OCID of the compartment where the secret containing the credentials will be created. 
	* `key_id` - (Required) (Updatable) OCID of the vault encryption key 
	* `vault_id` - (Required) (Updatable) OCID of the vault 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advisor_settings` - Optional Pre-Migration advisor settings. 
	* `is_ignore_errors` - True to not interrupt migration execution due to Pre-Migration Advisor errors. Default is false. 
	* `is_skip_advisor` - True to skip the Pre-Migration Advisor execution. Default is false. 
* `agent_id` - The OCID of the registered on-premises ODMS Agent. Only valid for Offline Migrations. 
* `compartment_id` - OCID of the compartment 
* `credentials_secret_id` - OCID of the Secret in the Oracle Cloud Infrastructure vault containing the Migration credentials. Used to store GoldenGate administrator user credentials. 
* `data_transfer_medium_details` - Data Transfer Medium details for the Migration. 
	* `database_link_details` - Optional details for creating a network database link from Oracle Cloud Infrastructure database to on-premise database. 
		* `name` - Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link, if the link does not already exist. 
		* `wallet_bucket` - In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
			* `bucket` - Bucket name. 
			* `namespace` - Namespace name of the object store bucket. 
	* `object_storage_details` - In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
		* `bucket` - Bucket name. 
		* `namespace` - Namespace name of the object store bucket.
* `data_transfer_medium_details_v2` - Optional additional properties for dump transfer in source or target host. 
	* `access_key_id` - AWS access key credentials identifier Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys 
	* `name` - Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link, if the link does not already exist. 
	* `object_storage_bucket` - In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Data Pump dump files for the migration. Additionally, it can be specified alongside a database link data transfer medium. 
		* `bucket` - Bucket name. 
		* `namespace` - Namespace name of the object store bucket. 
	* `region` - AWS region code where the S3 bucket is located. Region code should match the documented available regions: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions 
	* `secret_access_key` - AWS secret access key credentials Details: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys 
	* `type` - Type of the data transfer medium to use for the datapump
* `datapump_settings` - Optional settings for Data Pump Export and Import jobs 
	* `data_pump_parameters` - Optional parameters for Data Pump Export and Import. Refer to [Configuring Optional Initial Load Advanced Settings](https://docs.us.oracle.com/en/cloud/paas/database-migration/dmsus/working-migration-resources.html#GUID-24BD3054-FDF8-48FF-8492-636C1D4B71ED) 
		* `estimate` - Estimate size of dumps that will be generated. 
		* `exclude_parameters` - Exclude paratemers for Export and Import. 
		* `export_parallelism_degree` - Maximum number of worker processes that can be used for a Data Pump Export job. 
		* `import_parallelism_degree` - Maximum number of worker processes that can be used for a Data Pump Import job. For an Autonomous Database, ODMS will automatically query its CPU core count and set this property. 
		* `is_cluster` - Set to false to force Data Pump worker processes to run on one instance. 
		* `table_exists_action` - IMPORT: Specifies the action to be performed when data is loaded into a preexisting table. 
	* `export_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings. 
		* `name` - Name of directory object in database 
		* `path` - Absolute path of directory on database server 
	* `import_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings. 
		* `name` - Name of directory object in database 
		* `path` - Absolute path of directory on database server 
	* `job_mode` - Data Pump job mode. Refer to [Data Pump Export Modes ](https://docs.oracle.com/en/database/oracle/oracle-database/19/sutil/oracle-data-pump-export-utility.html#GUID-8E497131-6B9B-4CC8-AA50-35F480CAC2C4) 
	* `metadata_remaps` - Defines remapping to be applied to objects as they are processed. Refer to [METADATA_REMAP Procedure ](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D) 
		* `new_value` - Specifies the new value that oldValue should be translated into. 
		* `old_value` - Specifies the value which needs to be reset. 
		* `type` - Type of remap. Refer to [METADATA_REMAP Procedure ](https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D) 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Migration Display Name 
* `dump_transfer_details` - Optional additional properties for dump transfer. 
	* `shared_storage_mount_target_id` - Optional OCID of the shared storage mount target.
	* `source` - Optional additional properties for dump transfer in source or target host. Default kind is CURL 
		* `kind` - Type of dump transfer to use during migration in source or target host. Default kind is CURL 
		* `oci_home` - Path to the Oracle Cloud Infrastructure CLI installation in the node. 
		* `wallet_location` - Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node. 
	* `target` - Optional additional properties for dump transfer in source or target host. Default kind is CURL 
		* `kind` - Type of dump transfer to use during migration in source or target host. Default kind is CURL 
		* `oci_home` - Path to the Oracle Cloud Infrastructure CLI installation in the node. 
		* `wallet_location` - Directory path to Oracle Cloud Infrastructure SSL wallet location on Db server node. 
* `exclude_objects` - Database objects to exclude from migration. If 'includeObjects' are specified, only exclude object types can be specified with general wildcards (.*) for owner and objectName. 
	* `is_omit_excluded_table_from_replication` - Whether an excluded table should be omitted from replication. Only valid for database objects that have are of type TABLE and that are included in the exludeObjects. 
	* `object` - Name of the object (regular expression is allowed) 
	* `owner` - Owner of the object (regular expression is allowed) 
	* `type` - Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `executing_job_id` - OCID of the current ODMS Job in execution for the Migration, if any. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `golden_gate_details` - Details about Oracle GoldenGate Microservices. 
	* `hub` - Details about Oracle GoldenGate Microservices. 
		* `compute_id` - OCID of GoldenGate compute instance. 
		* `rest_admin_credentials` - Database Administrator Credentials details. 
			* `username` - Administrator username 
		* `source_container_db_admin_credentials` - Database Administrator Credentials details. 
			* `username` - Administrator username 
		* `source_db_admin_credentials` - Database Administrator Credentials details. 
			* `username` - Administrator username 
		* `source_microservices_deployment_name` - Name of GoldenGate deployment to operate on source database 
		* `target_db_admin_credentials` - Database Administrator Credentials details. 
			* `username` - Administrator username 
		* `target_microservices_deployment_name` - Name of GoldenGate deployment to operate on target database 
		* `url` - Oracle GoldenGate hub's REST endpoint. Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA 
	* `settings` - Optional settings for Oracle GoldenGate processes 
		* `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds. 
		* `extract` - Parameters for Extract processes. 
			* `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions. 
			* `performance_profile` - Extract performance. 
		* `replicat` - Parameters for Replicat processes. 
			* `map_parallelism` - Number of threads used to read trail files (valid for Parallel Replicat) 
			* `max_apply_parallelism` - Defines the range in which Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
			* `min_apply_parallelism` - Defines the range in which Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
* `golden_gate_service_details` - Details about Oracle GoldenGate GGS Deployment. 
	* `ggs_deployment` - Details about Oracle GoldenGate GGS Deployment. 
		* `deployment_id` - OCID of a GoldenGate Deployment 
		* `ggs_admin_credentials_secret_id` - OCID of a VaultSecret containing the Admin Credentials for the GGS Deployment 
	* `settings` - Optional settings for Oracle GoldenGate processes 
		* `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds. 
		* `extract` - Parameters for Extract processes. 
			* `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions. 
			* `performance_profile` - Extract performance. 
		* `replicat` - Parameters for Replicat processes. 
			* `map_parallelism` - Number of threads used to read trail files (valid for Parallel Replicat) 
			* `max_apply_parallelism` - Defines the range in which Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
			* `min_apply_parallelism` - Defines the range in which Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
            * `performance_profile` - Extract performance.
* `id` - The OCID of the resource 
* `include_objects` - Database objects to include from migration. 
	* `is_omit_excluded_table_from_replication` - Whether an excluded table should be omitted from replication. Only valid for database objects that have are of type TABLE and that are included in the exludeObjects. 
	* `object` - Name of the object (regular expression is allowed) 
	* `owner` - Owner of the object (regular expression is allowed) 
	* `type` - Type of object to exclude. If not specified, matching owners and object names of type TABLE would be excluded. 
* `lifecycle_details` - Additional status related to the execution and current state of the Migration. 
* `source_container_database_connection_id` - The OCID of the Source Container Database Connection. 
* `source_database_connection_id` - The OCID of the Source Database Connection. 
* `state` - The current state of the Migration resource. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_database_connection_id` - The OCID of the Target Database Connection. 
* `time_created` - The time the Migration was created. An RFC3339 formatted datetime string. 
* `time_last_migration` - The time of last Migration. An RFC3339 formatted datetime string. 
* `time_updated` - The time of the last Migration details update. An RFC3339 formatted datetime string. 
* `type` - Migration type. 
* `vault_details` - Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets 
	* `compartment_id` - OCID of the compartment where the secret containing the credentials will be created. 
	* `key_id` - OCID of the vault encryption key 
	* `vault_id` - OCID of the vault 
* `wait_after` - Name of a migration phase. The Job will wait after executing this phase until the Resume Job endpoint is called. 

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

