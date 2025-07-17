---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_source"
sidebar_current: "docs-oci-resource-ai_vision-stream_source"
description: |-
  Provides the Stream Source resource in Oracle Cloud Infrastructure Ai Vision service
---

# oci_ai_vision_stream_source
This resource provides the Stream Source resource in Oracle Cloud Infrastructure Ai Vision service.

Registration of new streamSource


## Example Usage

```hcl
resource "oci_ai_vision_stream_source" "test_stream_source" {
	#Required
	compartment_id = var.compartment_id
	stream_source_details {
		#Required
		camera_url = var.stream_source_stream_source_details_camera_url
		source_type = var.stream_source_stream_source_details_source_type
		stream_network_access_details {
			#Required
			private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
			stream_access_type = var.stream_source_stream_source_details_stream_network_access_details_stream_access_type
		}

		#Optional
		secret_id = oci_vault_secret.test_secret.id
	}

	#Optional
	defined_tags = var.stream_source_defined_tags
	display_name = var.stream_source_display_name
	freeform_tags = var.stream_source_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - (Optional) (Updatable) A human-friendly name for the streamSource.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `stream_source_details` - (Required) (Updatable) Details about a stream source
	* `camera_url` - (Required) (Updatable) url of camera
	* `secret_id` - (Optional) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret where credentials are stored in username:password format. 
	* `source_type` - (Required) (Updatable) Type of source Allowed values are:
		* RTSP 
	* `stream_network_access_details` - (Required) (Updatable) Details about a stream Connection type
		* `private_endpoint_id` - (Required) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private Endpoint 
		* `stream_access_type` - (Required) (Updatable) Type of access Allowed values are:
			* PRIVATE 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Source
	* `update` - (Defaults to 20 minutes), when updating the Stream Source
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Source


## Import

StreamSources can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_stream_source.test_stream_source "id"
```

