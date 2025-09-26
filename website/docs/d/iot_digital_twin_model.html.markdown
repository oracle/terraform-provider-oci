---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_model"
sidebar_current: "docs-oci-datasource-iot-digital_twin_model"
description: |-
  Provides details about a specific Digital Twin Model in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_model
This data source provides details about a specific Digital Twin Model resource in Oracle Cloud Infrastructure Iot service.

Retrieves the digital twin model identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_digital_twin_model" "test_digital_twin_model" {
	#Required
	digital_twin_model_id = oci_iot_digital_twin_model.test_digital_twin_model.id
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin model.


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

