---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_vision_private_endpoints"
sidebar_current: "docs-oci-datasource-ai_vision-vision_private_endpoints"
description: |-
  Provides the list of Vision Private Endpoints in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_vision_private_endpoints
This data source provides the list of Vision Private Endpoints in Oracle Cloud Infrastructure Ai Vision service.

Returns a list of visionPrivateEndpoints.


## Example Usage

```hcl
data "oci_ai_vision_vision_private_endpoints" "test_vision_private_endpoints" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.vision_private_endpoint_display_name
	id = var.vision_private_endpoint_id
	state = var.vision_private_endpoint_state
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

* `vision_private_endpoint_collection` - The list of vision_private_endpoint_collection.

### VisionPrivateEndpoint Reference

The following attributes are exported:

* `compartment_id` - A compartment identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - An optional description of the visionPrivateEndpoint.
* `display_name` - A human-friendly name for the visionPrivateEndpoint, which can be changed.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of private endpoint 
* `lifecycle_details` - A message describing the current state in more detail, that can provide actionable information if creation failed.
* `state` - The current state of the visionPrivateEndpoint.
* `subnet_id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of subnet 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `time_created` - When the visionPrivateEndpoint was created, as an RFC3339 datetime string.
* `time_updated` - When the visionPrivateEndpoint was updated, as an RFC3339 datetime string.

