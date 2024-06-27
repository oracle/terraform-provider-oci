---
subcategory: "Delegate Access Control"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_delegate_access_control_delegated_resource_access_request_histories"
sidebar_current: "docs-oci-datasource-delegate_access_control-delegated_resource_access_request_histories"
description: |-
  Provides the list of Delegated Resource Access Request Histories in Oracle Cloud Infrastructure Delegate Access Control service
---

# Data Source: oci_delegate_access_control_delegated_resource_access_request_histories
This data source provides the list of Delegated Resource Access Request Histories in Oracle Cloud Infrastructure Delegate Access Control service.

Returns a history of all status associated with the Delegated Resource Access RequestId.


## Example Usage

```hcl
data "oci_delegate_access_control_delegated_resource_access_request_histories" "test_delegated_resource_access_request_histories" {
	#Required
	delegated_resource_access_request_id = oci_delegate_access_control_delegated_resource_access_request.test_delegated_resource_access_request.id
}
```

## Argument Reference

The following arguments are supported:

* `delegated_resource_access_request_id` - (Required) Unique Delegated Resource Access Request identifier


## Attributes Reference

The following attributes are exported:

* `delegated_resource_access_request_history_collection` - The list of delegated_resource_access_request_history_collection.

### DelegatedResourceAccessRequestHistory Reference

The following attributes are exported:

* `items` - List of DelegatedResourceAccessRequestHistorySummary objects.
	* `comment` - Comment about the status change.
	* `request_status` - The current status of the Delegated Resource Access Request.
	* `state` - The current lifecycle state of the Delegated Resource Access Request.
	* `timestamp` - Time when the respective action happened in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'. 
	* `user_id` - ID of user who modified the Delegated Resource Access Request. For operator, this field is "Operator".

