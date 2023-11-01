---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_allowed_sqls"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_allowed_sqls"
description: |-
  Provides the list of Sql Firewall Allowed Sqls in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_allowed_sqls
This data source provides the list of Sql Firewall Allowed Sqls in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all SQL firewall allowed SQL statements.

The ListSqlFirewallAllowedSqls operation returns only the SQL firewall allowed SQL statements in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSqlFirewallPolicies on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_sql_firewall_allowed_sqls" "test_sql_firewall_allowed_sqls" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_firewall_allowed_sql_access_level
	compartment_id_in_subtree = var.sql_firewall_allowed_sql_compartment_id_in_subtree
	scim_query = var.sql_firewall_allowed_sql_scim_query
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(currentUser eq 'SCOTT') and (topLevel eq 'YES') 


## Attributes Reference

The following attributes are exported:

* `sql_firewall_allowed_sql_collection` - The list of sql_firewall_allowed_sql_collection.

### SqlFirewallAllowedSql Reference

The following attributes are exported:

* `items` - Array of SQL firewall allowed SQL statements.
	* `compartment_id` - The OCID of the compartment containing the SQL firewall allowed SQL.
	* `current_user` - The name of the user that SQL was executed as.
	* `db_user_name` - The database user name.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
	* `description` - The description of the SQL firewall allowed SQL.
	* `display_name` - The display name of the SQL firewall allowed SQL.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the SQL firewall allowed SQL.
	* `sql_accessed_objects` - The objects accessed by the SQL.
	* `sql_firewall_policy_id` - The OCID of the SQL firewall policy corresponding to the SQL firewall allowed SQL.
	* `sql_level` - Specifies the level of SQL included for this SQL firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
	* `sql_text` - The SQL text of the SQL firewall allowed SQL.
	* `state` - The current state of the SQL firewall allowed SQL.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_collected` - The time the the SQL firewall allowed SQL was collected from the target database, in the format defined by RFC3339.
	* `time_updated` - The last date and time the SQL firewall allowed SQL was updated, in the format defined by RFC3339.
	* `version` - Version of the associated SQL firewall policy. This identifies whether the allowed SQLs were added in the same batch or not.

