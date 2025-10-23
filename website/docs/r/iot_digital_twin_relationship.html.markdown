---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_relationship"
sidebar_current: "docs-oci-resource-iot-digital_twin_relationship"
description: |-
  Provides the Digital Twin Relationship resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_digital_twin_relationship
This resource provides the Digital Twin Relationship resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new digital twin relationship.


## Example Usage

```hcl
resource "oci_iot_digital_twin_relationship" "test_digital_twin_relationship" {
	#Required
	content_path = var.digital_twin_relationship_content_path
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id
	source_digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id
	target_digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id

	#Optional
	content = var.digital_twin_relationship_content
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.digital_twin_relationship_description
	display_name = var.digital_twin_relationship_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `content` - (Optional) (Updatable) The value(s) of the relationship properties defined in the source digital twin model.
* `content_path` - (Required) Its the name of the relationship that links two digital twin instances. Here, it is the relationship name of the source digital twin model. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the resource. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `source_digital_twin_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance. 
* `target_digital_twin_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `content` - The value(s) of the relationship properties defined in the source digital twin model.
* `content_path` - Its the name of the relationship that links two digital twin instances. Here, it is the relationship name of the source digital twin model. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `source_digital_twin_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance. 
* `state` - The current state of the digital twin relationship.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_digital_twin_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Digital Twin Relationship
	* `update` - (Defaults to 20 minutes), when updating the Digital Twin Relationship
	* `delete` - (Defaults to 20 minutes), when destroying the Digital Twin Relationship


## Import

DigitalTwinRelationships can be imported using the `id`, e.g.

```
$ terraform import oci_iot_digital_twin_relationship.test_digital_twin_relationship "id"
```

