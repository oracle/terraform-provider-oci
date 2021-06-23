---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_enrollment_status"
sidebar_current: "docs-oci-resource-optimizer-enrollment_status"
description: |-
  Provides the Enrollment Status resource in Oracle Cloud Infrastructure Optimizer service
---

# oci_optimizer_enrollment_status
This resource provides the Enrollment Status resource in Oracle Cloud Infrastructure Optimizer service.

Updates the enrollment status of the tenancy.


## Example Usage

```hcl
resource "oci_optimizer_enrollment_status" "test_enrollment_status" {
	#Required
	enrollment_status_id = oci_optimizer_enrollment_status.test_enrollment_status.id
	status = var.enrollment_status_status
}
```

## Argument Reference

The following arguments are supported:

* `enrollment_status_id` - (Required) The unique OCID associated with the enrollment status.
* `status` - (Required) (Updatable) The Cloud Advisor enrollment status.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `id` - The OCID of the enrollment status.
* `state` - The enrollment status' current state.
* `status` - The current Cloud Advisor enrollment status.
* `status_reason` - The reason for the enrollment status of the tenancy.
* `time_created` - The date and time the enrollment status was created, in the format defined by RFC3339.
* `time_updated` - The date and time the enrollment status was last updated, in the format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Enrollment Status
	* `update` - (Defaults to 20 minutes), when updating the Enrollment Status
	* `delete` - (Defaults to 20 minutes), when destroying the Enrollment Status


## Import

EnrollmentStatus can be imported using the `id`, e.g.

```
$ terraform import oci_optimizer_enrollment_status.test_enrollment_status "id"
```

