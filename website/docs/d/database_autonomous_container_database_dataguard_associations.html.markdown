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

Gets a list of the Autonomous Container Database dataguard associations for the specified Autonomous Container Database.


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

* `apply_lag` - The lag time between updates to the primary autonomous container database and application of the redo data on the standby autonomous container database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated container databases.  Example: `180 Mb per second` 
* `autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. 
* `id` - The OCID of the Autonomous Dataguard created for given Autonomous Container Database.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_container_database_dataguard_association_id` - The OCID of the peer Autonomous Container Database Dataguard Association.
* `peer_autonomous_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database. 
* `peer_lifecycle_state` - The current state of the Autonomous Dataguard.
* `peer_role` - The role of the Autonomous Dataguard enabled Autonomous Container Database.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the Autonomous Dataguard enabled Autonomous Container Database.
* `state` - The current state of the Autonomous Dataguard.
* `time_created` - The date and time the Autonomous DataGuard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.

