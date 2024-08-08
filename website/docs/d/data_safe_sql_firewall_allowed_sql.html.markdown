---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_allowed_sql"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_allowed_sql"
description: |-
  Provides details about a specific Sql Firewall Allowed Sql in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_allowed_sql
This data source provides details about a specific Sql Firewall Allowed Sql resource in Oracle Cloud Infrastructure Data Safe service.

Gets a SQL firewall allowed SQL by identifier.

## Example Usage

```hcl
data "oci_data_safe_sql_firewall_allowed_sql" "test_sql_firewall_allowed_sql" {
	#Required
	sql_firewall_allowed_sql_id = oci_data_safe_sql_firewall_allowed_sql.test_sql_firewall_allowed_sql.id
}
```

## Argument Reference

The following arguments are supported:

* `sql_firewall_allowed_sql_id` - (Required) The OCID of the sqlFirewallAllowedSql resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the SQL Firewall allowed SQL.
* `current_user` - The name of the user that SQL was executed as.
* `db_user_name` - The database user name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the SQL Firewall allowed SQL.
* `display_name` - The display name of the SQL Firewall allowed SQL.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the SQL Firewall allowed SQL.
* `sql_accessed_objects` - The objects accessed by the SQL.
* `sql_firewall_policy_id` - The OCID of the SQL Firewall policy corresponding to the SQL Firewall allowed SQL.
* `sql_level` - Specifies the level of SQL included for this SQL Firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
* `sql_text` - The SQL text of the SQL Firewall allowed SQL.
* `state` - The current state of the SQL Firewall allowed SQL.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_collected` - The time the the SQL Firewall allowed SQL was collected from the target database, in the format defined by RFC3339.
* `time_updated` - The last date and time the SQL Firewall allowed SQL was updated, in the format defined by RFC3339.
* `version` - Version of the associated SQL Firewall policy. This identifies whether the allowed SQLs were added in the same batch or not.

