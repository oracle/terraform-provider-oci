---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_gateway"
sidebar_current: "docs-oci-resource-apigateway-gateway"
description: |-
  Provides the Gateway resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_gateway
This resource provides the Gateway resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new gateway.


## Example Usage

```hcl
resource "oci_apigateway_gateway" "test_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	endpoint_type = "${var.gateway_endpoint_type}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.gateway_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint_type` - (Required) Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be accessible on a private IP address on the subnet.  Example: `PUBLIC` or `PRIVATE` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint_type` - Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be accessible on a private IP address on the subnet.  Example: `PUBLIC` or `PRIVATE` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for APIs deployed on the gateway.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `state` - The current state of the gateway.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created. 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Import

Gateways can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_gateway.test_gateway "id"
```

