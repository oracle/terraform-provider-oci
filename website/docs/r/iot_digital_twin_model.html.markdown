---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_model"
sidebar_current: "docs-oci-resource-iot-digital_twin_model"
description: |-
  Provides the Digital Twin Model resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_digital_twin_model
This resource provides the Digital Twin Model resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new digital twin model.


## Example Usage

```hcl
resource "oci_iot_digital_twin_model" "test_digital_twin_model" {
	#Required
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id
	spec = var.digital_twin_model_spec

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.digital_twin_model_description
	display_name = var.digital_twin_model_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the resource. If left blank, the description will be derived from the spec.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  If left blank, the display name will be derived from the spec. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `spec` - (Required) The specification of the digital twin model (DTDL).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. If left blank, the description will be derived from the spec.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  If left blank, the display name will be derived from the spec. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `spec_uri` - This is the DTMI (Digital Twin Model Identifier) of the digital twin model as defined in DTDL specification.
* `state` - The current state of the digital twin model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Digital Twin Model
	* `update` - (Defaults to 20 minutes), when updating the Digital Twin Model
	* `delete` - (Defaults to 20 minutes), when destroying the Digital Twin Model


## Import

DigitalTwinModels can be imported using the `id`, e.g.

```
$ terraform import oci_iot_digital_twin_model.test_digital_twin_model "id"
```

