---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_databases"
sidebar_current: "docs-oci-datasource-database_management-external_databases"
description: |-
Provides the list of External Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_databases
This data source provides the list of External Databases in Oracle Cloud Infrastructure Database Management service.

Lists the external databases in the specified compartment or in the specified DB system.

## Example Usage

```hcl
data "oci_database_management_external_databases" "test_external_databases" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.external_database_display_name
	external_database_id = oci_database_management_external_database.test_external_database.id
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.
* `external_database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database.
* `external_db_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.


## Attributes Reference

The following attributes are exported:

* `external_database_collection` - The list of external_database_collection.

### ExternalDatabase Reference

The following attributes are exported:

* `items` - An array of external databases.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `database_sub_type` - The subtype of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or Non-container Database.
	* `database_type` - The type of Oracle Database installation.
	* `db_management_config` - The configuration of the Database Management service.
		* `connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database connector.
		* `database_management_status` - The status of the Database Management service.
		* `license_model` - The Oracle license model that applies to the external database.
	* `db_system_info` - The basic information about an external DB system.
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
		* `exadata_infra_info` - The basic information about an external Exadata Infrastructure.
			* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
			* `display_name` - The user-friendly name for the Exadata Infrastructure. The name does not have to be unique.
			* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external Exadata Infrastructure.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
	* `db_unique_name` - The `DB_UNIQUE_NAME` of the external database.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The user-friendly name for the database. The name does not have to be unique.
	* `external_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent Container Database (CDB) if this is a Pluggable Database (PDB).
	* `external_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB home.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
	* `instance_details` - The list of database instances if the database is a RAC database.
		* `host_name` - The name of the host machine.
		* `instance_name` - The name of the database instance.
		* `instance_number` - The instance number of the database instance.
	* `state` - The current lifecycle state of the external database resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the external DB system was created.

