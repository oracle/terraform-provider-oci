---
subcategory: "Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_integration_oracle_managed_custom_endpoint"
sidebar_current: "docs-oci-resource-integration-oracle_managed_custom_endpoint"
description: |-
  Provides the Oracle Managed Custom Endpoint resource in Oracle Cloud Infrastructure Integration service
---

# oci_integration_oracle_managed_custom_endpoint
This resource provides the Oracle Managed Custom Endpoint resource in Oracle Cloud Infrastructure Integration service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/integration/latest/IntegrationInstance/AddOracleManagedCustomEndpoint

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/integration

Enables an Oracle managed custom endpoint on an existing Integration Instance. Oracle creates and renews the certificate and the DNS record for the custom hostname in the given OCI DNS zone. Destroying this resource removes the Oracle managed custom endpoint from the instance.

**Note:** For a customer managed custom endpoint (you supply the certificate in a Vault secret), use the `custom_endpoint` block on [`oci_integration_integration_instance`](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/integration_integration_instance) instead.

The integration instance must be allowed to manage the DNS zone and to obtain certificates. See the [example](https://github.com/oracle/terraform-provider-oci/tree/master/examples/integration) for a dynamic group matching the instance's IDCS application (`idcs_info[0].idcs_app_name`) and a policy granting it `manage dns-zones` and `manage dns-records`, plus `ENDORSE any-user TO MANAGE certificate-authority-family IN any-tenancy`.

`hostname` and `dns_zone_name` cannot be changed in place; to change them, replace the resource (for example with `terraform apply -replace=...`).

Operations on an Oracle managed custom endpoint and on a private endpoint outbound connection ([`oci_integration_private_endpoint_outbound_connection`](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/integration_private_endpoint_outbound_connection)) both put the integration instance into the `UPDATING` state. If you manage both for the same instance, order them with `depends_on` so they are not applied concurrently.

## Example Usage

```hcl
resource "oci_integration_oracle_managed_custom_endpoint" "test_oracle_managed_custom_endpoint" {
	#Required
	integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
	hostname = var.oracle_managed_custom_endpoint_hostname

	#Optional
	dns_zone_name = var.oracle_managed_custom_endpoint_dns_zone_name
}
```

## Argument Reference

The following arguments are supported:

* `integration_instance_id` - (Required) Unique Integration Instance identifier.
* `hostname` - (Required) Oracle managed custom hostname, in FQDN format. Must be within the DNS zone given by `dns_zone_name`.
* `dns_zone_name` - (Optional) Name of the OCI DNS zone in which Oracle creates the DNS record for the custom hostname.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `dns_type` - Type of DNS. Currently `OCI`.
* `dns_zone_name` - DNS Zone name.
* `hostname` - The custom hostname used for the integration instance URL, in FQDN format.
* `integration_instance_id` - Unique Integration Instance identifier.
* `managed_type` - Indicates if the custom endpoint is managed by Oracle or by the customer. `ORACLE_MANAGED` for endpoints created by this resource.
* `state` - The current state of the integration instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Oracle Managed Custom Endpoint
	* `update` - (Defaults to 1 hours), when updating the Oracle Managed Custom Endpoint
	* `delete` - (Defaults to 1 hours), when destroying the Oracle Managed Custom Endpoint

