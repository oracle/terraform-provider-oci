---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_event_report"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instance_event_report"
description: |-
  Provides details about a specific Managed Instance Event Report in Oracle Cloud Infrastructure OS Management service
---

# Data Source: oci_osmanagement_managed_instance_event_report
This data source provides details about a specific Managed Instance Event Report resource in Oracle Cloud Infrastructure OS Management service.

Get summary information about events on this instance.


## Example Usage

```hcl
data "oci_osmanagement_managed_instance_event_report" "test_managed_instance_event_report" {
	#Required
	compartment_id = var.compartment_id
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id

	#Optional
	latest_timestamp_greater_than_or_equal_to = var.managed_instance_event_report_latest_timestamp_greater_than_or_equal_to
	latest_timestamp_less_than = var.managed_instance_event_report_latest_timestamp_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `latest_timestamp_greater_than_or_equal_to` - (Optional) filter event occurrence. Selecting only those last occurred on or after given date in ISO 8601 format Example: 2017-07-14T02:40:00.000Z 
* `latest_timestamp_less_than` - (Optional) filter event occurrence. Selecting only those last occurred before given date in ISO 8601 format Example: 2017-07-14T02:40:00.000Z 
* `managed_instance_id` - (Required) Instance Oracle Cloud identifier (ocid)


## Attributes Reference

The following attributes are exported:

* `count` - count of events currently registered on the system.

