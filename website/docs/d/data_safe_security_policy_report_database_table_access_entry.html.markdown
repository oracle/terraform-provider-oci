---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_report_database_table_access_entry"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_report_database_table_access_entry"
description: |-
  Provides details about a specific Security Policy Report Database Table Access Entry in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_report_database_table_access_entry
This data source provides details about a specific Security Policy Report Database Table Access Entry resource in Oracle Cloud Infrastructure Data Safe service.

Gets a database table access entry object by identifier.

## Example Usage

```hcl
data "oci_data_safe_security_policy_report_database_table_access_entry" "test_security_policy_report_database_table_access_entry" {
	#Required
	database_table_access_entry_key = var.security_policy_report_database_table_access_entry_database_table_access_entry_key
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id
}
```

## Argument Reference

The following arguments are supported:

* `database_table_access_entry_key` - (Required) The unique key that identifies the table access object. This is a system-generated identifier.
* `security_policy_report_id` - (Required) The OCID of the security policy report resource.


## Attributes Reference

The following attributes are exported:

* `access_through_object` - A non-null value in this field indicates the object through which user has access to table, possible values could be table or view. 
* `access_type` - The type of the access the user has on the table, there can be one or more from SELECT, UPDATE, INSERT, OWNER or DELETE. 
* `are_all_tables_accessible` - Indicates whether the user has access to all the tables in the schema.
* `column_name` - If there are column level privileges on a table or view.
* `grant_from_role` - This can be empty in case of direct grant, in case of indirect grant, this attribute displays the name of the  role which is granted to the user though which the user has access to the table. 
* `grantee` - Grantee is the user who can access the table
* `grantor` - The one who granted this privilege.
* `is_access_constrained_by_database_vault` - Indicates whether the table access is constrained via Oracle Database Vault.
* `is_access_constrained_by_label_security` - Indicates whether the table access is constrained via Oracle Label Security.
* `is_access_constrained_by_real_application_security` - Indicates whether the table access is constrained via Real Application Security.
* `is_access_constrained_by_redaction` - Indicates whether the table access is constrained via Oracle Data Redaction.
* `is_access_constrained_by_sql_firewall` - Indicates whether the table access is constrained via Oracle Database SQL Firewall.
* `is_access_constrained_by_view` - Indicates whether the access is constrained on a table via a view.
* `is_access_constrained_by_virtual_private_database` - Indicates whether the table access is constrained via Virtual Private Database.
* `is_sensitive` - Indicates whether the table is marked as sensitive.
* `key` - The unique key that identifies the table access report. It is numeric and unique within a security policy report.
* `privilege` - Name of the privilege.
* `privilege_grantable` - Indicates whether the grantee can grant this privilege to other users. Privileges can be granted to a user or role with  GRANT_OPTION or ADMIN_OPTION 
* `privilege_type` - Type of the privilege user has, this includes System Privilege, Schema Privilege, Object Privilege, Column Privilege, Owner or Schema Privilege on a schema. 
* `table_name` - The name of the database table the user has access to.
* `table_schema` - The name of the schema the table belongs to.
* `target_id` - The OCID of the of the  target database.

