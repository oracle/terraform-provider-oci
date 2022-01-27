---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_alert_analytic"
sidebar_current: "docs-oci-datasource-data_safe-alert_analytic"
description: |-
  Provides details about a specific Alert Analytic in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_alert_analytic
This data source provides details about a specific Alert Analytic resource in Oracle Cloud Infrastructure Data Safe service.

Returns aggregation details of alerts.


## Example Usage

```hcl
data "oci_data_safe_alert_analytic" "test_alert_analytic" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.alert_analytic_access_level
	compartment_id_in_subtree = var.alert_analytic_compartment_id_in_subtree
	group_by = var.alert_analytic_group_by
	query_time_zone = var.alert_analytic_query_time_zone
	scim_query = var.alert_analytic_scim_query
	summary_field = var.alert_analytic_summary_field
	time_ended = var.alert_analytic_time_ended
	time_started = var.alert_analytic_time_started
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) A groupBy can only be used in combination with summaryField parameter. A groupBy value has to be a subset of the values mentioned in summaryField parameter. 
* `query_time_zone` - (Optional) Default time zone is UTC if no time zone provided. The date-time considerations of the resource will be in accordance with the specified time zone. 
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(timeCreated ge '2021-06-04T01-00-26') and (targetNames eq 'target_1') Supported fields: severity status alertType targetIds targetNames operationTime lifecycleState displayName timeCreated timeUpdated 
* `summary_field` - (Optional) Specifies a subset of summarized fields to be returned in the response.
* `time_ended` - (Optional) An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_started` - (Optional) An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 


## Attributes Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `count` - Total count of aggregated values.
	* `dimensions` - Details of aggregation dimension summarizing alerts.
		* `group_by` - GroupBy value used in aggregation.
	* `metric_name` - The name of the aggregation.
	* `time_ended` - The time at which the aggregation ended.
	* `time_started` - The time at which the aggregation started.

