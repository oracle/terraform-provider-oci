---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_asms"
sidebar_current: "docs-oci-datasource-database_management-cloud_asms"
description: |-
  Provides the list of Cloud Asms in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_asms
This data source provides the list of Cloud Asms in Oracle Cloud Infrastructure Database Management service.

Lists the ASMs in the specified cloud DB system.

## Example Usage

```hcl
data "oci_database_management_cloud_asms" "test_cloud_asms" {

	#Optional
	cloud_db_system_id = oci_database_management_cloud_db_system.test_cloud_db_system.id
	compartment_id = var.compartment_id
	display_name = var.cloud_asm_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `cloud_asm_collection` - The list of cloud_asm_collection.

### CloudAsm Reference

The following attributes are exported:

* `additional_details` - The additional details of the cloud ASM defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `cloud_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud connector.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the ASM is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud ASM.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud ASM. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `grid_home` - The directory in which ASM is installed. This is the same directory in which Oracle Grid Infrastructure is installed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM.
* `is_cluster` - Indicates whether the ASM is a cluster ASM or not.
* `is_flex_enabled` - Indicates whether Oracle Flex ASM is enabled or not.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `serviced_databases` - The list of databases that are serviced by the ASM.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the cloud database resides.
	* `database_sub_type` - The subtype of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `db_unique_name` - The unique name of the cloud database.
	* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas Oracle Cloud Infrastructure resource matching the discovered DB system component.
	* `disk_groups` - The list of ASM disk groups used by the database.
	* `display_name` - The user-friendly name for the database. The name does not have to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud database.
	* `is_managed` - Indicates whether the database is a Managed Database or not.
* `state` - The current lifecycle state of the cloud ASM.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud ASM was created.
* `time_updated` - The date and time the cloud ASM was last updated.
* `version` - The ASM version.

