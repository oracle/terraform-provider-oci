---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_report_database_view_access_entries"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_report_database_view_access_entries"
description: |-
  Provides the list of Security Policy Report Database View Access Entries in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_report_database_view_access_entries
This data source provides the list of Security Policy Report Database View Access Entries in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all database view access entries in Data Safe.

The ListDatabaseViewAccessEntries operation returns only the database view access objects for the specified security policy report.


## Example Usage

```hcl
data "oci_data_safe_security_policy_report_database_view_access_entries" "test_security_policy_report_database_view_access_entries" {
	#Required
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id

	#Optional
	scim_query = var.security_policy_report_database_view_access_entry_scim_query
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `scim_query` - (Optional) The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.)

	**Example:** query=(accessType eq 'SELECT') and (grantee eq 'ADMIN') 
* `security_policy_report_id` - (Required) The OCID of the security policy report resource.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `database_view_access_entry_collection` - The list of database_view_access_entry_collection.

### SecurityPolicyReportDatabaseViewAccessEntry Reference

The following attributes are exported:

* `access_type` - The type of the access the user has on the table, there can be one or more from SELECT, UPDATE, INSERT or DELETE. 
* `column_name` - If there are column level privileges on a table or view.
* `grant_from_role` - This can be empty in case of direct grant, in case of indirect grant, this attribute displays the name of the  role which is granted to the user though which the user has access to the table. 
* `grantee` - Grantee is the user who can access the table or view
* `grantor` - The user who granted the privilege.
* `is_access_constrained_by_database_vault` - Indicates whether the table access is constrained via Oracle Database Vault.
* `is_access_constrained_by_real_application_security` - Indicates whether the view access is constrained via Real Application Security.
* `is_access_constrained_by_redaction` - Indicates whether the view access is constrained via Oracle Data Redaction.
* `is_access_constrained_by_sql_firewall` - Indicates whether the view access is constrained via Oracle Database SQL Firewall.
* `is_access_constrained_by_virtual_private_database` - Indicates whether the view access is constrained via Virtual Private Database.
* `key` - The unique key that identifies the table access report. It is numeric and unique within a security policy report.
* `privilege` - The name of the privilege.
* `privilege_grantable` - Indicates whether the grantee can grant this privilege to other users. Privileges can be granted to a user or role with  GRANT_OPTION or ADMIN_OPTION 
* `privilege_type` - Type of the privilege user has, this includes System Privilege, Schema Privilege, Object Privilege, Column Privilege, Owner or Schema Privilege on a schema. 
* `table_name` - The name of the database table the user has access to.
* `table_schema` - The name of the schema the table belongs to.
* `target_id` - The OCID of the of the  target database.
* `view_name` - The name of the view.
* `view_schema` - The name of the schema.
* `view_text` - Definition of the view.

