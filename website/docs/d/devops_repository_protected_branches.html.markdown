---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_protected_branches"
sidebar_current: "docs-oci-datasource-devops-repository_protected_branches"
description: |-
  Provides the list of Repository Protected Branches in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_protected_branches
This data source provides the list of Repository Protected Branches in Oracle Cloud Infrastructure Devops service.

Returns a list of Protected Branches.


## Example Usage

```hcl
data "oci_devops_repository_protected_branches" "test_repository_protected_branches" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	name = var.repository_protected_branch_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) A filter to return only resources that match the given branch name.
* `repository_id` - (Required) Unique repository identifier.


## Attributes Reference

The following attributes are exported:

* `protected_branch_collection` - The list of protected_branch_collection.

### RepositoryProtectedBranch Reference

The following attributes are exported:

* `items` - List of objects describing protected branches
	* `branch_name` - Branch name inside a repository.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	* `protection_levels` - Protection level to be added on the branch.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`

