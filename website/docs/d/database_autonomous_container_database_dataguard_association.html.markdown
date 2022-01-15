---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database_dataguard_association"
sidebar_current: "docs-oci-datasource-database-autonomous_container_database_dataguard_association"
description: |-
  Provides details about a specific Autonomous Container Database Dataguard Association in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_database_dataguard_association
This data source provides details about a specific Autonomous Container Database Dataguard Association resource in Oracle Cloud Infrastructure Database service.

Gets an Autonomous Container Database enabled with Autonomous Data Guard associated with the specified Autonomous Container Database.


## Example Usage

```hcl
data "oci_database_autonomous_container_database_dataguard_association" "test_autonomous_container_database_dataguard_association" {
	#Required
	autonomous_container_database_dataguard_association_id = oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association.id
	autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_container_database_dataguard_association_id` - (Required) The Autonomous Container Database-Autonomous Data Guard association [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_container_database_id` - (Required) The Autonomous Container Database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synchronized between the associated Autonomous Container Databases.  Example: `180 Mb per second` 
* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. 
* `id` - The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
* `is_automatic_failover_enabled` - Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association 
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_container_database_dataguard_association_id` - The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
* `peer_autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database. 
* `peer_lifecycle_state` - The current state of Autonomous Data Guard.
* `peer_role` - The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled. 
* `protection_mode` - The protection mode of this Autonomous Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled. 
* `state` - The current state of Autonomous Data Guard.
* `time_created` - The date and time the Autonomous DataGuard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.
* `time_last_synced` - The date and time of the last update to the apply lag, apply rate, and transport lag values.
* `transport_lag` - The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database, as computed by the reporting database.  Example: `7 seconds` 

