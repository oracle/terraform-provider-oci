---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_source"
sidebar_current: "docs-oci-datasource-ai_vision-stream_source"
description: |-
  Provides details about a specific Stream Source in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_stream_source
This data source provides details about a specific Stream Source resource in Oracle Cloud Infrastructure Ai Vision service.

Get a  streamSource


## Example Usage

```hcl
data "oci_ai_vision_stream_source" "test_stream_source" {
	#Required
	stream_source_id = oci_ai_vision_stream_source.test_stream_source.id
}
```

## Argument Reference

The following arguments are supported:

* `stream_source_id` - (Required) StreamSource Id.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartm. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - display name.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamSource. 
* `state` - The current state of the streamSource.
* `stream_source_details` - Details about a stream source
	* `camera_url` - url of camera
	* `secret_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret where credentials are stored in username:password format. 
	* `source_type` - Type of source Allowed values are:
		* RTSP 
	* `stream_network_access_details` - Details about a stream Connection type
		* `private_endpoint_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private Endpoint 
		* `stream_access_type` - Type of access Allowed values are:
			* PRIVATE 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `time_created` - When the streamSource was created, as an RFC3339 datetime string.
* `time_updated` - When the streamSource was updated, as an RFC3339 datetime string.

