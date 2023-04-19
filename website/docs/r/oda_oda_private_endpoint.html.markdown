---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_private_endpoint"
sidebar_current: "docs-oci-resource-oda-oda_private_endpoint"
description: |-
  Provides the Oda Private Endpoint resource in Oracle Cloud Infrastructure Digital Assistant service
---

# oci_oda_oda_private_endpoint
This resource provides the Oda Private Endpoint resource in Oracle Cloud Infrastructure Digital Assistant service.

Starts an asynchronous job to create an ODA Private Endpoint.

To monitor the status of the job, take the `opc-work-request-id` response
header value and use it to call `GET /workRequests/{workRequestID}`.


## Example Usage

```hcl
resource "oci_oda_oda_private_endpoint" "test_oda_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.oda_private_endpoint_description
	display_name = var.oda_private_endpoint_display_name
	freeform_tags = {"bar-key"= "value"}
	nsg_ids = var.oda_private_endpoint_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint belongs to.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the ODA private endpoint.
* `display_name` - (Optional) (Updatable) User-defined name for the ODA private endpoint. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Example: `{"bar-key": "value"}` 
* `nsg_ids` - (Optional) (Updatable) List of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm)
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint belongs to.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the ODA private endpoint.
* `display_name` - User-defined name for the ODA private endpoint. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that was assigned when the ODA private endpoint was created.
* `nsg_ids` - List of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of [network security groups](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm)
* `state` - The current state of the ODA private endpoint.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private endpoint belongs to.
* `time_created` - When the resource was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the resource was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oda Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Oda Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Oda Private Endpoint


## Import

OdaPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_oda_oda_private_endpoint.test_oda_private_endpoint "id"
```

