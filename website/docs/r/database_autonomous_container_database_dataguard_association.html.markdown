---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_dataguard_association"
sidebar_current: "docs-oci-resource-database-autonomous_container_database_dataguard_association"
description: |-
  Provides the Autonomous Container Database Dataguard Association resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_container_database_dataguard_association
This resource provides the Autonomous Container Database Dataguard Association resource in Oracle Cloud Infrastructure Database service.

Create a new Autonomous Data Guard association. An Autonomous Data Guard association represents the replication relationship between the
specified Autonomous Container database and a peer Autonomous Container database. For more information, see [Using Oracle Data Guard](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/usingdataguard.htm).

All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
resource in the Console. For more information, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
resource "oci_database_autonomous_container_database_dataguard_association" "test_autonomous_container_database_dataguard_association" {
	#Required
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
	peer_autonomous_container_database_display_name = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_display_name
	peer_cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
	protection_mode = var.autonomous_container_database_dataguard_association_protection_mode

	#Optional
	fast_start_fail_over_lag_limit_in_seconds = var.autonomous_container_database_dataguard_association_fast_start_fail_over_lag_limit_in_seconds
	is_automatic_failover_enabled = var.autonomous_container_database_dataguard_association_is_automatic_failover_enabled
	peer_autonomous_container_database_backup_config {

		#Optional
		backup_destination_details {
			#Required
			type = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_backup_destination_details_type

			#Optional
			dbrs_policy_id = oci_identity_policy.test_policy.id
			id = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_backup_destination_details_id
			internet_proxy = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_backup_destination_details_internet_proxy
			vpc_password = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_password
			vpc_user = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_backup_destination_details_vpc_user
		}
		recovery_window_in_days = var.autonomous_container_database_dataguard_association_peer_autonomous_container_database_backup_config_recovery_window_in_days
	}
	peer_autonomous_container_database_compartment_id = oci_identity_compartment.test_compartment.id
	standby_maintenance_buffer_in_days = var.autonomous_container_database_dataguard_association_standby_maintenance_buffer_in_days
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `fast_start_fail_over_lag_limit_in_seconds` - (Optional) (Updatable) The lag time for my preference based on data loss tolerance in seconds.
* `peer_autonomous_container_database_backup_config` - (Optional) Backup options for the standby Autonomous Container Database. 
	* `backup_destination_details` - (Optional) Backup destination details.
		* `dbrs_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DBRS policy used for backup.
		* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - (Optional) Proxy URL to connect to object store.
		* `type` - (Required) Type of the database backup destination.
		* `vpc_password` - (Optional) For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - (Optional) For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - (Optional) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `peer_autonomous_container_database_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database will be created. 
* `peer_autonomous_container_database_display_name` - (Required) The display name for the peer Autonomous Container Database.
* `peer_cloud_autonomous_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
* `standby_maintenance_buffer_in_days` - (Optional) The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database. This value represents the number of days before scheduled maintenance of the primary database. 
* `is_automatic_failover_enabled` - (Optional) (Updatable) Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association. Input DataType: boolean. Example : `is_automatic_failover_enabled = true`.  
* `protection_mode` - (Optional) (Updatable) The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 



** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synchronized between the associated Autonomous Container Databases.  Example: `180 Mb per second` 
* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. 
* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
* `id` - The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association. Output DataType: boolean. Example : `is_automatic_failover_enabled = true`. 
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_container_database_dataguard_association_id` - The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
* `peer_autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database. 
* `peer_lifecycle_state` - The current state of Autonomous Data Guard.
* `peer_role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `protection_mode` - The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `state` - The current state of Autonomous Data Guard.
* `time_created` - The date and time the Autonomous DataGuard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.
* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database.  Example: `7 seconds` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Container Database Dataguard Association
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Container Database Dataguard Association
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Container Database Dataguard Association


## Import

AutonomousContainerDatabaseDataguardAssociations can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association "autonomousContainerDatabases/{autonomousContainerDatabaseId}/autonomousContainerDatabaseDataguardAssociations/{autonomousContainerDatabaseDataguardAssociationId}" 
```

