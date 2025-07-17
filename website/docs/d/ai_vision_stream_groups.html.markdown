---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_stream_groups"
sidebar_current: "docs-oci-datasource-ai_vision-stream_groups"
description: |-
  Provides the list of Stream Groups in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_stream_groups
This data source provides the list of Stream Groups in Oracle Cloud Infrastructure Ai Vision service.

Gets a list of the streamGroups in the specified compartment.


## Example Usage

```hcl
data "oci_ai_vision_stream_groups" "test_stream_groups" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.stream_group_display_name
	id = var.stream_group_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) The filter to find the device with the given identifier.


## Attributes Reference

The following attributes are exported:

* `stream_group_collection` - The list of stream_group_collection.

### StreamGroup Reference

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

