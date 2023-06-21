---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_iam_work_request"
sidebar_current: "docs-oci-datasource-identity-iam_work_request"
description: |-
  Provides details about a specific Iam Work Request in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_iam_work_request
This data source provides details about a specific Iam Work Request resource in Oracle Cloud Infrastructure Identity service.

Gets details on a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
iam work request id that can be provided in this API to track the current status of the operation.

- If workrequest exists, returns 202 ACCEPTED
- If workrequest does not exist, returns 404 NOT FOUND


## Example Usage

```hcl
data "oci_identity_iam_work_request" "test_iam_work_request" {
	#Required
	iam_work_request_id = oci_identity_iam_work_request.test_iam_work_request.id
}
```

## Argument Reference

The following arguments are supported:

* `iam_work_request_id` - (Required) The OCID of the IAM work request.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing this IAM work request.
* `id` - The OCID of the work request.
* `operation_type` - The asynchronous operation tracked by this IAM work request.
* `percent_complete` - How much progress the operation has made. 
* `resources` - The resources this work request affects.
	* `action_type` - The way in which this resource is affected by the work tracked in the work request. A resource being created, updated, or deleted will remain in the IN_PROGRESS state until work is complete for that resource at which point it will transition to CREATED, UPDATED, or DELETED, respectively. 
	* `entity_type` - The resource type the work request is affects.
	* `entity_uri` - The URI path that the user can do a GET on to access the resource metadata.
	* `identifier` - An OCID of the resource that the work request affects.
* `status` - Status of the work request
* `time_accepted` - Date and time the work was accepted, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `time_finished` - Date and time the work completed, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `time_started` - Date and time the work started, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

