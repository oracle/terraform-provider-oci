---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_endpoint"
sidebar_current: "docs-oci-resource-ai_language-endpoint"
description: |-
  Provides the Endpoint resource in Oracle Cloud Infrastructure Ai Language service
---

# oci_ai_language_endpoint
This resource provides the Endpoint resource in Oracle Cloud Infrastructure Ai Language service.

Creates a new endpoint and deploy the trained model


## Example Usage

```hcl
resource "oci_ai_language_endpoint" "test_endpoint" {
	#Required
	compartment_id = var.compartment_id
	model_id = oci_ai_language_model.test_model.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.endpoint_description
	display_name = var.endpoint_display_name
	freeform_tags = {"bar-key"= "value"}
	inference_units = var.endpoint_inference_units
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) compartment identifier for the endpoint
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the an endpoint.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `inference_units` - (Optional) (Updatable) Number of replicas required for this endpoint. This will be optional parameter. Default will be 1.
* `model_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the endpoint compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the endpoint.
* `display_name` - A user-friendly display name for the resource. It should be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier endpoint OCID of an endpoint that is immutable on creation.
* `inference_units` - Number of replicas required for this endpoint.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
* `model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model to associate with the endpoint.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the Endpoint.
* `state` - The state of the endpoint.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the endpoint was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the endpoint was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Endpoint


## Import

Endpoints can be imported using the `id`, e.g.

```
$ terraform import oci_ai_language_endpoint.test_endpoint "id"
```

