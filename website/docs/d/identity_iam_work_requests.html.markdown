---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_iam_work_requests"
sidebar_current: "docs-oci-datasource-identity-iam_work_requests"
description: |-
  Provides the list of Iam Work Requests in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_iam_work_requests
This data source provides the list of Iam Work Requests in Oracle Cloud Infrastructure Identity service.

List the IAM work requests in compartment

- If IAM workrequest  details are retrieved sucessfully, return 202 ACCEPTED.
- If any internal error occurs, return 500 INTERNAL SERVER ERROR.


## Example Usage

```hcl
data "oci_identity_iam_work_requests" "test_iam_work_requests" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	resource_identifier = var.iam_work_request_resource_identifier
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `resource_identifier` - (Optional) The identifier of the resource the work request affects.


## Attributes Reference

The following attributes are exported:

* `iam_work_requests` - The list of iam_work_requests.

### IamWorkRequest Reference

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

