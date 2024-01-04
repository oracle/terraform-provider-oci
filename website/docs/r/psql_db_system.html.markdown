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

Creates a new database system.


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
	apply_config = var.db_system_apply_config_type
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

	# Optional
	patch_operations {
		#Required
		operation = var.db_system_patch_operations_operation
		selection = var.db_system_patch_operations_selection

		#Optional
		value = var.db_system_patch_operations_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the database system.
* `config_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration associated with the database system.
* `apply_config` - (Optional) Whether a configuration update requires a restart of the database instance or a reload of the configuration. Some configuration changes require a restart of database instances to be applied. Apply config can be passed as `RESTART` or `RELOAD`
* `credentials` - (Optional) Initial database system credentials that the database system will be provisioned with. The password details are not visible on any subsequent operation, such as GET /dbSystems/{dbSystemId}. 
	* `password_details` - (Required) Details for the database system password. Password can be passed as `VaultSecretPasswordDetails` or `PlainTextPasswordDetails`. 
		* `password` - (Required when password_type=PLAIN_TEXT) The database system password.
		* `password_type` - (Required) The password type.
		* `secret_id` - (Required when password_type=VAULT_SECRET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret where the password is stored.
		* `secret_version` - (Required when password_type=VAULT_SECRET) The secret version of the stored password.
	* `username` - (Required) The database system administrator username.
* `db_version` - (Required) Version of database system software.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-provided description of a database system.
* `display_name` - (Required) (Updatable) A user-friendly display name for the database system. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `instance_count` - (Required) (Updatable when patch_operations are specified) Count of database instances nodes to be created in the database system. 
* `instance_memory_size_in_gbs` - (Optional) The total amount of memory available to each database instance node, in gigabytes.
* `instance_ocpu_count` - (Optional) The total number of OCPUs available to each database instance node.
* `instances_details` - (Optional) Details of database instances nodes to be created. This parameter is optional. If specified, its size must match `instanceCount`. 
	* `description` - (Optional) A user-provided description of the database instance node.
	* `display_name` - (Optional) Display name of the database instance node. Avoid entering confidential information.
	* `private_ip` - (Optional) Private IP in customer subnet that will be assigned to the database instance node. This value is optional. If the IP is not provided, the IP will be chosen from the available IP addresses in the specified subnet. 
* `management_policy` - (Optional) (Updatable) PostgreSQL database system management policy update details.
	* `backup_policy` - (Optional) (Updatable) PostgreSQL database system backup policy.
		* `backup_start` - (Required when kind=DAILY | MONTHLY | WEEKLY) (Updatable) Hour of the day when the backup starts.
		* `days_of_the_month` - (Required when kind=MONTHLY) (Updatable) Day of the month when the backup should start. To ensure that the backup runs monthly, the latest day of the month that you can use to schedule a backup is the the 28th day. 
		* `days_of_the_week` - (Required when kind=WEEKLY) (Updatable) The day of the week that the backup starts.
		* `kind` - (Optional) (Updatable) The kind of backup policy.
		* `retention_days` - (Optional) (Updatable) How many days the data should be stored after the database system deletion.
	* `maintenance_window_start` - (Optional) (Updatable) The start of the maintenance window. 
* `network_details` - (Required) Network details for the database system.
	* `nsg_ids` - (Optional) List of customer Network Security Group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the database system.
	* `primary_db_endpoint_private_ip` - (Optional) Private IP in customer subnet. The value is optional. If the IP is not provided, the IP will be chosen from the available IP addresses from the specified subnet. 
	* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer subnet associated with the database system.
* `patch_operations` - (Optional) (Updatable) For adding and removing from read replica database instances. Please remove the patch_operations after it is applied. Update the instance_count arrodrandly. Cannot be specified when creating the resource.
	* `operation` - (Required) The operation can be one of these values: `INSERT`, `REMOVE`. 
	* `selection` - (Required) In case of `INSERT`, selection is `instances`. In case of `REMOVE`, selection is `instances[?id == '${var.instance_id}']`.
	* `value` - (Required when operation=INSERT) Specify instance details such as displayName, description or privateIp. Example: `{"displayName": "value"}`.
* `shape` - (Required) The name of the shape for the database instance node. Use the /shapes API for accepted shapes. Example: `VM.Standard.E4.Flex` 
* `source` - (Optional) The source used to restore the database system.
	* `backup_id` - (Required when source_type=BACKUP) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database system backup.
	* `is_having_restore_config_overrides` - (Applicable when source_type=BACKUP) Deprecated. Don't use.
	* `source_type` - (Required) The source descriminator. 
* `storage_details` - (Required) (Updatable) Storage details of the database system.
	* `availability_domain` - (Optional) Specifies the availability domain of AD-local storage. If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified. If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified. 
	* `iops` - (Applicable when system_type=OCI_OPTIMIZED_STORAGE) (Updatable) Guaranteed input/output storage requests per second (IOPS) available to the database system.
	* `is_regionally_durable` - (Required) Specifies if the block volume used for the database system is regional or AD-local. If not specified, it will be set to false. If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified. If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified. 
	* `system_type` - (Required) Type of the database system.
* `system_type` - (Optional) Type of the database system.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

