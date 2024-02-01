---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_data_source"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_data_source"
description: |-
  Provides details about a specific Management Agent Data Source in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_data_source
This data source provides details about a specific Management Agent Data Source resource in Oracle Cloud Infrastructure Management Agent service.

Get Datasource details for given Id and given Management Agent.


## Example Usage

```hcl
data "oci_management_agent_management_agent_data_source" "test_management_agent_data_source" {
	#Required
	data_source_key = var.management_agent_data_source_data_source_key
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
}
```

## Argument Reference

The following arguments are supported:

* `data_source_key` - (Required) Data source type and name identifier.
* `management_agent_id` - (Required) Unique Management Agent identifier


## Attributes Reference

The following attributes are exported:

* `allow_metrics` - Comma separated metric name list. The complete set of desired scraped metrics. Use this property to limit the set of metrics uploaded if required.
* `compartment_id` - Compartment owning this DataSource.
* `connection_timeout` - Number in milliseconds. The timeout for connecting to the Prometheus Exporter's endpoint.
* `is_daemon_set` - If the Kubernetes cluster type is Daemon set then this will be set to true.
* `key` - Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
* `metric_dimensions` - The names of other user-supplied properties expressed as fixed values to be used as dimensions for every uploaded datapoint.
	* `name` - Name of the metric dimension
	* `value` - Value of the metric dimension
* `name` - Unique name of the DataSource.
* `namespace` - The Oracle Cloud Infrastructure monitoring namespace to which scraped metrics should be uploaded.
* `proxy_url` - The url of the network proxy that provides access to the Prometheus Exporter's endpoint (url required property).
* `read_data_limit` - Number in kilobytes. The limit on the data being sent, not to exceed the agent's fixed limit of 400 (KB).
* `read_timeout` - Number in milliseconds. The timeout for reading the response from the Prometheus Exporter's endpoint.
* `resource_group` - Oracle Cloud Infrastructure monitoring resource group to assign the metric to.
* `schedule_mins` - Number in minutes. The scraping occurs at the specified interval.
* `state` - State of the DataSource.
* `time_created` - The time the DataSource was created. An RFC3339 formatted datetime string
* `time_updated` - The time the DataSource data was last received. An RFC3339 formatted datetime string
* `type` - The type of the DataSource.
* `url` - The url through which the Prometheus Exporter publishes its metrics. (http only)

