---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_container_database_management"
sidebar_current: "docs-oci-resource-database-external_container_database_management"
description: |-
  Provides the External Container Database Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_external_container_database_management
This resource provides the External Container Database Management resource in Oracle Cloud Infrastructure Database service.

Enables Database Management Service for the external container database.
For more information about the Database Management Service, see
[Database Management Service](https://docs.cloud.oracle.com/iaas/Content/ExternalDatabase/Concepts/databasemanagementservice.htm).


## Example Usage

```hcl
resource "oci_database_external_container_database_management" "test_external_container_database_management" {
	#Required
	external_container_database_id = oci_database_external_container_database.test_external_container_database.id
	external_database_connector_id = oci_database_external_database_connector.test_external_database_connector.id
    #Optional
    license_mode = var.external_non_container_database_management_license_model
}
```

## Argument Reference

The following arguments are supported:

* `external_container_database_id` - (Required) The ExternalContainerDatabase [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `external_database_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `license_model` - (Optional) The Oracle license model that applies to the external database. Required only for enabling database management.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Container Database Management
	* `update` - (Defaults to 20 minutes), when updating the External Container Database Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Container Database Management


## Import

Import is not supported for this resource.

