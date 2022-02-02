---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_commit"
sidebar_current: "docs-oci-datasource-devops-repository_commit"
description: |-
  Provides details about a specific Repository Commit in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_commit
This data source provides details about a specific Repository Commit resource in Oracle Cloud Infrastructure Devops service.

Retrieves a repository's commit by commit ID.

## Example Usage

```hcl
data "oci_devops_repository_commit" "test_repository_commit" {
	#Required
	commit_id = oci_devops_commit.test_commit.id
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `commit_id` - (Required) A filter to return only resources that match the given commit ID.
* `repository_id` - (Required) Unique repository identifier.


## Attributes Reference

The following attributes are exported:

* `author_email` - Email of the author of the repository.
* `author_name` - Name of the author of the repository.
* `commit_id` - Commit hash pointed to by reference name.
* `commit_message` - The commit message.
* `committer_email` - Email of who creates the commit.
* `committer_name` - Name of who creates the commit.
* `parent_commit_ids` - An array of parent commit IDs of created commit.
* `time_created` - The time at which commit was created.
* `tree_id` - Tree information for the specified commit.

