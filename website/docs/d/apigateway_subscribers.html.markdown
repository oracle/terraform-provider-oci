---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_subscribers"
sidebar_current: "docs-oci-datasource-apigateway-subscribers"
description: |-
  Provides the list of Subscribers in Oracle Cloud Infrastructure API Gateway service
---

# Data Source: oci_apigateway_subscribers
This data source provides the list of Subscribers in Oracle Cloud Infrastructure API Gateway service.

Returns a list of subscribers.

## Example Usage

```hcl
data "oci_apigateway_subscribers" "test_subscribers" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.subscriber_display_name
	state = var.subscriber_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. Example: `ACTIVE` 


## Attributes Reference

The following attributes are exported:

* `subscriber_collection` - The list of subscriber_collection.

### Subscriber Reference

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

