---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_instance_invoke_raw_command"
sidebar_current: "docs-oci-resource-iot-digital_twin_instance_invoke_raw_command"
description: |-
  Provides the Digital Twin Instance Invoke Raw Command resource in Oracle Cloud Infrastructure Iot service
---

# oci_iot_digital_twin_instance_invoke_raw_command
This resource provides the Digital Twin Instance Invoke Raw Command resource in Oracle Cloud Infrastructure Iot service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iot/latest/DigitalTwinInstance/InvokeRawCommand

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/iot

Invokes the raw command on the specified digital twin instance.


## Example Usage

```hcl
resource "oci_iot_digital_twin_instance_invoke_raw_command" "test_digital_twin_instance_invoke_raw_command" {
	#Required
	digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id
	request_data_format = var.digital_twin_instance_invoke_raw_command_request_data_format
	request_endpoint = var.digital_twin_instance_invoke_raw_command_request_endpoint

	#Optional
	request_data = var.digital_twin_instance_invoke_raw_command_request_data
	request_data_content_type = var.digital_twin_instance_invoke_raw_command_request_data_content_type
	request_duration = var.digital_twin_instance_invoke_raw_command_request_duration
	response_duration = var.digital_twin_instance_invoke_raw_command_response_duration
	response_endpoint = var.digital_twin_instance_invoke_raw_command_response_endpoint
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin instance. 
* `request_data` - (Optional) base 64 encoded request data
* `request_data_content_type` - (Optional) Mime content type of data encoded using base64, default is application/octet-stream
* `request_data_format` - (Required) data format: json, binary, text
* `request_duration` - (Optional) Specified duration by which to send the request by.
* `request_endpoint` - (Required) Device endpoint where request should be forwarded to.
* `response_duration` - (Optional) Specified duration by which to receive the response by.
* `response_endpoint` - (Optional) Device endpoint from which response is expected to come.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Digital Twin Instance Invoke Raw Command
	* `update` - (Defaults to 20 minutes), when updating the Digital Twin Instance Invoke Raw Command
	* `delete` - (Defaults to 20 minutes), when destroying the Digital Twin Instance Invoke Raw Command


## Import

Import is not supported for this resource.

