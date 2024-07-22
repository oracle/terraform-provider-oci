---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_protected_branch_management"
sidebar_current: "docs-oci-resource-devops-repository_protected_branch_management"
description: |-
  Provides the Repository Protected Branch Management resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_repository_protected_branch_management
This resource provides the Repository Protected Branch Management resource in Oracle Cloud Infrastructure Devops service.

Creates a restriction on a branch that prevents certain actions on it.

## Example Usage

```hcl
resource "oci_devops_repository_protected_branch_management" "test_repository_protected_branch_management" {
	#Required
	branch_name = var.repository_protected_branch_management_branch_name
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	protection_levels = var.repository_protected_branch_management_protection_levels
}
```

## Argument Reference

The following arguments are supported:

* `branch_name` - (Required) Name of a branch to protect.
* `protection_levels` - (Optional) Level of protection to add on a branch.
* `repository_id` - (Required) Unique repository identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `branch_name` - Branch name inside a repository.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `protection_levels` - Protection levels to be added on the branch.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Repository Protected Branch Management
	* `update` - (Defaults to 20 minutes), when updating the Repository Protected Branch Management
	* `delete` - (Defaults to 20 minutes), when destroying the Repository Protected Branch Management


## Import

Import is not supported for this resource.

