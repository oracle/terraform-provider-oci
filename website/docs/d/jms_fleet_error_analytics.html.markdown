---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_error_analytics"
sidebar_current: "docs-oci-datasource-jms-fleet_error_analytics"
description: |-
  Provides the list of Fleet Error Analytics in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_error_analytics
This data source provides the list of Fleet Error Analytics in Oracle Cloud Infrastructure Jms service.

Returns a high level summary of FleetErrors.

## Example Usage

```hcl
data "oci_jms_fleet_error_analytics" "test_fleet_error_analytics" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.fleet_error_analytic_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments. 


## Attributes Reference

The following attributes are exported:

* `fleet_error_aggregation_collection` - The list of fleet_error_aggregation_collection.

### FleetErrorAnalytic Reference

The following attributes are exported:

* `items` - A list of FleetErrorAggregationSummary.
	* `fleet_error_aggregations` - List of fleet error aggregations.
		* `fleet_error_analytic_count` - Number of FleetErrors encountered for the specific reason.
		* `reason` - Enum that uniquely identifies the fleet error.
	* `healthy_fleet_count` - Count of fleets with no problems.

