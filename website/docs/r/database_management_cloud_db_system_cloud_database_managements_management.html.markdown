---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_system_cloud_database_managements_management"
sidebar_current: "docs-oci-resource-database_management-cloud_db_system_cloud_database_managements_management"
description: |-
  Provides the Cloud Db System Cloud Database Managements Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_db_system_cloud_database_managements_management
This resource provides the Cloud Db System Cloud Database Managements Management resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/CloudDbSystemCloudDatabaseManagementsManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement
Enables Database Management service for all the components of the specified
cloud DB system (except databases).


## Example Usage

```hcl
resource "oci_database_management_cloud_db_system_cloud_database_managements_management" "test_cloud_db_system_cloud_database_managements_management" {
	#Required
	cloud_db_system_id = oci_database_management_cloud_db_system.test_cloud_db_system.id
	enable_cloud_database_management = var.enable_cloud_database_management

	#Optional
	is_enabled = var.cloud_db_system_cloud_database_managements_management_is_enabled
	metadata = var.cloud_db_system_cloud_database_managements_management_metadata
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `is_enabled` - (Optional) The status of the associated service.
* `metadata` - (Optional) The associated service-specific inputs in JSON string format, which Database Management can identify.
* `enable_cloud_database_management` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Db System Cloud Database Managements Management
	* `update` - (Defaults to 20 minutes), when updating the Cloud Db System Cloud Database Managements Management
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Db System Cloud Database Managements Management
