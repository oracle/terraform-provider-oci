---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_migration"
sidebar_current: "docs-oci-datasource-database_migration-migration"
description: |-
Provides details about a specific Migration in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_migration
This data source provides details about a specific Migration resource in Oracle Cloud Infrastructure Database Migration service.

Display Migration details.

Note: If you wish to use the DMS deprecated API version /20210929 it is necessary to pin the Terraform Provider version to v5.47.0. Newer Terraform provider versions will not support the DMS deprecated API version /20210929

## Example Usage

```hcl
data "oci_database_migration_migration" "test_migration" {
	#Required
	migration_id = oci_database_migration_migration.test_migration.id
}
```

## Argument Reference

The following arguments are supported:

* `migration_id` - (Required) The OCID of the migration


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
* `ggs_details` - Details for Oracle GoldenGate Deployment (Internally managed by the service, not required and will be ignored if provided).
  * `acceptable_lag` - ODMS will monitor GoldenGate end-to-end latency until the lag time is lower than the specified value in seconds.
  * `extract` - Parameters for Extract processes.
    * `long_trans_duration` - Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running. If not specified, Extract will not generate a warning on long-running transactions.
    * `performance_profile` - Extract performance.
  * `ggs_deployment` - Details about Oracle GoldenGate GGS Deployment.
    * `deployment_id` - The OCID of the resource being referenced.
    * `ggs_admin_credentials_secret_id` - The OCID of the resource being referenced.
  * `replicat` - Parameters for Replicat processes.
    * `performance_profile` - Replicat performance.
* `hub_details` - Details for Oracle GoldenGate Marketplace  Instance / Deployment (Currently not supported for MySQL migrations).
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
