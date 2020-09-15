---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_work_request_log_entries"
sidebar_current: "docs-oci-datasource-containerengine-work_request_log_entries"
description: |-
  Provides the list of Work Request Log Entries in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_work_request_log_entries
This data source provides the list of Work Request Log Entries in Oracle Cloud Infrastructure Container Engine service.

Get the logs of a work request.

## Example Usage

```hcl
data "oci_containerengine_work_request_log_entries" "test_work_request_log_entries" {
	#Required
	compartment_id = var.compartment_id
	work_request_id = oci_containerengine_work_request.test_work_request.id
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

