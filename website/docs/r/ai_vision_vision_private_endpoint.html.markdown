---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_vision_private_endpoint"
sidebar_current: "docs-oci-resource-ai_vision-vision_private_endpoint"
description: |-
  Provides the Vision Private Endpoint resource in Oracle Cloud Infrastructure Ai Vision service
---

# oci_ai_vision_vision_private_endpoint
This resource provides the Vision Private Endpoint resource in Oracle Cloud Infrastructure Ai Vision service.

Create a new visionPrivateEndpoint.


## Example Usage

```hcl
resource "oci_ai_vision_vision_private_endpoint" "test_vision_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = var.vision_private_endpoint_defined_tags
	description = var.vision_private_endpoint_description
	display_name = var.vision_private_endpoint_display_name
	freeform_tags = var.vision_private_endpoint_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - (Optional) (Updatable) An optional description of the visionPrivateEndpoint.
* `display_name` - (Optional) (Updatable) A human-friendly name for the visionPrivateEndpoint, that can be changed.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `subnet_id` - (Required) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of subnet 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vision Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Vision Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Vision Private Endpoint


## Import

VisionPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_ai_vision_vision_private_endpoint.test_vision_private_endpoint "id"
```

