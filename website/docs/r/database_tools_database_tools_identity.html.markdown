---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_identity"
sidebar_current: "docs-oci-resource-database_tools-database_tools_identity"
description: |-
  Provides the Database Tools Identity resource in Oracle Cloud Infrastructure Database Tools service
---

# oci_database_tools_database_tools_identity
This resource provides the Database Tools Identity resource in Oracle Cloud Infrastructure Database Tools service.

Creates a new Database Tools identity.


## Example Usage

```hcl
resource "oci_database_tools_database_tools_identity" "test_database_tools_identity" {
	#Required
	compartment_id = var.compartment_id
	credential_key = var.database_tools_identity_credential_key
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	display_name = var.database_tools_identity_display_name
	type = var.database_tools_identity_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		type = var.database_tools_identity_locks_type

		#Optional
		message = var.database_tools_identity_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.database_tools_identity_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools identity.
* `credential_key` - (Required) The name of the credential object created in the Oracle Database.
* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `type` - (Required) (Updatable) The Database Tools identity type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools identity.
* `credential_key` - The name of the credential object created in the Oracle Database.
* `database_tools_connection_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools identity.
* `lifecycle_details` - A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `state` - The current state of the Database Tools identity.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools identity was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Database Tools identity was updated. An RFC3339 formatted datetime string.
* `type` - The Database Tools identity type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Identity
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Identity
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Identity


## Import

DatabaseToolsIdentities can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_database_tools_identity.test_database_tools_identity "id"
```

