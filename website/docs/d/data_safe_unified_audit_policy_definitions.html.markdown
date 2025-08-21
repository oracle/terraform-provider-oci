---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policy_definitions"
sidebar_current: "docs-oci-datasource-data_safe-unified_audit_policy_definitions"
description: |-
  Provides the list of Unified Audit Policy Definitions in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_unified_audit_policy_definitions
This data source provides the list of Unified Audit Policy Definitions in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all unified audit policy definitions in Data Safe.

The ListUnifiedAuditPolicyDefinitions operation returns only the unified audit policy definitions in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requester has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListUnifiedAuditPolicyDefinitions on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_unified_audit_policy_definitions" "test_unified_audit_policy_definitions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.unified_audit_policy_definition_access_level
	compartment_id_in_subtree = var.unified_audit_policy_definition_compartment_id_in_subtree
	display_name = var.unified_audit_policy_definition_display_name
	is_seeded = var.unified_audit_policy_definition_is_seeded
	state = var.unified_audit_policy_definition_state
	unified_audit_policy_category = var.unified_audit_policy_definition_unified_audit_policy_category
	unified_audit_policy_definition_id = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id
	unified_audit_policy_name = oci_data_safe_unified_audit_policy.test_unified_audit_policy.name
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `is_seeded` - (Optional) A boolean flag indicating to list seeded unified audit policy definitions. Set this parameter to get list of seeded unified audit policy definitions.
* `state` - (Optional) The current state of the unified audit policy definition.
* `unified_audit_policy_category` - (Optional) The category to which the unified audit policy definition belongs to.
* `unified_audit_policy_definition_id` - (Optional) An optional filter to return only resources that match the specified OCID of the unified audit policy definition resource.
* `unified_audit_policy_name` - (Optional) The name of the unified audit policy.


## Attributes Reference

The following attributes are exported:

* `unified_audit_policy_definition_collection` - The list of unified_audit_policy_definition_collection.

### UnifiedAuditPolicyDefinition Reference

The following attributes are exported:

* `audit_policy_category` - The category to which the unified audit policy belongs in the target database.
* `compartment_id` - The OCID of the compartment containing the unified audit policy definition.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the unified audit policy definition.
* `display_name` - The display name of the unified audit policy definition.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the unified audit policy definition.
* `is_seeded` - Signifies whether the unified audit policy definition is seeded or not.
* `lifecycle_details` - Details about the current state of the unified audit policy definition.
* `policy_definition_statement` - The definition of the unified audit policy to be provisioned in the target database.
* `policy_name` - The unified audit policy name in the target database.
* `state` - The current state of the unified audit policy definition.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the unified audit policy was created, in the format defined by RFC3339.
* `time_updated` - The last date and time the unified audit policy was updated, in the format defined by RFC3339.

