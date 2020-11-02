---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_enrollment_status"
sidebar_current: "docs-oci-datasource-optimizer-enrollment_status"
description: |-
  Provides details about a specific Enrollment Status in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_enrollment_status
This data source provides details about a specific Enrollment Status resource in Oracle Cloud Infrastructure Optimizer service.

Gets the Cloud Advisor enrollment status.


## Example Usage

```hcl
data "oci_optimizer_enrollment_status" "test_enrollment_status" {
	#Required
	enrollment_status_id = oci_optimizer_enrollment_status.test_enrollment_status.id
}
```

## Argument Reference

The following arguments are supported:

* `enrollment_status_id` - (Required) The unique OCID associated with the enrollment status.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `id` - The OCID of the enrollment status.
* `state` - The enrollment status' current state.
* `status` - The current Cloud Advisor enrollment status.
* `status_reason` - The reason for the enrollment status of the tenancy.
* `time_created` - The date and time the enrollment status was created, in the format defined by RFC3339.
* `time_updated` - The date and time the enrollment status was last updated, in the format defined by RFC3339.

