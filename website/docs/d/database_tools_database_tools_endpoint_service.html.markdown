---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_endpoint_service"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_endpoint_service"
description: |-
  Provides details about a specific Database Tools Endpoint Service in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_endpoint_service
This data source provides details about a specific Database Tools Endpoint Service resource in Oracle Cloud Infrastructure Database Tools service.

Gets a DatabaseToolsEndpointService by identifier

## Example Usage

```hcl
data "oci_database_tools_database_tools_endpoint_service" "test_database_tools_endpoint_service" {
	#Required
	database_tools_endpoint_service_id = oci_database_tools_database_tools_endpoint_service.test_database_tools_endpoint_service.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_endpoint_service_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a DatabaseToolsEndpointService.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the containing Compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the DatabaseToolsEndpointService.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DatabaseToolsEndpointService.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - A unique, non-changeable resource name.
* `state` - The current state of the DatabaseToolsEndpointService.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the DatabaseToolsEndpointService was created. An RFC3339 formatted datetime string
* `time_updated` - The time the DatabaseToolsEndpointService was updated. An RFC3339 formatted datetime string

