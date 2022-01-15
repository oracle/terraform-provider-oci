---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_dataguard_associations"
sidebar_current: "docs-oci-datasource-database-autonomous_database_dataguard_associations"
description: |-
  Provides the list of Autonomous Database Dataguard Associations in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_dataguard_associations
This data source provides the list of Autonomous Database Dataguard Associations in Oracle Cloud Infrastructure Database service.

Gets a list of the Autonomous Data Guard-enabled databases associated with the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database_dataguard_associations" "test_autonomous_database_dataguard_associations" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_database_dataguard_associations` - The list of autonomous_database_dataguard_associations.

### AutonomousDatabaseDataguardAssociation Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database. 
* `id` - The OCID of the Autonomous Dataguard created for Autonomous Container Database where given Autonomous Database resides in.
* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database. 
* `peer_autonomous_database_life_cycle_state` - The current state of Autonomous Data Guard.
* `peer_role` - The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled. 
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled. 
* `state` - The current state of Autonomous Data Guard.
* `time_created` - The date and time the Data Guard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.
* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database.  Example: `7 seconds` 

