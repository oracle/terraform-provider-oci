
---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_enrollment_statuses"
sidebar_current: "docs-oci-datasource-optimizer-enrollment_statuses"
description: |-
  Provides the list of Enrollment Statuses in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_enrollment_statuses
This data source provides the list of Enrollment Statuses in Oracle Cloud Infrastructure Optimizer service.

Lists the Cloud Advisor enrollment statuses.


## Example Usage

```hcl
data "oci_optimizer_enrollment_statuses" "test_enrollment_statuses" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	state = var.enrollment_status_state
	status = var.enrollment_status_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `state` - (Optional) A filter that returns results that match the lifecycle state specified. 
* `status` - (Optional) A filter that returns results that match the Cloud Advisor enrollment status specified. 


## Attributes Reference

The following attributes are exported:

* `enrollment_status_collection` - The list of enrollment_status_collection.

### EnrollmentStatus Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `id` - The OCID of the enrollment status.
* `state` - The enrollment status' current state.
* `status` - The current Cloud Advisor enrollment status.
* `status_reason` - The reason for the enrollment status of the tenancy.
* `time_created` - The date and time the enrollment status was created, in the format defined by RFC3339.
* `time_updated` - The date and time the enrollment status was last updated, in the format defined by RFC3339.

