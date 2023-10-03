---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_violation_analytics"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_violation_analytics"
description: |-
  Provides the list of Sql Firewall Violation Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_violation_analytics
This data source provides the list of Sql Firewall Violation Analytics in Oracle Cloud Infrastructure Data Safe service.

Returns the aggregation details of the SQL firewall violations.


## Example Usage

```hcl
data "oci_data_safe_sql_firewall_violation_analytics" "test_sql_firewall_violation_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_firewall_violation_analytic_access_level
	compartment_id_in_subtree = var.sql_firewall_violation_analytic_compartment_id_in_subtree
	group_by = var.sql_firewall_violation_analytic_group_by
	query_time_zone = var.sql_firewall_violation_analytic_query_time_zone
	scim_query = var.sql_firewall_violation_analytic_scim_query
	summary_field = var.sql_firewall_violation_analytic_summary_field
	time_ended = var.sql_firewall_violation_analytic_time_ended
	time_started = var.sql_firewall_violation_analytic_time_started
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

	**Example:** query=(operationTime ge '2021-06-04T01-00-26') and (violationAction eq 'BLOCKED') 
* `summary_field` - (Optional) Specifies a subset of summarized fields to be returned in the response.
* `time_ended` - (Optional) An optional filter to return audit events whose creation time in the database is less than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_started` - (Optional) An optional filter to return audit events whose creation time in the database is greater than and equal to the date-time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 


## Attributes Reference

The following attributes are exported:

* `sql_firewall_violation_analytics_collection` - The list of sql_firewall_violation_analytics_collection.

### SqlFirewallViolationAnalytic Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `dimensions` - The details of the aggregation dimensions used for summarizing SQL violations.
		* `client_ip` - The IP address of the host from which the session was spawned.
		* `client_os_user_name` - The name of the operating system user for the database session.
		* `client_program` - The application from which the SQL violation was generated. Examples SQL Plus or SQL Developer.
		* `db_user_name` - The name of the database user.
		* `operation` - The name of the action executed by the user on the target database, for example, ALTER, CREATE, DROP.
		* `operation_time` - The time of the SQL violation occurrence in the target database.
		* `sql_level` - Specifies the level of SQL included for this SQL firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
		* `target_id` - The OCID of the target database.
		* `target_name` - The name of the target database.
		* `violation_action` - The action taken for this SQL violation.
		* `violation_cause` - Indicates whether SQL or context violation.
	* `display_name` - The display name of aggregation field.
	* `metric_name` - The name of the aggregation.
	* `sql_firewall_violation_analytic_count` - Total count of aggregated value.
	* `time_ended` - The time at which the aggregation ended.
	* `time_started` - The time at which the aggregation started.

