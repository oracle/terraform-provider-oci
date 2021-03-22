---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_instance"
sidebar_current: "docs-oci-resource-oda-oda_instance"
description: |-
  Provides the Oda Instance resource in Oracle Cloud Infrastructure Digital Assistant service
---

# oci_oda_oda_instance
This resource provides the Oda Instance resource in Oracle Cloud Infrastructure Digital Assistant service.

Starts an asynchronous job to create a Digital Assistant instance.

To monitor the status of the job, take the `opc-work-request-id` response
header value and use it to call `GET /workRequests/{workRequestID}`.


## Example Usage

```hcl
resource "oci_oda_oda_instance" "test_oda_instance" {
	#Required
	compartment_id = var.compartment_id
	shape_name = "DEVELOPMENT"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.oda_instance_description
	display_name = var.oda_instance_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Identifier of the compartment.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the Digital Assistant instance.
* `display_name` - (Optional) (Updatable) User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `shape_name` - (Required) Shape or size of the instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Identifier of the compartment that the instance belongs to.
* `connector_url` - URL for the connector's endpoint.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the Digital Assistant instance.
* `display_name` - User-defined name for the Digital Assistant instance. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique immutable identifier that was assigned when the instance was created.
* `lifecycle_sub_state` - The current sub-state of the Digital Assistant instance.
* `shape_name` - Shape or size of the instance.
* `state` - The current state of the Digital Assistant instance.
* `state_message` - A message that describes the current state in more detail. For example, actionable information about an instance that's in the `FAILED` state. 
* `time_created` - When the Digital Assistant instance was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the Digital Assistance instance was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `web_app_url` - URL for the Digital Assistant web application that's associated with the instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oda Instance
	* `update` - (Defaults to 20 minutes), when updating the Oda Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Oda Instance


## Import

OdaInstances can be imported using the `id`, e.g.

```
$ terraform import oci_oda_oda_instance.test_oda_instance "id"
```

