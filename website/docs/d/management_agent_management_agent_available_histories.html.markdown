---
subcategory: "Management Agent"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_agent_management_agent_available_histories"
sidebar_current: "docs-oci-datasource-management_agent-management_agent_available_histories"
description: |-
  Provides the list of Management Agent Available Histories in Oracle Cloud Infrastructure Management Agent service
---

# Data Source: oci_management_agent_management_agent_available_histories
This data source provides the list of Management Agent Available Histories in Oracle Cloud Infrastructure Management Agent service.

Lists the availability history records of Management Agent

## Example Usage

```hcl
data "oci_management_agent_management_agent_available_histories" "test_management_agent_available_histories" {
	#Required
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id

	#Optional
	time_availability_status_ended_greater_than = var.management_agent_available_history_time_availability_status_ended_greater_than
	time_availability_status_started_less_than = var.management_agent_available_history_time_availability_status_started_less_than
}
```

## Argument Reference

The following arguments are supported:

* `management_agent_id` - (Required) Unique Management Agent identifier
* `time_availability_status_ended_greater_than` - (Optional) Filter to limit the availability history results to that of time after the input time including the boundary record. Defaulted to current date minus one year. The date and time to be given as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 
* `time_availability_status_started_less_than` - (Optional) Filter to limit the availability history results to that of time before the input time including the boundary record Defaulted to current date. The date and time to be given as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29. 


## Attributes Reference

The following attributes are exported:

* `availability_histories` - The list of availability_histories.

### ManagementAgentAvailableHistory Reference

The following attributes are exported:

* `availability_status` - The availability status of managementAgent
* `management_agent_id` - agent identifier
* `time_availability_status_ended` - The time till which the Management Agent was known to be in the availability status. An RFC3339 formatted datetime string
* `time_availability_status_started` - The time at which the Management Agent moved to the availability status. An RFC3339 formatted datetime string

