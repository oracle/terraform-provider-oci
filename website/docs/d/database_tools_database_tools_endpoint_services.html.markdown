---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_endpoint_services"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_endpoint_services"
description: |-
  Provides the list of Database Tools Endpoint Services in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_endpoint_services
This data source provides the list of Database Tools Endpoint Services in Oracle Cloud Infrastructure Database Tools service.

Returns a list of Database Tools endpoint services.


## Example Usage

```hcl
data "oci_database_tools_database_tools_endpoint_services" "test_database_tools_endpoint_services" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.database_tools_endpoint_service_display_name
	name = var.database_tools_endpoint_service_name
	state = var.database_tools_endpoint_service_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire specified display name.
* `name` - (Optional) A filter to return only resources that match the entire specified name.
* `state` - (Optional) A filter to return only resources their `lifecycleState` matches the specified `lifecycleState`.


## Attributes Reference

The following attributes are exported:

* `database_tools_endpoint_service_collection` - The list of database_tools_endpoint_service_collection.

### DatabaseToolsEndpointService Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools Endpoint Service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the Database Tools Endpoint Service.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools Endpoint Service.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - A unique, non-changeable resource name.
* `state` - The current state of the Database Tools Endpoint Service.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools Endpoint Service was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Database Tools Endpoint Service was updated. An RFC3339 formatted datetime string

