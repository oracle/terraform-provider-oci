---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_exadata_infrastructure_fleet_metric"
sidebar_current: "docs-oci-datasource-database_management-exadata_infrastructure_fleet_metric"
description: |-
  Provides details about a specific Exadata Infrastructure Fleet Metric in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_exadata_infrastructure_fleet_metric
This data source provides details about a specific Exadata Infrastructure Fleet Metric resource in Oracle Cloud Infrastructure Database Management service.

Gets the health metrics for a fleet of Exadata infrastructure in a compartment. 
The CompartmentId query parameters must be provided to retrieve the health metrics.


## Example Usage

```hcl
data "oci_database_management_exadata_infrastructure_fleet_metric" "test_exadata_infrastructure_fleet_metric" {
	#Required
	compare_baseline_time = var.exadata_infrastructure_fleet_metric_compare_baseline_time
	compare_target_time = var.exadata_infrastructure_fleet_metric_compare_target_time
	compartment_id = var.compartment_id

	#Optional
	compare_type = var.exadata_infrastructure_fleet_metric_compare_type
	filter_by_exadata_infrastructure_deployment_type = var.exadata_infrastructure_fleet_metric_filter_by_exadata_infrastructure_deployment_type
	filter_by_exadata_infrastructure_lifecycle_state = var.exadata_infrastructure_fleet_metric_filter_by_exadata_infrastructure_lifecycle_state
}
```

## Argument Reference

The following arguments are supported:

* `compare_baseline_time` - (Required) The baseline time for metrics comparison.
* `compare_target_time` - (Required) The target time for metrics comparison.
* `compare_type` - (Optional) The time window used for metrics comparison.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `filter_by_exadata_infrastructure_deployment_type` - (Optional) The filter used to filter the Exadata infrastructures in the fleet by a specific deployment type.
* `filter_by_exadata_infrastructure_lifecycle_state` - (Optional) The filter used to filter the Exadata infrastructure in the fleet by its lifecycle state. If the parameter is not provided, Exdata infrastructures in any state are returned. 


## Attributes Reference

The following attributes are exported:

* `compare_baseline_time` - The baseline date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". This is the date and time against which percentage change is calculated. 
* `compare_target_time` - The target date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". All the metrics are returned for the target date and time and the percentage change is calculated against the baseline date and time. 
* `compare_type` - The time window used for metrics comparison.
* `exadata_infrastructure_fleet_summary` - A summary of the inventory count grouped by Exadata infrastructure deployment type, and the metrics that describe the aggregated usage of CPU, storage, and so on of all Exadata infrastructures in the fleet. 
	* `aggregated_metrics` - A list of Exadata infrastructures present in the fleet and their usage metrics.
		* `baseline_value` - The metric aggregated value at the baseline date and time.
		* `dimensions` - The unique dimension key and values of the baseline metric.
			* `dimension_name` - The name of the dimension.
			* `dimension_value` - The value of the dimension.
		* `metric_name` - The name of the metric.
		* `percentage_change` - The percentage change in the metric aggregated value compared to the baseline value.
		* `target_value` - The metric aggregated value at the target date and time.
		* `unit` - The unit of the value.
	* `inventory` - A list of the Exadata infrastructures in the fleet. 
		* `deployment_type` - The infrastructure deployment type.
		* `inventory_count` - The number of Exadata infrastructures in the fleet.
		* `rack_size` - The size of the Exadata infrastructure.
* `fleet_exadata_infrastructures` - A list of the Exadata infrastructures present in the fleet and their usage metrics.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the Exadata infrastructure resides.
	* `deployment_type` - The Exadata infrastructure deployment type.
	* `infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	* `infrastructure_name` - The display name of the Exadata infrastructure.
	* `metrics` - A list of the health metrics like CPU, Storage, and Memory.
		* `baseline_value` - The baseline value of the metric.
		* `dimensions` - The dimensions of the metric.
			* `dimension_name` - The name of the dimension.
			* `dimension_value` - The value of the dimension.
		* `metric_name` - The name of the metric.
		* `percentage_change` - The percentage change in the metric aggregated value compared to the baseline value.
		* `target_value` - The target value of the metric.
		* `timestamp` - The data point date and time in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
		* `unit` - The unit of the value.
	* `number_of_db_systems` - The number of Database Systems created on the Exadata infrastructure.
	* `rack_size` - The size of the Exadata infrastructure.
	* `state` - The lifecycle state of the Exadata infrastructure.
	* `storage_server_count` - The number of storage server for the Exadata infrastructure.

