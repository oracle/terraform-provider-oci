---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_instance"
sidebar_current: "docs-oci-resource-iot-digital_twin_instance"
description: |-
  Provides the Digital Twin Instance resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_digital_twin_instance
This resource provides the Digital Twin Instance resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new digital twin instance.


## Example Usage

```hcl
resource "oci_iot_digital_twin_instance" "test_digital_twin_instance" {
	#Required
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id

	#Optional
	auth_id = oci_iot_auth.test_auth.id
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.digital_twin_instance_description
	digital_twin_adapter_id = oci_iot_digital_twin_adapter.test_digital_twin_adapter.id
	digital_twin_model_id = oci_iot_digital_twin_model.test_digital_twin_model.id
	digital_twin_model_spec_uri = var.digital_twin_instance_digital_twin_model_spec_uri
	display_name = var.digital_twin_instance_display_name
	external_key = var.digital_twin_instance_external_key
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `auth_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource (like VaultSecret, ClientCertificate etc.,) used to authenticate the digital twin instance.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the resource. 
* `digital_twin_adapter_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin adapter.
* `digital_twin_model_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model.
* `digital_twin_model_spec_uri` - (Optional) The URI of the digital twin model specification.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `external_key` - (Optional) (Updatable) A unique identifier for the physical entity (typically an IoT device) represented by the digital twin instance. This could be a Bluetooth address, Ethernet MAC address, or serial number, depending on the use case. If not provided, the system will automatically generate one. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `auth_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource (like VaultSecret, ClientCertificate etc.,) used to authenticate the digital twin instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `digital_twin_adapter_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin adapter.
* `digital_twin_model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model.
* `digital_twin_model_spec_uri` - The URI of the digital twin model specification.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `external_key` - A unique identifier for the physical entity (typically an IoT device) represented by the digital twin instance. This could be a Bluetooth address, Ethernet MAC address, or serial number, depending on the use case. If not provided, the system will automatically generate one. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `state` - The current state of the digital twin instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Digital Twin Instance
	* `update` - (Defaults to 20 minutes), when updating the Digital Twin Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Digital Twin Instance


## Import

DigitalTwinInstances can be imported using the `id`, e.g.

```
$ terraform import oci_iot_digital_twin_instance.test_digital_twin_instance "id"
```

