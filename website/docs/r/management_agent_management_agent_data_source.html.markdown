---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_data_source"
sidebar_current: "docs-oci-resource-management_agent-management_agent_data_source"
description: |-
  Provides the Management Agent Data Source resource in Oracle Cloud Infrastructure Management Agent service
---

# oci_management_agent_management_agent_data_source
This resource provides the Management Agent Data Source resource in Oracle Cloud Infrastructure Management Agent service.

Datasource creation request to given Management Agent.


## Example Usage

```hcl
resource "oci_management_agent_management_agent_data_source" "test_management_agent_data_source" {
	#Required
	compartment_id = var.compartment_id
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	name = var.management_agent_data_source_name
	type = var.management_agent_data_source_type
	url = var.management_agent_data_source_url

	#Optional
	allow_metrics = var.management_agent_data_source_allow_metrics
	connection_timeout = var.management_agent_data_source_connection_timeout
	metric_dimensions {
		#Required
		name = var.management_agent_data_source_metric_dimensions_name
		value = var.management_agent_data_source_metric_dimensions_value
	}
	namespace = var.management_agent_data_source_namespace
	proxy_url = var.management_agent_data_source_proxy_url
	read_data_limit_in_kilobytes = var.management_agent_data_source_read_data_limit_in_kilobytes
	read_timeout = var.management_agent_data_source_read_timeout
	resource_group = var.management_agent_data_source_resource_group
	schedule_mins = var.management_agent_data_source_schedule_mins
}
```

## Argument Reference

The following arguments are supported:

* `allow_metrics` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Comma separated metric name list. The complete set of desired scraped metrics. Use this property to limit the set of metrics uploaded if required.
* `compartment_id` - (Required) Compartment owning this DataSource.
* `connection_timeout` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Number in milliseconds. The timeout for connecting to the Prometheus Exporter's endpoint.
* `management_agent_id` - (Required) Unique Management Agent identifier
* `metric_dimensions` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) The names of other user-supplied properties expressed as fixed values to be used as dimensions for every uploaded datapoint.
	* `name` - (Required) (Updatable) Name of the metric dimension
	* `value` - (Required) (Updatable) Value of the metric dimension
* `name` - (Required) Unique name of the DataSource.
* `namespace` - (Required) The Oracle Cloud Infrastructure monitoring namespace to which scraped metrics should be uploaded.
* `proxy_url` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) The url of the network proxy that provides access to the Prometheus Exporter's endpoint (url required property).
* `read_data_limit_in_kilobytes` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Number in kilobytes. The limit on the data being sent, not to exceed the agent's fixed limit of 400 (KB).
* `read_timeout` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Number in milliseconds. The timeout for reading the response from the Prometheus Exporter's endpoint.
* `resource_group` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Oracle Cloud Infrastructure monitoring resource group to assign the metric to.
* `schedule_mins` - (Optional when type=PROMETHEUS_EMITTER) (Updatable) Number in minutes. The scraping occurs at the specified interval.
* `type` - (Required) (Updatable) The type of the DataSource. Support types: PROMETHEUS_EMITTER
* `url` - (Required when type=PROMETHEUS_EMITTER) (Updatable) The url through which the Prometheus Exporter publishes its metrics. (http only)


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Agent Data Source
	* `update` - (Defaults to 20 minutes), when updating the Management Agent Data Source
	* `delete` - (Defaults to 20 minutes), when destroying the Management Agent Data Source


## Import

ManagementAgentDataSources can be imported using the `id`, e.g.

```
$ terraform import oci_management_agent_management_agent_data_source.test_management_agent_data_source "managementAgents/{managementAgentId}/dataSources/{key}" 
```

