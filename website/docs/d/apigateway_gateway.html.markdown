---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_gateway"
sidebar_current: "docs-oci-datasource-apigateway-gateway"
description: |-
  Provides details about a specific Gateway in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_gateway
This data source provides details about a specific Gateway resource in Oracle Cloud Infrastructure API Gateway service.

Gets a gateway by identifier.

## Example Usage

```hcl
data "oci_apigateway_gateway" "test_gateway" {
	#Required
	gateway_id = "${oci_apigateway_gateway.test_gateway.id}"
}
```

## Argument Reference

The following arguments are supported:

* `gateway_id` - (Required) The ocid of the gateway.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.  Example: `My new resource`
* `endpoint_type` - Gateway endpoint type.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `hostname` - The hostname for APIs deployed on the gateway.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
* `state` - The current state of the gateway.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet in which related resources are created.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

