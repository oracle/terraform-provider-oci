---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_iam_work_request_logs"
sidebar_current: "docs-oci-datasource-identity-iam_work_request_logs"
description: |-
  Provides the list of Iam Work Request Logs in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_iam_work_request_logs
This data source provides the list of Iam Work Request Logs in Oracle Cloud Infrastructure Identity service.

Gets logs for a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
iam work request id that can be provided in this API to track the current status of the operation.

- If workrequest exists, returns 202 ACCEPTED
- If workrequest does not exist, returns 404 NOT FOUND


## Example Usage

```hcl
data "oci_identity_iam_work_request_logs" "test_iam_work_request_logs" {
	#Required
	iam_work_request_id = oci_identity_iam_work_request.test_iam_work_request.id
}
```

## Argument Reference

The following arguments are supported:

* `iam_work_request_id` - (Required) The OCID of the IAM work request.


## Attributes Reference

The following attributes are exported:

* `iam_work_request_logs` - The list of iam_work_request_logs.

### IamWorkRequestLog Reference

The following attributes are exported:

* `message` - A human-readable error string.
* `timestamp` - Date and time the log was written, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

