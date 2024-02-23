---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system"
sidebar_current: "docs-oci-resource-database_management-external_db_system"
description: |-
  Provides the External Db System resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_db_system
This resource provides the External Db System resource in Oracle Cloud Infrastructure Database Management service.

Creates an external DB system and its related resources.


## Example Usage

```hcl
resource "oci_database_management_external_db_system" "test_external_db_system" {
	#Required
	compartment_id = var.compartment_id
	db_system_discovery_id = oci_database_management_db_system_discovery.test_db_system_discovery.id

	#Optional
	database_management_config {
		#Required
		license_model = var.external_db_system_database_management_config_license_model
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.external_db_system_display_name
	freeform_tags = {"Department"= "Finance"}
	stack_monitoring_config {
		#Required
		is_enabled = var.external_db_system_stack_monitoring_config_is_enabled

		#Optional
		metadata = var.external_db_system_stack_monitoring_config_metadata
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the external DB system resides.
* `database_management_config` - (Optional) The details required to enable Database Management for an external DB system.
	* `license_model` - (Required) The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `stack_monitoring_config` - (Optional) The details of the associated service that will be enabled or disabled for an external DB System.
	* `is_enabled` - (Required) The status of the associated service.
	* `metadata` - (Optional) The associated service-specific inputs in JSON string format, which Database Management can identify.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_management_config` - The details required to enable Database Management for an external DB system.
	* `license_model` - The Oracle license model that applies to the external database. 
* `db_system_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system discovery.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `discovery_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used during the discovery of the DB system.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `home_directory` - The Oracle Grid home directory in case of cluster-based DB system and Oracle home directory in case of single instance-based DB system. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `is_cluster` - Indicates whether the DB system is a cluster DB system or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `stack_monitoring_config` - The configuration details of Stack Monitoring for an external DB system.
	* `is_enabled` - The status of the associated service.
	* `metadata` - The associated service-specific inputs in JSON string format, which Database Management can identify.
* `state` - The current lifecycle state of the external DB system resource.
* `time_created` - The date and time the external DB system was created.
* `time_updated` - The date and time the external DB system was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Db System
	* `update` - (Defaults to 20 minutes), when updating the External Db System
	* `delete` - (Defaults to 20 minutes), when destroying the External Db System


## Import

ExternalDbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_db_system.test_external_db_system "id"
```

