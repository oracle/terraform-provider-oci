---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_group"
sidebar_current: "docs-oci-resource-ai_vision-stream_group"
description: |-
  Provides the Stream Group resource in Oracle Cloud Infrastructure Ai Vision service
---

# oci_ai_vision_stream_group
This resource provides the Stream Group resource in Oracle Cloud Infrastructure Ai Vision service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/vision/latest/StreamGroup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/aiVision

Registration of new streamGroup


## Example Usage

```hcl
resource "oci_ai_vision_stream_group" "test_stream_group" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = var.stream_group_defined_tags
	display_name = var.stream_group_display_name
	freeform_tags = var.stream_group_freeform_tags
	is_enabled = var.stream_group_is_enabled
	stream_overlaps {

		#Optional
		overlapping_streams = var.stream_group_stream_overlaps_overlapping_streams
	}
	stream_source_ids = var.stream_group_stream_source_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - (Optional) (Updatable) A human-friendly name for the streamGroup.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `is_enabled` - (Optional) (Updatable) Stream
* `stream_overlaps` - (Optional) (Updatable) List of streamSource OCIDs where the streamSource overlaps in field of view.
	* `overlapping_streams` - (Optional) (Updatable) List of streamSource OCIDs.
* `stream_source_ids` - (Optional) (Updatable) List of streamSource OCIDs associated with the stream group


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `display_name` - A human-friendly name for the streamGroup.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the streamGroup. 
* `is_enabled` - Stream
* `state` - The current state of the streamGroup.
* `stream_overlaps` - List of streamSource OCIDs where the streamSource overlaps in field of view.
	* `overlapping_streams` - List of streamSource OCIDs.
* `stream_source_ids` - List of streamSource OCIDs associated with the stream group
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `time_created` - When the streamGroup was created, as an RFC3339 datetime string.
* `time_updated` - When the streamGroup was updated, as an RFC3339 datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Stream Group
	* `update` - (Defaults to 20 minutes), when updating the Stream Group
	* `delete` - (Defaults to 20 minutes), when destroying the Stream Group


## Import

StreamGroups can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_stream_group.test_stream_group "id"
```

