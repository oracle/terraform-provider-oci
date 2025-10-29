---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_identity"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_identity"
description: |-
  Provides details about a specific Database Tools Identity in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_identity
This data source provides details about a specific Database Tools Identity resource in Oracle Cloud Infrastructure Database Tools service.

Gets details of the specified Database Tools identity.

## Example Usage

```hcl
data "oci_database_tools_database_tools_identity" "test_database_tools_identity" {
	#Required
	database_tools_identity_id = oci_database_tools_database_tools_identity.test_database_tools_identity.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_identity_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools identity.


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

