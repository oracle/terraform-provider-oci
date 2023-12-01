---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_policy_management"
sidebar_current: "docs-oci-resource-data_safe-sql_firewall_policy_management"
description: |-
  Provides the Sql Firewall Policy Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sql_firewall_policy_management
This resource provides the Sql Firewall Policy Management resource in Oracle Cloud Infrastructure Data Safe service.

Updates the SQL firewall policy.

## Example Usage

```hcl
resource "oci_data_safe_sql_firewall_policy_management" "test_sql_firewall_policy_management" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_data_safe_target_database.test_target_database.id
	db_user_name = oci_identity_user.test_user.name
	
	#Optional
	allowed_client_ips = var.sql_firewall_policy_management_allowed_client_ips
	allowed_client_os_usernames = var.sql_firewall_policy_management_allowed_client_os_usernames
	allowed_client_programs = var.sql_firewall_policy_management_allowed_client_programs
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sql_firewall_policy_management_description
	display_name = var.sql_firewall_policy_management_display_name
	enforcement_scope = var.sql_firewall_policy_management_enforcement_scope
	freeform_tags = {"Department"= "Finance"}
	status = var.sql_firewall_policy_management_status
	violation_action = var.sql_firewall_policy_management_violation_action
	violation_audit = var.sql_firewall_policy_management_violation_audit
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the SQL collection.
* `target_id` - (Required) Unique target identifier.
* `db_user_name` - (Required) The database user name. 
* `allowed_client_ips` - (Optional) (Updatable) List of allowed ip addresses for the SQL firewall policy.
* `allowed_client_os_usernames` - (Optional) (Updatable) List of allowed operating system user names for the SQL firewall policy.
* `allowed_client_programs` - (Optional) (Updatable) List of allowed client programs for the SQL firewall policy.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the SQL firewall policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the SQL firewall policy.
* `display_name` - (Optional) (Updatable) The display name of the SQL firewall policy. The name does not have to be unique, and it is changeable.
* `enforcement_scope` - (Optional) (Updatable) Specifies the SQL firewall policy enforcement option.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `sql_firewall_policy_id` - (Required) The OCID of the SQL firewall policy resource.
* `status` - (Optional) (Updatable) Specifies whether the SQL firewall policy is enabled or disabled.
* `violation_action` - (Optional) (Updatable) Specifies the SQL firewall action based on detection of SQL firewall violations.
* `violation_audit` - (Optional) (Updatable) Specifies whether a unified audit policy should be enabled for auditing the SQL firewall policy violations.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sql Firewall Policy Management
	* `update` - (Defaults to 20 minutes), when updating the Sql Firewall Policy Management
	* `delete` - (Defaults to 20 minutes), when destroying the Sql Firewall Policy Management


## Import

Import is not supported for this resource.

