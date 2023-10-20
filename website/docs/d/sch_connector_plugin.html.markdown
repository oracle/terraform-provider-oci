---
subcategory: "Service Connector Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_sch_connector_plugin"
sidebar_current: "docs-oci-datasource-sch-connector_plugin"
description: |-
  Provides details about a specific Connector Plugin in Oracle Cloud Infrastructure Service Connector Hub service
---

# Data Source: oci_sch_connector_plugin
This data source provides details about a specific Connector Plugin resource in Oracle Cloud Infrastructure Service Connector Hub service.

Gets the specified connector plugin configuration information.


## Example Usage

```hcl
data "oci_sch_connector_plugin" "test_connector_plugin" {
	#Required
	connector_plugin_name = oci_sch_connector_plugin.test_connector_plugin.name
}
```

## Argument Reference

The following arguments are supported:

* `connector_plugin_name` - (Required) The name of the connector plugin. This name indicates the service to be called by the connector plugin. For example, `QueueSource` indicates the Queue service. 


## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `estimated_throughput` - The estimated throughput range (LOW, MEDIUM, HIGH). 
* `kind` - The plugin type discriminator. 
* `max_retention` - The estimated maximum period of time the data will be kept at the source. The duration is specified as a string in ISO 8601 format (P1D for one day or P30D for thrity days). 
* `name` - The service to be called by the connector plugin. Example: `QueueSource` 
* `schema` - Gets the specified connector plugin configuration information in OpenAPI specification format. 
* `state` - The current state of the service connector. 
* `time_created` - The date and time when this plugin became available. Format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2023-09-09T21:10:29.600Z` 

