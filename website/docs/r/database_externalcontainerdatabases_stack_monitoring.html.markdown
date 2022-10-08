---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_externalcontainerdatabases_stack_monitoring"
sidebar_current: "docs-oci-resource-database-externalcontainerdatabases_stack_monitoring"
description: |-
  Provides the Externalcontainerdatabases Stack Monitoring resource in Oracle Cloud Infrastructure Database service
---

# oci_database_externalcontainerdatabases_stack_monitoring
This resource provides the Externalcontainerdatabases Stack Monitoring resource in Oracle Cloud Infrastructure Database service.

Enable Stack Monitoring for the external container database.


## Example Usage

```hcl
resource "oci_database_externalcontainerdatabases_stack_monitoring" "test_externalcontainerdatabases_stack_monitoring" {
	#Required
	external_container_database_id = oci_database_external_container_database.test_external_container_database.id
	external_database_connector_id = oci_database_external_database_connector.test_external_database_connector.id
	enable_stack_monitoring = true
}
```

## Argument Reference

The following arguments are supported:

* `external_container_database_id` - (Required) The ExternalContainerDatabase [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `external_database_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails).
* `enable_stack_monitoring`  -  (Required) (Updatable) Enabling Stack Monitoring on External Container Databases . Requires boolean value "true" or "false".


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Externalcontainerdatabases Stack Monitoring
* `update` - (Defaults to 20 minutes), when updating the Externalcontainerdatabases Stack Monitoring
* `delete` - (Defaults to 20 minutes), when destroying the Externalcontainerdatabases Stack Monitoring


## Import

Import is not supported for this resource.

