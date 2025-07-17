---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_sources"
sidebar_current: "docs-oci-datasource-ai_vision-stream_sources"
description: |-
  Provides the list of Stream Sources in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_stream_sources
This data source provides the list of Stream Sources in Oracle Cloud Infrastructure Ai Vision service.

Gets a list of the streamSources in the specified compartment.


## Example Usage

```hcl
data "oci_ai_vision_stream_sources" "test_stream_sources" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.stream_source_display_name
	id = var.stream_source_id
	state = var.stream_source_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The filter to find the device with the given identifier.
* `state` - (Optional) The filter to match projects with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `stream_source_collection` - The list of stream_source_collection.

### StreamSource Reference

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

