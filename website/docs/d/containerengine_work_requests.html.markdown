---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_work_requests"
sidebar_current: "docs-oci-datasource-containerengine-work_requests"
description: |-
  Provides the list of Work Requests in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_work_requests
This data source provides the list of Work Requests in Oracle Cloud Infrastructure Container Engine service.

List all work requests in a compartment.

## Example Usage

```hcl
data "oci_containerengine_work_requests" "test_work_requests" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"
	resource_id = "${oci_containerengine_resource.test_resource.id}"
	resource_type = "${var.work_request_resource_type}"
	status = "${var.work_request_status}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The OCID of the cluster.
* `compartment_id` - (Required) The OCID of the compartment.
* `resource_id` - (Optional) The OCID of the resource associated with a work request
* `resource_type` - (Optional) Type of the resource associated with a work request
* `status` - (Optional) A work request status to filter on. Can have multiple parameters of this name.


## Attributes Reference

The following attributes are exported:

* `work_requests` - The list of work_requests.

### WorkRequest Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment in which the work request exists.
* `id` - The OCID of the work request.
* `operation_type` - The type of work the work request is doing.
* `resources` - The resources this work request affects.
	* `action_type` - The way in which this resource was affected by the work tracked by the work request.
	* `entity_type` - The resource type the work request affects.
	* `entity_uri` - The URI path on which the user can issue a GET request to access the resource metadata.
	* `identifier` - The OCID of the resource the work request affects.
* `status` - The current status of the work request.
* `time_accepted` - The time the work request was accepted.
* `time_finished` - The time the work request was finished.
* `time_started` - The time the work request was started.

