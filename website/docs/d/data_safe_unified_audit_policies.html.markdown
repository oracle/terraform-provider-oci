---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policies"
sidebar_current: "docs-oci-datasource-data_safe-unified_audit_policies"
description: |-
  Provides the list of Unified Audit Policies in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_unified_audit_policies
This data source provides the list of Unified Audit Policies in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all Unified Audit policies.

The ListUnifiedAuditPolicies operation returns only the Unified Audit policies in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requester has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a sub-compartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListUnifiedAuditPolicies on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and sub-compartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_unified_audit_policies" "test_unified_audit_policies" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.unified_audit_policy_access_level
	compartment_id_in_subtree = var.unified_audit_policy_compartment_id_in_subtree
	display_name = var.unified_audit_policy_display_name
	is_seeded = var.unified_audit_policy_is_seeded
	security_policy_id = oci_data_safe_security_policy.test_security_policy.id
	state = var.unified_audit_policy_state
	time_created_greater_than_or_equal_to = var.unified_audit_policy_time_created_greater_than_or_equal_to
	time_created_less_than = var.unified_audit_policy_time_created_less_than
	unified_audit_policy_definition_id = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id
	unified_audit_policy_id = oci_data_safe_unified_audit_policy.test_unified_audit_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `is_seeded` - (Optional) A boolean flag indicating to list seeded unified audit policies. Set this parameter to get list of seeded unified audit policies.
* `security_policy_id` - (Optional) An optional filter to return only resources that match the specified OCID of the security policy resource.
* `state` - (Optional) The current state of the Unified Audit policy.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `unified_audit_policy_definition_id` - (Optional) An optional filter to return only resources that match the specified OCID of the unified audit policy definition resource.
* `unified_audit_policy_id` - (Optional) An optional filter to return only resources that match the specified OCID of the Unified Audit policy resource.


## Attributes Reference

The following attributes are exported:

* `unified_audit_policy_collection` - The list of unified_audit_policy_collection.

### UnifiedAuditPolicy Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the unified audit policy.
* `conditions` - Lists the audit policy provisioning conditions.
	* `attribute_set_id` - The OCID of the attribute set.
	* `entity_selection` - Specifies whether to include or exclude the specified users or roles.
	* `entity_type` - The type of users or roles that the unified audit policy is applied to.
	* `operation_status` - The operation status that the policy must be enabled for.
	* `role_names` - List of roles that the policy must be enabled for.
	* `user_names` - The list of users that the unified audit policy is enabled for.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the unified audit policy.
* `display_name` - The display name of the unified audit policy.
* `enabled_entities` - Indicates on whom the audit policy is enabled.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the unified audit policy.
* `is_seeded` - Indicates whether the unified audit policy is seeded or not.
* `lifecycle_details` - The details of the current state of the unified audit policy in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the unified audit policy.
* `state` - The current state of the unified audit policy.
* `status` - Indicates whether the policy has been enabled or disabled.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the unified audit policy was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the unified audit policy was updated, in the format defined by RFC3339.
* `unified_audit_policy_definition_id` - The OCID of the associated unified audit policy definition.

