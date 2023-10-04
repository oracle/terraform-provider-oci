---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_ekms_private_endpoint"
sidebar_current: "docs-oci-resource-kms-ekms_private_endpoint"
description: |-
  Provides the Ekms Private Endpoint resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_ekms_private_endpoint
This resource provides the Ekms Private Endpoint resource in Oracle Cloud Infrastructure Kms service.

Create a new EKMS private endpoint used to connect to external key manager system

## Example Usage

```hcl
resource "oci_kms_ekms_private_endpoint" "test_ekms_private_endpoint" {
	#Required
	ca_bundle = var.ekms_private_endpoint_ca_bundle
	compartment_id = var.compartment_id
	display_name = var.ekms_private_endpoint_display_name
	external_key_manager_ip = var.ekms_private_endpoint_external_key_manager_ip
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	port = var.ekms_private_endpoint_port
}
```

## Argument Reference

The following arguments are supported:

* `ca_bundle` - (Required) CABundle to validate TLS certificate of the external key manager system in PEM format 
* `compartment_id` - (Required) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Display name of the EKMS private endpoint resource being created.
* `external_key_manager_ip` - (Required) External private IP to connect to from this EKMS private endpoint 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `port` - (Optional) The port of the external key manager system
* `subnet_id` - (Required) The OCID of subnet in which the EKMS private endpoint is to be created 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `ca_bundle` - CABundle to validate TLS certificate of the external key manager system in PEM format 
* `compartment_id` - Identifier of the compartment this EKMS private endpoint belongs to
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the EKMS private endpoint
* `external_key_manager_ip` - Private IP of the external key manager system to connect to from the EKMS private endpoint 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `port` - The port of the external key manager system
* `private_endpoint_ip` - The IP address in the customer's VCN for the EKMS private endpoint. This is taken from subnet
* `state` - The current state of the EKMS private endpoint resource.
* `subnet_id` - Subnet Identifier
* `time_created` - The time the EKMS private endpoint was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the EKMS private endpoint was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ekms Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Ekms Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Ekms Private Endpoint


## Import

EkmsPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_kms_ekms_private_endpoint.test_ekms_private_endpoint "id"
```

