---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_iam_work_request_errors"
sidebar_current: "docs-oci-datasource-identity-iam_work_request_errors"
description: |-
  Provides the list of Iam Work Request Errors in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_iam_work_request_errors
This data source provides the list of Iam Work Request Errors in Oracle Cloud Infrastructure Identity service.

Gets error details for a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
iam work request id that can be provided in this API to track the current status of the operation.

- If workrequest exists, returns 202 ACCEPTED
- If workrequest does not exist, returns 404 NOT FOUND


## Example Usage

```hcl
data "oci_identity_iam_work_request_errors" "test_iam_work_request_errors" {
	#Required
	iam_work_request_id = oci_identity_iam_work_request.test_iam_work_request.id
}
```

## Argument Reference

The following arguments are supported:

* `iam_work_request_id` - (Required) The OCID of the IAM work request.


## Attributes Reference

The following attributes are exported:

* `iam_work_request_errors` - The list of iam_work_request_errors.

### IamWorkRequestError Reference

The following attributes are exported:

* `code` - A machine-usable code for the error that occured.
* `message` - A human-readable error string.
* `timestamp` - The date and time the error occurred.

