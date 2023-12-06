---
subcategory: "Operator Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_operator_access_control_access_request_history"
sidebar_current: "docs-oci-datasource-operator_access_control-access_request_history"
description: |-
  Provides details about a specific Access Request History in Oracle Cloud Infrastructure Operator Access Control service
---

# Data Source: oci_operator_access_control_access_request_history
This data source provides details about a specific Access Request History resource in Oracle Cloud Infrastructure Operator Access Control service.

Returns a history of all status associated with the accessRequestId.


## Example Usage

```hcl
data "oci_operator_access_control_access_request_history" "test_access_request_history" {
	#Required
	access_request_id = oci_operator_access_control_access_request.test_access_request.id
}
```

## Argument Reference

The following arguments are supported:

* `access_request_id` - (Required) unique AccessRequest identifier


## Attributes Reference

The following attributes are exported:

* `items` - contains AccessRequestHistorySummary
	* `actions_list` - List of operator actions for which approvals were requested by the operator.
	* `description` - Reason or description about the cause of change.
	* `duration` - Duration for approval of request or extension depending on the type of action.
	* `is_auto_approved` - Whether the access request was automatically approved.
	* `state` - The current state of the AccessRequest.
	* `time_of_action` - Time when the respective action happened in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
	* `user_id` - Approver who modified the access request.

