---
subcategory: "Streaming"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_streaming_connect_harness"
sidebar_current: "docs-oci-datasource-streaming-connect_harness"
description: |-
  Provides details about a specific Connect Harness in Oracle Cloud Infrastructure Streaming service
---

# Data Source: oci_streaming_connect_harness
This data source provides details about a specific Connect Harness resource in Oracle Cloud Infrastructure Streaming service.

Gets detailed information about a connect harness.

## Example Usage

```hcl
data "oci_streaming_connect_harness" "test_connect_harness" {
	#Required
	connect_harness_id = "${oci_streaming_connect_harnes.test_connect_harnes.id}"
}
```

## Argument Reference

The following arguments are supported:

* `connect_harness_id` - (Required) The OCID of the connect harness. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the connect harness.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}' 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the connect harness.
* `lifecycle_state_details` - Any additional details about the current state of the connect harness.
* `name` - The name of the connect harness. Avoid entering confidential information.  Example: `JDBCConnector` 
* `state` - The current state of the connect harness.
* `time_created` - The date and time the connect harness was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z` 

