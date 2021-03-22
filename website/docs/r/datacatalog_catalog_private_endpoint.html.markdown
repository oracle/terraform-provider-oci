---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog_private_endpoint"
sidebar_current: "docs-oci-resource-datacatalog-catalog_private_endpoint"
description: |-
  Provides the Catalog Private Endpoint resource in Oracle Cloud Infrastructure Data Catalog service
---

# oci_datacatalog_catalog_private_endpoint
This resource provides the Catalog Private Endpoint resource in Oracle Cloud Infrastructure Data Catalog service.

Create a new private reverse connection endpoint.

## Example Usage

```hcl
resource "oci_datacatalog_catalog_private_endpoint" "test_catalog_private_endpoint" {
	#Required
	compartment_id = var.compartment_id
	dns_zones = var.catalog_private_endpoint_dns_zones
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.catalog_private_endpoint_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Display name of the private endpoint resource being created.
* `dns_zones` - (Required) (Updatable) List of DNS zones to be used by the data assets to be harvested. Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `subnet_id` - (Required) The OCID of subnet to which the reverse connection is to be created 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attached_catalogs` - The list of catalogs using the private reverse connection endpoint
* `compartment_id` - Identifier of the compartment this private endpoint belongs to
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the Private Reverse Connection Endpoint
* `dns_zones` - List of DNS zones to be used by the data assets to be harvested. Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `state` - The current state of the private endpoint resource.
* `subnet_id` - Subnet Identifier
* `time_created` - The time the private endpoint was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the private endpoint was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Catalog Private Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Catalog Private Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Catalog Private Endpoint


## Import

CatalogPrivateEndpoints can be imported using the `id`, e.g.

```
$ terraform import oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint "id"
```

