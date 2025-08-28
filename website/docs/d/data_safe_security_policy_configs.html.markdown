---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_configs"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_configs"
description: |-
  Provides the list of Security Policy Configs in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_configs
This data source provides the list of Security Policy Configs in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all security policy configurations in Data Safe.

The ListSecurityPolicyConfigs operation returns only the security policy configurations in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSecurityPolicyConfigs on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_security_policy_configs" "test_security_policy_configs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_policy_config_access_level
	compartment_id_in_subtree = var.security_policy_config_compartment_id_in_subtree
	display_name = var.security_policy_config_display_name
	security_policy_config_id = oci_data_safe_security_policy_config.test_security_policy_config.id
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id
	state = var.security_policy_config_state
	time_created_greater_than_or_equal_to = var.security_policy_config_time_created_greater_than_or_equal_to
	time_created_less_than = var.security_policy_config_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `security_policy_config_id` - (Optional) An optional filter to return only resources that match the specified OCID of the security policy configuration resource.
* `security_policy_id` - (Optional) An optional filter to return only resources that match the specified OCID of the security policy resource.
* `state` - (Optional) The current state of the security policy configuration resource.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `security_policy_config_collection` - The list of security_policy_config_collection.

### SecurityPolicyConfig Reference

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

