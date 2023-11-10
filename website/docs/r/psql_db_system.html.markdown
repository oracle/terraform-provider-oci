---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_db_system"
sidebar_current: "docs-oci-resource-psql-db_system"
description: |-
  Provides the Db System resource in Oracle Cloud Infrastructure Psql service
---

# oci_psql_db_system
This resource provides the Db System resource in Oracle Cloud Infrastructure Psql service.

Creates a new DbSystem.


## Example Usage

```hcl
resource "oci_psql_db_system" "test_db_system" {
	#Required
	compartment_id = var.compartment_id
	db_version = var.db_system_db_version
	display_name = var.db_system_display_name
	network_details {
		#Required
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		nsg_ids = var.db_system_network_details_nsg_ids
		primary_db_endpoint_private_ip = var.db_system_network_details_primary_db_endpoint_private_ip
	}
	shape = var.db_system_shape
	storage_details {
		#Required
		is_regionally_durable = var.db_system_storage_details_is_regionally_durable
		system_type = var.db_system_storage_details_system_type

		#Optional
		availability_domain = var.db_system_storage_details_availability_domain
		iops = var.db_system_storage_details_iops
	}

	#Optional
	config_id = oci_apm_config_config.test_config.id
	credentials {
		#Required
		password_details {
			#Required
			password_type = var.db_system_credentials_password_details_password_type

			#Optional
			password = var.db_system_credentials_password_details_password
			secret_id = oci_vault_secret.test_secret.id
			secret_version = var.db_system_credentials_password_details_secret_version
		}
		username = var.db_system_credentials_username
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.db_system_description
	freeform_tags = {"bar-key"= "value"}
	instance_count = var.db_system_instance_count
	instance_memory_size_in_gbs = var.db_system_instance_memory_size_in_gbs
	instance_ocpu_count = var.db_system_instance_ocpu_count
	instances_details {

		#Optional
		description = var.db_system_instances_details_description
		display_name = var.db_system_instances_details_display_name
		private_ip = var.db_system_instances_details_private_ip
	}
	management_policy {

		#Optional
		backup_policy {

			#Optional
			backup_start = var.db_system_management_policy_backup_policy_backup_start
			days_of_the_month = var.db_system_management_policy_backup_policy_days_of_the_month
			days_of_the_week = var.db_system_management_policy_backup_policy_days_of_the_week
			kind = var.db_system_management_policy_backup_policy_kind
			retention_days = var.db_system_management_policy_backup_policy_retention_days
		}
		maintenance_window_start = var.db_system_management_policy_maintenance_window_start
	}
	source {
		#Required
		source_type = var.db_system_source_source_type

		#Optional
		backup_id = oci_psql_backup.test_backup.id
		is_having_restore_config_overrides = var.db_system_source_is_having_restore_config_overrides
	}
	system_type = var.db_system_system_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier
* `config_id` - (Optional) Configuration identifier
* `credentials` - (Optional) Initial DbSystem credentials that the DbSystem will be provisioned with. The password details are not visible on any subsequent operation, such as GET /dbSystems/{dbSystemId}. 
	* `password_details` - (Required) Details for the DbSystem password. Password can be passed as `VaultSecretPasswordDetails`(Vault) or `PlainTextPasswordDetails`. 
		* `password` - (Required when password_type=PLAIN_TEXT) The dbSystem password.
		* `password_type` - (Required) Password type
		* `secret_id` - (Required when password_type=VAULT_SECRET) The OCID of secret where the password is stored.
		* `secret_version` - (Required when password_type=VAULT_SECRET) The secret version where the password is stored.
	* `username` - (Required) The DB system username.
* `db_version` - (Required) Version of DbSystem software.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of a DbSystem. This field should be input by the user.
* `display_name` - (Required) (Updatable) DbSystem display name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `instance_count` - (Optional) Count of DbInstances to be created in the DbSystem. 
* `instance_memory_size_in_gbs` - (Optional) The total amount of memory available to each DbInstance, in gigabytes.
* `instance_ocpu_count` - (Optional) The total number of OCPUs available to each DbInstance.
* `instances_details` - (Optional) Details of DbInstances to be created. Optional parameter. If specified, its size must match instanceCount. 
	* `description` - (Optional) Description of the DbInstance. This field should be input by the user.
	* `display_name` - (Optional) Display name of the DbInstance.
	* `private_ip` - (Optional) Private IP in customer subnet that will be assigned to the DbInstance. The value is optional. If the IP is not provided the IP will be chosen among the available IP addresses from the specified subnet. 
* `management_policy` - (Optional) (Updatable) Posgresql DB system management policy update details
	* `backup_policy` - (Optional) (Updatable) Posgresql DB system backup policy
		* `backup_start` - (Required when kind=DAILY | MONTHLY | WEEKLY) (Updatable) Hour of the day when backup starts.
		* `days_of_the_month` - (Required when kind=MONTHLY) (Updatable) Days of the month when backup should start. If the day is greater last day of the current month, then it will be triggered on the last day of the current month 
		* `days_of_the_week` - (Required when kind=WEEKLY) (Updatable) Weekly days
		* `kind` - (Optional) (Updatable) Backup policy kind
		* `retention_days` - (Optional) (Updatable) How many days the customers data should be stored after the db system deletion.
	* `maintenance_window_start` - (Optional) (Updatable) The start of the maintenance window. 
* `network_details` - (Required) DbSystem network details.
	* `nsg_ids` - (Optional) List of customer NetworkSecurityGroup identifiers
	* `primary_db_endpoint_private_ip` - (Optional) Private IP in customer subnet. The value is optional. If the IP is not provided the IP will be chosen among the available IP addresses from the specified subnet. 
	* `subnet_id` - (Required) Customer Subnet identifier
* `shape` - (Required) Shape of DbInstance. This name should match from with one of the available shapes from /shapes API.
* `source` - (Optional) New source is used to restore the DB system.
	* `backup_id` - (Required when source_type=BACKUP) DbSystem backup identifier.
	* `is_having_restore_config_overrides` - (Applicable when source_type=BACKUP) Restore the DB config overrides from backup. Default is false
	* `source_type` - (Required) The source descriminator. 
* `storage_details` - (Required) (Updatable) Storage details of the DbSystem.
	* `availability_domain` - (Optional) Specifies the availability domain of AD-local storage. If isRegionallyDurable is set to true, availabilityDomain should not be specified. If isRegionallyDurable is set to false, availabilityDomain must be specified. 
	* `iops` - (Applicable when system_type=OCI_OPTIMIZED_STORAGE) (Updatable) DbSystem Performance Unit
	* `is_regionally_durable` - (Required) Specifies if the block volume used for the DbSystem is regional or AD-local. If not specified, it will be set to false. If isRegionallyDurable is set to true, availabilityDomain should not be specified. If isRegionallyDurable is set to false, availabilityDomain must be specified. 
	* `system_type` - (Required) Type of the DbSystem.
* `system_type` - (Optional) Type of the DbSystem.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db System
	* `update` - (Defaults to 20 minutes), when updating the Db System
	* `delete` - (Defaults to 20 minutes), when destroying the Db System


## Import

DbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_psql_db_system.test_db_system "id"
```

