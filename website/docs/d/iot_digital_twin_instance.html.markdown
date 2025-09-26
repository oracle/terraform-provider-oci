---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_instance"
sidebar_current: "docs-oci-datasource-iot-digital_twin_instance"
description: |-
  Provides details about a specific Digital Twin Instance in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_instance
This data source provides details about a specific Digital Twin Instance resource in Oracle Cloud Infrastructure Iot service.

Retrieves the digital twin instance identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_digital_twin_instance" "test_digital_twin_instance" {
	#Required
	digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin instance. 


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

