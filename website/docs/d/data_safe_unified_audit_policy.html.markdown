---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policy"
sidebar_current: "docs-oci-datasource-data_safe-unified_audit_policy"
description: |-
  Provides details about a specific Unified Audit Policy in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_unified_audit_policy
This data source provides details about a specific Unified Audit Policy resource in Oracle Cloud Infrastructure Data Safe service.

Gets a Unified Audit policy by identifier.

## Example Usage

```hcl
data "oci_data_safe_unified_audit_policy" "test_unified_audit_policy" {
	#Required
	unified_audit_policy_id = oci_data_safe_unified_audit_policy.test_unified_audit_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `unified_audit_policy_id` - (Required) The OCID of the Unified Audit policy resource.


## Attributes Reference

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

