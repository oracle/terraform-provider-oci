---
subcategory: "Psa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psa_psa_work_requests"
sidebar_current: "docs-oci-datasource-psa-psa_work_requests"
description: |-
  Provides the list of Psa Work Requests in Oracle Cloud Infrastructure Psa service
---

# Data Source: oci_psa_psa_work_requests
This data source provides the list of Psa Work Requests in Oracle Cloud Infrastructure Psa service.

Lists the PrivateServiceAccess work requests in a compartment.


## Example Usage

```hcl
data "oci_psa_psa_work_requests" "test_psa_work_requests" {

	#Optional
	compartment_id = var.compartment_id
	resource_id = oci_cloud_guard_resource.test_resource.id
	status = var.psa_work_request_status
	work_request_id = oci_psa_psa_work_request.test_psa_work_request.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource affected by the work request.
* `status` - (Optional) A filter to return only the resources that match the given lifecycle state.
* `work_request_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.


## Attributes Reference

The following attributes are exported:

* `work_request_summary_collection` - The list of work_request_summary_collection.

### PsaWorkRequest Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
* `operation_type` - The asynchronous operation tracked by this work request.
* `percent_complete` - Shows the progress of the operation tracked by the work request, as a percentage of the total work that must be performed. 
* `resources` - The resources that are affected by the work request.
	* `action_type` - The way in which this resource is affected by the operation tracked in the work request. A resource being created, updated, or deleted remains in the IN_PROGRESS state until work is complete for that resource, at which point it transitions to CREATED, UPDATED, or DELETED, respectively. 
	* `entity_type` - The resource type that the work request affects.
	* `entity_uri` - The URI path that you can use for a GET request to access the resource metadata.
	* `identifier` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) or other unique identifier for the resource.
	* `metadata` - Additional information that helps to explain the resource.
* `status` - The status of the work request.
* `time_accepted` - The date and time the work request was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_finished` - The date and time the work request was finished, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_started` - The date and time the work request was started, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). 
* `time_updated` - The date and time the work request was updated, in the format defined by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

