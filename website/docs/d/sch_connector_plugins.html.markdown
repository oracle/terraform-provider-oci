---
subcategory: "Service Connector Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_sch_connector_plugins"
sidebar_current: "docs-oci-datasource-sch-connector_plugins"
description: |-
  Provides the list of Connector Plugins in Oracle Cloud Infrastructure Service Connector Hub service
---

# Data Source: oci_sch_connector_plugins
This data source provides the list of Connector Plugins in Oracle Cloud Infrastructure Service Connector Hub service.

Lists connector plugins according to the specified filter.


## Example Usage

```hcl
data "oci_sch_connector_plugins" "test_connector_plugins" {

	#Optional
	display_name = var.connector_plugin_display_name
	name = var.connector_plugin_name
	state = var.connector_plugin_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.  Example: `example_service_connector` 
* `name` - (Optional) A filter to return only resources that match the given connector plugin name ignoring case.  Example: `QueueSource` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state.  Example: `ACTIVE` 


## Attributes Reference

The following attributes are exported:

* `connector_plugin_collection` - The list of connector_plugin_collection.

### ConnectorPlugin Reference

The following attributes are exported:

* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `estimated_throughput` - The estimated throughput range (LOW, MEDIUM, HIGH). 
* `kind` - The plugin type discriminator. 
* `max_retention` - The estimated maximum period of time the data will be kept at the source. The duration is specified as a string in ISO 8601 format (P1D for one day or P30D for thrity days). 
* `name` - The service to be called by the connector plugin. Example: `QueueSource` 
* `schema` - Gets the specified connector plugin configuration information in OpenAPI specification format. 
* `state` - The current state of the service connector. 
* `time_created` - The date and time when this plugin became available. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2023-09-09T21:10:29.600Z` 

