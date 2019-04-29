---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_databases"
sidebar_current: "docs-oci-datasource-database-autonomous_container_databases"
description: |-
  Provides the list of Autonomous Container Databases in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_container_databases
This data source provides the list of Autonomous Container Databases in Oracle Cloud Infrastructure Database service.

Gets a list of the Autonomous Container Databases in the specified compartment.


## Example Usage

```hcl
data "oci_database_autonomous_container_databases" "test_autonomous_container_databases" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	autonomous_exadata_infrastructure_id = "${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}"
	availability_domain = "${var.autonomous_container_database_availability_domain}"
	display_name = "${var.autonomous_container_database_display_name}"
	state = "${var.autonomous_container_database_state}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Optional) The Autonomous Exadata Infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_container_databases` - The list of autonomous_container_databases.

### AutonomousContainerDatabase Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_id` - The OCID of the Autonomous Exadata Infrastructure.
* `availability_domain` - The availability domain of the Autonomous Container Database.
* `backup_config` - 
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `patch_model` - Database Patch model preference.
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `state` - The current state of the Autonomous Container Database.
* `time_created` - The date and time the Autonomous was created.

