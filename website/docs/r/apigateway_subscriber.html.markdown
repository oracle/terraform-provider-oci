---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_subscriber"
sidebar_current: "docs-oci-resource-apigateway-subscriber"
description: |-
  Provides the Subscriber resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_subscriber
This resource provides the Subscriber resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new subscriber.

## Example Usage

```hcl
resource "oci_apigateway_subscriber" "test_subscriber" {
	#Required
	clients {
		#Required
		name = var.subscriber_clients_name
		token = var.subscriber_clients_token
	}
	compartment_id = var.compartment_id
	usage_plans = var.subscriber_usage_plans

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.subscriber_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `clients` - (Required) (Updatable) The clients belonging to this subscriber.
	* `name` - (Required) (Updatable) The name of the client. Must be unique within a subscriber.
	* `token` - (Required) (Updatable) The token for the client. Must be unique within a tenancy.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `usage_plans` - (Required) (Updatable) An array of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of usage plan resources. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `clients` - The clients belonging to this subscriber.
	* `name` - The name of the client. Must be unique within a subscriber.
	* `token` - The token for the client. Must be unique within a tenancy.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `state` - The current state of the subscriber.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `usage_plans` - An array of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of usage plan resources. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Subscriber
	* `update` - (Defaults to 20 minutes), when updating the Subscriber
	* `delete` - (Defaults to 20 minutes), when destroying the Subscriber


## Import

Subscribers can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_subscriber.test_subscriber "id"
```

