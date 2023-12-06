---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_service_attachments"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_service_attachments"
description: |-
  Provides the list of Fusion Environment Service Attachments in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_service_attachments
This data source provides the list of Fusion Environment Service Attachments in Oracle Cloud Infrastructure Fusion Apps service.

Returns a list of service attachments.


## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_service_attachments" "test_fusion_environment_service_attachments" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

	#Optional
	display_name = var.fusion_environment_service_attachment_display_name
	service_instance_type = var.fusion_environment_service_attachment_service_instance_type
	state = var.fusion_environment_service_attachment_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `service_instance_type` - (Optional) A filter that returns all resources that match the specified lifecycle state.
* `state` - (Optional) A filter that returns all resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `service_attachment_collection` - The list of service_attachment_collection.

### FusionEnvironmentServiceAttachment Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Service Attachment Display name, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `is_sku_based` - Whether this service is provisioned due to the customer being subscribed to a specific SKU
* `service_instance_id` - The ID of the service instance created that can be used to identify this on the service control plane
* `service_instance_type` - Type of the serviceInstance.
* `service_url` - Public URL
* `state` - The current state of the ServiceInstance.
* `time_created` - The time the the ServiceInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the ServiceInstance was updated. An RFC3339 formatted datetime string

