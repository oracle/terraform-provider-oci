---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_groups"
sidebar_current: "docs-oci-datasource-database_management-managed_database_groups"
description: |-
  Provides the list of Managed Database Groups in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_groups
This data source provides the list of Managed Database Groups in Oracle Cloud Infrastructure Database Management service.

Gets the Managed Database Group for a specific ID or the list of Managed Database Groups in
a specific compartment. Managed Database Groups can also be filtered based on the name parameter.
Only one of the parameters, ID or name should be provided. If none of these parameters is provided,
all the Managed Database Groups in the compartment are listed.


## Example Usage

```hcl
data "oci_database_management_managed_database_groups" "test_managed_database_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.managed_database_group_id
	name = var.managed_database_group_name
	state = var.managed_database_group_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `id` - (Optional) The identifier of the resource. Only one of the parameters, id or name should be provided.
* `name` - (Optional) A filter to return only resources that match the entire name. Only one of the parameters, id or name should be provided
* `state` - (Optional) The lifecycle state of a resource.


## Attributes Reference

The following attributes are exported:

* `managed_database_group_collection` - The list of managed_database_group_collection.

### ManagedDatabaseGroup Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `description` - The information specified by the user about the Managed Database Group.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
* `managed_databases` - A list of Managed Databases in the Managed Database Group.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database resides.
	* `database_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `deployment_type` - The infrastructure used to deploy the Oracle Database.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	* `name` - The name of the Managed Database.
	* `time_added` - The date and time the Managed Database was added to the group.
	* `workload_type` - The workload type of the Autonomous Database.
* `name` - The name of the Managed Database Group.
* `state` - The current lifecycle state of the Managed Database Group.
* `time_created` - The date and time the Managed Database Group was created.
* `time_updated` - The date and time the Managed Database Group was last updated.

