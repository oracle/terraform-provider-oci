---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_connect_harnesses"
sidebar_current: "docs-oci-datasource-streaming-connect_harnesses"
description: |-
  Provides the list of Connect Harnesses in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_connect_harnesses
This data source provides the list of Connect Harnesses in Oracle Cloud Infrastructure Streaming service.

Lists the connectharness.

## Example Usage

```hcl
data "oci_streaming_connect_harnesses" "test_connect_harnesses" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.connect_harness_id
	name = var.connect_harness_name
	state = var.connect_harness_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `id` - (Optional) A filter to return only resources that match the given ID exactly. 
* `name` - (Optional) A filter to return only resources that match the given name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `connect_harness` - The list of connect_harness.

### ConnectHarness Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the connect harness.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the connect harness.
* `lifecycle_state_details` - Any additional details about the current state of the connect harness.
* `name` - The name of the connect harness. Avoid entering confidential information.  Example: `JDBCConnector` 
* `state` - The current state of the connect harness.
* `time_created` - The date and time the connect harness was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

