---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_systems"
sidebar_current: "docs-oci-datasource-psql-db_systems"
description: |-
  Provides the list of Db Systems in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_systems
This data source provides the list of Db Systems in Oracle Cloud Infrastructure Psql service.

Returns a list of database systems.


## Example Usage

```hcl
data "oci_psql_db_systems" "test_db_systems" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.db_system_display_name
	id = var.db_system_id
	state = var.db_system_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A unique identifier for the database system.
* `state` - (Optional) A filter to return only resources if their `lifecycleState` matches the given `lifecycleState`.


## Attributes Reference

The following attributes are exported:

* `db_system_collection` - The list of db_system_collection.

### DbSystem Reference

The following attributes are exported:

* `admin_username` - The database system administrator username.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the database system.
* `config_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration associated with the database system.
* `db_version` - The major and minor versions of the database system software.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the database system.
* `display_name` - A user-friendly display name for the database system. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - A unique identifier for the database system. Immutable on creation.
* `instance_count` - Count of instances, or nodes, in the database system.
* `instance_memory_size_in_gbs` - The total amount of memory available to each database instance node, in gigabytes.
* `instance_ocpu_count` - The total number of OCPUs available to each database instance node.
* `instances` - The list of instances, or nodes, in the database system.
	* `availability_domain` - The availability domain in which the database instance node is located.
	* `description` - Description of the database instance node.
	* `display_name` - A user-friendly display name for the database instance node. Avoid entering confidential information.
	* `id` - A unique identifier for the database instance node. Immutable on creation.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `state` - The current state of the database instance node.
	* `time_created` - The date and time that the database instance node was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
	* `time_updated` - The date and time that the database instance node was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_policy` - PostgreSQL database system management policy.
	* `backup_policy` - PostgreSQL database system backup policy.
		* `backup_start` - Hour of the day when the backup starts.
		* `days_of_the_month` - Day of the month when the backup should start. To ensure that the backup runs monthly, the latest day of the month that you can use to schedule a backup is the the 28th day. 
		* `days_of_the_week` - The day of the week that the backup starts.
		* `kind` - The kind of backup policy.
		* `retention_days` - How many days the data should be stored after the database system deletion.
	* `maintenance_window_start` - The start of the maintenance window. 
* `network_details` - Network details for the database system.
	* `nsg_ids` - List of customer Network Security Group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the database system.
	* `primary_db_endpoint_private_ip` - Private IP in customer subnet. The value is optional. If the IP is not provided, the IP will be chosen from the available IP addresses from the specified subnet. 
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer subnet associated with the database system.
* `shape` - The name of the shape for the database instance. Example: `VM.Standard.E4.Flex` 
* `source` - The source used to restore the database system.
	* `backup_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system backup.
	* `is_having_restore_config_overrides` - Deprecated. Don't use.
	* `source_type` - The source descriminator. 
* `state` - The current state of the database system.
* `storage_details` - Storage details of the database system.
	* `availability_domain` - Specifies the availability domain of AD-local storage. If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified. If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified. 
	* `iops` - Guaranteed input/output storage requests per second (IOPS) available to the database system.
	* `is_regionally_durable` - Specifies if the block volume used for the database system is regional or AD-local. If not specified, it will be set to false. If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified. If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified. 
	* `system_type` - Type of the database system.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `system_type` - Type of the database system.
* `time_created` - The date and time that the database system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time that the database system was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

