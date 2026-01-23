---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_db_system_maintenance_events"
sidebar_current: "docs-oci-datasource-mysql-db_system_maintenance_events"
description: |-
  Provides the list of Db System Maintenance Events in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_db_system_maintenance_events
This data source provides the list of Db System Maintenance Events in Oracle Cloud Infrastructure MySQL Database service.

List all the maintenance events.

## Example Usage

```hcl
data "oci_mysql_db_system_maintenance_events" "test_db_system_maintenance_events" {
	#Required
	db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id

	#Optional
	maintenance_action = var.db_system_maintenance_event_maintenance_action
	maintenance_status = var.db_system_maintenance_event_maintenance_status
	maintenance_type = var.db_system_maintenance_event_maintenance_type
	mysql_version_after_maintenance = var.db_system_maintenance_event_mysql_version_after_maintenance
	mysql_version_before_maintenance = var.db_system_maintenance_event_mysql_version_before_maintenance
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `maintenance_action` - (Optional) The nature of the maintenance event.
* `maintenance_status` - (Optional) The last status of the maintenance event.
* `maintenance_type` - (Optional) How the maintenance event was triggered.
* `mysql_version_after_maintenance` - (Optional) The MySQL version after the maintenance event.
* `mysql_version_before_maintenance` - (Optional) The MySQL version before the maintenance event.


## Attributes Reference

The following attributes are exported:

* `maintenance_events` - The list of maintenance_events.

### DbSystemMaintenanceEvent Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment the maintenance event belongs to.
* `db_system_id` - The OCID of the DB System this maintenance event is associated with.
* `maintenance_action` - The nature of the maintenance event.
* 
	DATABASE:  maintenance event causing a MySQL version upgrade. This may also include OS updates. OS_UPDATE: maintenance event causing an OS update. ONLINE_UPDATE: maintenance event causing downtime-free OS security patches. HARDWARE: hardware maintenance event affecting the DB System's VMs and BMs. 
* `maintenance_scope` - The MySQL instances operated during a maintenance event.

	ALL:             maintenance event targeting all MySQL instances in a DB System. ALL_BUT_PRIMARY: maintenance event targeting all MySQL instances in a highly available DB System except the primary group member. PRIMARY_ONLY:    maintenance event targeting only the primary group member in a highly available DB System. 
* `maintenance_status` - The last status of the maintenance event. 
* `maintenance_type` - How the maintenance event was triggered.

	AUTOMATIC:  maintenance event triggered as part of scheduled maintenance. MANUAL:     maintenance event triggered manually. SHAPE:      maintenance event triggered by a shape update. 
* `mysql_version_after_maintenance` - The MySQL version after the maintenance.
* `mysql_version_before_maintenance` - The MySQL version prior to the maintenance.
* `time_created` - The date and time the record was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_ended` - The date and time the maintenance event ended, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_mysql_switch_over_ended` - The date and time the DB System came back online during the maintenance, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_mysql_switch_over_started` - The date and time the DB System was initially down during the maintenance, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_started` - The date and time the maintenance event started, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

