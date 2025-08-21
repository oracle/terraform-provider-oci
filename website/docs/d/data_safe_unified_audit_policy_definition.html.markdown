---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unified_audit_policy_definition"
sidebar_current: "docs-oci-datasource-data_safe-unified_audit_policy_definition"
description: |-
  Provides details about a specific Unified Audit Policy Definition in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_unified_audit_policy_definition
This data source provides details about a specific Unified Audit Policy Definition resource in Oracle Cloud Infrastructure Data Safe service.

Gets a unified audit policy definition by the specified OCID of the unified audit policy definition resource.

## Example Usage

```hcl
data "oci_data_safe_unified_audit_policy_definition" "test_unified_audit_policy_definition" {
	#Required
	unified_audit_policy_definition_id = oci_data_safe_unified_audit_policy_definition.test_unified_audit_policy_definition.id
}
```

## Argument Reference

The following arguments are supported:

* `unified_audit_policy_definition_id` - (Required) The OCID of the unified audit policy definition resource.


## Attributes Reference

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

