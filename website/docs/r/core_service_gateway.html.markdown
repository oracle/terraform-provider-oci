---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_service_gateway"
sidebar_current: "docs-oci-resource-core-service_gateway"
description: |-
  Provides the Service Gateway resource in Oracle Cloud Infrastructure Core service
---

# oci_core_service_gateway
This resource provides the Service Gateway resource in Oracle Cloud Infrastructure Core service.

Creates a new service gateway in the specified compartment.

For the purposes of access control, you must provide the OCID of the compartment where you want
the service gateway to reside. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the service gateway, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_service_gateway" "test_service_gateway" {
	#Required
	compartment_id = "${var.compartment_id}"
	services {
		#Required
		service_id = "${oci_core_service.test_service.id}"
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "foo-value"}
	display_name = "${var.service_gateway_display_name}"
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)  of the compartment to contain the Service Gateway. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `services` - (Required) (Updatable) List of the Service OCIDs. These are the Services which will be enabled on the Service Gateway. This list can be empty.
	* `service_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service. 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `block_traffic` - Whether the service gateway blocks all traffic through it. The default is `false`. When this is `true`, traffic is not routed to any services, regardless of route rules.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the service gateway. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "foo-value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service gateway.
* `services` - List of the services enabled on this service gateway. The list can be empty. You can enable a particular service by using [AttachServiceId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/ServiceGateway/AttachServiceId). 
	* `service_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service. 
	* `service_name` - The name of the service.
* `state` - The service gateway's current state.
* `time_created` - The date and time the service gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the service gateway belongs to. 

## Import

ServiceGateways can be imported using the `id`, e.g.

```
$ terraform import oci_core_service_gateway.test_service_gateway "id"
```

