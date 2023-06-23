---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system_stack_monitorings_management"
sidebar_current: "docs-oci-resource-database_management-external_db_system_stack_monitorings_management"
description: |-
  Provides the External Db System Stack Monitorings Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_db_system_stack_monitorings_management
This resource provides the External Db System Stack Monitorings Management resource in Oracle Cloud Infrastructure Database Management service.

Enables Stack Monitoring for all the components of the specified
external DB system (except databases).


## Example Usage

```hcl
resource "oci_database_management_external_db_system_stack_monitorings_management" "test_external_db_system_stack_monitorings_management" {
	#Required
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
	enable_stack_monitoring = var.enable_stack_monitoring

	#Optional
	is_enabled = var.external_db_system_stack_monitorings_management_is_enabled
	metadata = var.external_db_system_stack_monitorings_management_metadata
}
```

## Argument Reference

The following arguments are supported:

* `external_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
* `is_enabled` - (Optional) The status of the associated service.
* `metadata` - (Optional) The associated service-specific inputs in JSON string format, which Database Management can identify.
* `enable_stack_monitoring` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Db System Stack Monitorings Management
	* `update` - (Defaults to 20 minutes), when updating the External Db System Stack Monitorings Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Db System Stack Monitorings Management
