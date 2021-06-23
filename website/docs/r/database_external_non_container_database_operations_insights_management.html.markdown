---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_non_container_database_operations_insights_management"
sidebar_current: "docs-oci-resource-database-external_non_container_database_operations_insights_management"
description: |-
  Provides the External Non Container Database Operations Insights Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_external_non_container_database_operations_insights_management
This resource provides the External Non Container Database Operations Insights Management resource in Oracle Cloud Infrastructure Database service.

Enable Operations Insights for the external non-container database.
When deleting this resource block , we call disable if it was in enabled state .

## Example Usage

```hcl
resource "oci_database_external_non_container_database_operations_insights_management" "test_external_non_container_database_operations_insights_management" {
	#Required
	external_database_connector_id = oci_database_external_database_connector.test_external_database_connector.id
	external_non_container_database_id = oci_database_external_non_container_database.test_external_non_container_database.id
    enable_operations_insights = true
}
```

## Argument Reference

The following arguments are supported:

* `external_database_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `external_non_container_database_id` - (Required) The external non-container database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `enable_operations_insights`  -  (Required) (Updatable) Enabling OPSI on External non-container Databases . Requires boolean value "true" or "false".

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Non Container Database Operations Insights Management
	* `update` - (Defaults to 20 minutes), when updating the External Non Container Database Operations Insights Management
	* `delete` - (Defaults to 20 minutes), when destroying the External Non Container Database Operations Insights Management


## Import

Import is not supported for this resource.

