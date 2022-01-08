---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_data_guard_associations"
sidebar_current: "docs-oci-datasource-database-data_guard_associations"
description: |-
  Provides the list of Data Guard Associations in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_data_guard_associations
This data source provides the list of Data Guard Associations in Oracle Cloud Infrastructure Database service.

Lists all Data Guard associations for the specified database.


## Example Usage

```hcl
data "oci_database_data_guard_associations" "test_data_guard_associations" {
	#Required
	database_id = oci_database_database.test_database.id
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `data_guard_associations` - The list of data_guard_associations.

### DataGuardAssociation Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the reporting database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Data Guard association.
* `is_active_data_guard_enabled` - True if active Data Guard is enabled.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_data_guard_association_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
* `peer_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated peer database.
* `peer_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home containing the associated peer database. 
* `peer_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system containing the associated peer database. 
* `peer_role` - The role of the peer database in this Data Guard association.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the reporting database in this Data Guard association.
* `state` - The current state of the Data Guard association.
* `time_created` - The date and time the Data Guard association was created.
* `transport_type` - The redo transport type used by this Data Guard association.  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation. 

