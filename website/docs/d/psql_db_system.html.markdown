---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system"
sidebar_current: "docs-oci-datasource-psql-db_system"
description: |-
  Provides details about a specific Db System in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_db_system
This data source provides details about a specific Db System resource in Oracle Cloud Infrastructure Psql service.

Gets a DbSystem by identifier

## Example Usage

```hcl
data "oci_psql_db_system" "test_db_system" {
	#Required
	db_system_id = oci_psql_db_system.test_db_system.id

	#Optional
	excluded_fields = var.db_system_excluded_fields
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) unique DbSystem identifier
* `excluded_fields` - (Optional) A filter to exclude DB config  when this query param is set to OverrideDbConfig


## Attributes Reference

The following attributes are exported:

* `admin_username` - The DB system username.
* `compartment_id` - Compartment identifier
* `config_id` - Configuration identifier
* `db_version` - The major and minor versions of the DbSystem software.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the DbSystem.
* `display_name` - DbSystem display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `instance_count` - Count of DbInstances in the DbSystem.
* `instance_memory_size_in_gbs` - The total amount of memory available to each DbInstance, in gigabytes.
* `instance_ocpu_count` - The total number of OCPUs available to each DbInstance.
* `instances` - The list of DbInstances in the DbSystem.
	* `availability_domain` - The availability domain in which the DbInstance is placed.
	* `description` - Description of the DbInstance.
	* `display_name` - Display name of the DbInstance.
	* `id` - Unique identifier that is immutable on creation.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	* `state` - The current state of the DbInstance.
	* `time_created` - The time the the DbInstance was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time the DbInstance was updated. An RFC3339 formatted datetime string.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `management_policy` - PostgreSQL DB system management policy
	* `backup_policy` - Posgresql DB system backup policy
		* `backup_start` - Hour of the day when backup starts.
		* `days_of_the_month` - Days of the month when backup should start. If the day is greater last day of the current month, then it will be triggered on the last day of the current month 
		* `days_of_the_week` - Weekly days
		* `kind` - Backup policy kind
		* `retention_days` - How many days the customers data should be stored after the db system deletion.
	* `maintenance_window_start` - The start of the maintenance window. 
* `network_details` - DbSystem network details.
	* `nsg_ids` - List of customer NetworkSecurityGroup identifiers
	* `primary_db_endpoint_private_ip` - Private IP in customer subnet. The value is optional. If the IP is not provided the IP will be chosen among the available IP addresses from the specified subnet. 
	* `subnet_id` - Customer Subnet identifier
* `shape` - Shape of dbInstance.
* `source` - New source is used to restore the DB system.
	* `backup_id` - DbSystem backup identifier.
	* `is_having_restore_config_overrides` - Restore the DB config overrides from backup. Default is false
	* `source_type` - The source descriminator. 
* `state` - The current state of the DbSystem.
* `storage_details` - Storage details of the DbSystem.
	* `availability_domain` - Specifies the availability domain of AD-local storage. If isRegionallyDurable is set to true, availabilityDomain should not be specified. If isRegionallyDurable is set to false, availabilityDomain must be specified. 
	* `iops` - DbSystem Performance Unit
	* `is_regionally_durable` - Specifies if the block volume used for the DbSystem is regional or AD-local. If not specified, it will be set to false. If isRegionallyDurable is set to true, availabilityDomain should not be specified. If isRegionallyDurable is set to false, availabilityDomain must be specified. 
	* `system_type` - Type of the DbSystem.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `system_type` - Type of the DbSystem.
* `time_created` - The time the the DbSystem was created. An RFC3339 formatted datetime string
* `time_updated` - The time the DbSystem was updated. An RFC3339 formatted datetime string

