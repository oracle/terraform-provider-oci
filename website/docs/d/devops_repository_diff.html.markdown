---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_diff"
sidebar_current: "docs-oci-datasource-devops-repository_diff"
description: |-
  Provides details about a specific Repository Diff in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_diff
This data source provides details about a specific Repository Diff resource in Oracle Cloud Infrastructure Devops service.

Gets the line-by-line difference between files on different commits.


## Example Usage

```hcl
data "oci_devops_repository_diff" "test_repository_diff" {
	#Required
	base_version = var.repository_diff_base_version
	file_path = var.repository_diff_file_path
	repository_id = oci_devops_repository.test_repository.id
	target_version = var.repository_diff_target_version

	#Optional
	is_comparison_from_merge_base = var.repository_diff_is_comparison_from_merge_base
}
```

## Argument Reference

The following arguments are supported:

* `base_version` - (Required) The branch to compare changes against.
* `file_path` - (Required) Path to a file within a repository.
* `is_comparison_from_merge_base` - (Optional) Boolean to indicate whether to use merge base or most recent revision.
* `repository_id` - (Required) Unique repository identifier.
* `target_version` - (Required) The branch where changes are coming from.


## Attributes Reference

The following attributes are exported:

* `are_conflicts_in_file` - Indicates whether the changed file contains conflicts.
* `changes` - List of changed section in the file.
	* `base_line` - Line number in base version where changes begin.
	* `base_span` - Number of lines chunk spans in base version.
	* `diff_sections` - List of difference section.
		* `lines` - The lines within changed section.
			* `base_line` - The number of a line in the base version.
			* `conflict_marker` - Indicates whether a line in a conflicted section of the difference is from the base version, the target version, or if its just a marker indicating the beginning, middle, or end of a conflicted section.
			* `line_content` - The contents of a line.
			* `target_line` - The number of a line in the target version.
		* `type` - Type of change.
	* `target_line` - Line number in target version where changes begin.
	* `target_span` - Number of lines chunk spans in target version.
* `is_binary` - Indicates whether the file is binary.
* `is_large` - Indicates whether the file is large.
* `new_id` - The ID of the changed object on the target version.
* `new_path` - The path on the target version to the changed object.
* `old_id` - The ID of the changed object on the base version.
* `old_path` - The path on the base version to the changed object.

