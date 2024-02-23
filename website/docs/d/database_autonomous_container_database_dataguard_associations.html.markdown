---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_dataguard_associations"
sidebar_current: "docs-oci-datasource-database-autonomous_container_database_dataguard_associations"
description: |-
  Provides the list of Autonomous Container Database Dataguard Associations in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_database_dataguard_associations
This data source provides the list of Autonomous Container Database Dataguard Associations in Oracle Cloud Infrastructure Database service.

Gets a list of the Autonomous Container Databases with Autonomous Data Guard-enabled associated with the specified Autonomous Container Database.


## Example Usage

```hcl
data "oci_database_autonomous_container_database_dataguard_associations" "test_autonomous_container_database_dataguard_associations" {
	#Required
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_container_database_dataguard_associations` - The list of autonomous_container_database_dataguard_associations.

### AutonomousContainerDatabaseDataguardAssociation Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synchronized between the associated Autonomous Container Databases.  Example: `180 Mb per second` 
* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. 
* `fast_start_fail_over_lag_limit_in_seconds` - The lag time for my preference based on data loss tolerance in seconds.
* `id` - The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association. Output DataType: boolean. Example : is_automatic_failover_enabled = true.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_container_database_dataguard_association_id` - The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
* `peer_autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database. 
* `peer_lifecycle_state` - The current state of the Autonomous Container Database.
* `peer_role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `protection_mode` - The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled. 
* `state` - The current state of Autonomous Data Guard.
* `time_created` - The date and time the Autonomous DataGuard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.
* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database.  Example: `7 seconds` 

