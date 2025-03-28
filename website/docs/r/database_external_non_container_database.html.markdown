---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_non_container_database"
sidebar_current: "docs-oci-resource-database-external_non_container_database"
description: |-
  Provides the External Non Container Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_external_non_container_database
This resource provides the External Non Container Database resource in Oracle Cloud Infrastructure Database service.

Creates a new ExternalNonContainerDatabase resource


## Example Usage

```hcl
resource "oci_database_external_non_container_database" "test_external_non_container_database" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.external_non_container_database_display_name

	#Optional
	defined_tags = var.external_non_container_database_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the external database. The name does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `character_set` - The character set of the external database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_configuration` - The Oracle Database configuration
* `database_edition` - The Oracle Database Edition that applies to all the databases on the DB system. Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE. 
* `database_management_config` - The configuration of the Database Management service.
	* `database_management_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
	* `database_management_status` - The status of the Database Management service.
	* `license_model` - The Oracle license model that applies to the external database. 
* `database_version` - The Oracle Database version.
* `db_id` - The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud. 
* `db_packs` - The database packs licensed for the external Oracle Database.
* `db_unique_name` - The `DB_UNIQUE_NAME` of the external database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the external database. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure external database resource. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `ncharacter_set` - The national character of the external database.
* `operations_insights_config` - The configuration of Operations Insights for the external database
	* `operations_insights_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
	* `operations_insights_status` - The status of Operations Insights
* `stack_monitoring_config` - The configuration of Stack Monitoring for the external database.
	* `stack_monitoring_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
	* `stack_monitoring_status` - The status of Stack Monitoring.
* `state` - The current state of the Oracle Cloud Infrastructure external database resource.
* `time_created` - The date and time the database was created.
* `time_zone` - The time zone of the external database. It is a time zone offset (a character type in the format '[+|-]TZH:TZM') or a time zone region name, depending on how the time zone value was specified when the database was created / last altered. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Non Container Database
	* `update` - (Defaults to 20 minutes), when updating the External Non Container Database
	* `delete` - (Defaults to 20 minutes), when destroying the External Non Container Database


## Import

ExternalNonContainerDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_external_non_container_database.test_external_non_container_database "id"
```

