---
subcategory: "Budget"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_budget_cost_anomaly_event_analytics"
sidebar_current: "docs-oci-datasource-budget-cost_anomaly_event_analytics"
description: |-
  Provides the list of Cost Anomaly Event Analytics in Oracle Cloud Infrastructure Budget service
---

# Data Source: oci_budget_cost_anomaly_event_analytics
This data source provides the list of Cost Anomaly Event Analytics in Oracle Cloud Infrastructure Budget service.

Gets a list of Cost Anomaly Events analytics summary - aggregated metrics for a given time period.


## Example Usage

```hcl
data "oci_budget_cost_anomaly_event_analytics" "test_cost_anomaly_event_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cost_anomaly_monitor_id = oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor.id
	cost_impact = var.cost_anomaly_event_analytic_cost_impact
	cost_impact_percentage = var.cost_anomaly_event_analytic_cost_impact_percentage
	name = var.cost_anomaly_event_analytic_name
	region = var.cost_anomaly_event_analytic_region
	target_tenant_id = oci_budget_target_tenant.test_target_tenant.id
	time_anomaly_event_end_date = var.cost_anomaly_event_analytic_time_anomaly_event_end_date
	time_anomaly_event_start_date = var.cost_anomaly_event_analytic_time_anomaly_event_start_date
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `cost_anomaly_monitor_id` - (Optional) The cost monitor ocid.
* `cost_impact` - (Optional) cost impact (absolute) of the anomaly event.
* `cost_impact_percentage` - (Optional) cost impact (percentage) of the anomaly event.
* `name` - (Optional) Unique, non-changeable resource name. 
* `region` - (Optional) region of the anomaly event.
* `target_tenant_id` - (Optional) The target tenantId ocid filter param.
* `time_anomaly_event_end_date` - (Optional) endDate for anomaly event date.
* `time_anomaly_event_start_date` - (Optional) startDate for anomaly event date.


## Attributes Reference

The following attributes are exported:

* `cost_anomaly_event_analytic_collection` - The list of cost_anomaly_event_analytic_collection.

### CostAnomalyEventAnalytic Reference

The following attributes are exported:

* `items` - The list of CostAnomalyEvent Analytic summary.
	* `average_cost_impact` - The average cost impact of the anomaly events in the given time period.
	* `average_cost_variance` - The average cost variance of the anomaly events in the given time period.
	* `cost_anomaly_event_analytic_count` - The number of cost anomaly events in the given time period.

