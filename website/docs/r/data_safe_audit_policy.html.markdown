---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_policy"
sidebar_current: "docs-oci-resource-data_safe-audit_policy"
description: |-
  Provides the Audit Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_audit_policy
This resource provides the Audit Policy resource in Oracle Cloud Infrastructure Data Safe service.

Updates the audit policy.

## Example Usage

```hcl
resource "oci_data_safe_audit_policy" "test_audit_policy" {
	#Required
	audit_policy_id = oci_data_safe_audit_policy.test_audit_policy.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.audit_policy_description
	display_name = var.audit_policy_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `audit_policy_id` - (Required) Unique audit policy identifier.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the audit policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the audit policy.
* `display_name` - (Optional) (Updatable) The display name of the audit policy. The name does not have to be unique, and it is changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `provision_trigger` - (Optional) (Updatable) An optional property when incremented triggers Provision. Could be set to any integer value.
* `retrieve_from_target_trigger` - (Optional) (Updatable) An optional property when incremented triggers Retrieve From Target. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `audit_conditions` - Lists the audit policy provisioning conditions for the target database.
	* `audit_policy_name` - Indicates the audit policy name. Refer to the [documentation](https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827) for seeded audit policy names. For custom policies, refer to the user-defined policy name created in the target database. 
	* `enable_conditions` - Indicates the users/roles in the target database for which the audit policy is enforced, and the success/failure event condition to generate the audit event..
		* `entity_names` - List of users or roles that the policy must be enabled for.
		* `entity_selection` - The entity include or exclude selection.
		* `entity_type` - The entity type that the policy must be enabled for.
		* `operation_status` - The operation status that the policy must be enabled for.
	* `is_data_safe_service_account_audited` - Indicates whether the Data Safe user activity on the target database will be audited by the policy.
	* `is_priv_users_managed_by_data_safe` - Indicates whether the privileged user list is managed by Data Safe.
	* `is_enabled` - Indicates whether the policy has to be enabled or disabled in the target database. Set this to true if you want the audit policy to be enabled in the target database. If the seeded audit policy is not already created in the database, the provisioning creates and enables them. If this is set to false, the policy will be disabled in the target database.
* `audit_specifications` - Represents all available audit policy specifications relevant for the target database. For more details on available audit polcies, refer to [documentation](https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827). 
	* `audit_policy_category` - The category to which the audit policy belongs.
	* `audit_policy_name` - Indicates the audit policy name. Refer to the [documentation](https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827) for seeded audit policy names. For custom policies, refer to the user-defined policy name created in the target database. 
	* `database_policy_names` - Indicates the names of corresponding database policy ( or policies) in the target database.
	* `enable_status` - Indicates whether the policy has been enabled, disabled or partially enabled in the target database. The status is PARTIALLY_ENABLED if any of the constituent database audit policies is not enabled.
	* `enabled_entities` - Indicates on whom the audit policy is enabled.
	* `is_created` - Indicates whether the policy is already created on the target database.
	* `is_enabled_for_all_users` - Indicates whether the policy by default is enabled for all users with no flexibility to alter the enablement conditions.
	* `is_seeded_in_data_safe` - Indicates whether the audit policy is one of the seeded policies provided by Oracle Data Safe.
	* `is_seeded_in_target` - Indicates whether the audit policy is one of the predefined policies provided by Oracle Database.
	* `is_view_only` - Indicates whether the audit policy is available for provisioning/ de-provisioning from Oracle Data Safe, or is only available for displaying the current provisioning status from the target.
	* `partially_enabled_msg` - Provides information about the policy that has been only partially enabled.
* `compartment_id` - The OCID of the compartment containing the audit policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the audit policy.
* `display_name` - The display name of the audit policy.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the audit policy.
* `is_data_safe_service_account_excluded` - Option provided to users at the target to indicate whether the Data Safe service account has to be excluded while provisioning the audit policies.
* `lifecycle_details` - Details about the current state of the audit policy in Data Safe.
* `state` - The current state of the audit policy.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the target for which the audit policy is created.
* `time_created` - The time the the audit policy was created, in the format defined by RFC3339.
* `time_last_provisioned` - Indicates the last provisioning time of audit policies on the target, in the format defined by RFC3339.
* `time_last_retrieved` - The time when the audit policies was last retrieved from this target, in the format defined by RFC3339.
* `time_updated` - The last date and time the audit policy was updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Audit Policy
	* `update` - (Defaults to 20 minutes), when updating the Audit Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Audit Policy


## Import

AuditPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_audit_policy.test_audit_policy "id"
```

