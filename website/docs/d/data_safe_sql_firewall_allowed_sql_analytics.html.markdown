---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_allowed_sql_analytics"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_allowed_sql_analytics"
description: |-
  Provides the list of Sql Firewall Allowed Sql Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_allowed_sql_analytics
This data source provides the list of Sql Firewall Allowed Sql Analytics in Oracle Cloud Infrastructure Data Safe service.

Returns the aggregation details of all SQL Firewall allowed SQL statements.

The ListSqlFirewallAllowedSqlAnalytics operation returns the aggregates of the SQL Firewall allowed SQL statements in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSqlFirewallAllowedSqlAnalytics on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_sql_firewall_allowed_sql_analytics" "test_sql_firewall_allowed_sql_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_firewall_allowed_sql_analytic_access_level
	compartment_id_in_subtree = var.sql_firewall_allowed_sql_analytic_compartment_id_in_subtree
	group_by = var.sql_firewall_allowed_sql_analytic_group_by
	scim_query = var.sql_firewall_allowed_sql_analytic_scim_query
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) The group by parameter to summarize the allowed SQL aggregation.
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(currentUser eq 'SCOTT') and (topLevel eq 'YES') 


## Attributes Reference

The following attributes are exported:

* `sql_firewall_allowed_sql_analytics_collection` - The list of sql_firewall_allowed_sql_analytics_collection.

### SqlFirewallAllowedSqlAnalytic Reference

The following attributes are exported:

* `items` - The aggregated data point items.
	* `dimensions` - The dimensions available for SQL Firewall allow SQL analytics.
		* `db_user_name` - The database user name.
		* `sql_firewall_policy_id` - The OCID of the SQL Firewall policy corresponding to the SQL Firewall allowed SQL.
		* `sql_level` - Specifies the level of SQL included for this SQL Firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
		* `state` - The current state of the SQL Firewall allowed SQL.
	* `sql_firewall_allowed_sql_analytic_count` - The total count of the aggregated metric.

