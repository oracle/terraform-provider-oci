---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_diffs"
sidebar_current: "docs-oci-datasource-devops-repository_diffs"
description: |-
  Provides the list of Repository Diffs in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_diffs
This data source provides the list of Repository Diffs in Oracle Cloud Infrastructure Devops service.

Compares two revisions and lists the differences. Supports comparison between two refs or commits.


## Example Usage

```hcl
data "oci_devops_repository_diffs" "test_repository_diffs" {
	#Required
	base_version = var.repository_diff_base_version
	repository_id = oci_devops_repository.test_repository.id
	target_version = var.repository_diff_target_version

	#Optional
	is_comparison_from_merge_base = var.repository_diff_is_comparison_from_merge_base
}
```

## Argument Reference

The following arguments are supported:

* `base_version` - (Required) The commit or ref name to compare changes against
* `is_comparison_from_merge_base` - (Optional) boolean for whether to use merge base or most recent revision
* `repository_id` - (Required) unique Repository identifier.
* `target_version` - (Required) The commit or ref name where changes are coming from


## Attributes Reference

The following attributes are exported:

* `diff_collection` - The list of diff_collection.

### RepositoryDiff Reference

The following attributes are exported:

* `are_conflicts_in_file` - Indicates whether the changed file contains conflicts.
* `changes` - List of changed section in the file.
	* `base_line` - Line number in base version where changes begin.
	* `base_span` - Number of lines chunk spans in base version.
	* `diff_sections` - List of DiffSection.
		* `lines` - The lines within changed section.
			* `base_line` - The number of a line in the base version.
			* `conflict_marker` - Indicates whether a line in a conflicted section of the diff is from the base version, the target version, or if its just a marker indicating the beginning, middle, or end of a conflicted section.
			* `line_content` - The contents of a line.
			* `target_line` - The number of a line in the target version.
		* `type` - Type of change.
	* `target_line` - Line number in target version where changes begin.
	* `target_span` - Number of lines chunk spans in target version.
* `is_binary` - Indicates whether the file is binary.
* `is_large` - Indicates whether the file is large.
* `new_id` - The ID of the changed object on the targetVersion.
* `new_path` - The path on the targetVersion to the changed object.
* `old_id` - The ID of the changed object on the baseVersion.
* `old_path` - The path on the baseVersion to the changed object.

