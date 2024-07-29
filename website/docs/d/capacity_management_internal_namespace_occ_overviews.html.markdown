---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_namespace_occ_overviews"
sidebar_current: "docs-oci-datasource-capacity_management-internal_namespace_occ_overviews"
description: |-
  Provides the list of Internal Namespace Occ Overviews in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_namespace_occ_overviews
This data source provides the list of Internal Namespace Occ Overviews in Oracle Cloud Infrastructure Capacity Management service.

Lists an overview of all resources in that namespace in a given time interval.


## Example Usage

```hcl
data "oci_capacity_management_internal_namespace_occ_overviews" "test_internal_namespace_occ_overviews" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.internal_namespace_occ_overview_namespace
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id

	#Optional
	from = var.internal_namespace_occ_overview_from
	to = var.internal_namespace_occ_overview_to
	workload_type = var.internal_namespace_occ_overview_workload_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `from` - (Optional) The month corresponding to this date would be considered as the starting point of the time period against which we would like to perform an aggregation.
* `namespace` - (Required) The namespace by which we would filter the list.
* `occ_customer_group_id` - (Required) The customer group ocid by which we would filter the list.
* `to` - (Optional) The month corresponding to this date would be considered as the ending point of the time period against which we would like to perform an aggregation.
* `workload_type` - (Optional) Workload type using the resources in an availability catalog can be filtered.


## Attributes Reference

The following attributes are exported:

* `occ_overview_collection` - The list of occ_overview_collection.

### InternalNamespaceOccOverview Reference

The following attributes are exported:

* `items` - An array of overview summary.
	* `capacity_requests_blob` - A raw json blob containing all the capacity requests corresponding to the resource name
	* `compartment_id` - The OCID of the compartment from which the api call is made. This will be used for authorizing the request.
	* `period_value` - The name of the month along with year for which this summary corresponds to.
	* `resource_name` - The name of the resource for which we have aggregated the value.
	* `total_available` - The quantity of the resource which is available at the end of the period of aggregationDetails model in consideration.
	* `total_cancelled` - The quantity of the resource which is cancelled by the customer. Once the capacity request was submitted, the customer can still cancel it. This field sums up those values.
	* `total_demanded` - The quantity of the resource which is demanded by customers via capacity requests against the resource name at the end of the time period in consideration for overview.
	* `total_rejected` - The quantity of the resource which is rejected by Oracle.
	* `total_supplied` - The quantity of the resource which is supplied by Oracle to the customer against the resource name at the end of the time period in consideration.
	* `total_unfulfilled` - The quantity of the resource which Oracle was unable to supply. For a given capacity request, Oracle sometimes cannot supply the entire value demanded by the customer. In such cases a partial value is provided, thereby leaving behind a portion of unfulfilled values. This field sums that up.
	* `unit` - The unit e.g SERVER in which the above values like totalAvailable, totalSupplied etc is measured.
	* `workload_type_breakdown_blob` - A raw json blob containing breakdown of totalAvailable, totalDemanded, totalSupplied, totalRejected, totalCancelled and totalUnfulfilled by workload types

