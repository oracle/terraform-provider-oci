---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_service_attachment"
sidebar_current: "docs-oci-resource-fusion_apps-fusion_environment_service_attachment"
description: |-
  Provides the Fusion Environment Service Attachment resource in Oracle Cloud Infrastructure Fusion Apps service
---

# oci_fusion_apps_fusion_environment_service_attachment
This resource provides the Fusion Environment Service Attachment resource in Oracle Cloud Infrastructure Fusion Apps service.

Attaches a service instance to the fusion pod.


## Example Usage

```hcl
resource "oci_fusion_apps_fusion_environment_service_attachment" "test_fusion_environment_service_attachment" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
	service_instance_id = oci_core_instance.test_instance.id
	service_instance_type = var.fusion_environment_service_attachment_service_instance_type
}
```

## Argument Reference

The following arguments are supported:

* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `service_instance_id` - (Required) The service instance OCID of the instance being attached
* `service_instance_type` - (Required) Type of the ServiceInstance being attached.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fusion Environment Service Attachment
	* `update` - (Defaults to 20 minutes), when updating the Fusion Environment Service Attachment
	* `delete` - (Defaults to 20 minutes), when destroying the Fusion Environment Service Attachment


## Import

FusionEnvironmentServiceAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_fusion_apps_fusion_environment_service_attachment.test_fusion_environment_service_attachment "fusionEnvironments/{fusionEnvironmentId}/serviceAttachments/{serviceAttachmentId}" 
```

