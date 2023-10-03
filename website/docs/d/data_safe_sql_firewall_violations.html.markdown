---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_violations"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_violations"
description: |-
  Provides the list of Sql Firewall Violations in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_violations
This data source provides the list of Sql Firewall Violations in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all the SQL firewall violations captured by the firewall.


## Example Usage

```hcl
data "oci_data_safe_sql_firewall_violations" "test_sql_firewall_violations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sql_firewall_violation_access_level
	compartment_id_in_subtree = var.sql_firewall_violation_compartment_id_in_subtree
	scim_query = var.sql_firewall_violation_scim_query
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(operationTime ge '2021-06-04T01-00-26') and (violationAction eq 'BLOCKED') 


## Attributes Reference

The following attributes are exported:

* `sql_firewall_violations_collection` - The list of sql_firewall_violations_collection.

### SqlFirewallViolation Reference

The following attributes are exported:

* `items` - Array of SQL violation summary.
	* `client_ip` - The IP address of the host machine from which the session was generated.
	* `client_os_user_name` - The name of the operating system user for the database session.
	* `client_program` - The application from which the SQL violation was generated. Examples include SQL Plus or SQL Developer.
	* `compartment_id` - The OCID of the compartment containing the SQL violation.
	* `current_db_user_name` - The name of the user that SQL was executed as.
	* `db_user_name` - The name of the database user.
	* `id` - The OCID of the SQL violation.
	* `operation` - The name of the action executed by the user on the target database. For example, ALTER, CREATE, DROP.
	* `operation_time` - The time of the SQL violation occurrence in the target database.
	* `sql_accessed_objects` - The objects accessed by the SQL.
	* `sql_level` - Specifies the level of SQL for this violation. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
	* `sql_text` - The SQL text caught by the firewall.
	* `target_id` - The OCID of the target database.
	* `target_name` - The name of the target database.
	* `time_collected` - The timestamp when this SQL violation was collected from the target database by Data Safe.
	* `violation_action` - The action taken for this SQL violation.
	* `violation_cause` - Indicates whether SQL or context violation.

