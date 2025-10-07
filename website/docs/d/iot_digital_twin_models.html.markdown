---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_models"
sidebar_current: "docs-oci-datasource-iot-digital_twin_models"
description: |-
  Provides the list of Digital Twin Models in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_models
This data source provides the list of Digital Twin Models in Oracle Cloud Infrastructure Iot service.

Retrieves a list of digital twin models within the specified IoT domain.


## Example Usage

```hcl
data "oci_iot_digital_twin_models" "test_digital_twin_models" {
	#Required
	iot_domain_id = oci_iot_iot_domain.test_iot_domain.id

	#Optional
	display_name = var.digital_twin_model_display_name
	id = var.digital_twin_model_id
	spec_uri_starts_with = var.digital_twin_model_spec_uri_starts_with
	state = var.digital_twin_model_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) Filter resources whose display name matches the specified value. 
* `id` - (Optional) Filter resources by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a valid OCID of the resource type. 
* `iot_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain in which to list digital twin resources. 
* `spec_uri_starts_with` - (Optional) Filters resources by spec URI prefix. For example, to search all versions of the `dtmi:example:device;1` model, pass the prefix without the version: `dtmi:example:device`. 
* `state` - (Optional) Filter resources whose lifecycleState matches the specified value. 


## Attributes Reference

The following attributes are exported:

* `digital_twin_model_collection` - The list of digital_twin_model_collection.

### DigitalTwinModel Reference

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

