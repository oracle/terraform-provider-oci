---
layout: "oci"
page_title: "OCI: oci_containerengine_work_request_log_entries"
sidebar_current: "docs-oci-datasource-containerengine-work_request_log_entries"
description: |-
Provides a list of WorkRequestLogEntries
---
# Data Source: oci_containerengine_work_request_log_entries
The WorkRequestLogEntries data source allows access to the list of OCI work_request_log_entries

Get the logs of a work request.

## Example Usage

```hcl
data "oci_containerengine_work_request_log_entries" "test_work_request_log_entries" {
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

* `work_request_log_entries` - The list of work_request_log_entries.

### WorkRequestLogEntry Reference

The following attributes are exported:

* `message` - The description of an action that occurred.
* `timestamp` - The date and time the log entry occurred.

