---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_policy"
sidebar_current: "docs-oci-datasource-data_safe-sql_firewall_policy"
description: |-
  Provides details about a specific Sql Firewall Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sql_firewall_policy
This data source provides details about a specific Sql Firewall Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets a SQL firewall policy by identifier.

## Example Usage

```hcl
data "oci_data_safe_sql_firewall_policy" "test_sql_firewall_policy" {
	#Required
	sql_firewall_policy_id = oci_data_safe_sql_firewall_policy.test_sql_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `sql_firewall_policy_id` - (Required) The OCID of the SQL firewall policy resource.


## Attributes Reference

The following attributes are exported:

* `allowed_client_ips` - The list of allowed ip addresses for the SQL firewall policy.
* `allowed_client_os_usernames` - The list of allowed operating system user names for the SQL firewall policy.
* `allowed_client_programs` - The list of allowed client programs for the SQL firewall policy.
* `compartment_id` - The OCID of the compartment containing the SQL firewall policy.
* `db_user_name` - The database user name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the SQL firewall policy.
* `display_name` - The display name of the SQL firewall policy.
* `enforcement_scope` - Specifies the SQL firewall policy enforcement option.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the SQL firewall policy.
* `lifecycle_details` - Details about the current state of the SQL firewall policy in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the SQL firewall policy.
* `sql_level` - Specifies the level of SQL included for this SQL firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
* `state` - The current state of the SQL firewall policy.
* `status` - Specifies whether the SQL firewall policy is enabled or disabled.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the SQL firewall policy was created, in the format defined by RFC3339.
* `time_updated` - The date and time the SQL firewall policy was last updated, in the format defined by RFC3339.
* `violation_action` - Specifies the mode in which the SQL firewall policy is enabled.
* `violation_audit` - Specifies whether a unified audit policy should be enabled for auditing the SQL firewall policy violations.

