---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_pluggable_database_management"
sidebar_current: "docs-oci-resource-database-external_pluggable_database_management"
description: |-
  Provides the External Pluggable Database Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_external_pluggable_database_management
This resource provides the External Pluggable Database Management resource in Oracle Cloud Infrastructure Database service.

Enable Database Management Service for the external pluggable database.
For more information about the Database Management Service, see
[Database Management Service](https://docs.cloud.oracle.com/iaas/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).


## Example Usage

```hcl
resource "oci_database_external_pluggable_database_management" "test_external_pluggable_database_management" {
	#Required
	external_database_connector_id = oci_database_external_database_connector.test_external_database_connector.id
	external_pluggable_database_id = oci_database_external_pluggable_database.test_external_pluggable_database.id
}
```

## Argument Reference

The following arguments are supported:

* `external_database_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `external_pluggable_database_id` - (Required) The ExternalPluggableDatabaseId [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Pluggable Database Management
	* `update` - (Defaults to 20 minutes), when updating the External Pluggable Database Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Pluggable Database Management


## Import

Import is not supported for this resource.

