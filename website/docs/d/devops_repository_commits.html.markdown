---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_commits"
sidebar_current: "docs-oci-datasource-devops-repository_commits"
description: |-
  Provides the list of Repository Commits in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_commits
This data source provides the list of Repository Commits in Oracle Cloud Infrastructure Devops service.

Returns a list of Commits.


## Example Usage

```hcl
data "oci_devops_repository_commits" "test_repository_commits" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	author_name = var.repository_commit_author_name
	commit_message = var.repository_commit_commit_message
	exclude_ref_name = var.repository_commit_exclude_ref_name
	file_path = var.repository_commit_file_path
	ref_name = var.repository_commit_ref_name
	timestamp_greater_than_or_equal_to = var.repository_commit_timestamp_greater_than_or_equal_to
	timestamp_less_than_or_equal_to = var.repository_commit_timestamp_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `author_name` - (Optional) A filter to return any commits that are pushed by the requested author.
* `commit_message` - (Optional) A filter to return any commits that contains the given message.
* `exclude_ref_name` - (Optional) A filter to exclude commits that match the given Ref name.
* `file_path` - (Optional) A filter to return only commits that affect any of the specified paths.
* `ref_name` - (Optional) A filter to return only resources that match the given Ref name.
* `repository_id` - (Required) unique Repository identifier.
* `timestamp_greater_than_or_equal_to` - (Optional) A filter to return commits only created after the specified timestamp value.
* `timestamp_less_than_or_equal_to` - (Optional) A filter to return commits only created before the specified timestamp value.


## Attributes Reference

The following attributes are exported:

* `repository_commit_collection` - The list of repository_commit_collection.

### RepositoryCommit Reference

The following attributes are exported:

* `author_email` - The email of the author of the repository.
* `author_name` - The name of the author of the repository.
* `commit_id` - Commit hash pointed to by Ref name
* `commit_message` - The commit message.
* `committer_email` - The email of who create the commit.
* `committer_name` - The name of who create the commit.
* `parent_commit_ids` - An array of parent commit ids of created commit.
* `time_created` - The time at which commit was created.
* `tree_id` - Tree information for the specified commit

