---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_protected_databases"
sidebar_current: "docs-oci-datasource-recovery-protected_databases"
description: |-
  Provides the list of Protected Databases in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_protected_databases
This data source provides the list of Protected Databases in Oracle Cloud Infrastructure Recovery service.

Lists the protected databases based on the specified parameters.


## Example Usage

```hcl
data "oci_recovery_protected_databases" "test_protected_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.protected_database_display_name
	id = var.protected_database_id
	protection_policy_id = oci_recovery_protection_policy.test_protection_policy.id
	recovery_service_subnet_id = oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id
	state = var.protected_database_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire 'displayname' given.
* `id` - (Optional) The protected database OCID.
* `protection_policy_id` - (Optional) The protection policy OCID.
* `recovery_service_subnet_id` - (Optional) The recovery service subnet OCID.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `protected_database_collection` - The list of protected_database_collection.

### ProtectedDatabase Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the protected database.
* `database_id` - The OCID of the protected database.
* `database_size` - The size of the protected database. XS - Less than 5GB, S - 5GB to 50GB, M - 50GB to 500GB, L - 500GB to 1TB, XL - 1TB to 5TB, XXL - Greater than 5TB.
* `db_unique_name` - The dbUniqueName for the protected database in Recovery Service. You cannot change the unique name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - The protected database name. You can change the displayName. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `health` - Indicates the protection status of the database.

	A 'PROTECTED' status indicates that Recovery Service can ensure database recovery to any point in time within the entire recovery window. The potential data loss exposure since the last backup is:
	* Less than 10 seconds, if Real-time data protection is enabled
	* Less than 70 minutes if Real-time data protection is disabled

	A 'WARNING' status indicates that Recovery Service can ensure database recovery within the current recovery window - 1 day. The potential data loss exposure since the last backup is:
	* Greater than 10 seconds, if Real-time data protection is enabled
	* Greater than 60 minutes, if if Real-time data protection is disabled

	An 'ALERT' status indicates that Recovery Service cannot recover the database within the current recovery window.  
* `health_details` - A message describing the current health of the protected database.
* `id` - The OCID of the protected database.
* `is_read_only_resource` - Indicates whether the protected database is created by Recovery Service or created manually. Set to <b>TRUE</b> for a service-defined protected database. When you enable the OCI-managed automatic backups option for a database and set Recovery Service as the backup destination, then Recovery Service creates the associated protected database resource. Set to <b>FALSE</b> for a user-defined protected database. 
* `is_redo_logs_shipped` - The value TRUE indicates that the protected database is configured to use Real-time data protection, and redo-data is sent from the protected database to Recovery Service. Real-time data protection substantially reduces the window of potential data loss that exists between successive archived redo log backups. For this to be effective, additional configuration is needed on client side. 
* `lifecycle_details` - Detailed description about the current lifecycle state of the protected database. For example, it can be used to provide actionable information for a resource in a Failed state.
* `metrics` - Backup performance and storage utilization metrics for the protected database.
	* `backup_space_estimate_in_gbs` - The estimated backup storage space, in gigabytes, required to meet the recovery window goal, including foot print and backups for the protected database.
	* `backup_space_used_in_gbs` - Backup storage space, in gigabytes, utilized by the protected database. Oracle charges for the total storage used.
	* `current_retention_period_in_seconds` - Number of seconds backups are currently retained for this database.
	* `db_size_in_gbs` - The estimated space, in gigabytes, consumed by the protected database. The database size is based on the size of the data files in the catalog, and does not include archive logs.
	* `is_redo_logs_enabled` - The value TRUE indicates that the protected database is configured to use Real-time data protection, and redo-data is sent from the protected database to Recovery Service. Real-time data protection substantially reduces the window of potential data loss that exists between successive archived redo log backups. 
	* `minimum_recovery_needed_in_days` - Number of days of redo/archive to be applied to recover database.
	* `retention_period_in_days` - The maximum number of days to retain backups for a protected database.
	* `unprotected_window_in_seconds` - This is the time window when there is data loss exposure. The point after which recovery is impossible unless additional redo is available.  This is the time we received the last backup or last redo-log shipped. 
* `policy_locked_date_time` - An RFC3339 formatted datetime string that specifies the exact date and time for the retention lock to take effect and permanently lock the retention period defined in the policy.

	The retention lock feature controls whether Recovery Service strictly preserves backups for the duration defined in a policy. Retention lock is useful to enforce recovery window compliance and to prevent unintentional modifications to protected database backups.  Recovery Service enforces a 14-day delay before the retention lock set for a policy can take effect. 
* `protection_policy_id` - The OCID of the protection policy associated with the protected database.
* `recovery_service_subnets` - List of recovery service subnet resources associated with the protected database.
	* `recovery_service_subnet_id` - Recovery Service Subnet Identifier.
	* `state` - The current state of the Recovery Service Subnet.
* `state` - The current state of the Protected Database.
* `subscription_id` - The OCID of the cloud service subscription to which the protected database is linked.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_created` - An RFC3339 formatted datetime string that indicates the created time for a protected database. For example: '2020-05-22T21:10:29.600Z' 
* `time_updated` - An RFC3339 formatted datetime string that indicates the last updated time for a protected database. For example: '2020-05-22T21:10:29.600Z' 
* `vpc_user_name` - The virtual private catalog (VPC) user credentials that authenticates the protected database to access Recovery Service.

