---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_mirror_record"
sidebar_current: "docs-oci-datasource-devops-repository_mirror_record"
description: |-
  Provides details about a specific Repository Mirror Record in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_mirror_record
This data source provides details about a specific Repository Mirror Record resource in Oracle Cloud Infrastructure Devops service.

Returns either current mirror record or last successful mirror record for a specific mirror repository


## Example Usage

```hcl
data "oci_devops_repository_mirror_record" "test_repository_mirror_record" {
	#Required
	mirror_record_type = var.repository_mirror_record_mirror_record_type
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `mirror_record_type` - (Required) The field of mirror record type. Only one mirror record type may be provided. current - The current mirror record. lastSuccessful - The last successful mirror record 
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `mirror_status` - Mirror status of current mirror entry. QUEUED - Mirroring Queued RUNNING - Mirroring is Running PASSED - Mirroring Passed FAILED - Mirroring Failed 
* `time_ended` - Time that the mirror operation ended or null if it hasn't yet ended.
* `time_enqueued` - The time to enqueue a mirror operation.
* `time_started` - The time to start a mirror operation.
* `work_request_id` - Workrequest Id to track current mirror operation

