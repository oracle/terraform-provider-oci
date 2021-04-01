---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migrations"
sidebar_current: "docs-oci-datasource-database_migration-migrations"
description: |-
  Provides the list of Migrations in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migrations
This data source provides the list of Migrations in Oracle Cloud Infrastructure Database Migration service.

List all Migrations.


## Example Usage

```hcl
data "oci_database_migration_migrations" "test_migrations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.migration_display_name
	lifecycle_details = var.migration_lifecycle_details
	state = var.migration_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `lifecycle_details` - (Optional) The lifecycle detailed status of the Migration. 
* `state` - (Optional) The current state of the Database Migration Deployment. 


## Attributes Reference

The following attributes are exported:

* `migration_collection` - The list of migration_collection.

### Migration Reference

The following attributes are exported:

* `agent_id` - The OCID of the registered On-Prem ODMS Agent. Required for Offline Migrations. 
* `compartment_id` - OCID of the compartment 
* `credentials_secret_id` - OCID of the Secret in the Oracle Cloud Infrastructure vault containing the Migration credentials. Used to store Golden Gate admin user credentials. 
* `data_transfer_medium_details` - Data Transfer Medium details for the Migration. 
	* `database_link_details` - Optional details for creating a network database link from Oracle Cloud Infrastructure database to on-premise database. 
		* `name` - Name of database link from Oracle Cloud Infrastructure database to on-premise database. ODMS will create link, if the link does not already exist. 
	* `object_storage_details` - In lieu of a network database link, Oracle Cloud Infrastructure Object Storage bucket will be used to store Datapump dump files for the migration. 
		* `bucket` - Bucket name. 
		* `namespace` - Namespace name of the object store bucket. 
* `datapump_settings` - Optional settings for Datapump Export and Import jobs 
	* `data_pump_parameters` - Optional parameters for Datapump Export and Import. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-62324358-2F26-4A94-B69F-1075D53FA96D__BABDECJE 
		* `estimate` - Estimate size of dumps that will be generated. 
		* `exclude_parameters` - Exclude paratemers for export and import. 
		* `export_parallelism_degree` - Maximum number of worker processes that can be used for a Datapump Export job. 
		* `import_parallelism_degree` - Maximum number of worker processes that can be used for a Datapump Import job. For an Autonomous Database, ODMS will automatically query its CPU core count and set this property. 
		* `is_cluster` - False to force datapump worker process to run on one instance. 
		* `table_exists_action` - IMPORT: Specifies the action to be performed when data is loaded into a preexisting table. 
	* `export_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings. 
		* `name` - Name of directory object in database 
		* `path` - Absolute path of directory on database server 
	* `import_directory_object` - Directory object details, used to define either import or export directory objects in Data Pump Settings. 
		* `name` - Name of directory object in database 
		* `path` - Absolute path of directory on database server 
	* `job_mode` - DataPump job mode. Refer to docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-92C2CB46-8BC9-414D-B62E-79CD788C1E62__BABBDEHD 
	* `metadata_remaps` - Defines remapping to be applied to objects as they are processed. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D. 
		* `new_value` - Specifies the new value that oldValue should be translated into. 
		* `old_value` - Specifies the value which needs to be reset. 
		* `type` - Type of remap. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D__BABDJGDI 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Migration Display Name 
* `exclude_objects` - Database objects to exclude from migration. 
	* `object` - Name of the object (regular expression is allowed) 
	* `owner` - Owner of the object (regular expression is allowed) 
* `executing_job_id` - OCID of the current ODMS Job in execution for the Migration, if any. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `golden_gate_details` - Details about Oracle GoldenGate Microservices. 
	* `hub` - Details about Oracle GoldenGate Microservices. 
		* `compute_id` - OCID of Golden Gate compute instance. 
		* `rest_admin_credentials` - Database Admin Credentials details. 
			* `username` - Admin username 
		* `source_container_db_admin_credentials` - Database Admin Credentials details. 
			* `username` - Admin username 
		* `source_db_admin_credentials` - Database Admin Credentials details. 
			* `username` - Admin username 
		* `source_microservices_deployment_name` - Name of Microservices deployment to operate on source DB 
		* `target_db_admin_credentials` - Database Admin Credentials details. 
			* `username` - Admin username 
		* `target_microservices_deployment_name` - Name of Microservices deployment to operate on target DB 
		* `url` - Oracle GoldenGate hub's REST endpoint. Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA 
	* `settings` - Optional settings for Oracle GoldenGate processes 
		* `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds. 
		* `extract` - Parameters for Extract processes. 
			* `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions. 
			* `performance_profile` - Extract performance. 
		* `replicat` - Parameters for Replicat processes. 
			* `map_parallelism` - Number of threads used to read trail files (valid for Parallel Replicat) 
			* `max_apply_parallelism` - Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
			* `min_apply_parallelism` - Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat) 
* `id` - The OCID of the resource 
* `lifecycle_details` - Additional status related to the execution and current state of the Migration. 
* `source_container_database_connection_id` - The OCID of the Source Container Database Connection. 
* `source_database_connection_id` - The OCID of the Source Database Connection. 
* `state` - The current state of the Migration Resource. 
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

