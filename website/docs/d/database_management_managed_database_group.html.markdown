---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_group"
sidebar_current: "docs-oci-datasource-database_management-managed_database_group"
description: |-
  Provides details about a specific Managed Database Group in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_group
This data source provides details about a specific Managed Database Group resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the Managed Database Group specified by managedDatabaseGroupId.


## Example Usage

```hcl
data "oci_database_management_managed_database_group" "test_managed_database_group" {
	#Required
	managed_database_group_id = oci_database_management_managed_database_group.test_managed_database_group.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.


## Attributes Reference

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

