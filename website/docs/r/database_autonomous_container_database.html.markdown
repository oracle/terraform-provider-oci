---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database"
sidebar_current: "docs-oci-resource-database-autonomous_container_database"
description: |-
  Provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_container_database
This resource provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service.

Create a new Autonomous Container Database in the specified Autonomous Exadata Infrastructure.


## Example Usage

```hcl
resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
	#Required
	autonomous_exadata_infrastructure_id = "${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}"
	display_name = "${var.autonomous_container_database_display_name}"
	patch_model = "${var.autonomous_container_database_patch_model}"

	#Optional
	backup_config {

		#Optional
		recovery_window_in_days = "${var.autonomous_container_database_backup_config_recovery_window_in_days}"
	}
	compartment_id = "${var.compartment_id}"
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	service_level_agreement_type = "${var.autonomous_container_database_service_level_agreement_type}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Required) The OCID of the Autonomous Exadata Infrastructure.
* `backup_config` - (Optional) (Updatable) 
	* `recovery_window_in_days` - (Optional) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) The display name for the Autonomous Container Database.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `patch_model` - (Required) (Updatable) Database Patch model preference.
* `service_level_agreement_type` - (Optional) The service level agreement type of the Autonomous Container Database. The default is STANDARD. For a Mission Critical Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

AutonomousContainerDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_container_database.test_autonomous_container_database "id"
```

