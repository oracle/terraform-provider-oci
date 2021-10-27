---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_mirrorrecord_history"
sidebar_current: "docs-oci-datasource-devops-repository_mirrorrecord_history"
description: |-
  Provides details about a specific Repository Mirrorrecord History in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_mirrorrecord_history
This data source provides details about a specific Repository Mirrorrecord History resource in Oracle Cloud Infrastructure Devops service.

Returns a list of mirror entry in history within 30 days


## Example Usage

```hcl
data "oci_devops_repository_mirrorrecord_history" "test_repository_mirrorrecord_history" {
	#Required
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `items` - List of mirror entry objects.
	* `end_time` - The time complete a mirror operation.
	* `enqueue_time` - The time to enqueue a mirror operation.
	* `mirror_status` - Mirror status of current mirror entry.
	* `start_time` - The time to start a mirror operation.
	* `work_request_id` - Workrequest Id to track current mirror operation

