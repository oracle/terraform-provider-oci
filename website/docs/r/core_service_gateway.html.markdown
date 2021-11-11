---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_service_gateway"
sidebar_current: "docs-oci-resource-core-service_gateway"
description: |-
  Provides the Service Gateway resource in Oracle Cloud Infrastructure Core service
---

# oci_core_service_gateway
This resource provides the Service Gateway resource in Oracle Cloud Infrastructure Core service.

Creates a new service gateway in the specified compartment.

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want
the service gateway to reside. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the service gateway, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = var.compartment_id
	services {
		#Required
		service_id = data.oci_core_services.test_services.services.0.id
	}
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.service_gateway_display_name
	freeform_tags = {"Department"= "Finance"}
	route_table_id = oci_core_route_table.test_route_table.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the service gateway. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `route_table_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the service gateway will use.

	If you don't specify a route table here, the service gateway is created without an associated route table. The Networking service does NOT automatically associate the attached VCN's default route table with the service gateway.

	For information about why you would associate a route table with a service gateway, see [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm). 
* `services` - (Required) (Updatable) List of the OCIDs of the [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) objects to enable for the service gateway. This list can be empty if you don't want to enable any `Service` objects when you create the gateway. You can enable a `Service` object later by using either [AttachServiceId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ServiceGateway/AttachServiceId) or [UpdateServiceGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ServiceGateway/UpdateServiceGateway).

	For each enabled `Service`, make sure there's a route rule with the `Service` object's `cidrBlock` as the rule's destination and the service gateway as the rule's target. See [Route Table](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/RouteTable/). 
	* `service_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/). 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `block_traffic` - Whether the service gateway blocks all traffic through it. The default is `false`. When this is `true`, traffic is not routed to any services, regardless of route rules.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the service gateway. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service gateway. 
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the service gateway is using. For information about why you would associate a route table with a service gateway, see [Transit Routing: Private Access to Oracle Services](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm). 
* `services` - List of the [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) objects enabled for this service gateway. The list can be empty. You can enable a particular `Service` by using [AttachServiceId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ServiceGateway/AttachServiceId) or [UpdateServiceGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ServiceGateway/UpdateServiceGateway). 
	* `service_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service. 
	* `service_name` - The name of the service.
* `state` - The service gateway's current state.
* `time_created` - The date and time the service gateway was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the service gateway belongs to. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Service Gateway
	* `update` - (Defaults to 20 minutes), when updating the Service Gateway
	* `delete` - (Defaults to 20 minutes), when destroying the Service Gateway


## Import

ServiceGateways can be imported using the `id`, e.g.

```
$ terraform import oci_core_service_gateway.test_service_gateway "id"
```

