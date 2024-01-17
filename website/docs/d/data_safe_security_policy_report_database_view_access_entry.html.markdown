---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_report_database_view_access_entry"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_report_database_view_access_entry"
description: |-
  Provides details about a specific Security Policy Report Database View Access Entry in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_report_database_view_access_entry
This data source provides details about a specific Security Policy Report Database View Access Entry resource in Oracle Cloud Infrastructure Data Safe service.

Gets a database view access object by identifier.

## Example Usage

```hcl
data "oci_data_safe_security_policy_report_database_view_access_entry" "test_security_policy_report_database_view_access_entry" {
	#Required
	database_view_access_entry_key = var.security_policy_report_database_view_access_entry_database_view_access_entry_key
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id
}
```

## Argument Reference

The following arguments are supported:

* `database_view_access_entry_key` - (Required) The unique key that identifies the view access object. This is a system-generated identifier.
* `security_policy_report_id` - (Required) The OCID of the security policy report resource.


## Attributes Reference

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

