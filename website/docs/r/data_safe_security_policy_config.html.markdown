---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_config"
sidebar_current: "docs-oci-resource-data_safe-security_policy_config"
description: |-
  Provides the Security Policy Config resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_policy_config
This resource provides the Security Policy Config resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new security policy configuration resource.


## Example Usage

```hcl
resource "oci_data_safe_security_policy_config" "test_security_policy_config" {
	#Required
	compartment_id = var.compartment_id
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.security_policy_config_description
	display_name = var.security_policy_config_display_name
	firewall_config {

		#Optional
		exclude_job = var.security_policy_config_firewall_config_exclude_job
		status = var.security_policy_config_firewall_config_status
		violation_log_auto_purge = var.security_policy_config_firewall_config_violation_log_auto_purge
	}
	freeform_tags = {"Department"= "Finance"}
	unified_audit_policy_config {

		#Optional
		exclude_datasafe_user = var.security_policy_config_unified_audit_policy_config_exclude_datasafe_user
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the security policy configuration.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the security policy.
* `display_name` - (Optional) (Updatable) The display name of the security policy configuration. The name does not have to be unique, and it is changeable.
* `firewall_config` - (Optional) (Updatable) Details to update the SQL Firewall configuration. 
	* `exclude_job` - (Optional) (Updatable) Specifies whether the firewall should include or exclude the database internal job activities.
	* `status` - (Optional) (Updatable) Specifies whether the firewall is enabled or disabled.
	* `violation_log_auto_purge` - (Optional) (Updatable) Specifies whether Data Safe should automatically purge the violation logs  from the database after collecting the violation logs and persisting them in Data Safe. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `security_policy_id` - (Required) The OCID of the security policy corresponding to the security policy configuration.
* `unified_audit_policy_config` - (Optional) (Updatable) The unified audit policy related configurations. 
	* `exclude_datasafe_user` - (Optional) (Updatable) Specifies whether the Data Safe service account on the target database should be excluded in the unified audit policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security policy configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy configuration.
* `display_name` - The display name of the security policy configuration.
* `firewall_config` - The SQL Firewall related configurations. 
	* `exclude_job` - Specifies whether the firewall should include or exclude the database internal job activities.
	* `status` - Specifies if the firewall is enabled or disabled.
	* `time_status_updated` - The date and time the firewall configuration was last updated, in the format defined by RFC3339.
	* `violation_log_auto_purge` - Specifies whether Data Safe should automatically purge the violation logs  from the database after collecting the violation logs and persisting on Data Safe. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy configuration.
* `lifecycle_details` - Details about the current state of the security policy configuration.
* `security_policy_id` - The OCID of the security policy corresponding to the security policy configuration.
* `state` - The current state of the security policy configuration.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the security policy configuration was created, in the format defined by RFC3339.
* `time_updated` - The date and time the security policy configuration was last updated, in the format defined by RFC3339.
* `unified_audit_policy_config` - The unified audit policy related configurations. 
	* `exclude_datasafe_user` - Specifies whether the Data Safe service account on the target database should be excluded in the unified audit policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Policy Config
	* `update` - (Defaults to 20 minutes), when updating the Security Policy Config
	* `delete` - (Defaults to 20 minutes), when destroying the Security Policy Config


## Import

SecurityPolicyConfigs can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_security_policy_config.test_security_policy_config "id"
```

