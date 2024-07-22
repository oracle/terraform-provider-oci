---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_project_repository_setting"
sidebar_current: "docs-oci-datasource-devops-project_repository_setting"
description: |-
  Provides details about a specific Project Repository Setting in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_project_repository_setting
This data source provides details about a specific Project Repository Setting resource in Oracle Cloud Infrastructure Devops service.

Retrieves a project's repository settings details.

## Example Usage

```hcl
data "oci_devops_project_repository_setting" "test_project_repository_setting" {
	#Required
	project_id = oci_devops_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) Unique project identifier.


## Attributes Reference

The following attributes are exported:

* `approval_rules` - List of approval rules which must be statisfied before pull requests which match the rules can be merged
	* `items` - List of approval rules.
		* `destination_branch` - Branch name where pull requests targeting the branch must satisfy the approval rule. This value being null means the rule applies to all pull requests
		* `min_approvals_count` - Minimum number of approvals which must be provided by the reviewers specified in the list before the rule can be satisfied
		* `name` - Name which is used to uniquely identify an approval rule.
		* `reviewers` - List of users who must provide approvals up to the minApprovalsCount specified in the rule. An empty list means the approvals can come from any user.
			* `principal_id` - the OCID of the principal
			* `principal_name` - the name of the principal
			* `principal_state` - The state of the principal, it can be active or inactive or suppressed for emails
			* `principal_type` - the type of principal
* `merge_settings` - Enabled and disabled merge strategies for a project or repository, also contains a default strategy.
	* `allowed_merge_strategies` - List of merge strategies which are allowed for a Project or Repository.
	* `default_merge_strategy` - Default type of merge strategy associated with the a Project or Repository.

