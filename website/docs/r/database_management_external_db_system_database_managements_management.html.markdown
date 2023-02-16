---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system_database_managements_management"
sidebar_current: "docs-oci-resource-database_management-external_db_system_database_managements_management"
description: |-
  Provides the External Db System Database Managements Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_db_system_database_managements_management
This resource provides the External Db System Database Managements Management resource in Oracle Cloud Infrastructure Database Management service.

Enables Database Management service for all the components of the specified
external DB system (except databases).


## Example Usage

```hcl
resource "oci_database_management_external_db_system_database_managements_management" "test_external_db_system_database_managements_management" {
	#Required
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
	enable_database_management = var.enable_database_management

	#Optional
	license_model = var.external_db_system_database_managements_management_license_model
}
```

## Argument Reference

The following arguments are supported:

* `external_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `license_model` - (Optional) The Oracle license model that applies to the external database. 
* `enable_database_management` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Db System Database Managements Management
	* `update` - (Defaults to 20 minutes), when updating the External Db System Database Managements Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Db System Database Managements Management
