---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_mirror_records"
sidebar_current: "docs-oci-datasource-devops-repository_mirror_records"
description: |-
  Provides the list of Repository Mirror Records in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_mirror_records
This data source provides the list of Repository Mirror Records in Oracle Cloud Infrastructure Devops service.

Returns a list of mirror entry in history within 30 days


## Example Usage

```hcl
data "oci_devops_repository_mirror_records" "test_repository_mirror_records" {
	#Required
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `repository_mirror_record_collection` - The list of repository_mirror_record_collection.

### RepositoryMirrorRecord Reference

The following attributes are exported:

* `mirror_status` - Mirror status of current mirror entry. QUEUED - Mirroring Queued RUNNING - Mirroring is Running PASSED - Mirroring Passed FAILED - Mirroring Failed 
* `time_ended` - Time that the mirror operation ended or null if it hasn't yet ended.
* `time_enqueued` - The time to enqueue a mirror operation.
* `time_started` - The time to start a mirror operation.
* `work_request_id` - Workrequest Id to track current mirror operation

