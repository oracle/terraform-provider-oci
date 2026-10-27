---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_private_endpoint_outbound_connection"
sidebar_current: "docs-oci-resource-integration-private_endpoint_outbound_connection"
description: |-
  Provides the Private Endpoint Outbound Connection resource in Oracle Cloud Infrastructure Integration service
---

# oci_integration_private_endpoint_outbound_connection
This resource provides the Private Endpoint Outbound Connection resource in Oracle Cloud Infrastructure Integration service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/integration/latest/IntegrationInstance/ChangePrivateEndpointOutboundConnection

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/integration

Adds a Private Endpoint Outbound Connection to an existing Integration Instance. The private endpoint is created in the given subnet and lets integrations reach private resources in that VCN (for example databases or APIs that are not reachable from the internet). Destroying this resource removes the private endpoint outbound connection from the instance.

**Note:** The `private_endpoint_outbound_connection` block on [`oci_integration_integration_instance`](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/integration_integration_instance) is a read-only attribute. Setting it on the instance fails with `Can't configure a value for "private_endpoint_outbound_connection": its value will be decided automatically based on the result of applying this configuration.` Use this resource instead.

An integration instance can have at most one private endpoint outbound connection. `subnet_id` and `nsg_ids` cannot be changed in place; to change them, replace the resource (for example with `terraform apply -replace=...`).

Operations on the private endpoint and on an Oracle managed custom endpoint ([`oci_integration_oracle_managed_custom_endpoint`](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/integration_oracle_managed_custom_endpoint)) both put the integration instance into the `UPDATING` state. If you manage both for the same instance, order them with `depends_on` so they are not applied concurrently.

## Example Usage

```hcl
resource "oci_integration_private_endpoint_outbound_connection" "test_private_endpoint_outbound_connection" {
	#Required
	integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}
```

## Argument Reference

The following arguments are supported:

* `integration_instance_id` - (Required) Unique Integration Instance identifier.
* `subnet_id` - (Required) Customer Private Network VCN Subnet OCID.
* `nsg_ids` - (Optional) One or more Network security group Ids.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `integration_instance_id` - Unique Integration Instance identifier.
* `nsg_ids` - One or more Network security group Ids.
* `state` - The current state of the integration instance.
* `subnet_id` - Customer Private Network VCN Subnet OCID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Private Endpoint Outbound Connection
	* `update` - (Defaults to 1 hours), when updating the Private Endpoint Outbound Connection
	* `delete` - (Defaults to 1 hours), when destroying the Private Endpoint Outbound Connection

