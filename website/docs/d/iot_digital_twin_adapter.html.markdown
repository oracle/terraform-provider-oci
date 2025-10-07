---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_adapter"
sidebar_current: "docs-oci-datasource-iot-digital_twin_adapter"
description: |-
  Provides details about a specific Digital Twin Adapter in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_adapter
This data source provides details about a specific Digital Twin Adapter resource in Oracle Cloud Infrastructure Iot service.

Retrieves the digital twin adapter identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_digital_twin_adapter" "test_digital_twin_adapter" {
	#Required
	digital_twin_adapter_id = oci_iot_digital_twin_adapter.test_digital_twin_adapter.id
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_adapter_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin adapter. 


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `digital_twin_model_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the digital twin model.
* `digital_twin_model_spec_uri` - The URI of the digital twin model specification.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `inbound_envelope` - Payload containing device-specific metadata and optional value mappings used to interpret or transform that metadata. This structure includes the device endpoint, the actual payload, and an optional envelope mapping that applies [JQ](https://stedolan.github.io/jq/) expressions to extract or reshape the data as needed. 
	* `envelope_mapping` - Maps the metadata fields from the inbound payload using JQ. These mappings allow you to extract specific metadata such as timestamps using JQ expressions. 
		* `time_observed` - JQ expression to extract the observation timestamp from the payload. If not specified, the system will default to using `timeReceived` as the timestamp.  Example: For payload `{"time": "<timestamp>","temp": 65,"hum": 55}` 'timeObserved' can be mapped as [JQ Expression](https://jqplay.org/) `$.time`. 
	* `reference_endpoint` - The device endpoint. 
	* `reference_payload` - Reference payload structure template received from IoT device. This payload must specify its content type using the `dataFormat` property. 
		* `data` - JSON raw data.
		* `data_format` - Data format of the payload.
* `inbound_routes` - list of routes
	* `condition` - A boolean expression used to determine whether the following transformation should be processed for the incoming payload. This expression is typically based on fields defined at the inbound Envelope and is evaluated before applying the `payloadMapping`. 
	* `description` - Meaningful write up about the inbound route. 
	* `payload_mapping` - A set of key-value JQ expressions used to transform the incoming payload into a shape compatible with the digital twin model's context or schema.

		The keys are target fields (in the digital twin model), and values are JQ expressions pointing to data in the reference payload.

		Example: Given payload: { "time": "<timestamp>", "temp": 65, "hum": 55 } And mapping: { "temperature": "$.temp", "humidity": "$.hum", "timeObserved": "$.time" } The output will be: { "temperature": 65, "humidity": 55, "timeObserved": "<timestamp>" } 
	* `reference_payload` - Reference payload structure template received from IoT device. This payload must specify its content type using the `dataFormat` property. 
		* `data` - JSON raw data.
		* `data_format` - Data format of the payload.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `state` - The current state of the digital twin adapter.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

