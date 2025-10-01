---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_databases_estimate_cost_savings"
sidebar_current: "docs-oci-datasource-database-autonomous_databases_estimate_cost_savings"
description: |-
  Provides the list of Autonomous Databases Estimate Cost Savings in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_databases_estimate_cost_savings
This data source provides the list of Autonomous Databases Estimate Cost Savings in Oracle Cloud Infrastructure Database service.

Gets the estimate cost savings of the Autonomous Database.

## Example Usage

```hcl
data "oci_database_autonomous_databases_estimate_cost_savings" "test_autonomous_databases_estimate_cost_savings" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	is_cpu_autoscale = var.autonomous_databases_estimate_cost_saving_is_cpu_autoscale
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_cpu_autoscale` - (Required) If provided as true, cost estimate with cpu autoscaling.


## Attributes Reference

The following attributes are exported:

* `estimate_cost_savings_summary_collection` - The list of estimate_cost_savings_summary_collection.

### AutonomousDatabasesEstimateCostSaving Reference

The following attributes are exported:

* `items` - List of estimate cost saving summary.
	* `cost_savings_with_elastic_pool` - Estimated cost savings in percentage with elastic pool utilization.
	* `estimated_usage_without_elastic_pool` - CPU cost for a given time period under regular billing plan, in ECPU hours.
	* `is_cpu_autoscale` - Indicates if CPU autoscaling is applied.
	* `time_ended` - The epoch time at which cost aggregation ends.
	* `time_started` - The epoch time at which cost aggregation starts.
	* `usage_with_elastic_pool` - CPU cost for a given time period under elastic pool billing plan, in ECPU hours.

