variable "branch_name" {
	default = "<branch_name>"
}
variable "repository_id" {
}
variable "project_id" {
}
variable "repository_setting_merge_checks_last_build_succeeded"{
	default = "DISABLED"
}
variable "repository_setting_approval_rules_items_min_approvals_count" {
	default = 1
}

resource "oci_devops_repository_protected_branch_management" "test_repository_protected_branch_management" {
	#Required
	branch_name = var.branch_name
	repository_id = var.repository_id

	#Optional
#	protection_levels = <protection_levels>
}

resource "oci_devops_project_repository_setting" "test_project_repository_setting" {
	#Required
	project_id = var.project_id

	#Optional
	approval_rules {
		#Required
		items {
			#Required
			min_approvals_count = var.repository_setting_approval_rules_items_min_approvals_count
			name = "ApprovalRuleName"

			#Optional
			destination_branch = var.branch_name
			reviewers {
				#Required
				principal_id = <reviewers>
			}
		}
	}
	merge_settings {
		#Required
		allowed_merge_strategies = ["MERGE_COMMIT", "FAST_FORWARD"]
		default_merge_strategy = "MERGE_COMMIT"
	}
}

resource "oci_devops_repository_setting" "test_repository_setting" {
	#Required
	repository_id = var.repository_id

	#Optional
	approval_rules {
		#Required
		items {
			#Required
			min_approvals_count = var.repository_setting_approval_rules_items_min_approvals_count
			name = "ApprovalName"

			#Optional
			destination_branch = var.branch_name
			reviewers {
				#Required
				principal_id = "<principal_id>"
			}
		}
	}
	merge_checks {
		#Required
		last_build_succeeded = var.repository_setting_merge_checks_last_build_succeeded
	}
	merge_settings {
                #Required
                allowed_merge_strategies = ["MERGE_COMMIT", "FAST_FORWARD"]
                default_merge_strategy = "MERGE_COMMIT"
        }
}
