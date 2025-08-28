---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_config"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_config"
description: |-
  Provides details about a specific Security Policy Config in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_config
This data source provides details about a specific Security Policy Config resource in Oracle Cloud Infrastructure Data Safe service.

Gets a security policy configuration by identifier.

## Example Usage

```hcl
data "oci_data_safe_security_policy_config" "test_security_policy_config" {
	#Required
	security_policy_config_id = oci_data_safe_security_policy_config.test_security_policy_config.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_config_id` - (Required) The OCID of the security policy configuration resource.


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

