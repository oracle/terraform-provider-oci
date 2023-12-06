---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_metric_data"
sidebar_current: "docs-oci-datasource-monitoring-metric_data"
description: |-
  Provides the list of Metric Data in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_metric_data
This data source provides the list of Metric Data in Oracle Cloud Infrastructure Monitoring service.

Returns aggregated data that match the criteria specified in the request. Compartment OCID required.
For more information, see
[Querying Metric Data](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-landing.htm)
and
[Creating a Query](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric.htm).
For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

Transactions Per Second (TPS) per-tenancy limit for this operation: 10.


## Example Usage

```hcl
data "oci_monitoring_metric_data" "test_metric_data" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.metric_data_namespace
	query = var.metric_data_query

	#Optional
	compartment_id_in_subtree = var.metric_data_compartment_id_in_subtree
	end_time = var.metric_data_end_time
	resolution = var.metric_data_resolution
	resource_group = var.metric_data_resource_group
	start_time = var.metric_data_start_time
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment.  Example: `ocid1.compartment.oc1..exampleuniqueID` 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `end_time` - (Optional) The end of the time range to use when searching for metric data points. Format is defined by RFC3339. The response excludes metric data points for the endTime. Default value: the timestamp representing when the call was sent.  Example: `2019-02-01T02:02:29.600Z` 
* `namespace` - (Required) The source service or application to use when searching for metric data points to aggregate.  Example: `oci_computeagent` 
* `query` - (Required) The Monitoring Query Language (MQL) expression to use when searching for metric data points to aggregate. The query must specify a metric, statistic, and interval. Supported values for interval depend on the specified time range. More interval values are supported for smaller time ranges. You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.

	Construct your query to avoid exceeding limits on returned data. See [MetricData Reference](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/20180401/MetricData).

	For details about Monitoring Query Language (MQL), see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm). For available dimensions, review the metric definition for the supported service. See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

	Example: `CpuUtilization[1m].sum()` 
* `resolution` - (Optional) The time between calculated aggregation windows. Use with the query interval to vary the frequency for returning aggregated data points. For example, use a query interval of 5 minutes with a resolution of 1 minute to retrieve five-minute aggregations at a one-minute frequency. The resolution must be equal or less than the interval in the query. The default resolution is 1m (one minute). Supported values: `1m`-`60m`, `1h`-`24h`, `1d`.  Example: `5m` 
* `resource_group` - (Optional) Resource group that you want to match. A null value returns only metric data that has no resource groups. The specified resource group must exist in the definition of the posted metric. Only one resource group can be applied per metric. A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).  Example: `frontend-fleet` 
* `start_time` - (Optional) The beginning of the time range to use when searching for metric data points. Format is defined by RFC3339. The response includes metric data points for the startTime. Default value: the timestamp 3 hours before the call was sent.  Example: `2019-02-01T01:02:29.600Z` 


## Attributes Reference

The following attributes are exported:

* `metric_data` - The list of metric_data.

### MetricData Reference

The following attributes are exported:

* `aggregated_datapoints` - The list of timestamp-value pairs returned for the specified request. Metric values are rolled up to the start time specified in the request. For important limits information related to data points, see MetricData Reference at the top of this page. 
	* `timestamp` - The date and time associated with the value of this data point. Format defined by RFC3339.  Example: `2019-02-01T01:02:29.600Z` 
	* `value` - Numeric value of the metric.  Example: `10.4` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources that the aggregated data was returned from. 
* `dimensions` - Qualifiers provided in the definition of the returned metric. Available dimensions vary by metric namespace. Each dimension takes the form of a key-value pair.  Example: `"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"` 
* `metadata` - The references provided in a metric definition to indicate extra information about the metric.  Example: `"unit": "bytes"` 
* `name` - The name of the metric.  Example: `CpuUtilization` 
* `namespace` - The reference provided in a metric definition to indicate the source service or application that emitted the metric.  Example: `oci_computeagent` 
* `resolution` - The time between calculated aggregation windows. Use with the query interval to vary the frequency for returning aggregated data points. For example, use a query interval of 5 minutes with a resolution of 1 minute to retrieve five-minute aggregations at a one-minute frequency. The resolution must be equal or less than the interval in the query. The default resolution is 1m (one minute). Supported values: `1m`-`60m`, `1h`-`24h`, `1d`.  Example: `5m` 
* `resource_group` - Resource group provided with the posted metric. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric. A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).  Example: `frontend-fleet` 

