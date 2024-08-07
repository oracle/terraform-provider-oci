---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_project_repository_setting"
sidebar_current: "docs-oci-resource-devops-project_repository_setting"
description: |-
  Provides the Project Repository Setting resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_project_repository_setting
This resource provides the Project Repository Setting resource in Oracle Cloud Infrastructure Devops service.

Updates the repository settings for a project.

## Example Usage

```hcl
resource "oci_devops_project_repository_setting" "test_project_repository_setting" {
	#Required
	project_id = oci_devops_project.test_project.id

	#Optional
	approval_rules {
		#Required
		items {
			#Required
			min_approvals_count = var.project_repository_setting_approval_rules_items_min_approvals_count
			name = var.project_repository_setting_approval_rules_items_name

			#Optional
			destination_branch = var.project_repository_setting_approval_rules_items_destination_branch
			reviewers {
				#Required
				principal_id = oci_devops_principal.test_principal.id
			}
		}
	}
	merge_settings {
		#Required
		allowed_merge_strategies = var.project_repository_setting_merge_settings_allowed_merge_strategies
		default_merge_strategy = var.project_repository_setting_merge_settings_default_merge_strategy
	}
}
```

## Argument Reference

The following arguments are supported:

* `approval_rules` - (Optional) (Updatable) List of approval rules which must be statisfied before pull requests which match the rules can be merged
	* `items` - (Required) (Updatable) List of approval rules.
		* `destination_branch` - (Optional) (Updatable) Branch name where pull requests targeting the branch must satisfy the approval rule. This value being null means the rule applies to all pull requests
		* `min_approvals_count` - (Required) (Updatable) Minimum number of approvals which must be provided by the reviewers specified in the list before the rule can be satisfied
		* `name` - (Required) (Updatable) Name which is used to uniquely identify an approval rule.
		* `reviewers` - (Optional) (Updatable) List of users who must provide approvals up to the minApprovalsCount specified in the rule. An empty list means the approvals can come from any user.
			* `principal_id` - (Required) (Updatable) Pull Request reviewer id
* `merge_settings` - (Optional) (Updatable) Enabled and disabled merge strategies for a project or repository, also contains a default strategy.
	* `allowed_merge_strategies` - (Required) (Updatable) List of merge strategies which are allowed for a Project or Repository.
	* `default_merge_strategy` - (Required) (Updatable) Default type of merge strategy associated with the a Project or Repository.
* `project_id` - (Required) Unique project identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Project Repository Setting
	* `update` - (Defaults to 20 minutes), when updating the Project Repository Setting
	* `delete` - (Defaults to 20 minutes), when destroying the Project Repository Setting


## Import

ProjectRepositorySettings can be imported using the `id`, e.g.

```
$ terraform import oci_devops_project_repository_setting.test_project_repository_setting "projects/{projectId}/repositorySettings" 
```

