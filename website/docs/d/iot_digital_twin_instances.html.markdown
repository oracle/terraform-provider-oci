---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_instances"
sidebar_current: "docs-oci-datasource-iot-digital_twin_instances"
description: |-
  Provides the list of Digital Twin Instances in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_instances
This data source provides the list of Digital Twin Instances in Oracle Cloud Infrastructure Iot service.

Retrieves a list of digital twin instances within the specified IoT domain.


## Example Usage

```hcl
data "oci_iot_digital_twin_instances" "test_digital_twin_instances" {
	#Required
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id

	#Optional
	digital_twin_model_id = oci_iot_digital_twin_model.test_digital_twin_model.id
	digital_twin_model_spec_uri = var.digital_twin_instance_digital_twin_model_spec_uri
	display_name = var.digital_twin_instance_display_name
	id = var.digital_twin_instance_id
	state = var.digital_twin_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_model_id` - (Optional) Filter resources that match the specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model. 
* `digital_twin_model_spec_uri` - (Optional) Filter resources that match the specified URI (DTMI) of the digital twin model. 
* `display_name` - (Optional) Filter resources whose display name matches the specified value. 
* `id` - (Optional) Filter resources by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type. 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources. 
* `state` - (Optional) Filter resources whose lifecycleState matches the specified value. 


## Attributes Reference

The following attributes are exported:

* `digital_twin_instance_collection` - The list of digital_twin_instance_collection.

### DigitalTwinInstance Reference

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

