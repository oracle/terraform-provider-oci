---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_vision_private_endpoint"
sidebar_current: "docs-oci-datasource-ai_vision-vision_private_endpoint"
description: |-
  Provides details about a specific Vision Private Endpoint in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_vision_private_endpoint
This data source provides details about a specific Vision Private Endpoint resource in Oracle Cloud Infrastructure Ai Vision service.

Get a visionPrivateEndpoint by identifier.

## Example Usage

```hcl
data "oci_ai_vision_vision_private_endpoint" "test_vision_private_endpoint" {
	#Required
	vision_private_endpoint_id = oci_ai_vision_vision_private_endpoint.test_vision_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `vision_private_endpoint_id` - (Required) Vision private endpoint Id.


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

