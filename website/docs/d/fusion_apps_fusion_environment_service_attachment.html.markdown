---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_service_attachment"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_service_attachment"
description: |-
  Provides details about a specific Fusion Environment Service Attachment in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_service_attachment
This data source provides details about a specific Fusion Environment Service Attachment resource in Oracle Cloud Infrastructure Fusion Apps service.

Gets a Service Attachment by identifier

## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_service_attachment" "test_fusion_environment_service_attachment" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	service_attachment_id = oci_fusion_apps_service_attachment.test_service_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `service_attachment_id` - (Required) OCID of the Service Attachment


## Attributes Reference

The following attributes are exported:

* `action` - Action
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

