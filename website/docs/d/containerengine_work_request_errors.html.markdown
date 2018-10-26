---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_work_request_errors"
sidebar_current: "docs-oci-datasource-containerengine-work_request_errors"
description: |-
  Provides the list of Work Request Errors in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_work_request_errors
This data source provides the list of Work Request Errors in Oracle Cloud Infrastructure Container Engine service.

Get the errors of a work request.

## Example Usage

```hcl
data "oci_containerengine_work_request_errors" "test_work_request_errors" {
	#Required
	compartment_id = "${var.compartment_id}"
	work_request_id = "${oci_containerengine_work_request.test_work_request.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `work_request_id` - (Required) The OCID of the work request.


## Attributes Reference

The following attributes are exported:

* `work_request_errors` - The list of work_request_errors.

### WorkRequestError Reference

The following attributes are exported:

* `code` - A short error code that defines the error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
* `message` - A human-readable error string.
* `timestamp` - The date and time the error occurred.

